package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func main() {

	users := []*User{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "Dave"},
		{5, "Eve"},
	}

	fmt.Println("До удаления:")
	printSlice(users)

	i := 2

	copy(users[i:], users[i+1:])

	users[len(users)-1] = nil

	users = users[:len(users)-1]

	fmt.Println("\nПосле удаления (удален Charlie):")
	printSlice(users)

	fmt.Printf("\nLen: %d, Cap: %d\n", len(users), cap(users))
}

func printSlice(users []*User) {
	for i, u := range users {
		fmt.Printf("[%d] %s (ID: %d)\n", i, u.Name, u.ID)
	}
}
