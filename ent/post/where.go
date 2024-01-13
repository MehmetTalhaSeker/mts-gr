// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/MehmetTalhaSeker/mts-gr/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// ShortDescr applies equality check predicate on the "short_descr" field. It's identical to ShortDescrEQ.
func ShortDescr(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldShortDescr, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTitle, v))
}

// ShortDescrEQ applies the EQ predicate on the "short_descr" field.
func ShortDescrEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldShortDescr, v))
}

// ShortDescrNEQ applies the NEQ predicate on the "short_descr" field.
func ShortDescrNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldShortDescr, v))
}

// ShortDescrIn applies the In predicate on the "short_descr" field.
func ShortDescrIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldShortDescr, vs...))
}

// ShortDescrNotIn applies the NotIn predicate on the "short_descr" field.
func ShortDescrNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldShortDescr, vs...))
}

// ShortDescrGT applies the GT predicate on the "short_descr" field.
func ShortDescrGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldShortDescr, v))
}

// ShortDescrGTE applies the GTE predicate on the "short_descr" field.
func ShortDescrGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldShortDescr, v))
}

// ShortDescrLT applies the LT predicate on the "short_descr" field.
func ShortDescrLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldShortDescr, v))
}

// ShortDescrLTE applies the LTE predicate on the "short_descr" field.
func ShortDescrLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldShortDescr, v))
}

// ShortDescrContains applies the Contains predicate on the "short_descr" field.
func ShortDescrContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldShortDescr, v))
}

// ShortDescrHasPrefix applies the HasPrefix predicate on the "short_descr" field.
func ShortDescrHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldShortDescr, v))
}

// ShortDescrHasSuffix applies the HasSuffix predicate on the "short_descr" field.
func ShortDescrHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldShortDescr, v))
}

// ShortDescrEqualFold applies the EqualFold predicate on the "short_descr" field.
func ShortDescrEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldShortDescr, v))
}

// ShortDescrContainsFold applies the ContainsFold predicate on the "short_descr" field.
func ShortDescrContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldShortDescr, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldBody, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}
