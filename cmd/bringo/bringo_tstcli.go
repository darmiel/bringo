package main

import (
	"encoding/json"
	"fmt"
	"github.com/darmiel/bringo/pkg/bringo"
	"os"
	"os/signal"
	"syscall"
)

type creds struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func main() {
	fmt.Println("bringo! Test CLI")

	var err error

	// parse creds
	var data []byte
	if data, err = os.ReadFile("creds.json"); err != nil {
		panic(err)
	}
	var c *creds
	if err = json.Unmarshal(data, &c); err != nil {
		panic(err)
	}

	fmt.Println("  >> Logging in with email:", c.Email, "...")

	auth, err := bringo.NewWithLogin(c.Email, c.Password)
	fmt.Printf("  >> Auth obj (%v): %+v\n", err, auth)
	fmt.Println("  >> expires:", auth.Dog.Expires)

	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	_ = <-sc

	fmt.Println("stopping ...")
	auth.Close()
	fmt.Println("should be closed now.")
}
