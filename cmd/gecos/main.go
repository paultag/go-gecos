package main

import (
	"fmt"
	"os/user"
	"pault.ag/go/gecos"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	entry, err := gecos.Lookup(current)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uid:      %s\n", current.Uid)
	fmt.Printf("Gid:      %s\n", current.Gid)
	fmt.Printf("Username: %s\n", current.Username)
	fmt.Printf("Name:     %s\n", current.Name)
	fmt.Printf("Home:     %s\n", current.HomeDir)
	fmt.Printf("Room:     %s\n", entry.Room)
	fmt.Printf("Office:   %s\n", entry.OfficePhone)
	fmt.Printf("Home:     %s\n", entry.HomePhone)
	fmt.Printf("Other:    %s\n", entry.Other)
}
