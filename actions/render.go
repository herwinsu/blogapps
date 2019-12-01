package actions

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// "csrf": func() template.HTML {
			// 	return template.HTML("<input name=\"authenticity_token\" value=\"<%= authenticity_token %>\" type=\"hidden\">")
			// },
			"csrf": func(ctx plush.HelperContext) (template.HTML, error) {
				tok, ok := ctx.Value("authenticity_token").(string)
				if !ok {
					return "", fmt.Errorf("expected CSRF token got %T", ctx.Value("authenticity_token"))
				}
				t := tags.New("input", tags.Options{
					"value": tok,
					"type":  "hidden",
					"name":  "authenticity_token",
				})
				return t.HTML(), nil
			},
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
		},
	})
}
