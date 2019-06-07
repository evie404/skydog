package main

import (
	"context"
	"fmt"

	_ "github.com/golang/protobuf/ptypes/wrappers"
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
	for _, msg := range messages {
		fmt.Printf("%s\n", msg.String())
	}
}
