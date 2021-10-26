package main

import (
	"io"
	"strconv"

	mapset "github.com/deckarep/golang-set"
	"github.com/github/git-sizer/git"
)

//  Global OID and ID tables
var (
	IDs             = NewIDs()
	Id_hash         = make(map[int32]string)
	Hash_id         = make(map[string]int32)
	SKIPPED_COMMITS = mapset.NewSet()
)

type RepoFilter struct {
	repo    *Repository
	opts    Options
	input   io.PipeWriter
	output  io.PipeReader
	parser  RepoParser
	targets map[git.OID]string // blob oid => file name
}

func (filter *RepoFilter) tweak_blob(blob *Blob) {
	for target := range filter.targets {
		if target.String() == blob.original_oid {
			// set new id to 0
			blob.ele.skip(0)
		}
	}
}

func (filter *RepoFilter) tweak_commit(commit *Commit, helper *Helper_info) {

	// 如果没有from, 但是有filechange，则可能是first commit
	// if commit.first_parent() == 0 {
	// 	fmt.Println("DEBUG: this is a init commit!")
	// }
	// 如果有from，但是没有filechange， 则可能是merge commit
	// if len(commit.filechanges) == 0 {
	// 	fmt.Println("DEBUG: this is a merge commit!")
	// }

	// orig_parents := helper.orig_parents
	// parents := commit.parents
	old_1st_parent := commit.first_parent()

	// 如果filechange中，查询不到from-id，则需要删除该条记录，
	// 如果整个filechange都没了，则需要删除该commit
	filter_filechange(commit)

	if len(commit.filechanges) == 0 {
		commit.skip(old_1st_parent)
	}

	// 如果 from-id 在ID-hash中能够查询到，则正常，否则说明parent commit被删了
	// 或者，如果from-id在Skipped-commit中能够查询到，则也需要skip
	if SKIPPED_COMMITS.Contains(old_1st_parent) {
		commit.skip(old_1st_parent)
	}
}

func filter_filechange(commit *Commit) {
	newfilechanges := make([]FileChange, 0)
	for _, filechange := range commit.filechanges {

		// *if id is 40-byte hash-1 then its a subdirectory
		if len(filechange.blob_id) == 40 {
			continue
		}
		// #TODO，在Skip_blob中查询应该更快
		if filechange.changetype == "M" {
			id, _ := strconv.Atoi(filechange.blob_id)
			if _, ok := Id_hash[int32(id)]; !ok {
				continue
			}
		}
		// otherwise, keep it in newfilechange
		newfilechanges = append(newfilechanges, filechange)
	}
	commit.filechanges = newfilechanges
}

func (filter *RepoFilter) tweak_reset(reset *Reset) {
	Lasted_commit[reset.ref] = reset.from
	Lasted_orig_commit[reset.ref] = reset.from
}

func (filter *RepoFilter) tweak_tag(tag *Tag) {
	// the tag may have no parents, if so skip it
	if SKIPPED_COMMITS.Contains(tag.from_ref) == false {
		tag.ele.base.dumped = false
		tag.ele.skip(0)
	}
}
