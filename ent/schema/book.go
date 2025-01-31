package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(100).
			NotEmpty(),
		field.Text("body").
			NotEmpty(),
		field.Int("price").
			Positive(),
		field.Text("thoughts").
			NotEmpty(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
