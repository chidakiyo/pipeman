package main

import (
	"os"
	"fmt"
	"os/exec"
	"time"
)

func main() {

	argSize := len(os.Args)

	if argSize < 1 {
		fmt.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}

	for i :=  1 ; i <= argSize -1 ; i++ {
		c := os.Args[i]
		fmt.Printf("== BEGIN == [%s]\n", c)
		sTime := time.Now().Unix()
		o, err := exec.Command("sh", "-c", c).Output()
		if err != nil {
			fmt.Println("err : " + err.Error())
			os.Exit(1)
		} else {
			fmt.Println(string(o))
		}
		eTime := time.Now().Unix()
		fmt.Printf("== END == [%s] %d sec. \n", c, eTime - sTime)
	}

}
