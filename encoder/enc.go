package encoder

type Person map[string]interface{}

type Proto struct {
	field_name  string
	field_value []FieldValue
}

type FieldValue struct {
	field_name    string
	field_seq_num int32
	field_type    int32
}

//	person := Person{
//		Name:  "John Doe",
//		ID:    1234,
//		Email: "john@example.com",
//	}
//
// { age: 20, name: "JAck" , unknown: "hii"}
//
//	{
//		"Person" : [
//			{
//				field_name: "Name",
//				field_seq_num: "",
//				field_type: 0,
//			},
//			{
//				field_name: "Id",
//				field_seq_num: "",
//				field_type: 1,
//			},
//		]
//	}
//
// bytes[]
// iterate over Person Array
// ()
// value
// bytes.push([tag, value])
func en(user Person, personProto Proto) []byte {
	var result []byte

	for _, value := range personProto.field_value {
		// data := user[value.field_name]
		tag := value.field_seq_num<<3 | value.field_type
		result = append(result, byte(tag))

	}

	return result

}
