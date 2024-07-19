package routes

import (
	"net/http"
	"github.com/SergioCB20/RutaEdu-Backend/pkg/handlers"
)

func SetRoutes() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("/reader",func(w http.ResponseWriter, r *http.Request) {
		if(r.Method == "POST"){
			handlers.ReadCurriculumGrid(w,r)
		}
	})
    return router
}