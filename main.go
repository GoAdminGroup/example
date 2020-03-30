package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // adapter
	"github.com/GoAdminGroup/go-admin/context"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                    // theme

	"github.com/GoAdminGroup/example/tables"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()

	// customize a plugin
	// 自己定制一个插件👇

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	// you can also add config like:
	// 您也可以像下面这样的方式去引入数据库👇
	//
	// import "github.com/GoAdminGroup/go-admin/modules/config"
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

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", tables.GetUserTable).
		AddGenerator("external", tables.GetExternalTable).
		AddPlugins(examplePlugin).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// customize your index pages
	// 下面这样定制您的首页👇

	eng.HTML("GET", "/admin", func(ctx *context.Context) (panel types.Panel, e error) {
		return DashboardPage()
	})

	_ = r.Run(":9033")
}
