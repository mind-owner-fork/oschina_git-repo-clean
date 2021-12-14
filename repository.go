package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"path/filepath"
)

type Repository struct {
	path   string
	gitBin string
	gitDir string
	opts   Options
}

type HistoryRecord struct {
	oid        string
	objectSize uint64
	objectName string
}

type BlobList []HistoryRecord

var Blob_size_list = make(map[string]string)

func (repo Repository) GetBlobName(oid string) (string, error) {
	cmd := exec.Command(repo.gitBin, "-C", repo.path, "rev-list", "--objects", "--all")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return "", err
	}
	blobname := ""
	buf := bufio.NewReader(out)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		// drop LF
		line = line[:len(line)-1]

		if len(line) <= 41 {
			continue
		}
		texts := strings.Split(line, " ")
		if texts[0] == oid {
			// blob(file) name maybe lile: "bad dir/readme.md
			blobname = strings.Join(texts[1:], " ")
		}
	}
	return blobname, nil
}

func parseBatchHeader(header string) (objectid, objecttype, objectsize string, err error) {
	// drop LF
	header = header[:len(header)-1]

	infos := strings.Split(header, " ")
	if infos[len(infos)-1] == "missing" {
		return "", "", "", errors.New("got missing object")
	}
	return infos[0], infos[1], infos[2], nil
}

func GetBlobSize(repo Repository) error {

	cmd := exec.Command(repo.gitBin, "-C", repo.path, "cat-file", "--batch-all-objects",
		"--batch-check=%(objectname) %(objecttype) %(objectsize)")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return err
	}

	buf := bufio.NewReader(out)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		objectid, objecttype, objectsize, err := parseBatchHeader(line)
		if err != nil {
			return err
		}
		if objecttype == "blob" {
			Blob_size_list[objectid] = objectsize
		}
	}
	return nil
}

func ScanRepository(repo Repository) (BlobList, error) {

	var empty BlobList
	var blobs BlobList

	if repo.opts.verbose {
		PrintLocalWithGreenln("start scanning")
	}

	for objectid, objectsize := range Blob_size_list {
		limit, err := UnitConvert(repo.opts.limit)
		if err != nil {
			return empty, fmt.Errorf(LocalPrinter().Sprintf(
				"convert uint error: %s", err))
		}
		size, err := strconv.ParseUint(objectsize, 10, 0)
		if err != nil {
			return empty, fmt.Errorf(LocalPrinter().Sprintf(
				"parse uint error: %s", err))
		}
		if size > limit {
			name, err := repo.GetBlobName(objectid)
			if err != nil {
				if err != io.EOF {
					return empty, fmt.Errorf(LocalPrinter().Sprintf(
						"run GetBlobName error: %s", err))
				}
			}
			if len(repo.opts.types) != 0 || repo.opts.types != "*" {
				var pattern string
				if strings.HasSuffix(name, "\"") {
					pattern = "." + repo.opts.types + "\"$"
				} else {
					pattern = "." + repo.opts.types + "$"
				}
				if matches := Match(pattern, name); len(matches) == 0 {
					// matched none, skip
					continue
				}
			}
			// append this record blob into slice
			blobs = append(blobs, HistoryRecord{objectid, size, name})
			// sort according by size
			sort.Slice(blobs, func(i, j int) bool {
				return blobs[i].objectSize > blobs[j].objectSize
			})
			// remain first [op.number] blobs
			if len(blobs) > int(repo.opts.number) {
				blobs = blobs[:repo.opts.number]
				// break
			}
		}
	}
	return blobs, nil
}

// check if the current repository is bare repo
func IsBare(gitbin, path string) (bool, error) {

	cmd := exec.Command(gitbin, "-C", path, "rev-parse", "--is-bare-repository")
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(LocalPrinter().Sprintf(
			"could not run 'git rev-parse --is-bare-repository': %s", err),
		)
	}

	return string(bytes.TrimSpace(out)) == "true", nil
}

// check if the current repository is shallow repo, need Git version 2.15.0 or newer
func IsShallow(gitbin, path string) (bool, error) {
	cmd := exec.Command(gitbin, "-C", path, "rev-parse", "--is-shallow-repository")
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(LocalPrinter().Sprintf(
			"could not run 'git rev-parse --is-shallow-repository': %s", err),
		)
	}

	return string(bytes.TrimSpace(out)) == "true", nil
}

// check if the current repository is flesh clone.
func IsFresh(gitbin, path string) (bool, error) {
	cmd := exec.Command(gitbin, "-C", path, "reflog", "show")
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(LocalPrinter().Sprintf(
			"could not run 'git reflog show': %s", err),
		)
	}
	return strings.Count(string(out), "\n") < 2, nil
}

// check if Git-LFS has installed in host machine
func HasLFSEnv(gitbin, path string) (bool, error) {
	cmd := exec.Command(gitbin, "lfs", "version")
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(LocalPrinter().Sprintf("could not run 'git lfs version': %s", err))
	}

	return strings.Contains(string(out), "git-lfs"), nil
}

// get git version string
func GitVersion(gitbin, path string) (string, error) {
	cmd := exec.Command(gitbin, "version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(LocalPrinter().Sprintf("could not run 'git version': %s", err))
	}
	matches := Match("[0-9]+.[0-9]+.[0-9]+.?[0-9]?", string(out))
	if len(matches) == 0 {
		return "", errors.New(LocalPrinter().Sprintf("match git version wrong"))
	}
	return matches[0], nil
}

