insert into usuarios (nome, nick, email, senha)
values 
("Usuario 1", "Usuario 1", "usuario1@gmail.com", "$2a$10$hdrP/ZyD8tNlyZq4dlQghOBLOBPl2tWrifda530VoghhLGZoXK88G"), -- usuario1
("Usuario 2", "Usuario 2", "usuario2@gmail.com", "$2a$10$hdrP/ZyD8tNlyZq4dlQghOBLOBPl2tWrifda530VoghhLGZoXK88G"), -- usuario2
("Usuario 3", "Usuario 3", "usuario3@gmail.com", "$2a$10$hdrP/ZyD8tNlyZq4dlQghOBLOBPl2tWrifda530VoghhLGZoXK88G"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa e a publicação do Usuário 1 Oba!", 1),
("Publicação do Usuário 2", "Essa e a publicação do Usuário 2 Oba!", 2),
("Publicação do Usuário 3", "Essa e a publicação do Usuário 3 Oba!", 3);