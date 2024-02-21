package project

import (
	"context"

	"github.com/hf/simple-admin-cost-api/ent/predicate"
	"github.com/hf/simple-admin-cost-api/ent/project"
	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
	"github.com/hf/simple-admin-cost-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectListLogic) GetProjectList(req *types.ProjectListReq) (*types.ProjectListResp, error) {
	var predicates []predicate.Project
	if req.Name != nil {
		predicates = append(predicates, project.NameContains(*req.Name))
	}
	if req.Code != nil {
		predicates = append(predicates, project.CodeContains(*req.Code))
	}
	data, err := l.svcCtx.DB.Project.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.ProjectListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.ProjectInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreateTime.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdateTime.UnixMilli()),
				},
				Name:       &v.Name,
				Code:       &v.Code,
				CreateBy:   &v.CreateBy,
				CreateTime: pointy.GetUnixMilliPointer(v.CreateTime.UnixMilli()),
				UpdateBy:   &v.UpdateBy,
				UpdateTime: pointy.GetUnixMilliPointer(v.UpdateTime.UnixMilli()),
				TenantId:   &v.TenantID,
				Deleted:    &v.Deleted,
			})
	}

	return resp, nil
}
