package controller

import (
	"github.com/kataras/iris"
	"lbps/dto"
	"lbps/service"
)

func DeviceInfoAction(ctx iris.Context) {
	var err error
	var params dto.DeviceInfo
	//var users []*models.DeviceInfo
	var counts int
	// 绑定参数
	if err = params.Bind(ctx); err != nil {
		ctx.JSON(err)
		return
	}

	// 根据ID获取用户
	if counts, err = service.QueryDeviceInfo(params.Status); err != nil {
		ctx.JSON(err)
		return
	}

	//if len(users) == 0 {
	//	ctx.JSON(errors.NoDataError())
	//	return
	//}

	ctx.JSON(iris.Map{"code": 1000, "data": counts })
}