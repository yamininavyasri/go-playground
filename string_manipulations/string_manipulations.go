package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "  Hello, GoLang World!  "

	// 1. Basic Info
	fmt.Println("Original String:", str)
	fmt.Println("Length:", len(str))

	// 2. Trim spaces
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed:", trimmed)

	// 3. Change case
	fmt.Println("To Upper:", strings.ToUpper(trimmed))
	fmt.Println("To Lower:", strings.ToLower(trimmed))
	fmt.Println("Title Case:", strings.Title(trimmed)) // Deprecated but works

	// 4. Contains, HasPrefix, HasSuffix
	fmt.Println("Contains 'Go':", strings.Contains(trimmed, "Go"))
	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(trimmed, "Hello"))
	fmt.Println("HasSuffix 'World!':", strings.HasSuffix(trimmed, "World!"))

	// 5. Replace
	replaced := strings.ReplaceAll(trimmed, "World", "Universe")
	fmt.Println("Replaced:", replaced)

	// 6. Split and Join
	split := strings.Split(trimmed, " ")
	fmt.Println("Split:", split)
	joined := strings.Join(split, "-")
	fmt.Println("Joined:", joined)

	// 7. Count occurrences
	fmt.Println("Count of 'l':", strings.Count(trimmed, "l"))

	// 8. Index and LastIndex
	fmt.Println("First index of 'o':", strings.Index(trimmed, "o"))
	fmt.Println("Last index of 'o':", strings.LastIndex(trimmed, "o"))

	// 9. Compare strings
	fmt.Println("Compare 'go' and 'Go':", strings.Compare("go", "Go")) // 1 if first > second
	fmt.Println("EqualFold (case-insensitive):", strings.EqualFold("go", "Go"))

	// 10. Repeat
	fmt.Println("Repeat 'Go' 3 times:", strings.Repeat("Go", 3))

	// 11. Fields (split by whitespace)
	fmt.Println("Fields:", strings.Fields("  Go is fun  "))

	// 12. Rune-based iteration (for Unicode)
	fmt.Println("\nIterating characters:")
	for i, r := range trimmed {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}

	// 13. Filter only letters
	letters := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, trimmed)
	fmt.Println("Only letters:", letters)

	// 14. String concatenation
	part1 := "Go"
	part2 := "Lang"
	fmt.Println("Concatenated:", part1+part2)

	// 15. Convert string to byte slice and back
	bytes := []byte(trimmed)
	fmt.Println("Bytes:", bytes)
	backToString := string(bytes)
	fmt.Println("Back to String:", backToString)
}
