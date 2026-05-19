package main

import (
	"fmt"
	"protobuffs/endec"
	"protobuffs/personpb"
)

func main() {

	person := &personpb.Person{
		Name: "Mahboob",
		Id:   10,
	}

	fields := []endec.Field{
		{
			FieldName:   "name",
			FieldSeqNum: 1,
			FieldType:   "string",
		},
		{
			FieldName:   "id",
			FieldSeqNum: 2,
			FieldType:   "int32",
		},
	}

	arr1 := endec.Encoder(person, fields)

	fmt.Printf("% x\n", arr1)

	testPerson := endec.Decoder(arr1, person, fields)

	fmt.Println(testPerson)

}
