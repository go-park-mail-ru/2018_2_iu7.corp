package main

import (
	"2018_2_iu7.corp/profile-service/services/rpc"
	"context"
	"fmt"
)

func main() {
	client, err := rpc.CreateClient()
	if err != nil {
		panic(err)
	}

	rsp, err := (*client).GetProfile(context.TODO(), &rpc.Credentials{Username: "admin", Password: "1234"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.GetID())
}
