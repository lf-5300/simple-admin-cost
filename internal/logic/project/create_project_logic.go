package project

import (
	"context"

	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
	"github.com/hf/simple-admin-cost-api/internal/utils/dberrorhandler"

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
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
