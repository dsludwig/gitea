// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	api "code.gitea.io/sdk/gitea"
	"code.gitea.io/git"

	// "code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	// "code.gitea.io/gitea/routers/repo"
)

// GetTree get a tree (directory) by path on a repository
func GetTree(ctx *context.APIContext) {
	if !ctx.Repo.HasAccess() {
		ctx.Status(404)
		return
	}

	tree, err := ctx.Repo.Commit.SubTree(ctx.Repo.TreePath)
	if err != nil {
		ctx.NotFoundOrServerError("Repo.Commit.SubTree", git.IsErrNotExist, err)
		return
	}

	entries, err := tree.ListEntries()
	if err != nil {
		ctx.Handle(500, "ListEntries", err)
		return
	}
	entries.Sort()

	moreInfo := make([]*api.TreeEntry, len(entries))
	for i, entry := range entries {
		moreInfo[i] = &api.TreeEntry {
			entry.Name(),
			entry.ID.String(),
			entry.Size(),
			api.ObjectType(entry.Type),
		}
	}

	ctx.JSON(200, moreInfo)
}
