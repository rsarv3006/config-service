package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AppConfig holds the schema definition for the AppConfig entity.
type AppConfig struct {
	ent.Schema
}

// Fields of the AppConfig.
func (AppConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("created_at").Default(time.Now),
		field.String("app_name"),
		field.Int("version"),
		field.Enum("status").Values("active", "inactive").Default("inactive"),
		field.JSON("config", map[string]interface{}{}),
	}
}

// Edges of the AppConfig.
func (AppConfig) Edges() []ent.Edge {
	return nil
}
