package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Task Name | 任务名称"),
		field.String("task_group").Comment("Task Group | 任务分组"),
		field.String("cron_expression").Comment("Cron expression | 定时任务表达式"),
		field.String("pattern").Comment("Cron Pattern | 任务的模式 （用于区分和确定要执行的任务）"),
		field.String("payload").Comment("The data used in cron (JSON string) | 任务需要的数据(JSON 字符串)"),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task_logs", TaskLog.Type),
	}
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pattern").Unique(),
	}
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Task Configuration Table | 任务配置表"),
		entsql.Annotation{Table: "sys_tasks"},
	}
}
