package main

import (
	"fmt"

	"simplewal/internal/wal"
)

func main() {

	key := "name"
	value := "Samin"

	savedResult := wal.Save(key, value)
	fmt.Printf("Saved: %s\n", savedResult)
}
