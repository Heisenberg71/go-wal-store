package main

import (
	"fmt"
	"wal/internal/wal"
)

func main() {
	fmt.Println("WAL!")

	key1 := "name"
	value1 := "Samin"

	key2 := "age"
	value2 := "26"

	wal.Add(key1, value1)
	wal.Add(key2, value2)
}
