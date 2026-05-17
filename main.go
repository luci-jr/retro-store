package main

import (
	"net/http"
	"retro_store/routes"
)

func main() {
	routes.Rotas()

	// Configura o Go para servir os arquivos da pasta "images" na rota "/images/"
	fsImages := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fsImages))

	// Configura o Go para servir arquivos estáticos (CSS, JS) da pasta "static"
	fsStatic := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))
	
	http.ListenAndServe(":8000", nil)
}
