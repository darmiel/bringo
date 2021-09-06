// +build cli

package main

import (
	"flag"
	"fmt"
	"github.com/darmiel/bringo/internal/bringotest"
	"strings"
	"time"
)

var (
	DisplayedLists string
	Delay          time.Duration
)

func init() {
	flag.StringVar(&DisplayedLists, "l", "", "Lists to display")
	flag.DurationVar(&Delay, "d", 10*time.Second, "Refresh delay")
	flag.Parse()
}

func main() {
	// check displayed lists not empty
	if DisplayedLists == "" {
		fmt.Println("Displayed lists cannot be empty.")
		return
	}

	lists := strings.Split(DisplayedLists, ",")
	for i, v := range lists {
		lists[i] = strings.TrimSpace(v)
	}

	fmt.Println("Logging in ...")
	var auth = bringotest.GetAuth()

	// get list uuids
	fmt.Println("Finding corresponding UUIDs ...")
	metas, err := auth.GetListMetas()
	if err != nil {
		fmt.Println("Cannot get metas:")
		panic(err)
	}

	var uuids = make(map[string]string)
	for _, m := range metas {
		// check if lists contains m
		for _, l := range lists {
			if m.Name == l {
				uuids[m.UUID] = l
				break
			}
		}
	}

	for {
		// collect all items and after fetching each list, display output
		// to prevent "scoll-flickering"
		var builder strings.Builder

		// request list items
	ml:
		for u, l := range uuids {
			list, err := auth.GetList(u)
			if err != nil {
				continue ml
			}

			// Header
			// [ Home ]
			builder.WriteString("[ ")
			builder.WriteString(l)
			builder.WriteString(" ]")
			builder.WriteRune('\n')

			for j, p := range list.Purchase {
				// Tree index
				if j != len(list.Purchase)-1 {
					builder.WriteString("   ├ ")
				} else {
					builder.WriteString("   └ ")
				}

				// Item Name
				builder.WriteString(p.Name)

				if p.Specification != "" {
					builder.WriteString(" (")
					builder.WriteString(p.Specification)
					builder.WriteString(")")
				}
				builder.WriteRune('\n')
			}
			builder.WriteRune('\n')
		}

		// "clear" screen
		fmt.Println(strings.Repeat("\n", 30))

		// print output
		fmt.Print(builder.String())

		fmt.Println("-- refreshing in", Delay)
		// refresh delay
		time.Sleep(Delay)
	}
}
