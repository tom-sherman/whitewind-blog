//go:build exclude
// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package app_autogen

// schema: app.bsky.feed.repost

import (
	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
)

func init() {
	util.RegisterType("app.bsky.feed.repost", &FeedRepost{})
} //
// RECORDTYPE: FeedRepost
type FeedRepost struct {
	LexiconTypeID string                         `json:"$type,const=app.bsky.feed.repost" cborgen:"$type,const=app.bsky.feed.repost"`
	CreatedAt     string                         `json:"createdAt" cborgen:"createdAt"`
	Subject       *comatprototypes.RepoStrongRef `json:"subject" cborgen:"subject"`
}
