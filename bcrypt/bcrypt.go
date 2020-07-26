package main

import (
	"fmt"
	// importação do pacote crypto
	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "minhasenha"
	senhaerrada := "suasenha"

	// GenerateFromPassword encripta um slice de bytes em um hash
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))
	// ==> $2a$10$S3H8wX4DMKAn3b9PF2N.zemg/i3xLsTcpUJwTF6NQJ2gTXgCtzqo2

	// CompareHashAndPassword compara o hash com um password para saber se aquele hash
	// é proveniente daquela senha, caso contrário retorna um erro, quando obtem sucesso retorna nil
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	// ===> nil
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))
	// ===> crypto/bcrypt: hashedPassword is not the hash of the given password
}
