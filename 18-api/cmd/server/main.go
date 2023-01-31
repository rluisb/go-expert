package main

import "github.com/rluisb/curso-go/18-api/configs"

func main()  {
    config, err := configs.LoadConfig("./")
    if err != nil {
        panic(err)
    }
    print(config)
}