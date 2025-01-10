package animal

import (
	"fmt"
	"net/http"
	"strconv"
)

type Animal struct {
	Id          int64 `xml:",attr"`
	Name        string
	Description string
}

type SingleAnimal struct {
	Animal Animal
}

type AnimalCollection struct {
	Animals []Animal
}

type DeleteAnimalData struct {
	Id int64
}

type Animals []Animal

var collectionOfAnimals Animals = Animals{
	Animal{1, "Cat", "House Cat"},
	Animal{2, "Dog", "House Dog"},
	Animal{3, "Lion", "Big Cat"},
	Animal{4, "Tiger", "Cat with strips"},
	Animal{5, "Meerkat", "Fun little creature"},
	Animal{6, "Charmander", "Fire type pokemon"},
	Animal{7, "Seadra", "Water pokemon with a cool name"},
	Animal{8, "Pikachu", "Ash's companion"},
	Animal{9, "Chimpanzee", "Often confused as a Monkey"},
	Animal{10, "Gorilla", "Rare animal that live in Central Africa"},
}

func ListAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}
	Encode(w, r, AnimalCollection{Animals: collectionOfAnimals}, http.StatusOK)
}

func GetAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest)
		return
	}
	if int64(len(collectionOfAnimals)) <= id {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusNotFound)}, http.StatusNotFound)
		return
	}
	animal := collectionOfAnimals[id]
	Encode(w, r, SingleAnimal{Animal: animal}, http.StatusOK)
}

func PostAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}

	singleAnimal := &SingleAnimal{}
	err := Decode(singleAnimal, r)
	if err != nil {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest)
		return
	}

	collectionOfAnimals = append(collectionOfAnimals, singleAnimal.Animal)
	Encode(w, r, *singleAnimal, http.StatusOK)
}

func DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest)
		return
	}

	if int64(len(collectionOfAnimals)) <= id {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusNotFound)}, http.StatusNotFound)
		return
	}

	collectionOfAnimals = append(collectionOfAnimals[:id], collectionOfAnimals[id+1:]...)

	Encode(w, r, Message{
		Message: fmt.Sprintf("'%d' has been deleted.", id),
	}, http.StatusOK)
}

func PatchAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusMethodNotAllowed)}, http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest)
		return
	}

	singleAnimal := &SingleAnimal{}
	err = Decode(singleAnimal, r)
	if err != nil {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest)
		return
	}

	if int64(len(collectionOfAnimals)) <= id {
		Encode(w, r, HttpStatus{Status: http.StatusText(http.StatusNotFound)}, http.StatusNotFound)
		return
	}

	collectionOfAnimals[int(id)] = singleAnimal.Animal

	Encode(w, r, Message{fmt.Sprintf("'%d' has been updated", id)}, http.StatusOK)
}
