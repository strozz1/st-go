package main

import ("io"

"fmt")

var stIdCounter=0
type STManager struct{
    Global *SymbolTable
    reservedWords string
    atributtes string
    writer io.Writer

}

// Creates a new SymbolTable Manager.
// @filename: file path where the Table will be written to.
// Initializes the global SymbolTable,.
func CreateSTManager(writer io.Writer)STManager{
    return STManager{
        Global: createST("Global Table",nil),
        writer: writer,
    }
}
// Writes ST to the file specified
// @st: symbol table to write to the file
func (m *STManager) Write(st *SymbolTable){
    fmt.Println("a")
    fmt.Fprintf(m.writer,"%v #%d:\n\r",st.name,st.id)
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
func (st* SymbolTable) NewScope(name string)*SymbolTable{
    st.inner=createST(name,st)
    return st.inner
}

