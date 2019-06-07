package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/rickypai/stardog/protobufs/datadog"
	"github.com/stripe/skycfg"
)

func main() {
	ctx := context.Background()
	config, err := skycfg.Load(ctx, "hello.sky")
	if err != nil {
		panic(err)
	}
	messages, err := config.Main(ctx)
	if err != nil {
		panic(err)
	}

	j, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(j))
}
