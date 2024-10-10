package handlers

import (
	"library/internal/templates"
	"library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"library/internal/utils/ui"
	"net/http"
	"net/url"

	"github.com/a-h/templ"
)

type AboutHandLer struct{}

func NewAboutHandler() *AboutHandLer {
	return &AboutHandLer{}
}

func (h *AboutHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.About()
	hxBoostedHeader := r.Header.Get(requestHeaders.HxBoosted)

	var err error
	if hxBoostedHeader == "true" {
		originUrl, _ := url.Parse(r.Header.Get(requestHeaders.HxCurrentURL))

		oobSwaps := []templ.Component{
			templates.DisabledNavbarLink(ui.IdAnchorAbout, ui.TextAnchorAbout, true),
		}
		if anchor, anchorExists := ui.PathToAnchor[originUrl.Path]; anchorExists {
			oobSwaps = append(oobSwaps, templates.EnabledNavbarLink(anchor.Id, anchor.Text, originUrl.Path, true))
		}

		err = templates.ContentsWithTitle(c, ui.TitleAbout, oobSwaps).Render(r.Context(), w)
	} else {
		err = templates.Layout(c, ui.TitleAbout, "/about").Render(r.Context(), w)
	}

	if err != nil {
		errors.ServerError(r.Context(), w, err, "Error rendering template")
	}
}
