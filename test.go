package main

import (
	"go-st/st"
	"os"
)

func main() {
	mng := st.CreateSTManager(os.Stdout, false)
	mng.CreateAttribute("despl", st.DESC_DESPL, st.T_INTEGER)
	mng.CreateAttribute("numero de parametros", st.DESC_NUM_PARAM, st.T_INTEGER)
	mng.CreateAttribute("tipo de parametros", st.DESC_TIPO_PARAM, st.T_ARRAY)
	mng.CreateAttribute("modo de parametros", st.DESC_MODO_PARAM, st.T_ARRAY)
	mng.CreateAttribute("tipo de retorno", st.DESC_TIPO_RETORNO, st.T_STRING)
	mng.CreateAttribute("etiqueta", st.DESC_ETIQ_FUNCION,st.T_STRING)
	mng.CreateAttribute("parametro", st.DESC_PARAM,st.T_INTEGER)
	mng.CreateAttribute("dimension", "dimension",st.T_INTEGER)
	mng.CreateAttribute("elem", "elem",st.T_STRING)


	var1 := mng.AddGlobalEntry("x")
	var1.SetType(st.LEXEME_INTEGER)
	mng.SetEntryAttribute(var1, "despl", 0)
	mng.SetEntryAttribute(var1, "parametro", 0)
	
	var2 := mng.AddGlobalEntry("y")
	var2.SetType(st.LEXEME_INTEGER)
	mng.SetEntryAttribute(var2, "despl", 4)
	mng.SetEntryAttribute(var2, "parametro", 0)

	var3 := mng.AddGlobalEntry("res")
	var3.SetType(st.LEXEME_INTEGER)
	mng.SetEntryAttribute(var3,"despl",8)
	mng.SetEntryAttribute(var3,"parametro",0)
	
	func1:=mng.AddGlobalEntry("suma")
	func1.SetType(st.LEXEME_FUNCTION)
	mng.SetEntryAttribute(func1,"numero de parametros",2)
	tipos:=[]string{"int","int"}
	mng.SetEntryAttribute(func1,"tipo de parametros",tipos)
	modos:=[]string{"0","0"}
	mng.SetEntryAttribute(func1,"modo de parametros",modos)
	mng.SetEntryAttribute(func1,"tipo de retorno","int")
	mng.SetEntryAttribute(func1,"etiqueta","sumatorio")

	mng.CreateLocalTable("suma")

	local1:=mng.AddLocalEntry("a")
	local1.SetType(st.LEXEME_INTEGER)
	mng.SetEntryAttribute(local1,"despl",0)
	mng.SetEntryAttribute(local1,"parametro",1)

	local2:=mng.AddLocalEntry("b")
	local2.SetType(st.LEXEME_INTEGER)

	mng.SetEntryAttribute(local2,"despl",4)
	mng.SetEntryAttribute(local2,"parametro",1)

	local3:=mng.AddLocalEntry("res")
	local3.SetType(st.LEXEME_INTEGER)

	mng.SetEntryAttribute(local3,"despl",8)
	mng.SetEntryAttribute(local3,"parametro",0)


	mng.Write(mng.Local)
	mng.Write(mng.Global)
}
