package main

import (
	"bufio"
	"os"
)

func main(){

    mng:=CreateSTManager(bufio.NewWriter(os.Stdout))
    mng.Write(mng.Global)
}
