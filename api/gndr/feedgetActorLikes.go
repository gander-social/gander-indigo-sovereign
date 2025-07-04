// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.feed.getActorLikes

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// FeedGetActorLikes_Output is the output of a gndr.app.feed.getActorLikes call.
type FeedGetActorLikes_Output struct {
	Cursor *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Feed   []*FeedDefs_FeedViewPost `json:"feed" cborgen:"feed"`
}

// FeedGetActorLikes calls the XRPC method "gndr.app.feed.getActorLikes".
func FeedGetActorLikes(ctx context.Context, c util.LexClient, actor string, cursor string, limit int64) (*FeedGetActorLikes_Output, error) {
	var out FeedGetActorLikes_Output

	params := map[string]interface{}{}
	params["actor"] = actor
	if cursor != "" {
		params["cursor"] = cursor
	}
	if limit != 0 {
		params["limit"] = limit
	}
	if err := c.LexDo(ctx, util.Query, "", "gndr.app.feed.getActorLikes", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
