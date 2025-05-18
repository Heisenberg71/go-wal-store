package wal

import (
	"fmt"
	"log"
	"os"
	"wal/proto"

	"google.golang.org/protobuf/proto"
)

func Add(name string, age string) {

	wal := &proto.Wal{
		Key:   name,
		Value: age,
	}

	data, err := proto.Marshal(wal)
	if err != nil {
		log.Fatalf("Marshal error: %v", err)
	}

	fmt.Println("Serialized data: %s", data)
	os.WriteFile("keyValue.bin", data, 0644)
}
