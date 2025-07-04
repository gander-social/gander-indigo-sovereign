// Copied from indigo:api/atproto/repolistRecords.go

package agnostic

// schema: com.atproto.repo.getRecord

import (
	"context"
	"encoding/json"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// RepoGetRecord_Output is the output of a com.atproto.repo.getRecord call.
type RepoGetRecord_Output struct {
	Cid *string `json:"cid,omitempty" cborgen:"cid,omitempty"`
	Uri string  `json:"uri" cborgen:"uri"`
	//  NOTE: changed from lex decoder to json.RawMessage
	Value *json.RawMessage `json:"value" cborgen:"value"`
}

// RepoGetRecord calls the XRPC method "com.atproto.repo.getRecord".
//
// cid: The CID of the version of the record. If not specified, then return the most recent version.
// collection: The NSID of the record collection.
// repo: The handle or DID of the repo.
// rkey: The Record Key.
func RepoGetRecord(ctx context.Context, c util.LexClient, cid string, collection string, repo string, rkey string) (*RepoGetRecord_Output, error) {
	var out RepoGetRecord_Output

	params := map[string]interface{}{
		"collection": collection,
		"repo":       repo,
		"rkey":       rkey,
	}
	if cid != "" {
		params["cid"] = cid
	}
	if err := c.LexDo(ctx, util.Query, "", "com.atproto.repo.getRecord", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
