package project

import (
	"context"

	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
	"github.com/hf/simple-admin-cost-api/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectByIdLogic {
	return &GetProjectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectByIdLogic) GetProjectById(req *types.IDReq) (*types.ProjectInfoResp, error) {
	data, err := l.svcCtx.DB.Project.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.ProjectInfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
        },
        Data: types.ProjectInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &data.ID,
				CreatedAt:    pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(data.UpdatedAt.UnixMilli()),
            },
			CreateBy:	&data.CreateBy,
			UpdateBy:	&data.UpdateBy,
			Name:	&data.Name,
			Code:	&data.Code,
        },
	}, nil
}

