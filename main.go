package main

import (
	"fmt"
	"net/http"
)

//main: função principal do programa
func main() {
	//Criamos um server em Go
	mux := http.NewServeMux()
	//Criamos uma rota nova, que vai ser chamada em /devas utilizando como resposta o método Desenvolvedoras
	mux.HandleFunc("/devas", Desenvolvedoras)

	//Configuramos para nosso server responder na porta 5000, ou seja, http://localhost:5000/[rota]

	fmt.Println("Servindo na porta 5000")
	http.ListenAndServe(":5000", mux)
}

//Desenvolvedoras : Rota que retorna algumas das deves fodas que eu conheço
func Desenvolvedoras(w http.ResponseWriter, r *http.Request) {
	// Adicionamos um header ao nosso response, informando que o content type retornado vai ser em json
	w.Header().Set("Content-Type", "application/json")

	//criamos uma listagem, estática, de objetos.
	//Aqui é construida a listagem que vamos retornar na nossa API
	jsonStr := `[{"name":"Silvana","age":22,"favlang":"Vue", "twitter":"@silcrana"},
				{"name":"Marcela","age":75,"favlang":"Golang","twitter":"@masisiliani"},
				{"name":"Jessica Temporal","age":25,"favlang":"Python","twitter":"@jessTemporal"},
				{"name":"Marylly","age":25,"favlang":"Java","twitter":"@MaryllyOficial"},
				{"name":"Lauren","age":23,"favlang":"Go","twitter":"@laurienmf"},
				{"name":"Erika","age":23,"favlang":"Go","twitter":"@erikones_"},
				{"name":"Jessica Nathany","age":27,"favlang":".NET","twitter":"@JessicaNathanyF"}]`

	//escrevemos a listagem do nosso retorno no nosso response
	w.Write([]byte(jsonStr))
}
