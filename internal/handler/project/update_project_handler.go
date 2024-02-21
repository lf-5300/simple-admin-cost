package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/hf/simple-admin-cost-api/internal/logic/project"
	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
)

// swagger:route post /project/update project UpdateProject
//
// Update project information | 更新Project
//
// Update project information | 更新Project
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ProjectInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewUpdateProjectLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProject(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
