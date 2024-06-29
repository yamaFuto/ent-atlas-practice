package schema

import (
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asaskevich/govalidator"
	
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			NotEmpty(),
		field.Int("age").
			Positive(),
		field.String("email").
			NotEmpty().
			Unique().
			Validate(func(email string) error {
				if !govalidator.IsEmail(email) {
					return errors.New("invalid email format")
				}
				return nil
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
