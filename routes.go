package fieldarchive

import "net/http"

func Routes(a *API) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("POST /renders", a.Render)
	m.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
	return m
}
