package main

import (
	"AnimalApiTestServer/animal"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/list-animals", animal.ListAnimal)
	http.HandleFunc("/animal/{id}", animal.GetAnimal)
	http.HandleFunc("/post-animal", animal.PostAnimal)
	http.HandleFunc("/delete-animal", animal.DeleteAnimal)
	http.HandleFunc("/patch-animal/{id}", animal.PatchAnimal)

	log.Fatal(http.ListenAndServe(":18080", nil))
}
