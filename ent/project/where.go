// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/hf/simple-admin-cost-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreateBy applies equality check predicate on the "create_by" field. It's identical to CreateByEQ.
func CreateBy(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreateBy, v))
}

// UpdateBy applies equality check predicate on the "update_by" field. It's identical to UpdateByEQ.
func UpdateBy(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdateBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldUpdatedAt))
}

// CreateByEQ applies the EQ predicate on the "create_by" field.
func CreateByEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreateBy, v))
}

// CreateByNEQ applies the NEQ predicate on the "create_by" field.
func CreateByNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreateBy, v))
}

// CreateByIn applies the In predicate on the "create_by" field.
func CreateByIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreateBy, vs...))
}

// CreateByNotIn applies the NotIn predicate on the "create_by" field.
func CreateByNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreateBy, vs...))
}

// CreateByGT applies the GT predicate on the "create_by" field.
func CreateByGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreateBy, v))
}

// CreateByGTE applies the GTE predicate on the "create_by" field.
func CreateByGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreateBy, v))
}

// CreateByLT applies the LT predicate on the "create_by" field.
func CreateByLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreateBy, v))
}

// CreateByLTE applies the LTE predicate on the "create_by" field.
func CreateByLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreateBy, v))
}

// CreateByContains applies the Contains predicate on the "create_by" field.
func CreateByContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldCreateBy, v))
}

// CreateByHasPrefix applies the HasPrefix predicate on the "create_by" field.
func CreateByHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldCreateBy, v))
}

// CreateByHasSuffix applies the HasSuffix predicate on the "create_by" field.
func CreateByHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldCreateBy, v))
}

// CreateByEqualFold applies the EqualFold predicate on the "create_by" field.
func CreateByEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldCreateBy, v))
}

// CreateByContainsFold applies the ContainsFold predicate on the "create_by" field.
func CreateByContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldCreateBy, v))
}

// UpdateByEQ applies the EQ predicate on the "update_by" field.
func UpdateByEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdateBy, v))
}

// UpdateByNEQ applies the NEQ predicate on the "update_by" field.
func UpdateByNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdateBy, v))
}

// UpdateByIn applies the In predicate on the "update_by" field.
func UpdateByIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdateBy, vs...))
}

// UpdateByNotIn applies the NotIn predicate on the "update_by" field.
func UpdateByNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdateBy, vs...))
}

// UpdateByGT applies the GT predicate on the "update_by" field.
func UpdateByGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdateBy, v))
}

// UpdateByGTE applies the GTE predicate on the "update_by" field.
func UpdateByGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdateBy, v))
}

// UpdateByLT applies the LT predicate on the "update_by" field.
func UpdateByLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdateBy, v))
}

// UpdateByLTE applies the LTE predicate on the "update_by" field.
func UpdateByLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdateBy, v))
}

// UpdateByContains applies the Contains predicate on the "update_by" field.
func UpdateByContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldUpdateBy, v))
}

// UpdateByHasPrefix applies the HasPrefix predicate on the "update_by" field.
func UpdateByHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldUpdateBy, v))
}

// UpdateByHasSuffix applies the HasSuffix predicate on the "update_by" field.
func UpdateByHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldUpdateBy, v))
}

// UpdateByIsNil applies the IsNil predicate on the "update_by" field.
func UpdateByIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldUpdateBy))
}

// UpdateByNotNil applies the NotNil predicate on the "update_by" field.
func UpdateByNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldUpdateBy))
}

// UpdateByEqualFold applies the EqualFold predicate on the "update_by" field.
func UpdateByEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldUpdateBy, v))
}

// UpdateByContainsFold applies the ContainsFold predicate on the "update_by" field.
func UpdateByContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldUpdateBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldCode, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(sql.NotPredicates(p))
}
