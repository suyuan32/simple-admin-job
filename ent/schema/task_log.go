package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
			Comment("Task Started Time | 任务启动时间"),
		field.Time("finished_at").Comment("Task Finished Time | 任务完成时间"),
		field.Uint8("result").Comment("The Task Process Result | 任务执行结果"),
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
		entsql.WithComments(true),
		schema.Comment("Task Log Table | 任务日志表"),
		entsql.Annotation{Table: "sys_task_logs"},
	}
}
