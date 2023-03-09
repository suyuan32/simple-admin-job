package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
		field.Int("entry_id").Comment("The entry ID of the task | 任务启动返回的ID"),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_tasks"},
	}
}
