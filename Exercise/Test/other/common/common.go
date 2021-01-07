package common

import "fmt"

func NewLoginer() Loginer{
	return defaultLogin(0)
}
type Loginer interface {
	Login()
}
type defaultLogin int
func (d defaultLogin) Login(){
	fmt.Println("login in...")
}
