// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/MehmetTalhaSeker/mts-gr/ent/meta"
	"github.com/MehmetTalhaSeker/mts-gr/ent/post"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// MetaEdge is the edge representation of Meta.
type MetaEdge struct {
	Node   *Meta  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// MetaConnection is the connection containing edges to Meta.
type MetaConnection struct {
	Edges      []*MetaEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *MetaConnection) build(nodes []*Meta, pager *metaPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Meta
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Meta {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Meta {
			return nodes[i]
		}
	}
	c.Edges = make([]*MetaEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &MetaEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// MetaPaginateOption enables pagination customization.
type MetaPaginateOption func(*metaPager) error

// WithMetaOrder configures pagination ordering.
func WithMetaOrder(order *MetaOrder) MetaPaginateOption {
	if order == nil {
		order = DefaultMetaOrder
	}
	o := *order
	return func(pager *metaPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultMetaOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithMetaFilter configures pagination filter.
func WithMetaFilter(filter func(*MetaQuery) (*MetaQuery, error)) MetaPaginateOption {
	return func(pager *metaPager) error {
		if filter == nil {
			return errors.New("MetaQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type metaPager struct {
	reverse bool
	order   *MetaOrder
	filter  func(*MetaQuery) (*MetaQuery, error)
}

func newMetaPager(opts []MetaPaginateOption, reverse bool) (*metaPager, error) {
	pager := &metaPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMetaOrder
	}
	return pager, nil
}

func (p *metaPager) applyFilter(query *MetaQuery) (*MetaQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *metaPager) toCursor(m *Meta) Cursor {
	return p.order.Field.toCursor(m)
}

func (p *metaPager) applyCursors(query *MetaQuery, after, before *Cursor) (*MetaQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultMetaOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *metaPager) applyOrder(query *MetaQuery) *MetaQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultMetaOrder.Field {
		query = query.Order(DefaultMetaOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *metaPager) orderExpr(query *MetaQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultMetaOrder.Field {
			b.Comma().Ident(DefaultMetaOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Meta.
func (m *MetaQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...MetaPaginateOption,
) (*MetaConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newMetaPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if m, err = pager.applyFilter(m); err != nil {
		return nil, err
	}
	conn := &MetaConnection{Edges: []*MetaEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = m.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if m, err = pager.applyCursors(m, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		m.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := m.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	m = pager.applyOrder(m)
	nodes, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// MetaOrderField defines the ordering field of Meta.
type MetaOrderField struct {
	// Value extracts the ordering value from the given Meta.
	Value    func(*Meta) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) meta.OrderOption
	toCursor func(*Meta) Cursor
}

// MetaOrder defines the ordering of Meta.
type MetaOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *MetaOrderField `json:"field"`
}

// DefaultMetaOrder is the default ordering of Meta.
var DefaultMetaOrder = &MetaOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &MetaOrderField{
		Value: func(m *Meta) (ent.Value, error) {
			return m.ID, nil
		},
		column: meta.FieldID,
		toTerm: meta.ByID,
		toCursor: func(m *Meta) Cursor {
			return Cursor{ID: m.ID}
		},
	},
}

// ToEdge converts Meta into MetaEdge.
func (m *Meta) ToEdge(order *MetaOrder) *MetaEdge {
	if order == nil {
		order = DefaultMetaOrder
	}
	return &MetaEdge{
		Node:   m,
		Cursor: order.Field.toCursor(m),
	}
}

// PostEdge is the edge representation of Post.
type PostEdge struct {
	Node   *Post  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// PostConnection is the connection containing edges to Post.
type PostConnection struct {
	Edges      []*PostEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *PostConnection) build(nodes []*Post, pager *postPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Post
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Post {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Post {
			return nodes[i]
		}
	}
	c.Edges = make([]*PostEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PostEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PostPaginateOption enables pagination customization.
type PostPaginateOption func(*postPager) error

// WithPostOrder configures pagination ordering.
func WithPostOrder(order *PostOrder) PostPaginateOption {
	if order == nil {
		order = DefaultPostOrder
	}
	o := *order
	return func(pager *postPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPostOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPostFilter configures pagination filter.
func WithPostFilter(filter func(*PostQuery) (*PostQuery, error)) PostPaginateOption {
	return func(pager *postPager) error {
		if filter == nil {
			return errors.New("PostQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type postPager struct {
	reverse bool
	order   *PostOrder
	filter  func(*PostQuery) (*PostQuery, error)
}

func newPostPager(opts []PostPaginateOption, reverse bool) (*postPager, error) {
	pager := &postPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPostOrder
	}
	return pager, nil
}

func (p *postPager) applyFilter(query *PostQuery) (*PostQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *postPager) toCursor(po *Post) Cursor {
	return p.order.Field.toCursor(po)
}

func (p *postPager) applyCursors(query *PostQuery, after, before *Cursor) (*PostQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultPostOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *postPager) applyOrder(query *PostQuery) *PostQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultPostOrder.Field {
		query = query.Order(DefaultPostOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *postPager) orderExpr(query *PostQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPostOrder.Field {
			b.Comma().Ident(DefaultPostOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Post.
func (po *PostQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PostPaginateOption,
) (*PostConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPostPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if po, err = pager.applyFilter(po); err != nil {
		return nil, err
	}
	conn := &PostConnection{Edges: []*PostEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = po.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if po, err = pager.applyCursors(po, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		po.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := po.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	po = pager.applyOrder(po)
	nodes, err := po.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// PostOrderField defines the ordering field of Post.
type PostOrderField struct {
	// Value extracts the ordering value from the given Post.
	Value    func(*Post) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) post.OrderOption
	toCursor func(*Post) Cursor
}

// PostOrder defines the ordering of Post.
type PostOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *PostOrderField `json:"field"`
}

// DefaultPostOrder is the default ordering of Post.
var DefaultPostOrder = &PostOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &PostOrderField{
		Value: func(po *Post) (ent.Value, error) {
			return po.ID, nil
		},
		column: post.FieldID,
		toTerm: post.ByID,
		toCursor: func(po *Post) Cursor {
			return Cursor{ID: po.ID}
		},
	},
}

// ToEdge converts Post into PostEdge.
func (po *Post) ToEdge(order *PostOrder) *PostEdge {
	if order == nil {
		order = DefaultPostOrder
	}
	return &PostEdge{
		Node:   po,
		Cursor: order.Field.toCursor(po),
	}
}