package main

import ("io"

"fmt")

var stIdCounter=0
type STManager struct{
    Global *SymbolTable
    reservedWords []string
    attributes []Attribute
    writer io.Writer

}

// Creates a new SymbolTable Manager.
// @filename: file path where the Table will be written to.
// Initializes the global SymbolTable,.
func CreateSTManager(writer io.Writer)STManager{
    return STManager{
        Global: createST("Global Table",nil),
        writer: writer,
		attributes: []Attribute{},
    }
}
// Writes ST to the file specified
// @st: symbol table to write to the file
func (m *STManager) Write(st *SymbolTable){
	st.Write(m.writer)
}

type SymbolTable struct{
    id int
    name string
    table map[string]SymbolEntry
    inner *SymbolTable
    parent *SymbolTable
}

func createST(name string,parent *SymbolTable) *SymbolTable{
    stIdCounter++
    return &SymbolTable{
        id: stIdCounter,
        name: name,
        table: map[string]SymbolEntry{},
        inner: nil,
        parent: parent,
    }

}

// Creates a new scope for the current SymbolTable.
// The function creates a new SymbolTable and sets its to the inner field of the parent.
func (s* SymbolTable) NewScope(name string)*SymbolTable{
    s.inner=createST(name,s)
    return s.inner
}



func (s* SymbolTable) AddSymbol(lex string){
	_,err:=s.table[lex]
	if !err{
		return
	}
	s.table[lex]=NewSymbolEntry(lex)
}

func (s* SymbolTable) RemoveSybol(){

}


// Writes the SymbolTable in the Specified format for PDL.
// @input: io.Writer 
func (s *SymbolTable) Write(w io.Writer){
    fmt.Fprintf(w,"%v #%d:\n\r",s.name,s.id)
	for _,i:=range s.table{
		i.Write(w)
	}
}
