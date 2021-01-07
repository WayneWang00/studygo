package iptest

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	area := string(Find("8.8.8.8"))
	fmt.Println(area)
	if area != "GOOGLE\tGOOGLE" {
		t.Error("ip find area error!")
	}
	area1 := string(Find("14.116.139.99"))
	fmt.Println(area1)
}
