package controller

import "net/http"

/* CriarUsuario - Insere usuario */
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

/* BuscarUsuarios - Busca todos os usuario */
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando os Usuários"))
}

/* BuscarUsuarioPorId - busca usuario por id */
func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário"))
}

/* AtualizarUsuario - atualiza usuario */
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário"))
}

/* DeletarUsuario - deleta usuario */
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário"))
}
