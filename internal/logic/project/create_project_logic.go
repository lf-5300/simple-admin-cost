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

type CreateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProjectLogic) CreateProject(req *types.ProjectInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Project.Create().
		SetNotNilName(req.Name).
		SetNotNilCode(req.Code).
		SetNotNilCreateBy(req.CreateBy).
		SetNotNilCreateTime(pointy.GetTimePointer(req.CreateTime, 0)).
		SetNotNilUpdateBy(req.UpdateBy).
		SetNotNilUpdateTime(pointy.GetTimePointer(req.UpdateTime, 0)).
		SetNotNilTenantID(req.TenantId).
		SetNotNilDeleted(req.Deleted).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
