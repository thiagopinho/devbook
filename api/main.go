package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//gerar token

// func init() {
// 	key := make([]byte, 64)

// 	if _, erro := rand.Read(key); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)

// }

func main() {
	config.Carregar()
	fmt.Println(config.SecretKey)
	fmt.Printf("Escutando na porta %d", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
