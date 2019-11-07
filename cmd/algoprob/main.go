package main

import (
	"bufio"
	"os"

	"fmt"

	"errors"

	"strconv"

	"github.com/tangit86/al-go-prob/pkg/riddles"
)

var cmds map[string]runner

type runner func(input *bufio.Reader, output *bufio.Writer)

func init() {
	cmds = make(map[string]runner, 10)
	cmds["candies"] = candiesRunner
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	stdout := os.Stdout
	defer stdout.Close()

	writer := bufio.NewWriter(stdout)

	cmd, err := handleArgs(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	runner := cmds[cmd]

	runner(reader, writer)
	writer.Flush()
}

func candiesRunner(input *bufio.Reader, output *bufio.Writer) {
	i := 0
	var n int = 0

	var arr []int

	for {
		line, _, err := input.ReadLine()
		if err != nil {
			break
		}

		lineNum, _ := strconv.Atoi(string(line))

		if i == 0 {
			n = lineNum
			arr = make([]int, n)

		} else {
			arr[i-1] = lineNum
		}
		i++

		if i > n {
			break
		}
	}

	res := riddles.Candies(n, arr)

	_, err := output.WriteString(fmt.Sprint(res))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleArgs(args []string) (string, error) {

	if len(args) < 2 || cmds[args[1]] == nil {
		return "", errors.New("No command found.\n Use: algoprob <cmd>")
	}

	return args[1], nil
}