// convert version string to int number for compare. e.g. convert 2.33.0 to 2330
func GitVersionConvert(version string) int {
	var vstr string
	dict := strings.Split(version, ".")
	if len(dict) == 3 {
		vstr = dict[0] + dict[1] + dict[2]
	}
	if len(dict) == 4 {
		vstr = dict[0] + dict[1] + dict[2] + dict[3]
	}
	vstr = strings.TrimSuffix(vstr, "\n")
	ret, err := strconv.Atoi(vstr)
	if err != nil {
		return 0
	}
	return ret
}

// get current branch
func GetCurrentBranch(gitbin, path string) (string, error) {
	cmd := exec.Command(gitbin, "-C", path, "symbolic-ref", "HEAD", "--short")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(LocalPrinter().Sprintf("could not run 'git symbolic-ref HEAD --short': %s", err))
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

// get current status
func GetCurrentStatus(gitbin, path string) {
	cmd := exec.Command(gitbin, "-C", path, "status", "-z")
	out, err := cmd.Output()
	if err != nil {
		PrintLocalWithRedln("could not run 'git status'")
	}
	st := string(out)
	if st == "" {
		PrintLocalWithGreenln("git status clean")
	}
	fmt.Println(st)
}

func GetDatabaseSize(gitbin, path string) string {
	path = filepath.Join(path, ".git")
	cmd := exec.Command("du", "-hs", path)
	out, err := cmd.Output()
	if err != nil {
		PrintLocalWithRedln("could not run 'du -hs .git'")
	}
	return strings.TrimSuffix(string(out), "\n")
}

// get repo GC url if the repo is hosted on Gitee.com
func GetGiteeGCWeb(gitbin, path string) string {
	cmd := exec.Command(gitbin, "-C", path, "config", "--get", "remote.origin.url")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	url := string(out)
	if url == "" {
		return ""
	}
	if strings.Contains(url, "gitee.com") {
		if strings.HasPrefix(url, "git@") {
			url = strings.TrimPrefix(url, "git@")
			url = "https://" + strings.Replace(url, ":", "/", 1)
		}
		url = strings.TrimSuffix(url, ".git\n") + "/settings#git-gc"
	} else {
		return ""
	}
	return url
}

func (repo *Repository) BackUp(gitbin, path string) {
	PrintLocalWithGreenln("start preparing repository data")
	// #TODO specify backup directory by option
	dst := "../backup.bak"
	// check if the same directory exist
	_, err := os.Stat(dst)
	if err == nil {
		ok := AskForOverride()
		if !ok {
			PrintLocalWithYellowln("backup canceled")
			return
		} else {
			os.RemoveAll(dst)
		}
	}
	PrintLocalWithGreen("start backup")
	cmd := exec.Command(gitbin, "-C", path, "clone", "--no-local", path, dst)
	_, err = cmd.Output()
	if err != nil {
		PrintLocalWithRedln("clone error")
		fmt.Println(err)
		return
	}
	abs_path, err := filepath.Abs(dst)
	if err != nil {
		PrintLocalWithRedln("run filepach.Abs error")
	}
	ft := LocalPrinter().Sprintf("backup done! Backup file path is: %s", abs_path)
	PrintGreenln(ft)
}

func (repo *Repository) PushRepo(gitbin, path string) error {
	cmd := exec.Command(gitbin, "-C", path, "push", "origin", "--all", "--force", "--porcelain")
	out1, err := cmd.Output()
	if err != nil {
		PrintLocalWithRedln("Push failed")
		return err
	}
	PrintYellowln(strings.TrimSuffix(string(out1), "\n"))

	cmd2 := exec.Command(gitbin, "-C", path, "push", "origin", "--tags", "--force")
	out2, err := cmd2.Output()
	if err != nil {
		PrintLocalWithRedln("Push failed")
		return err
	}
	PrintYellowln(strings.TrimSuffix(string(out2), "\n"))
	PrintLocalWithYellowln("Done")
	return nil
}

func NewRepository(path string) (*Repository, error) {
	// Find the `git` executable to be used:
	gitBin, err := findGitBin()
	if err != nil {
		return nil, fmt.Errorf(
			"could not find 'git' executable (is it in your PATH?): %v", err,
		)
	}
	gitdir, err := GitDir(gitBin, path)
	if err != nil {
		return &Repository{}, err
	}

	if bare, err := IsBare(gitBin, path); err != nil || bare {
		return &Repository{}, err
	}

	if shallow, err := IsShallow(gitBin, path); err != nil || shallow {
		return &Repository{}, err
	}

	return &Repository{
		path:   path,   // working dir
		gitDir: gitdir, // .git dir
		gitBin: gitBin,
	}, nil
}

func (repo *Repository) Close() error {
	return nil
}

func (repo *Repository) CleanUp() {
	// clean up
	PrintLocalWithGreenln("file cleanup is complete. Start cleaning the repository")

	fmt.Println("running git reset --hard")
	cmd1 := exec.Command(repo.gitBin, "-C", repo.path, "reset", "--hard")
	cmd1.Stdout = os.Stdout
	err := cmd1.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd1.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("running git reflog expire --expire=now --all")
	cmd2 := exec.Command(repo.gitBin, "-C", repo.path, "reflog", "expire", "--expire=now", "--all")
	cmd2.Stderr = os.Stderr
	cmd2.Stdout = os.Stdout
	err = cmd2.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd2.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("running git gc --prune=now")
	cmd3 := exec.Command(repo.gitBin, "-C", repo.path, "gc", "--prune=now")
	cmd3.Stderr = os.Stderr
	cmd3.Stdout = os.Stdout
	err = cmd3.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd3.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
