package project

import (
	"context"
	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
	"github.com/hf/simple-admin-cost-api/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProjectLogic) UpdateProject(req *types.ProjectInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.Project.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilCode(req.Code).
		SetNotNilCreateBy(req.CreateBy).
		SetNotNilCreateTime(pointy.GetTimePointer(req.CreateTime, 0)).
		SetNotNilUpdateBy(req.UpdateBy).
		SetNotNilUpdateTime(pointy.GetTimePointer(req.UpdateTime, 0)).
		SetNotNilTenantID(req.TenantId).
		SetNotNilDeleted(req.Deleted).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
