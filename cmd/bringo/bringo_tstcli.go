package main

import (
	"encoding/json"
	"fmt"
	"github.com/darmiel/bringo/pkg/bringo"
	"log"
	"os"
)

type creds struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func main() {
	fmt.Println("bringo! Test CLI")

	var (
		err  error
		data []byte
		auth *bringo.AuthBringo
	)

	// from cache?
	if data, err = os.ReadFile("cache.json"); err == nil {
		fmt.Println("[d] loading from cache ...")
		if err = json.Unmarshal(data, &auth); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("[d] cannot read from cache:", err)
		if data, err = os.ReadFile("creds.json"); err != nil {
			panic(err)
		}
		var c *creds
		if err = json.Unmarshal(data, &c); err != nil {
			panic(err)
		}
		fmt.Println("  >> Logging in with email:", c.Email, "...")
		if auth, err = bringo.NewWithLogin(c.Email, c.Password); err != nil {
			panic(err)
		}
		if data, err = json.Marshal(auth); err == nil {
			fmt.Println("[d] saving cache ...", string(data))
			if err = os.WriteFile("cache.json", data, 777); err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("  >> Auth obj (%v): %+v\n", err, auth)

	// retrieving lists:
	fmt.Println("<< Loading lists ...")
	lists, err := auth.GetLists()
	if err != nil {
		log.Fatalln(err)
		return
	}
	for _, l := range lists {
		fmt.Printf(">> %s\n", l.Name)
		fmt.Println("  Purchase:")
		for _, i := range l.Purchase {
			fmt.Println("    (", i.Name, "::", i.Specification, ")")
		}
		fmt.Println("  Recently:")
		for _, r := range l.Recently {
			fmt.Println("    (", r.Name, "::", r.Specification, ")")
		}
	}

}
