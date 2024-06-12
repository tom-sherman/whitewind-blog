// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package whtwnd_autogen

// schema: com.whtwnd.blog.entry

import (
	"github.com/bluesky-social/indigo/lex/util"
)

func init() {
	util.RegisterType("com.whtwnd.blog.entry", &BlogEntry{})
} //
// RECORDTYPE: BlogEntry
type BlogEntry struct {
	LexiconTypeID string                   `json:"$type,const=com.whtwnd.blog.entry" cborgen:"$type,const=com.whtwnd.blog.entry"`
	Content       string                   `json:"content" cborgen:"content"`
	CreatedAt     *string                  `json:"createdAt,omitempty" cborgen:"createdAt,omitempty"`
	Ogp           *BlogDefs_Ogp            `json:"ogp,omitempty" cborgen:"ogp,omitempty"`
	Theme         *string                  `json:"theme,omitempty" cborgen:"theme,omitempty"`
	Title         *string                  `json:"title,omitempty" cborgen:"title,omitempty"`
	Tracker       *string                  `json:"tracker,omitempty" cborgen:"tracker,omitempty"`
}
