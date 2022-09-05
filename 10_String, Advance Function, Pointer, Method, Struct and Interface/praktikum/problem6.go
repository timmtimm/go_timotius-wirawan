package main

import (
	"fmt"
	// "math"
)

type student struct {
   name       string
   nameEncode string
   score      int
}

type Chiper interface {
   Encode() string
   Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, char := range s.name {
		if char >= 'A' && char <= 'M' {
			nameEncode += string(78 + (77 - int(char)))
		} else if char >= 'N' && char <= 'Z' {
			nameEncode += string(78 - (int(char) - 77))
		} else if char >= 'a' && char <= 'm' {
			nameEncode += string(110 + (109 - int(char)))
		} else {
			nameEncode += string(110 - (int(char) - 109))
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, char := range s.nameEncode {
		if char >= 'A' && char <= 'M' {
			nameDecode += string(78 + (77 - int(char)))
		} else if char >= 'N' && char <= 'Z' {
			nameDecode += string(78 - (int(char) - 77))
		} else if char >= 'a' && char <= 'm' {
			nameDecode += string(110 + (109 - int(char)))
		} else {
			nameDecode += string(110 - (int(char) - 109))
		}
	}

   return nameDecode
}

func main() {
	var menu int

	var a = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}