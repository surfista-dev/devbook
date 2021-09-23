package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//CarregarTelaDeLogin renderizar tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}

//CarregarPaginaDeCadastroUsuario carrega pagina de cadastro do usuários
func CarregarPaginaDeCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "cadastro.html", nil)
}
