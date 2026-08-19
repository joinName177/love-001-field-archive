package fieldarchive

import (
	"encoding/json"
	"errors"
	"net/http"
)

type API struct{ render *RenderService }

func NewAPI(r *RenderService) *API { return &API{r} }
func (a *API) Render(w http.ResponseWriter, r *http.Request) {
	var q RenderRequest
	if e := json.NewDecoder(r.Body).Decode(&q); e != nil {
		http.Error(w, "bad request", 400)
		return
	}
	out, e := a.render.Render(r.Context(), q)
	if e != nil {
		if errors.Is(e, ErrRenderCancelled) {
			http.Error(w, "render cancelled", 499)
			return
		}
		http.Error(w, "render unavailable", 500)
		return
	}
	json.NewEncoder(w).Encode(out)
}
