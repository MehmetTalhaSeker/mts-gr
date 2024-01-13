// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/MehmetTalhaSeker/mts-gr/ent/meta"
	"github.com/MehmetTalhaSeker/mts-gr/ent/post"
	"github.com/MehmetTalhaSeker/mts-gr/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	metaFields := schema.Meta{}.Fields()
	_ = metaFields
	// metaDescKey is the schema descriptor for key field.
	metaDescKey := metaFields[0].Descriptor()
	// meta.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	meta.KeyValidator = metaDescKey.Validators[0].(func(string) error)
	// metaDescValue is the schema descriptor for value field.
	metaDescValue := metaFields[1].Descriptor()
	// meta.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	meta.ValueValidator = metaDescValue.Validators[0].(func(string) error)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[0].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(time.Time)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[1].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescShortDescr is the schema descriptor for short_descr field.
	postDescShortDescr := postFields[2].Descriptor()
	// post.ShortDescrValidator is a validator for the "short_descr" field. It is called by the builders before save.
	post.ShortDescrValidator = postDescShortDescr.Validators[0].(func(string) error)
	// postDescBody is the schema descriptor for body field.
	postDescBody := postFields[3].Descriptor()
	// post.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	post.BodyValidator = postDescBody.Validators[0].(func(string) error)
}
