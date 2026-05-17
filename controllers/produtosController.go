package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"retro_store/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	totalProdutos := models.BuscaDeProdutos()
	templ.ExecuteTemplate(w, "Index", totalProdutos)
}

func Produto(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Produto", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preço:", err)
		}
		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	templ.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		
		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter id:", err)
		}
		
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preço:", err)
		}
		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade:", err)
		}
		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}