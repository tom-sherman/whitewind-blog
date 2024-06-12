//go:build exclude
// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package app_autogen

// schema: app.bsky.actor.getSuggestions

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ActorGetSuggestions_Output is the output of a app.bsky.actor.getSuggestions call.
type ActorGetSuggestions_Output struct {
	Actors []*ActorDefs_ProfileView `json:"actors" cborgen:"actors"`
	Cursor *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
}

// ActorGetSuggestions calls the XRPC method "app.bsky.actor.getSuggestions".
func ActorGetSuggestions(ctx context.Context, c *xrpc.Client, cursor string, limit int64) (*ActorGetSuggestions_Output, error) {
	var out ActorGetSuggestions_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.actor.getSuggestions", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
