package main

import (
	"encoding/json"
	"fmt"
)

type Person1 struct {
	UserName       string   `json:"user_name"`
	FavoriteNumber int      `json:"favorite_number"`
	Interests      []string `json:"interests"`
}

func decode(arr []byte) Person1 {
	i := 0
	person := Person1{}
	for i < len(arr) {
		num := arr[i]

		buffType := num & 0b111
		buffTag := num >> 3
		fmt.Println(num, buffTag, buffType)

		switch buffType {
		case 2:
			length := int(arr[i+1])
			stringArr := arr[i+2 : i+2+length]
			str := string(stringArr)
			if i == 0 {
				person.UserName = str
			} else {
				person.Interests = append(person.Interests, str)
			}
			fmt.Println(str)
			i = i + 2 + length
			break
		case 0:
			bit := 0
			num := 0
			for arr[i+1]>>7 != 0 {
				curr := int(arr[i+1]) & 0x7F
				num = num + curr<<bit

				bit += 7
				fmt.Println(num)
				i++

			}

			curr := int(arr[i+1]) & 0x7F
			num = num + curr<<bit
			bit += 7
			fmt.Println(num)
			i = i + 2
			person.FavoriteNumber = num

			break

		default:
			//

		}

	}

	return person

}

func main() {

	arr := []byte{
		0x0a, 0x06, 0x4d, 0x61, 0x72, 0x74, 0x69, 0x6e,
		0x10, 0xb9, 0x0a, 0x1a, 0x0b, 0x64, 0x61, 0x79,
		0x64, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
		0x1a, 0x07, 0x68, 0x61, 0x63, 0x6b, 0x69, 0x6e,
		0x67,
	}
	testPerson := decode(arr)
	b, err := json.MarshalIndent(testPerson, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	person := Person{
		Name:  "John Doe",
		ID:    1234,
		Email: "john@example.com",
	}
	fmt.Println(en(person))

}
