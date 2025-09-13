package main

import (
	"go-st/st"
	"os"
)

func main(){
    mng:=st.CreateSTManager(os.Stdout,true)
	mng.CreateAttribute("entero",st.T_INTEGER,"entero")
	var1:=mng.AddGlobalEntry("var1")
	var1.SetType(st.LEXEME_INTEGER)
	mng.AddEntryAttribute(var1,"despl")
	mng.SetEntryAttrValue(var1,"despl",23)
	
    mng.Write(mng.Global)
}
