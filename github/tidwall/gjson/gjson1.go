package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var json2 = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}`

func main() {
	if gjson.Valid(json2) {
		fmt.Println("valid success")
	}
	lastname := gjson.Parse(json2).Get("name").Get("last")
	lastname1 := gjson.Get(json2, "name").Get("last")
	lastname2 := gjson.Get(json2, "name.last")
	fmt.Println("lastname: ", lastname)
	fmt.Println("lastname1: ", lastname1)
	fmt.Println("lastname2: ", lastname2)
	value := gjson.Get(json2, "name.last")
	value1 := gjson.Get(json2, "age")
	value2 := gjson.Get(json2, "children")
	value3 := gjson.Get(json2, "children.#")
	value4 := gjson.Get(json2, "children.1")
	value5 := gjson.Get(json2, "child*.2")
	value6 := gjson.Get(json2, "c?ildren.0")
	value7 := gjson.Get(json2, "fav\\.movie")
	value8 := gjson.Get(json2, "friends.#.first")
	value9 := gjson.Get(json2, "friends.1.last")
	fmt.Println("name.last: ", value)
	fmt.Println("age: ", value1)
	fmt.Println("children: ", value2)
	fmt.Println("children.#: ", value3)
	fmt.Println("children.1: ", value4)
	fmt.Println("child*.2: ", value5)
	fmt.Println("c?ildren.0: ", value6)
	fmt.Println("fav\\.movie: ", value7)
	fmt.Println("friends.#.first: ", value8)
	fmt.Println("friends.1.last: ", value9)
	value10 := gjson.Get(json2, `friends.#[last=="Murphy"].first`)
	value11 := gjson.Get(json2, `friends.#[last=="Murphy"]#.first`)
	value12 := gjson.Get(json2, `friends.#[age>45]#.last`)
	value13 := gjson.Get(json2, `friends.#[first%"D*"].last`)
	fmt.Println(`friends.#[last=="Murphy"].first: `, value10)
	fmt.Println(`friends.#[last=="Murphy"]#.first: `, value11)
	fmt.Println(`friends.#[age>45]#.last: `, value12)
	fmt.Println(`friends.#[first%"D*"].last: `, value13)
}
