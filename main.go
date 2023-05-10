package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/joho/godotenv"
)

type FrontPage struct {
	Title string
}

func main() {
	tmpl := make(map[string]*template.Template)
	tmpl["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d, err := deta.New()
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return
	}

	db, err := base.New(d, "base_name")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err, db)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("INDEX")
		// fmt.Println(t.DefinedTemplates())
		data := FrontPage{Title: "Hello"}
		tmpl["index"].ExecuteTemplate(w, "layout.html", data)
		// template.ExecuteTemplate(w, "index.html", nil)
		// http.ServeFile(w, r, "templates/index.html")
		// w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}
