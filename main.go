/*
 * @Author: Junlang
 * @Date: 2021-07-18 00:42:25
 * @LastEditTime: 2021-07-23 17:04:50
 * @LastEditors: Junlang
 * @FilePath: /openscore/main.go
 */
package main

import (
	_ "openscore/routers"

	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/plugins/cors"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{



		AllowOrigins: []string{"*"},
		//AllowMethods:     []string{"GET", "PUT", "PATCH", "POST"},
		AllowMethods: []string{"GET", "PUT", "PATCH", "POST", "OPTIONS"},
		// AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Headers", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	// beego.SetStaticPath("/static", "web/build/static")
	// beego.InsertFilter("/", beego.BeforeRouter, routers.TransparentStatic) // must has this for default page
	// beego.InsertFilter("/*", beego.BeforeRouter, routers.TransparentStatic)

	beego.BConfig.WebConfig.Session.SessionName = "openscore_session_id"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 365



//	http.Handle("/d/", http.StripPrefix("/d/", http.FileServer(http.Dir("d")))) // 正确

	//mux := http.NewServeMux()
	//
	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/usr/workspace/src/open-ct/"))))
	//
	//mux.Handle("/c/", http.StripPrefix("/c/", http.FileServer(http.Dir("C:\\Users\\chen\\go\\src\\openscore\\img\\"))))
	//
	//mux.Handle("/d/", http.FileServer((http.Dir("d"))))
	//
	//mux.Handle("/e/", http.StripPrefix("/e/", http.FileServer(http.Dir("e:"))))
	//
	//http.HandleFunc("/", myWeb)
	//fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
	//_ = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Println("服务器开启错误: ", err)

		beego.Run()
	}

//func myWeb(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
//	//w.Header().Set("content-type", "application/json") //返回数据格式是json
//
//	query := r.URL.Query();
//	var url string;
//	var name string;
//	if len(query["url"]) > 0{
//		url = query["url"][0]
//	}else {
//		fmt.Fprintf(w, "");
//		return ;
//	}
//
//	if len(query["name"]) > 0{
//		name = query["name"][0];
//		w.Header().Set("content-type", "application/octet-stream")
//		w.Header().Set("Content-Disposition", "attachment; filename=" + name)
//	}
//
//	fmt.Println(url,name);
//	pix := ReadImgData(url);
//	w.Write(pix )
//	//fmt.Fprintf(w, "这是一个开始");
//}
//
////获取C的图片数据
//func ReadImgData(url string) []byte {
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	pix, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//	return pix
//}
