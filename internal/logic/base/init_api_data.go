package base

import (
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	// Project

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/project/create"),
		Description: pointy.GetPointer("apiDesc.createProject"),
		ApiGroup:    pointy.GetPointer("project"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/project/update"),
		Description: pointy.GetPointer("apiDesc.updateProject"),
		ApiGroup:    pointy.GetPointer("project"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/project/delete"),
		Description: pointy.GetPointer("apiDesc.deleteProject"),
		ApiGroup:    pointy.GetPointer("project"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/project/list"),
		Description: pointy.GetPointer("apiDesc.getProjectList"),
		ApiGroup:    pointy.GetPointer("project"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/project"),
		Description: pointy.GetPointer("apiDesc.getProjectById"),
		ApiGroup:    pointy.GetPointer("project"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	return err
}
