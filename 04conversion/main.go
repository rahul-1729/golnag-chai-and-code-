package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcom to our Pizza")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	reader.ReadString()
}
