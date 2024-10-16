package main

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

func main() {
	secret, err := keyring.Get("myservice", "myuser")
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			panic(err)
		}
		err = keyring.Set("myservice", "myuser", "very-secret")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(secret)
}
