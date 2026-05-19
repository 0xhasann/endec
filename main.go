package main

import (
	"fmt"
	"protobuffs/endec"
)

func main() {

	person := endec.Person{
		Name: "Mahboob",
		Id:   10,
	}

	schema := endec.ProtoStruct{
		"Person": {
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
		},
	}

	arr1 := endec.Encoder(person, schema)
	// fmt.Println(arr1)
	fmt.Printf("% x\n", arr1)

	testPerson := endec.Decoder(arr1)

	fmt.Println(testPerson)

}
