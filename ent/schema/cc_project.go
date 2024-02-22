package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/hf/simple-admin-cost-api/ent/schema/localmixin"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("项目名称"),
		field.String("code").
			Annotations(entsql.WithComments(true)).
			Comment("项目编码"),
	}
}
func (Project) Edges() []ent.Edge {
	return nil
}
func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		localmixin.BaseField{},
		localmixin.SoftDeleteMixin{},
	}
}
func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "cc_project"}}
}
