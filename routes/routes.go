package routes

import (
	"net/http"
	"retro_store/controllers"
)

func Rotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/produto", controllers.Produto)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
