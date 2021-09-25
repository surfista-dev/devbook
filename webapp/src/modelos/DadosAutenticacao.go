package modelos

//DadosAtenticacao contem tokem e id do usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
