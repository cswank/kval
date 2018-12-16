package main

import (
	"fmt"
)

func CliCreate(args []string) error {
	fmt.Printf("Create database %v\n", args)

	if len(args) != 2 {
		return fmt.Errorf("Syntax: create|c <database>")
	}
	return nil
}

func CliRemove(args []string) error {
	fmt.Printf("Remove database %v\n", args)

	if len(args) != 2 {
		return fmt.Errorf("Syntax: remove|r <database>")
	}
	return nil
}

func CliSet(args []string) error {
	fmt.Printf("Set key-value pair %v\n", args)

	if len(args) != 3 {
		return fmt.Errorf("Syntax: set|s <key> <value>")
	}
	return nil
}

func CliGet(args []string) error {
	fmt.Printf("Get value for key %v\n", args)

	if len(args) != 2 {
		return fmt.Errorf("Syntax: get|g <key>")
	}
	return nil
}

func CliDel(args []string) error {
	fmt.Printf("Delete key-value pair %v\n", args)

	if len(args) != 2 {
		return fmt.Errorf("Syntax: del|d <key>")
	}
	return nil
}

func CliHelp(args []string) {
	fmt.Printf("Unknown command %v (add something helpful here)\n", args)
}
