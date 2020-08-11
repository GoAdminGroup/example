package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
)

func GetTableContent(ctx *context.Context) (types.Panel, error) {

	comp := template.Get(config.GetTheme())

	table := comp.DataTable().
		SetInfoList([]map[string]types.InfoItem{
			{
				"id":     {Content: "0"},
				"name":   {Content: "Jack"},
				"gender": {Content: "men"},
				"age":    {Content: "20"},
			},
			{
				"id":     {Content: "1"},
				"name":   {Content: "Jane"},
				"gender": {Content: "women"},
				"age":    {Content: "23"},
			},
		}).
		SetPrimaryKey("id").
		SetThead(types.Thead{
			{Head: "ID", Field: "id"},
			{Head: "Name", Field: "name"},
			{Head: "Gender", Field: "gender"},
			{Head: "Age", Field: "age"},
		})

	allBtns := make(types.Buttons, 0)

	// Add a ajax button action
	allBtns = append(allBtns, types.GetDefaultButton("Btn Here", icon.ArrowLeft, action.Ajax("ajax_id",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			fmt.Println("ctx request", ctx.FormValue("id"))
			return true, "ok", nil
		})))

	btns, btnsJs := allBtns.Content()
	table = table.SetButtons(btns).SetActionJs(btnsJs)

	cbs := make(types.Callbacks, 0)
	for _, btn := range allBtns {
		cbs = append(cbs, btn.GetAction().GetCallbacks())
	}

	body := table.GetContent()

	return types.Panel{
		Content: comp.Box().
			SetBody(body).
			SetNoPadding().
			SetHeader(table.GetDataTableHeader()).
			WithHeadBorder().
			SetFooter(paginator.Get(paginator.Config{
				Size:         50,
				PageSizeList: []string{"10", "20", "30", "50"},
				Param:        parameter.GetParam(ctx.Request.URL, 10),
			}).GetContent()).
			GetContent(),
		Title:       "Table",
		Description: "table example",
		Callbacks:   cbs,
	}, nil
}
