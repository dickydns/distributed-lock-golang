package main

import (
	"fmt"
	"time"
	"distributed-lock-golang/helper"
)

func main() {

	//load env
	helper.LoadEnv()

	//load redis helper
	redis := helper.NewRedisHelper()
	if err := redis.Ping(); err != nil {
		panic(err)
	}

	lock := helper.NewDistributedLock(redis)
	event:= "create:voucer"
	for i := 1; i <= 3; i++ {
		result, err := lock.Execute(
			event,
			30*time.Second,
			func() (interface{}, error) {
				fmt.Println("Generating Event...", i)
				time.Sleep(3* time.Second)
				return "",nil
			},
		)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
	fmt.Println("Application Finished")
}