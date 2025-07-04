// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.signature.findRelatedAccounts

import (
	"context"

	comatprototypes "github.com/gander-social/gander-indigo-sovereign/api/atproto"
	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// SignatureFindRelatedAccounts_Output is the output of a tools.ozone.signature.findRelatedAccounts call.
type SignatureFindRelatedAccounts_Output struct {
	Accounts []*SignatureFindRelatedAccounts_RelatedAccount `json:"accounts" cborgen:"accounts"`
	Cursor   *string                                        `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
}

// SignatureFindRelatedAccounts_RelatedAccount is a "relatedAccount" in the tools.ozone.signature.findRelatedAccounts schema.
type SignatureFindRelatedAccounts_RelatedAccount struct {
	Account      *comatprototypes.AdminDefs_AccountView `json:"account" cborgen:"account"`
	Similarities []*SignatureDefs_SigDetail             `json:"similarities,omitempty" cborgen:"similarities,omitempty"`
}

// SignatureFindRelatedAccounts calls the XRPC method "tools.ozone.signature.findRelatedAccounts".
func SignatureFindRelatedAccounts(ctx context.Context, c util.LexClient, cursor string, did string, limit int64) (*SignatureFindRelatedAccounts_Output, error) {
	var out SignatureFindRelatedAccounts_Output

	params := map[string]interface{}{}
	if cursor != "" {
		params["cursor"] = cursor
	}
	params["did"] = did
	if limit != 0 {
		params["limit"] = limit
	}
	if err := c.LexDo(ctx, util.Query, "", "tools.ozone.signature.findRelatedAccounts", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
