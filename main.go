package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

var kvaldir string

func main() {
	fmt.Println("kval (C) Frank W Miller")

	/* Get path for kval directory */
	usr, _ := user.Current()
	kvaldir = usr.HomeDir + "/.kval"
	fmt.Printf("kval directory = %v\n", kvaldir)

	/* Check if kval directory exists */
	_, err := os.Stat(kvaldir)
	if os.IsNotExist(err) {
		fmt.Printf("Create kval directory %v", kvaldir)
		err = os.Mkdir(kvaldir, 0777)
		if err != nil {
			log.Fatal(" failed\n")
		}
		fmt.Printf("\n")
	}
	/* Command line client loop */
	stdin := bufio.NewReader(os.Stdin)
	var s1 string
	for {
		fmt.Printf("kval> ")
		s1, _ = stdin.ReadString('\n')
		s2 := strings.TrimSpace(s1)

		collapsespace := regexp.MustCompile(`\s+`)
		s3 := collapsespace.ReplaceAllString(s2, " ")

		if len(s3) > 0 {
			s4 := strings.Split(s3, " ")
			var err error
			switch s4[0] {
			case "quit":
				os.Exit(0)
			case "create", "c":
				err = CliCreate(s4)
			case "remove", "r":
				err = CliRemove(s4)
			case "set", "s":
				err = CliSet(s4)
			case "get", "g":
				err = CliGet(s4)
			case "del", "d":
				err = CliDel(s4)
			default:
				CliHelp(s4)
			}

			if err != nil {
				CliHelp(s4)
			}
		}
	}
}
