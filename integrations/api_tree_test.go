// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	api "code.gitea.io/sdk/gitea"

	"github.com/stretchr/testify/assert"
)

func testAPIGetTree(t *testing.T, treePath string, exists bool, entries []*api.TreeEntry) {
	prepareTestEnv(t)

	session := loginUser(t, "user2")
	req := NewRequestf(t, "GET", "/api/v1/repos/user2/repo1/trees/%s", treePath)
	resp := session.MakeRequest(t, req, NoExpectedStatus)
	if !exists {
		assert.EqualValues(t, http.StatusNotFound, resp.HeaderCode)
		return
	}
	assert.EqualValues(t, http.StatusOK, resp.HeaderCode)
	fmt.Print(bytes.NewBuffer(resp.Body))
	var trees []*api.TreeEntry
	DecodeJSON(t, resp, &trees)

	if assert.EqualValues(t, len(entries), len(trees)) {
		for i, tree := range trees {
			assert.EqualValues(t, entries[i], tree)
		}
	}
}

func TestAPIGetTree(t *testing.T) {
	for _, test := range []struct {
		TreePath string
		Exists     bool
		Entries  []*api.TreeEntry
	}{
		{"master", true, []*api.TreeEntry{
			&api.TreeEntry{
				Name: "README.md",
				ID: "4b4851ad51df6a7d9f25c979345979eaeb5b349f",
				Type: "blob",
				Size: 30,
			},
		}},
		{"master/doesnotexist", false, []*api.TreeEntry{}},
		{"feature/1", true, []*api.TreeEntry{
			&api.TreeEntry{
				Name: "README.md",
				ID: "4b4851ad51df6a7d9f25c979345979eaeb5b349f",
				Type: "blob",
				Size: 30,
			},
		}},
		{"feature/1/doesnotexist", false, []*api.TreeEntry{}},
	} {
		testAPIGetTree(t, test.TreePath, test.Exists, test.Entries)
	}
}

