package main

import "fmt"

func main() {
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(c())
}
func a() (result int) {
	defer func() {
		result=result+1
		fmt.Println("里面", result)
	}()
	fmt.Println("外面", result)
	return 8
}
func b() (r int) {
	t := 5
	defer func() {
		//t = t + 5
		t++
		fmt.Println(&t,"里面", t)
		fmt.Println(&r,"外面", r)
	}()
	fmt.Println(&t,"外面", t)

	return t
}
func c() (r int,t int) {

	defer func(r int) {
		t++
		r = r + 5
		fmt.Println(&r)
	}(r)
	fmt.Println(&r)
	return 4,t
}
