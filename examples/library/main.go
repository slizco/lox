package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/slizco/lox"
)

// In this example, we use a standard lox.Bank struct
// to represent our library. In this library, we have
// one copy of each book, so only one book can be checked
// out at a time. We use the notion of "locking" to represent
// checking out a book, and the notion of "unlocking" to
// represent returning a book.

var (
	library = lox.NewBank()

	bookA = "One Flew Over the Cuckoo's Nest"
	bookB = "All the Light We Cannot See"
)

type reader string

func (_ reader) checkout(book string) {
	library.Lock(book)
}

func (_ reader) checkIn(book string) {
	library.Unlock(book)
}

func main() {
	pip := reader("Pip")
	agatha := reader("Agatha")

	var wg sync.WaitGroup // needed just to show print statements
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Printf("%s is looking for %s.\n", pip, bookA)
		pip.checkout(bookA)
		fmt.Printf("%s has checked out %s.\n", pip, bookA)
		fmt.Printf("%s is reading %s.\n", pip, bookA)
		time.Sleep(time.Second)
		pip.checkIn(bookA)
		fmt.Printf("%s has returned %s.\n", pip, bookA)
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("%s is looking for %s.\n", agatha, bookA)
		agatha.checkout(bookA)
		fmt.Printf("%s has checked out %s.\n", agatha, bookA)
		fmt.Printf("%s is reading %s.\n", agatha, bookA)
		time.Sleep(time.Second)
		agatha.checkIn(bookA)
		fmt.Printf("%s has returned %s.\n", agatha, bookA)
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("%s is looking for %s.\n", agatha, bookB)
		agatha.checkout(bookB)
		fmt.Printf("%s has checked out %s.\n", agatha, bookB)
		fmt.Printf("%s is reading %s.\n", agatha, bookB)
		time.Sleep(time.Second)
		agatha.checkIn(bookB)
		fmt.Printf("%s has returned %s.\n", agatha, bookB)
	}()

	wg.Wait()
}
