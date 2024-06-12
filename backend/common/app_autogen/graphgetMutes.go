//go:build exclude
// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package app_autogen

// schema: app.bsky.graph.getMutes

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// GraphGetMutes_Output is the output of a app.bsky.graph.getMutes call.
type GraphGetMutes_Output struct {
	Cursor *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Mutes  []*ActorDefs_ProfileView `json:"mutes" cborgen:"mutes"`
}

// GraphGetMutes calls the XRPC method "app.bsky.graph.getMutes".
func GraphGetMutes(ctx context.Context, c *xrpc.Client, cursor string, limit int64) (*GraphGetMutes_Output, error) {
	var out GraphGetMutes_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.graph.getMutes", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
