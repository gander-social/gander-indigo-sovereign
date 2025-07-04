// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.unspecced.getSuggestionsSkeleton

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// UnspeccedGetSuggestionsSkeleton_Output is the output of a gndr.app.unspecced.getSuggestionsSkeleton call.
type UnspeccedGetSuggestionsSkeleton_Output struct {
	Actors []*UnspeccedDefs_SkeletonSearchActor `json:"actors" cborgen:"actors"`
	Cursor *string                              `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	// recId: Snowflake for this recommendation, use when submitting recommendation events.
	RecId *int64 `json:"recId,omitempty" cborgen:"recId,omitempty"`
	// relativeToDid: DID of the account these suggestions are relative to. If this is returned undefined, suggestions are based on the viewer.
	RelativeToDid *string `json:"relativeToDid,omitempty" cborgen:"relativeToDid,omitempty"`
}

// UnspeccedGetSuggestionsSkeleton calls the XRPC method "gndr.app.unspecced.getSuggestionsSkeleton".
//
// relativeToDid: DID of the account to get suggestions relative to. If not provided, suggestions will be based on the viewer.
// viewer: DID of the account making the request (not included for public/unauthenticated queries). Used to boost followed accounts in ranking.
func UnspeccedGetSuggestionsSkeleton(ctx context.Context, c util.LexClient, cursor string, limit int64, relativeToDid string, viewer string) (*UnspeccedGetSuggestionsSkeleton_Output, error) {
	var out UnspeccedGetSuggestionsSkeleton_Output

	params := map[string]interface{}{}
	if cursor != "" {
		params["cursor"] = cursor
	}
	if limit != 0 {
		params["limit"] = limit
	}
	if relativeToDid != "" {
		params["relativeToDid"] = relativeToDid
	}
	if viewer != "" {
		params["viewer"] = viewer
	}
	if err := c.LexDo(ctx, util.Query, "", "gndr.app.unspecced.getSuggestionsSkeleton", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
