package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/google/uuid"
)

type Books struct {
	uuid     string
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

func remove(slice []Books, s int) []Books {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	books := []Books{}
	var (
		command int
		uuid    string
	)

	var found bool = false

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
			for _, book := range books {
				fmt.Println("ID:\t", book.uuid)
				fmt.Println("Title:\t", book.title)
				fmt.Println("Price:\t", book.price)
				fmt.Println("Category:\t", book.category, "\n===")
			}
			fmt.Println("")
		case 2:
			title, price, category := scanner()
			uuid := generateUUID()

			books = append(books, Books{uuid, title, price, category})

			sort.Slice(books, func(i, j int) bool {
				return books[i].title < books[j].title
			})

			fmt.Printf("Book added!\n\n")
		case 3:
			fmt.Print("Enter ID: ")
			fmt.Scan(&uuid)

			for index, book := range books {
				if book.uuid == uuid {
					found = true
					title, price, category := scanner()
					books = remove(books, index)
					books = append(books, Books{uuid, title, price, category})

					sort.Slice(books, func(i, j int) bool {
						return books[i].title < books[j].title
					})

					fmt.Printf("Book updated!\n\n")
				}
			}

			if !found {
				fmt.Printf("Book is not exist!\n\n")
			}
		case 4:
			fmt.Print("Enter ID: ")
			fmt.Scan(&uuid)

			for index, book := range books {
				if book.uuid == uuid {
					found = true
					books = remove(books, index)

					fmt.Printf("Book deleted!\n\n")
				}
			}

			if !found {
				fmt.Printf("Book is not exist!\n\n")
			}
		case 5:
			fmt.Println("Bye...")
		default:
			fmt.Printf("\nWrong command!\n\n")
		}
	}
}
