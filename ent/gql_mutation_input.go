// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateMetaInput represents a mutation input for creating metaslice.
type CreateMetaInput struct {
	Key   string
	Value string
}

// Mutate applies the CreateMetaInput on the MetaMutation builder.
func (i *CreateMetaInput) Mutate(m *MetaMutation) {
	m.SetKey(i.Key)
	m.SetValue(i.Value)
}

// SetInput applies the change-set in the CreateMetaInput on the MetaCreate builder.
func (c *MetaCreate) SetInput(i CreateMetaInput) *MetaCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreatedAt  *time.Time
	Title      string
	ShortDescr string
	Body       string
}

// Mutate applies the CreatePostInput on the PostMutation builder.
func (i *CreatePostInput) Mutate(m *PostMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	m.SetTitle(i.Title)
	m.SetShortDescr(i.ShortDescr)
	m.SetBody(i.Body)
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	Title      *string
	ShortDescr *string
	Body       *string
}

// Mutate applies the UpdatePostInput on the PostMutation builder.
func (i *UpdatePostInput) Mutate(m *PostMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.ShortDescr; v != nil {
		m.SetShortDescr(*v)
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdate builder.
func (c *PostUpdate) SetInput(i UpdatePostInput) *PostUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdateOne builder.
func (c *PostUpdateOne) SetInput(i UpdatePostInput) *PostUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
