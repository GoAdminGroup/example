package main

import (
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// customize a plugin

	examplePlugin := example.NewExample()

	// you can also add config like:
	//
	// import "github.com/chenhg5/go-admin/modules/config""
	//
	// cfg := config.Config{
	//	 Databases: config.DatabaseList{
	//		"default": {
	//			Host:       "127.0.0.1",
	//			Port:       "3306",
	//			User:       "root",
	//			Pwd:        "root",
	//			Name:       "godmin",
	//			MaxIdleCon: 50,
	//			MaxOpenCon: 150,
	//			Driver:     db.DriverMysql,
	//		},
	//	},
	//	UrlPrefix: "admin",
	//	IndexUrl:  "/",
	//	Debug:     true,
	//	Language:  language.CN,
	// }
	//
	// eng.AddConfig(cfg)

	if err := eng.AddConfigFromJson("./config.json").
		AddPlugins(adminPlugin, examplePlugin).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// customize your pages
	// http://localhost:9033/admin/custom

	r.GET("/admin/custom", func(ctx *gin.Context) {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	_ = r.Run(":9033")
}
