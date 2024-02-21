package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/hf/simple-admin-cost-api/internal/logic/project"
	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
)

// swagger:route post /project/list project GetProjectList
//
// Get project list | 获取Project列表
//
// Get project list | 获取Project列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ProjectListReq
//
// Responses:
//  200: ProjectListResp

func GetProjectListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewGetProjectListLogic(r.Context(), svcCtx)
		resp, err := l.GetProjectList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
