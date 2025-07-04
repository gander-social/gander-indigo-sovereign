// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.unspecced.getSuggestedUsers

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// UnspeccedGetSuggestedUsers_Output is the output of a gndr.app.unspecced.getSuggestedUsers call.
type UnspeccedGetSuggestedUsers_Output struct {
	Actors []*ActorDefs_ProfileView `json:"actors" cborgen:"actors"`
}

// UnspeccedGetSuggestedUsers calls the XRPC method "gndr.app.unspecced.getSuggestedUsers".
//
// category: Category of users to get suggestions for.
func UnspeccedGetSuggestedUsers(ctx context.Context, c util.LexClient, category string, limit int64) (*UnspeccedGetSuggestedUsers_Output, error) {
	var out UnspeccedGetSuggestedUsers_Output

	params := map[string]interface{}{}
	if category != "" {
		params["category"] = category
	}
	if limit != 0 {
		params["limit"] = limit
	}
	if err := c.LexDo(ctx, util.Query, "", "gndr.app.unspecced.getSuggestedUsers", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
