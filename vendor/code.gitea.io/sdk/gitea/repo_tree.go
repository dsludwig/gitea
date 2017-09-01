// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
)


// ObjectType git object type
type ObjectType string

const (
	// ObjectCommit commit object type
	ObjectCommit ObjectType = "commit"
	// ObjectTree tree object type
	ObjectTree ObjectType = "tree"
	// ObjectBlob blob object type
	ObjectBlob ObjectType = "blob"
	// ObjectTag tag object type
	ObjectTag ObjectType = "tag"
)

// Tree represents a repository tree (directory) entry.
type TreeEntry struct {
	Name   string             `json:"name"`
	ID     string             `json:"sha1"`
	Size   int64              `json:"size"`
	Type   ObjectType         `json:"type"`
}

// GetRepoBranch get one tree's information of one repository
func (c *Client) GetRepoTree(user, repo, treePath string) (*TreeEntry, error) {
	t := new(TreeEntry)
	return t, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/trees/%s", user, repo, treePath), nil, nil, &t)
}
