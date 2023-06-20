package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TaskLog holds the schema definition for the TaskLog entity.
type TaskLog struct {
	ent.Schema
}

// Fields of the TaskLog.
func (TaskLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("started_at").Immutable().
			Default(time.Now).
			Comment("Task Started Time | 任务启动时间").
			Annotations(entsql.WithComments(true)),
		field.Time("finished_at").Comment("Task Finished Time | 任务完成时间").
			Annotations(entsql.WithComments(true)),
		field.Uint8("result").Comment("The Task Process Result | 任务执行结果").
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the TaskLog.
func (TaskLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tasks", Task.Type).Ref("task_logs").Unique(),
	}
}

func (TaskLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_task_logs"},
	}
}
