package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	_, err = c.Do("SET", "username", "nick")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "username"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got username %v \n", username)
	}
}
