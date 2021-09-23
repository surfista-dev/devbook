$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();

    if ($('#senha').val() != $('#confirma-senha').val()) {
        alert("As senha não coincidem");
        return;
    } 

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function (){
        alert("Usuário cadastrado com sucesso");
    }).fail(function (erro){
        console.log(erro);
        alert("Erro ao cadastra Usuário");
    });
}