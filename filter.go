package main

import (
	"os"
	"strconv"
)

type RepoFilter struct {
	repo      *Repository
	scanned   []string // file's oid provided by scanner
	filepaths []string // files(or dir) provided by user
}

func (filter *RepoFilter) tweak_blob(blob *Blob) {
	for _, target := range filter.scanned {
		if target == blob.original_oid {
			// set new id to 0
			blob.ele.skip(0)
		}
	}
}

func (filter *RepoFilter) tweak_commit(commit *Commit, helper *Helper_info) {

	// 如果没有from, 但是有filechange，则可能是first commit
	// 如果有from，但是没有filechange， 则可能是merge commit

	// 如果存在from，且from：0, 说明是从第一个blob就开始删除了，后面都是0
	if len(commit.parents) != 0 && commit.parents[0] == 0 {
		commit.skip(0)
	}

	old_1st_parent := commit.first_parent()

	filter_filechange(commit, filter)

	if len(commit.filechanges) == 0 {
		commit.skip(old_1st_parent)
	}

	// 如果 from-id 在ID-hash中能够查询到，则正常，否则说明parent commit被删了
	// 或者，如果from-id在Skipped-commit中能够查询到，则也需要skip
	if SKIPPED_COMMITS.Contains(old_1st_parent) {
		commit.skip(old_1st_parent)
	}
}

func filter_filechange(commit *Commit, filter *RepoFilter) {
	newfilechanges := make([]FileChange, 0)
	matched := false
	for _, filechange := range commit.filechanges {
		if filter.repo.opts.scan && len(filter.scanned) != 0 {
			// filter by blob oid
			for _, target := range filter.scanned {
				if len(filechange.blob_id) == 40 {
					if target == filechange.blob_id {
						Branch_changed.Add(filechange.branch)
						matched = true
						break
					}
				} else {
					if filechange.changetype == "M" {
						id, _ := strconv.Atoi(filechange.blob_id)
						if _, ok := ID_HASH[int32(id)]; !ok {
							Branch_changed.Add(filechange.branch)
							matched = true
							break
						}
					}
				}
			}
		} else if !filter.repo.opts.scan && filter.repo.opts.limit != "" && filter.repo.opts.types == "*" {
			// filter by blob size threshold
			objectsize := Blob_size_list[filechange.blob_id]
			// set bitsize to 64, means max single blob size is 4 GiB
			size, _ := strconv.ParseUint(objectsize, 10, 64)
			limit, err := UnitConvert(filter.repo.opts.limit)
			if err != nil {
				ft := LocalPrinter().Sprintf("convert uint error: %s", err)
				PrintRedln(ft)
				os.Exit(1)
			}
			if size > limit {
				Branch_changed.Add(filechange.branch)
				matched = true
				break
			}
		} else if !filter.repo.opts.scan && len(filter.filepaths) != 0 {
			// filter by blob name or directory
			for _, filepath := range filter.filepaths {
				matches := Match(filepath, filechange.filepath)
				if len(matches) != 0 {
					matched = true
					Branch_changed.Add(filechange.branch)
					break
				}
			}
		}
		if matched {
			// skip this file
			continue
		}
		// otherwise, keep it in newfilechange
		newfilechanges = append(newfilechanges, filechange)
	}
	commit.filechanges = newfilechanges
}

func (filter *RepoFilter) tweak_reset(reset *Reset) {
	if SKIPPED_COMMITS.Contains(reset.from) == true {
		reset.base.dumped = false
		reset.base.skip()
	}
}

func (filter *RepoFilter) tweak_tag(tag *Tag) {
	// the tag may have no parent, if so skip it
	if SKIPPED_COMMITS.Contains(tag.from_ref) == true {
		tag.ele.base.dumped = false
		tag.ele.skip(0)
	}
}
