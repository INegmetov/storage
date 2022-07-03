package main

import (
	"fmt"

	"github.com/inegmetov/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello wold!", st)
}
