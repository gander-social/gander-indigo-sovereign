// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.checkAccountStatus

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// ServerCheckAccountStatus_Output is the output of a com.atproto.server.checkAccountStatus call.
type ServerCheckAccountStatus_Output struct {
	Activated          bool   `json:"activated" cborgen:"activated"`
	ExpectedBlobs      int64  `json:"expectedBlobs" cborgen:"expectedBlobs"`
	ImportedBlobs      int64  `json:"importedBlobs" cborgen:"importedBlobs"`
	IndexedRecords     int64  `json:"indexedRecords" cborgen:"indexedRecords"`
	PrivateStateValues int64  `json:"privateStateValues" cborgen:"privateStateValues"`
	RepoBlocks         int64  `json:"repoBlocks" cborgen:"repoBlocks"`
	RepoCommit         string `json:"repoCommit" cborgen:"repoCommit"`
	RepoRev            string `json:"repoRev" cborgen:"repoRev"`
	ValidDid           bool   `json:"validDid" cborgen:"validDid"`
}

// ServerCheckAccountStatus calls the XRPC method "com.atproto.server.checkAccountStatus".
func ServerCheckAccountStatus(ctx context.Context, c util.LexClient) (*ServerCheckAccountStatus_Output, error) {
	var out ServerCheckAccountStatus_Output
	if err := c.LexDo(ctx, util.Query, "", "com.atproto.server.checkAccountStatus", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
