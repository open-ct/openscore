package filter

import (
	"strconv"

	"github.com/astaxie/beego/context"
)

func AuthScore(ctx *context.Context) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))

	typeId := ctx.Input.Session("userType").(int64)

	if typeId != 2 && typeId != 1 {
		ctx.Output.JSON("typeId can't match :"+strconv.Itoa(int(typeId)), true, true)
		return
	}

}

func AuthSupervisor(ctx *context.Context) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))

	typeId := ctx.Input.Session("userType").(int64)

	if typeId != 1 {
		ctx.Output.JSON("typeId can't match:"+strconv.Itoa(int(typeId)), true, true)
		return
	}
}
