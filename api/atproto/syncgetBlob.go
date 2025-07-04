// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.getBlob

import (
	"bytes"
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// SyncGetBlob calls the XRPC method "com.atproto.sync.getBlob".
//
// cid: The CID of the blob to fetch
// did: The DID of the account.
func SyncGetBlob(ctx context.Context, c util.LexClient, cid string, did string) ([]byte, error) {
	buf := new(bytes.Buffer)

	params := map[string]interface{}{}
	params["cid"] = cid
	params["did"] = did
	if err := c.LexDo(ctx, util.Query, "", "com.atproto.sync.getBlob", params, nil, buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
