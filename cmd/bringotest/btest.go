// +build cli

package main

import (
	"fmt"
	"github.com/darmiel/bringo/internal/bringotest"
)

func main() {
	var err error
	fmt.Println("bringo! Test CLI")

	var auth = bringotest.GetAuth()
	if auth == nil {
		panic("auth was nil")
	}

	// retrieving lists:
	fmt.Println("<< Loading lists ...")

	metas, err := auth.GetListMetas()
	if err != nil {
		panic(err)
	}

	for _, m := range metas {
		if m.Name == "Test" {
			if err := auth.SaveItemByMeta(m, "This is a Test", ""); err != nil {
				fmt.Println("WARN: Cannot save meta:")
				panic(err)
			}
		}
	}

	/*
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
	*/

}
