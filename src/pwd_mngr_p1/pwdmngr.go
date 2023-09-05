package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const pwd_db = "password.db"

func main() {
	var args []string
	args = os.Args
	// if 'add is provided as the first param then call the store function
	if args[1] == "add" {
		store(args[2], args[3], args[4])
		// if 'get' is provided first then call retrieve function
	} else if args[1] == "get" {
		retrieve(args[2])
	} else {
		fmt.Println("Invalid operation", args[1])
	}
}

func store(platform string, username string, password string) {
	// ^ params the function is taking in along with their type
	entry := platform + "," + username + "," + password + "\n"
	f, err := os.OpenFile(pwd_db, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// OpenFile returns 1.filehandle + 2.error (if any)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(entry)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func retrieve(platform string) {
	f, err := os.Open(pwd_db)
	if err != nil {
		fmt.Println(err)
		return
	}
	// input is defined and initlialized for reading the file
	// bufio is an inbuilt package that does buddered IO operations => get 'scanner' from that which can read input and break into lines that are being used in the code
	input := bufio.NewScanner(f)
	for input.Scan() {
		entry := strings.Split(input.Text(), ",")
		if entry[0] == platform {
			fmt.Println(entry[1], entry[2])
			return
		}
	}
	fmt.Printf("Platform %s not known\n", platform)
}

// Terminal Commands to build and test
// ❯ go build pwdmngr.go
// ❯ ./pwdmngr add twitter realjohn '!facebook'
// 27 bytes written
// ❯ ./pwdmngr add facebook john '!twitter'
// 23 bytes written
// ❯ ./pwdmngr get facebook
// john !twitter
// ❯ ./pwdmngr get twitter
// realjohn !facebook
// ❯ ./pwdmngr get google
// Platform google not known
// ❯ ./pwdmngr put facebook
// Invalid operation  put
