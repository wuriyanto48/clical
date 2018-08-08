package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/wuriyanto48/clical/cmd"
)

func main() {

	flag.Usage = func() {
		fmt.Println("    add", "command for add operation with two arguments x and y (add x y)")
		fmt.Println("    sub", "command for sub operation with two arguments x and y (sub x y)")
		fmt.Println("    mul", "command for mul operation with two arguments x and y (mul x y)")
		fmt.Println("    div", "command for div operation with two arguments x and y (div x y)")
		fmt.Println("    -h or --help for show help")
	}

	flag.Parse()

	command := cmd.NewCommander()

	/*
		using arguments
	*/
	// if len(os.Args) <= 1 {
	// 	flag.Usage()
	// 	os.Exit(0)
	// }

	// args := os.Args[1:]

	// arg := strings.Join(args, " ")

	// result := command.Exec(arg)

	/*
		using bufio scanner
	*/

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input your command..")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}

		result := command.Exec(line)
		fmt.Println(result)
		fmt.Println("")
	}

}
