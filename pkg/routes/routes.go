package routes

import (
	"net/http"
	"github.com/SergioCB20/RutaEdu-Backend/pkg/handlers"
	"github.com/SergioCB20/RutaEdu-Backend/pkg/middlewares"
)

func SetRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/reader",func(w http.ResponseWriter, r *http.Request) {
		if(r.Method == "POST"){
			handlers.ReadCurriculumGrid(w,r)
		}else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	corsRouter := middlewares.CorsMiddleware(router)

    return corsRouter
}