package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Books struct {
	title    string
	price    int
	category string
}

func generateUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

func scanner() (string, int, string) {
	stdin := bufio.NewReader(os.Stdin)
	var (
		title    string
		price    int
		category string
		err      error
	)

	fmt.Printf("Enter title: ")
	fmt.Scan(&title)

	fmt.Printf("Enter price: ")
	for price, err = fmt.Scan(&price); err != nil; {
		value, err := fmt.Scan(&price)
		if value == 0 {
			stdin.ReadString('\n')
			fmt.Print("\nWrong input!\n\n")
		} else {
			break
		}
		fmt.Printf("Enter price: ")

		if err == error(nil) {
			break
		}
	}

	fmt.Printf("Enter category: ")
	fmt.Scan(&category)
	return title, price, category
}

func main() {
	books := map[string]Books{}
	var (
		command int
		uuid    string
	)

	for command != 5 {
		fmt.Println("===BOOK MANAGEMENT===")
		fmt.Println("1. Get all books")
		fmt.Println("2. Create a book")
		fmt.Println("3. Update a book")
		fmt.Println("4. Delete a book")
		fmt.Println("5. Exit")

		fmt.Printf("Choose your menu: ")
		fmt.Scan(&command)

		switch command {
		case 1:
			fmt.Println("All books\n===")
			for key, book := range books {
				fmt.Println("ID:\t", key)
				fmt.Println("Title:\t", book.title)
				fmt.Println("Price:\t", book.price)
				fmt.Println("Category:\t", book.category, "\n===")
			}
			fmt.Println("")
		case 2:
			title, price, category := scanner()
			uuid := generateUUID()

			books[uuid] = Books{title, price, category}

			fmt.Printf("Book added!\n\n")
		case 3:
			fmt.Print("Enter ID: ")
			fmt.Scan(&uuid)

			if _, isExist := books[uuid]; isExist {
				title, price, category := scanner()
				books[uuid] = Books{title, price, category}
				fmt.Printf("Book updated!\n\n")
			} else {
				fmt.Printf("Book is not exist!\n\n")
			}
		case 4:
			fmt.Print("Enter ID: ")
			fmt.Scan(&uuid)

			if _, isExist := books[uuid]; isExist {
				delete(books, uuid)
				fmt.Printf("Book deleted!\n\n")
			} else {
				fmt.Printf("Book is not exist!\n\n")
			}
		case 5:
			fmt.Println("Bye...")
		default:
			fmt.Printf("\nWrong command!\n\n")
		}
	}
}
