package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("created_at").Default(time.Now),
		field.Enum("role").Values("admin", "access").Default("access"),
		field.String("app_name"),
		field.Enum("status").Values("active", "inactive").Default("active"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
