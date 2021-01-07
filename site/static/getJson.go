package static

import (
	"fmt"
	"os"
)

func GetFile() {
	str, err := os.Open("./static/cdn.json")
	if err != nil {
		fmt.Println("open Error: ", err)
		return
	}
	fmt.Printf("open Value: %+v\n", str)
}
