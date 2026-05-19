package endec

import (
	"bytes"
	"protobuffs/protobuffs/personpb"
)

type Field struct {
	FieldName   string
	FieldSeqNum int
	FieldType   string
}

type ProtoStruct map[string][]Field

func Decoder(arr []byte) *personpb.Person {
	person := &personpb.Person{}
	i := 0
	for i < len(arr) {
		num := arr[i]

		buffType := num & 0b111
		// buffTag := num >> 3

		switch buffType {
		case 2:
			length := int(arr[i+1])
			stringArr := arr[i+2 : i+2+length]
			str := string(stringArr)
			person.Name = str
			i = i + 2 + length
			break
		case 0:
			bit := 0
			num := 0
			for arr[i+1]>>7 != 0 {
				curr := int(arr[i+1]) & 0x7F
				num = num + curr<<bit

				bit += 7
				i++

			}

			curr := int(arr[i+1]) & 0x7F
			num = num + curr<<bit
			bit += 7
			i = i + 2
			person.Id = int32(num)

			break

		default:

		}

	}

	return person

}

func Encoder(user *personpb.Person, p ProtoStruct) []byte {
	fields := p["Person"]

	var buf bytes.Buffer

	for _, field := range fields {
		switch field.FieldName {

		case "name":
			buf.WriteByte(byte(field.FieldSeqNum<<3 | 2))
			nameBytes := []byte(user.Name)
			buf.WriteByte(byte(len(nameBytes)))
			buf.Write(nameBytes)

		case "id":
			buf.WriteByte(byte(field.FieldSeqNum<<3 | 0))
			buf.WriteByte(byte(user.Id))
		}
	}

	return buf.Bytes()
}
