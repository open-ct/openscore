package filter

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/open-ct/openscore/pkg/token"
	"log"
)

func Auth(ctx *context.Context) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))

	authorization := ctx.Input.Header("Authorization")

	if len(authorization) == 0 {
		ctx.Output.JSON("cant get Authorization", true, true)
		// ctx.Abort(400)
		return
	}

	res, err := token.ResolveToken(authorization)
	if err != nil {
		log.Println(err)
		return
		// api.ResponseError("cant resolve Authorization", err)
	}

	ctx.Output.Session("userId", res.Id)
	ctx.Output.Session("typeId", res.TypeId)
}
