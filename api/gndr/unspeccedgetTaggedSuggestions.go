// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.unspecced.getTaggedSuggestions

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// UnspeccedGetTaggedSuggestions_Output is the output of a gndr.app.unspecced.getTaggedSuggestions call.
type UnspeccedGetTaggedSuggestions_Output struct {
	Suggestions []*UnspeccedGetTaggedSuggestions_Suggestion `json:"suggestions" cborgen:"suggestions"`
}

// UnspeccedGetTaggedSuggestions_Suggestion is a "suggestion" in the gndr.app.unspecced.getTaggedSuggestions schema.
type UnspeccedGetTaggedSuggestions_Suggestion struct {
	Subject     string `json:"subject" cborgen:"subject"`
	SubjectType string `json:"subjectType" cborgen:"subjectType"`
	Tag         string `json:"tag" cborgen:"tag"`
}

// UnspeccedGetTaggedSuggestions calls the XRPC method "gndr.app.unspecced.getTaggedSuggestions".
func UnspeccedGetTaggedSuggestions(ctx context.Context, c util.LexClient) (*UnspeccedGetTaggedSuggestions_Output, error) {
	var out UnspeccedGetTaggedSuggestions_Output

	params := map[string]interface{}{}
	if err := c.LexDo(ctx, util.Query, "", "gndr.app.unspecced.getTaggedSuggestions", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
