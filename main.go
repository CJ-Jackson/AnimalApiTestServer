package main

import (
	"AnimalApiTestServer/animal"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", animal.GetDocumentation)
	http.HandleFunc("/doc/openapi", animal.GetOpenAPIDocumentation)

	http.HandleFunc("/animal/delete/{id}", animal.DeleteAnimal)
	http.HandleFunc("/animal/get/{id}", animal.GetAnimal)
	http.HandleFunc("/animal/list", animal.ListAnimal)
	http.HandleFunc("/animal/post", animal.PostAnimal)
	http.HandleFunc("/animal/update/{id}", animal.PatchAnimal)

	log.Fatal(http.ListenAndServe(":18080", nil))
}
