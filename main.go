package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/devas", Desenvolvedores)

	http.ListenAndServe(":5000", mux)
}

func Desenvolvedoras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonStr := `[{"name":"Silvana","age":29,"favlang":"Vue", "twitter":"@silcrana"}, 
				{"name":"Marcela","age":31,"favlang":"Golang","twitter":"@masisiliani"},
				{"name":"Jessica Temporal","age":25,"favlang":"Python","twitter":"@jessTemporal"},
				{"name":"Marylly","age":25,"favlang":"Java","twitter":"@MaryllyOficial"},				
				{"name":"Lauren","age":23,"favlang":"Go","twitter":"@laurienmf"},			
				{"name":"Erika","age":23,"favlang":"Go","twitter":"@erikones_"},
				{"name":"Jessica Nathany","age":27,"favlang":".NET","twitter":"@JessicaNathanyF"}]`

	w.Write([]byte(jsonStr))
}
