//go:build exclude
// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package com_autogen

// schema: com.atproto.server.createSession

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ServerCreateSession_Input is the input argument to a com.atproto.server.createSession call.
type ServerCreateSession_Input struct {
	// identifier: Handle or other identifier supported by the server for the authenticating user.
	Identifier string `json:"identifier" cborgen:"identifier"`
	Password   string `json:"password" cborgen:"password"`
}

// ServerCreateSession_Output is the output of a com.atproto.server.createSession call.
type ServerCreateSession_Output struct {
	AccessJwt      string       `json:"accessJwt" cborgen:"accessJwt"`
	Did            string       `json:"did" cborgen:"did"`
	DidDoc         *interface{} `json:"didDoc,omitempty" cborgen:"didDoc,omitempty"`
	Email          *string      `json:"email,omitempty" cborgen:"email,omitempty"`
	EmailConfirmed *bool        `json:"emailConfirmed,omitempty" cborgen:"emailConfirmed,omitempty"`
	Handle         string       `json:"handle" cborgen:"handle"`
	RefreshJwt     string       `json:"refreshJwt" cborgen:"refreshJwt"`
}

// ServerCreateSession calls the XRPC method "com.atproto.server.createSession".
func ServerCreateSession(ctx context.Context, c *xrpc.Client, input *ServerCreateSession_Input) (*ServerCreateSession_Output, error) {
	var out ServerCreateSession_Output
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.server.createSession", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
