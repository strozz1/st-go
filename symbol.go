package main

import (
	"fmt"
	"io"
)

type LexemeType string

const (
	LEXEME_FUNCTION  = "Funcion"
	LEXEME_PROCEDURE = "Procedimiento"
	LEXEME_INTEGER   = "Entero"
	LEXEME_STRING    = "Cadena"
	LEXEME_REAL      = "Real"
	LEXEME_LOGIC     = "Logico"
	LEXEME_POINTER   = "Puntero"
	LEXEME_VECTOR    = "Vector"
	LEXEME_NONE ="none"
)

type Entry struct {
	Lexeme     string      //lexeme
	Type       LexemeType  //lexeme type 
	//DUDA: el tupo no seria un atributo?
	Id         int         //id of Entry
	Attributes []Attribute //attribute list
}

// Creates new symbol entry from the original lexeme.
// Returns the new Entry object with attributes init.
func NewEntry(lex string) Entry {
	return Entry{
		Attributes: []Attribute{},
		Lexeme: lex,
	}

}

// sets the type of the lexeme
// IF an invalid lexem type is provided, an error is returned
func (e *Entry) SetType(t LexemeType) error {
	switch t {
	case LEXEME_FUNCTION:
		break
	case LEXEME_PROCEDURE:
		break
	case LEXEME_INTEGER:
		break
	case LEXEME_STRING:
		break
	case LEXEME_REAL:
		break
	case LEXEME_LOGIC:
		break
	case LEXEME_POINTER:
		break
	case LEXEME_VECTOR:
		break
	default:
		return fmt.Errorf("Error: Invalid Lexem type:[%v]", t)
	}
	e.Type = t
	return nil
}

// Writes the SymbolEntry from the ST to the specified Writer with
// PDL specified format
func (s *Entry) Write(w io.Writer) {
	fmt.Fprintf(w, "* LEXEMA: '%v'\r\n", s.Lexeme)
	fmt.Fprintln(w,"\tAtributos:")
	
}

// Adds an attribute to the Entry.
// Returns error if the attribute is already present
func (s *Entry) AddAtribute(name string,desc string,aType AttributeType)error{
	if contains(s.Attributes,name){
		return fmt.Errorf("Error: Attribute '%v' already exists",name)
	}
	s.Attributes=append(s.Attributes,Attribute{
		Type: aType,
		Desc: desc,
		Name:name,
	})
	return nil
}
func  contains(array []Attribute,val string)bool{
	for _,b:=range array{
		if val==b.Name{
			return true
		}
	}
	return false
}
