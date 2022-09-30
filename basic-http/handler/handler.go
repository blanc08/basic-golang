package handler

import (
	"basic-http/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Hellow Root"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]string{
	// 	"title":          "Golang basic http web",
	// 	"welcomeMessage": "Halo, this is a basic web example made by GO",
	// }

	data := []entity.Product{
		{ID: 1, Name: "Asus", Price: 20, Stock: 2},
		{ID: 1, Name: "HP", Price: 15, Stock: 4},
		{ID: 1, Name: "Macbool", Price: 25, Stock: 5},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow Word"))
}

func BagusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow Bagus"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Product number : %d", idNum)

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title": "Golang basic http web",
		"id":    id,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}
