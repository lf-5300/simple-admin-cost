package localmixin

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

func init() {
	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(0)
	options.WorkerIdBitLength = 10 // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
	// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。

	// 保存参数（必须的操作，否则以上设置都不能生效）：
	idgen.SetIdGenerator(options)
	// 以上初始化过程只需全局一次，且必须在第2步之前设置。
}

// BaseField 定义了审计字段的 Mixin。
type BaseField struct {
	mixin.Schema
}

// Fields 返回 Mixin 的字段。
func (BaseField) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Default(uint64(idgen.NextId())).
			Immutable(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Comment("Create Time | 创建日期").
			Annotations(entsql.WithComments(true)),
		field.Time("updated_at").
			Optional().
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Comment("Update Time | 修改日期").
			Annotations(entsql.WithComments(true)),
		field.String("create_by").
			Immutable().
			Comment("创建人").
			Annotations(entsql.WithComments(true)),
		field.String("update_by").
			Optional().
			Comment("修改人").
			Annotations(entsql.WithComments(true)),
	}
}

// Hooks 是 Mixin 的钩子。
func (BaseField) Hooks() []ent.Hook {
	return []ent.Hook{
		// 你可以在这里添加钩子，例如设置 created_by 字段
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if m.Op().Is(ent.OpCreate) {
					userId, ok := ctx.Value("userId").(string)
					if !ok || userId == "" {
						// 在实际应用中，这里可能会返回错误或者默认用户
						userId = "0"
					}
					if err := m.SetField("create_by", userId); err != nil {
						return nil, err
					}
				}
				if m.Op().Is(ent.OpUpdateOne) {
					userId, ok := ctx.Value("userId").(string)
					if !ok || userId == "" {
						// 在实际应用中，这里可能会返回错误或者默认用户
						userId = "0"
					}
					if err := m.SetField("update_by", userId); err != nil {
						return nil, err
					}
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}
