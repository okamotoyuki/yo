package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("#include \"textflag.h\"")
	fmt.Println()
	fmt.Println("TEXT ·run(SB), NOSPLIT, $0")
	fmt.Println("\tMOVQ\t$" + os.Args[1] + ", ret+0(FP)")
	fmt.Println("\tRET")
}
