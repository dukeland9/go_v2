package main

import (
	"fmt"
	//"time"
)

import (
	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:33331",
	})
	fmt.Printf("%v\n", client.Ping())
	//fmt.Printf("%#v\n", client.Set("haha", "great!", 0))
	fmt.Printf("%v\n", client.Get("haha"))
	fmt.Printf("%#v\n", client.SetNX("haha", "hoho", 0))
	fmt.Printf("%#v\n", client.SetNX("hoho", "Yes", 0))
	//fmt.Printf("%#v\n", client.Get("haha"))
	//fmt.Printf("%#v\n", "Sleep for 10 seconds...")
	//time.Sleep(1e10)
	//fmt.Printf("%#v\n", client.Get("haha"))
	//fmt.Printf("%#v\n", client.GetSet("haha", "hoho"))
	//fmt.Printf("%#v\n", client.Get("haha"))
}
