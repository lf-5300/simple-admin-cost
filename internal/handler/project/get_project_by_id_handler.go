package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/hf/simple-admin-cost-api/internal/logic/project"
	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
)

// swagger:route post /project project GetProjectById
//
// Get project by ID | 通过ID获取Project
//
// Get project by ID | 通过ID获取Project
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: ProjectInfoResp

func GetProjectByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewGetProjectByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetProjectById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
