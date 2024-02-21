package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Annotations(entsql.WithComments(true)).
			Comment("主键"),
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("项目名称"),
		field.String("code").
			Annotations(entsql.WithComments(true)).
			Comment("项目编码"),
		field.Uint64("create_by").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("创建人"),
		field.Time("create_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("创建时间"),
		field.Uint64("update_by").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("更新人"),
		field.Time("update_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("更新时间"),
		field.Uint64("tenant_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("租户号"),
		field.Bool("deleted").
			Annotations(entsql.WithComments(true)).
			Comment("是否删除")}
}
func (Project) Edges() []ent.Edge {
	return nil
}
func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "cc_project"}}
}
