package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFormContent(ctx *context.Context) (types.Panel, error) {

	components := template.Get(config.GetTheme())

	col1 := components.Col().GetContent()
	btn1 := components.Button().SetType("submit").
		SetContent(language.GetFromHtml("Save")).
		SetThemePrimary().
		SetOrientationRight().
		SetLoadingText(icon.Icon("fa-spinner fa-spin", 2) + `Save`).
		GetContent()
	btn2 := components.Button().SetType("reset").
		SetContent(language.GetFromHtml("Reset")).
		SetThemeWarning().
		SetOrientationLeft().
		GetContent()
	col2 := components.Col().SetSize(types.SizeMD(8)).
		SetContent(btn1 + btn2).GetContent()

	var panel = types.NewFormPanel()
	panel.AddField("Name", "name", db.Varchar, form.Text)
	panel.AddField("Age", "age", db.Int, form.Number)
	panel.AddField("HomePage", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")
	panel.AddField("Email", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")
	panel.AddField("Birthday", "birthday", db.Varchar, form.Datetime).FieldDefault("2010-09-05")
	panel.AddField("Password", "password", db.Varchar, form.Password)
	panel.AddField("IP", "ip", db.Varchar, form.Ip)
	panel.AddField("Certificate", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	})
	panel.AddField("Money", "currency", db.Int, form.Currency)
	panel.AddField("Content", "content", db.Text, form.RichText).
		FieldDefault(`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`)

	panel.AddField("Website", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("The Website will not be able to access after closing, the admin system still can login").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	panel.AddField("Fruit", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "Apple", Value: "apple"},
			{Text: "Banana", Value: "banana"},
			{Text: "Watermelon", Value: "watermelon"},
			{Text: "Pear", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"Pear"}
		})
	panel.AddField("Gender", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Men", Value: "0"},
			{Text: "Women", Value: "1"},
		})
	panel.AddField("Drink", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "Beer", Value: "beer"},
			{Text: "Juice", Value: "juice"},
			{Text: "Water", Value: "water"},
			{Text: "Red Bull", Value: "red bull"},
		}).FieldDefault("beer")
	panel.AddField("Work Experience", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "two years", Value: "0"},
			{Text: "three years", Value: "1"},
			{Text: "four years", Value: "2"},
			{Text: "five years", Value: "3"},
		}).FieldDefault("beer")
	panel.SetTabGroups(types.TabGroups{
		{"name", "age", "homepage", "email", "birthday", "password", "ip", "certificate", "currency", "content"},
		{"website", "fruit", "gender", "drink", "experience"},
	})
	panel.SetTabHeaders("input", "select")

	fields, headers := panel.GroupField()

	aform := components.Form().
		SetTabHeaders(headers).
		SetTabContents(fields).
		SetPrefix(config.PrefixFixSlash()).
		SetUrl("/admin/form/update").
		SetTitle("Form").
		SetHiddenFields(map[string]string{
			form2.PreviousKey: "/admin",
		}).
		SetOperationFooter(col1 + col2)

	return types.Panel{
		Content: components.Box().
			SetHeader(aform.GetDefaultBoxHeader()).
			WithHeadBorder().
			SetBody(aform.GetContent()).
			GetContent(),
		Title:       "Form",
		Description: "form example",
	}, nil
}
