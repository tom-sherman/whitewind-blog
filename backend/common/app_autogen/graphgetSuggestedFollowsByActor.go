//go:build exclude
// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package app_autogen

// schema: app.bsky.graph.getSuggestedFollowsByActor

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// GraphGetSuggestedFollowsByActor_Output is the output of a app.bsky.graph.getSuggestedFollowsByActor call.
type GraphGetSuggestedFollowsByActor_Output struct {
	Suggestions []*ActorDefs_ProfileView `json:"suggestions" cborgen:"suggestions"`
}

// GraphGetSuggestedFollowsByActor calls the XRPC method "app.bsky.graph.getSuggestedFollowsByActor".
func GraphGetSuggestedFollowsByActor(ctx context.Context, c *xrpc.Client, actor string) (*GraphGetSuggestedFollowsByActor_Output, error) {
	var out GraphGetSuggestedFollowsByActor_Output

	params := map[string]interface{}{
		"actor": actor,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.graph.getSuggestedFollowsByActor", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
