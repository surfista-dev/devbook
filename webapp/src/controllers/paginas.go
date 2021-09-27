package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
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

//CarregarPaginaPrincipal carrega pagina principal com publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, erro)
	utils.ExecutarTemplates(w, "home.html", nil)
}
