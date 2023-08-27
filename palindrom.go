package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// Menghapus spasi dan karakter non-alphanumeric dari string
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	// Membalikkan string
	reversed := ""
	for _, char := range s {
		reversed = string(char) + reversed
	}

	result := s == reversed
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan string: ")
	scanner.Scan()
	input := scanner.Text()

	if isPalindrome(input) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
