package main

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/imroc/req"
)

func main() {
	callGoogle()
	//callHashNode()
}

func callGoogle() {
	cmd := "call_google"
	hystrix.ConfigureCommand(cmd, hystrix.CommandConfig{
		Timeout:               1000, // 500ms
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	output := make(chan bool, 1)
	errors := hystrix.Go(cmd, func() error {
		r := req.New()
		resp, err := r.Get("https://google.com")
		if err != nil {
			output <- false
			return err
		}
		fmt.Printf("get resp: %+v\n", resp)
		output <- true
		return nil
	}, nil)

	select {
	case out := <-output:
		fmt.Println("call success: ", out)
	case err := <-errors:
		fmt.Printf("call %s error: %v\n", cmd, err) // call call_google error: hystrix: timeout
	}
}

func callHashNode() {
	cmd := "call_hashnode"
	hystrix.ConfigureCommand(cmd, hystrix.CommandConfig{
		Timeout:               500, // 500ms
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	output := make(chan bool, 1)
	errors := hystrix.Go(cmd, func() error {
		r := req.New()
		resp, err := r.Get("https://codeforfun.hashnode.dev/hystrix-in-go")
		if err != nil {
			output <- false
			return err
		}
		fmt.Printf("get resp: %+v\n", resp)
		output <- true
		return nil
	}, nil)

	select {
	case out := <-output:
		fmt.Println("call success: ", out)
	case err := <-errors:
		fmt.Printf("call %s error: %v\n", cmd, err) // call call_google error: hystrix: timeout
	}
}
