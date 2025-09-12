package main

import (
	"os"
)

func main(){
    mng:=CreateSTManager(os.Stdout)
	
    mng.Write(mng.Global)
}
