package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// função para gerar o secret base 64 para assinar o token
// func init(){
// 	key := make([]byte, 64)
// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}
// 	base64String := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(base64String)
// }

func main() {
	config.Load()
	r := router.Generate()
	fmt.Printf("Server running in %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}