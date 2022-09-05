package main

import "fmt"

type Student struct {
	name  []string
	score []int
}


func (s Student) Avarage() float64 {
	sum := 0
	for _, value := range s.score { sum += value }
	return float64(sum / len(s.score))
}


func (s Student) Min() (min int, name string) {
	indexMin, min := 0, 999999
	for index, value := range s.score {
		if value < min {
			min = value
			indexMin = index
		}
	}
	return s.score[indexMin], s.name[indexMin]
}


func (s Student) Max() (max int, name string) {
	indexMax, max := 0, -999999
	for index, value := range s.score {
		if value > max {
			max = value
			indexMax = index
		}
	}
	return s.score[indexMax], s.name[indexMax]
}


func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}