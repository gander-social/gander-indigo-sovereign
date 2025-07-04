// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.refreshSession

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// ServerRefreshSession_Output is the output of a com.atproto.server.refreshSession call.
type ServerRefreshSession_Output struct {
	AccessJwt  string       `json:"accessJwt" cborgen:"accessJwt"`
	Active     *bool        `json:"active,omitempty" cborgen:"active,omitempty"`
	Did        string       `json:"did" cborgen:"did"`
	DidDoc     *interface{} `json:"didDoc,omitempty" cborgen:"didDoc,omitempty"`
	Handle     string       `json:"handle" cborgen:"handle"`
	RefreshJwt string       `json:"refreshJwt" cborgen:"refreshJwt"`
	// status: Hosting status of the account. If not specified, then assume 'active'.
	Status *string `json:"status,omitempty" cborgen:"status,omitempty"`
}

// ServerRefreshSession calls the XRPC method "com.atproto.server.refreshSession".
func ServerRefreshSession(ctx context.Context, c util.LexClient) (*ServerRefreshSession_Output, error) {
	var out ServerRefreshSession_Output
	if err := c.LexDo(ctx, util.Procedure, "", "com.atproto.server.refreshSession", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
