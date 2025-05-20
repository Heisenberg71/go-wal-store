package wal

import (
	"fmt"
	"log"
	"os"
	"simplewal/proto"

	"google.golang.org/protobuf/proto"
)

func Add(a int, b int) int {
	return a + b
}

func Save(key string, value string) string {

	wal := &proto.WAL_Entry{
		isCheckpoint: proto.Bool(true),
	}

	data, err := proto.Marshal(wal)
	if err != nil {
		log.Fatalf("Marshal error: %v", err)
	}

	fmt.Println("Serialized data: %s", data)
	os.WriteFile("person.bin", data, 0644)

	return "key: " + key + " value: " + value
}
