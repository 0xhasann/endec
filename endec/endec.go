package endec

import (
	"bytes"
	"protobuffs/personpb"
	"reflect"
	"strings"
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

		default:

		}

	}

	return person

}

func Encoder(person any, fields []Field) []byte {
	v := reflect.ValueOf(person).Elem()

	var buf bytes.Buffer

	for _, field := range fields {
		fieldName := strings.ToUpper(field.FieldName[:1]) + field.FieldName[1:]
		value := v.FieldByName(fieldName)
		switch field.FieldType {

		case "string":
			str := value.String()
			nameBytes := []byte(str)

			buf.WriteByte(byte(field.FieldSeqNum<<3 | 2))

			buf.WriteByte(byte(len(nameBytes)))
			buf.Write(nameBytes)

		case "int32":
			num := value.Int()
			buf.WriteByte(byte(field.FieldSeqNum << 3)) // int32 wire type is 0 so ignoring it.
			buf.WriteByte(byte(num))
		}
	}

	return buf.Bytes()
}
