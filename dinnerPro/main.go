package main

import (
	"dinnerPro/dinnergacha"
	"net/http"
	"text/template"
)

type GachaData struct {
	Menu   string
	Result string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	// DinnerMenuGachaの実行
	menu, result := dinnergacha.DinnerMenuGacha()

	gachaData := GachaData{
		Menu:   menu,
		Result: result,
	}

	t, err := template.ParseFiles("result.html")

	if err != nil {
		panic(err.Error())
	}

	if err := t.Execute(w, gachaData); err != nil {
		panic(err.Error())
	}
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/result/", resultHandler)
	http.ListenAndServe(":8000", nil)
}
