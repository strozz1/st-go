package st

import (
	"fmt"
	"io"
)

type SymbolTable struct {
	id     int
	name   string
	table  map[string]*Entry
}


// Interal function to create a new SymbolTable.
//Initializes the table and assings an id
func createST(name string) *SymbolTable {
	stIdCounter++
	return &SymbolTable{
		id:     stIdCounter,
		name:   name,
		table:  map[string]*Entry{},
	}

}

func (s *SymbolTable) GetEntry(name string) (*Entry,bool){
	a,ok:=s.table[name]
	return a,ok
}

// Adds a new Symbol/Entry to the table. If it already exists returns a nil.
func (s *SymbolTable) AddEntry(lex string) *Entry {
	_, err := s.table[lex]
	if err {
		if debug {
			fmt.Printf("DEBUG: Failed to insert already existing Symbol '%v' on table [%v]\n\r", lex, s.name)
		}
		return nil
	}
	s.table[lex] = NewEntry(lex)
	a, _ := s.table[lex]
	if debug{
		fmt.Printf("DEBUG: Added new entry '%v' to table '%v'\n\r",lex,s.name)
	}
	return a
}

func (s *SymbolTable) RemoveEntry(lex string) {
	delete(s.table, lex)
	if debug{
		fmt.Printf("DEBUG: Removed entry '%v' from table '%v'\n\r",lex,s.name)
	}
}

// Writes the SymbolTable in the Specified format for PDL.
// @input: io.Writer
func (s *SymbolTable) Write(w io.Writer) {
	if debug {
		fmt.Printf("DEBUG: Writing table '%v' to output\n\r", s.name)
	}
	fmt.Fprintf(w, "%v #%d:\n\r", s.name, s.id)
	for _, i := range s.table {
		i.Write(w)
	}
}

