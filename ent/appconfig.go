// Code generated by ent, DO NOT EDIT.

package ent

import (
	"RjsConfigService/ent/appconfig"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AppConfig is the model entity for the AppConfig schema.
type AppConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// AppName holds the value of the "app_name" field.
	AppName string `json:"app_name,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// Status holds the value of the "status" field.
	Status appconfig.Status `json:"status,omitempty"`
	// Config holds the value of the "config" field.
	Config       map[string]interface{} `json:"config,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppConfig) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appconfig.FieldConfig:
			values[i] = new([]byte)
		case appconfig.FieldVersion:
			values[i] = new(sql.NullInt64)
		case appconfig.FieldAppName, appconfig.FieldStatus:
			values[i] = new(sql.NullString)
		case appconfig.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case appconfig.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppConfig fields.
func (ac *AppConfig) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appconfig.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case appconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case appconfig.FieldAppName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_name", values[i])
			} else if value.Valid {
				ac.AppName = value.String
			}
		case appconfig.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ac.Version = int(value.Int64)
			}
		case appconfig.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ac.Status = appconfig.Status(value.String)
			}
		case appconfig.FieldConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ac.Config); err != nil {
					return fmt.Errorf("unmarshal field config: %w", err)
				}
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppConfig.
// This includes values selected through modifiers, order, etc.
func (ac *AppConfig) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// Update returns a builder for updating this AppConfig.
// Note that you need to call AppConfig.Unwrap() before calling this method if this AppConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AppConfig) Update() *AppConfigUpdateOne {
	return NewAppConfigClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AppConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AppConfig) Unwrap() *AppConfig {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppConfig is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AppConfig) String() string {
	var builder strings.Builder
	builder.WriteString("AppConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_name=")
	builder.WriteString(ac.AppName)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", ac.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ac.Status))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(fmt.Sprintf("%v", ac.Config))
	builder.WriteByte(')')
	return builder.String()
}

// AppConfigs is a parsable slice of AppConfig.
type AppConfigs []*AppConfig
