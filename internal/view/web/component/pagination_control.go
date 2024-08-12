package component

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type PaginationControlParams struct {
	ID                 string
	Name               string
	Label              string
	Required           bool
	HelpText           string
	Children           []gomponents.Node
	HelpButtonChildren []gomponents.Node
}

func PaginationControl(params PaginationControlParams) gomponents.Node {
	id := params.ID
	if id == "" {
		id = "pagination-control-" + uuid.NewString()
	}

	return html.Div(
		components.Classes{
			"form-control w-full": true,
		},
		html.Div(
			html.Class("label flex justify-start"),
			html.Label(
				html.For(id),
				html.Class("flex justify-start items-center space-x-1"),
				SpanText(params.Label),
				gomponents.If(
					params.Required,
					lucide.Asterisk(html.Class("text-error")),
				),
			),
			gomponents.If(
				len(params.HelpButtonChildren) > 0,
				HelpButtonModal(HelpButtonModalParams{
					ModalTitle: params.Label,
					Children:   params.HelpButtonChildren,
				}),
			),
		),
		html.Div(
			html.Class("join"),
			gomponents.Group(params.Children),
		),
		gomponents.If(
			params.HelpText != "",
			html.Label(
				html.Class("label"),
				html.For(id),
				SpanText(params.HelpText),
			),
		),
	)
}
