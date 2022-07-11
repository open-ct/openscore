/*
 * @Author: Junlang
 * @Date: 2021-07-18 00:42:25
 * @LastEditTime: 2021-07-23 17:04:50
 * @LastEditors: Junlang
 * @FilePath: /openscore/main.go
 */
package main

import (
	"openscore/routers"

	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/plugins/cors"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{

		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "PATCH", "POST", "OPTIONS"},
		// AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Headers", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.SetStaticPath("/static", "web/build/static")
	// beego.InsertFilter("/", beego.BeforeRouter, routers.TransparentStatic) // must has this for default page
	// beego.InsertFilter("/*", beego.BeforeRouter, routers.TransparentStatic)
	beego.InsertFilter("*", beego.BeforeRouter, routers.StaticFilter)
	beego.BConfig.WebConfig.Session.SessionName = "openscore_session_id"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 365

	beego.Run()
}
