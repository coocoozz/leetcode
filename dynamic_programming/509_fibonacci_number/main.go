package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("input n: ")

	r := bufio.NewReader(os.Stdin)
	nStr, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("fail to read string from keyboard. err:%v\n", err)
		return
	}

	n, err := strconv.Atoi(strings.TrimSpace(nStr))
	if err != nil {
		fmt.Printf("fail to convert input to number. err:%v\n", err)
		return
	}

	now := time.Now()
	res1 := fib(n)
	res1Elapsed := time.Now().Sub(now)
	fmt.Printf("n: %d, res1: %d, elapsed: %v\n", n, res1, res1Elapsed)

	now = time.Now()
	res2 := fibWithMem(n)
	res2Elapsed := time.Now().Sub(now)
	fmt.Printf("n: %d, res2: %d, elapsed: %v\n", n, res2, res2Elapsed)

	now = time.Now()
	res3 := fibWithBT(n)
	res3Elapsed := time.Now().Sub(now)
	fmt.Printf("n: %d, res3: %d, elapsed: %v\n", n, res3, res3Elapsed)
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

var mem map[int]int = make(map[int]int)

func fibWithMem(n int) int {
	if val, exist := mem[n]; exist {
		return val
	} else if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	mem[n] = fib(n-1) + fib(n-2)
	return mem[n]
}

func fibWithBT(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	res[2] = 1

	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}
