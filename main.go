package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/devs", Desenvolvedores)

	http.ListenAndServe(":5000", mux)
}

func Desenvolvedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonStr := `[{"Nome":"Silvana","Idade":29,"Especialidade":"Vue"}, 
				{"Nome":"Marcela","Idade":31,"Especialidade":"GoLang"},
				{"Name":"Jessica","Idade":27,"Especialidade":".NET"}]`

	w.Write([]byte(jsonStr))
}
