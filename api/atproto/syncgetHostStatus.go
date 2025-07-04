// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.getHostStatus

import (
	"context"

	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// SyncGetHostStatus_Output is the output of a com.atproto.sync.getHostStatus call.
type SyncGetHostStatus_Output struct {
	// accountCount: Number of accounts on the server which are associated with the upstream host. Note that the upstream may actually have more accounts.
	AccountCount *int64 `json:"accountCount,omitempty" cborgen:"accountCount,omitempty"`
	Hostname     string `json:"hostname" cborgen:"hostname"`
	// seq: Recent repo stream event sequence number. May be delayed from actual stream processing (eg, persisted cursor not in-memory cursor).
	Seq    *int64  `json:"seq,omitempty" cborgen:"seq,omitempty"`
	Status *string `json:"status,omitempty" cborgen:"status,omitempty"`
}

// SyncGetHostStatus calls the XRPC method "com.atproto.sync.getHostStatus".
//
// hostname: Hostname of the host (eg, PDS or relay) being queried.
func SyncGetHostStatus(ctx context.Context, c util.LexClient, hostname string) (*SyncGetHostStatus_Output, error) {
	var out SyncGetHostStatus_Output

	params := map[string]interface{}{}
	params["hostname"] = hostname
	if err := c.LexDo(ctx, util.Query, "", "com.atproto.sync.getHostStatus", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
