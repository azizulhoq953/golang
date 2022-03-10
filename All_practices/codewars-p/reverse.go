package kata

import "fmt"

func Solution(word string) string {
	byte_word := []rune(word)
	for i, j := 0, len(byte_word)-1; i < j; i, j = i+1, j-1 {
		byte_word[i], byte_word[j] = byte_word[j], byte_word[i]
	}
	return string(byte_word)
}

func main() {
	fmt.Println(Solution("World"))
	fmt.Println(Solution("word"))

}
