package main

import (
	"fmt"
	"log"

	"github.com/inegmetov/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File is restored", restoredFile)
}
