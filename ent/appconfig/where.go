// Code generated by ent, DO NOT EDIT.

package appconfig

import (
	"RjsConfigService/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// AppName applies equality check predicate on the "app_name" field. It's identical to AppNameEQ.
func AppName(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldAppName, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldVersion, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// AppNameEQ applies the EQ predicate on the "app_name" field.
func AppNameEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldAppName, v))
}

// AppNameNEQ applies the NEQ predicate on the "app_name" field.
func AppNameNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldAppName, v))
}

// AppNameIn applies the In predicate on the "app_name" field.
func AppNameIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldAppName, vs...))
}

// AppNameNotIn applies the NotIn predicate on the "app_name" field.
func AppNameNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldAppName, vs...))
}

// AppNameGT applies the GT predicate on the "app_name" field.
func AppNameGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldAppName, v))
}

// AppNameGTE applies the GTE predicate on the "app_name" field.
func AppNameGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldAppName, v))
}

// AppNameLT applies the LT predicate on the "app_name" field.
func AppNameLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldAppName, v))
}

// AppNameLTE applies the LTE predicate on the "app_name" field.
func AppNameLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldAppName, v))
}

// AppNameContains applies the Contains predicate on the "app_name" field.
func AppNameContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldAppName, v))
}

// AppNameHasPrefix applies the HasPrefix predicate on the "app_name" field.
func AppNameHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldAppName, v))
}

// AppNameHasSuffix applies the HasSuffix predicate on the "app_name" field.
func AppNameHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldAppName, v))
}

// AppNameEqualFold applies the EqualFold predicate on the "app_name" field.
func AppNameEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldAppName, v))
}

// AppNameContainsFold applies the ContainsFold predicate on the "app_name" field.
func AppNameContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldAppName, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(sql.NotPredicates(p))
}
