package main

import (
	fa "github.com/joinName177/love-001-field-archive"
	"log"
	"net/http"
)

func main() {
	c := fa.LoadConfig()
	s, e := fa.OpenJSONStore(c.DataPath)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	log.Fatal(http.ListenAndServe(c.Address, fa.Routes(fa.NewApplication(s, fa.InlineRenderer{}))))
}
