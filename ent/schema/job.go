package schema

import "entgo.io/ent"

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return nil
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
