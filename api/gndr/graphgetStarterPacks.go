// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.graph.getStarterPacks

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// GraphGetStarterPacks_Output is the output of a gndr.app.graph.getStarterPacks call.
type GraphGetStarterPacks_Output struct {
	StarterPacks []*GraphDefs_StarterPackViewBasic `json:"starterPacks" cborgen:"starterPacks"`
}

// GraphGetStarterPacks calls the XRPC method "gndr.app.graph.getStarterPacks".
func GraphGetStarterPacks(ctx context.Context, c util.LexClient, uris []string) (*GraphGetStarterPacks_Output, error) {
	var out GraphGetStarterPacks_Output

	params := map[string]interface{}{}
	params["uris"] = uris
	if err := c.LexDo(ctx, util.Query, "", "gndr.app.graph.getStarterPacks", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
