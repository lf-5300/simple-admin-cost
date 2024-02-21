package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/hf/simple-admin-cost-api/internal/logic/project"
	"github.com/hf/simple-admin-cost-api/internal/svc"
	"github.com/hf/simple-admin-cost-api/internal/types"
)

// swagger:route post /project/delete project DeleteProject
//
// Delete project information | 删除Project信息
//
// Delete project information | 删除Project信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewDeleteProjectLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProject(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
