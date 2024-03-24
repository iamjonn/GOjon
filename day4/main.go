package main

//endpoint para listar usuários

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// user é uma struct que representa um usuário
type user struct {
	ID    int
	Name  string
	email string
}

func main() {
	mux := http.NewServeMux()                  // cria um novo multiplexador
	mux.HandleFunc("/users", listUsersHandler) // adiciona a função listUsersHandler ao multiplexador

	http.ListenAndServe(":3000", mux) // inicia o servidor na porta 8080
}

func listUsersHandler(w http. ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "user.db") // abre a conexão com o banco de dados
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // se der erro, retorna um erro 500
		return
	}
	defer db.Close() // fecha a conexão com o banco de dados depois de usar

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // se der erro, retorna um erro 500
		return
	}
	defer rows.Close() // fecha o cursor depois de usar

	users := []*user{} // cria um slice de ponteiros para usuários
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // se der erro, retorna um erro 500
			return
		}
		users = append(users, &u) // adiciona o ponteiro para o usuário ao slice
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // se der erro, retorna um erro 500
		return
	}

}
