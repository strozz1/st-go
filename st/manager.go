package st

import (
	"fmt"
	"io"
)

var debug bool

// ID counter for the tables
var stIdCounter = 0

type STManager struct {
	Global        *SymbolTable
	Local         *SymbolTable
	ReservedWords []string
	//Defines the set of available attributes. This attributes are only the
	//template, so when you want to use it for an Entry you need to take it and modify it with the
	//corresponding values
	Attributes map[string]Attribute
	Writer     io.Writer
}

// Creates a new SymbolTable Manager.
// @filename: file path where the Table will be written to.
// Initializes the global SymbolTable.
func CreateSTManager(writer io.Writer, debugger bool) STManager {
	debug = debugger
	if debug {
		fmt.Printf("DEBUG: Initializing STManager\n\r")
	}

	return STManager{
		Global:     createST("Global Table"),
		Writer:     writer,
		Attributes: map[string]Attribute{},
	}
}

// Creates an attribute with the specified name, type & description.
// 'AttributeType' is an int & 'AttributeDesc' is a string
func (m *STManager) CreateAttribute(name string, d string, t AttributeType) {
	if m.containsAttribute(name) {
		if debug {
			fmt.Printf("DEBUG: Can't create attribute, it already exists [%v]\n\r", name)
			return
		}
	}
	m.Attributes[name] = NewAttribute(name, t, d)
	if debug {
		fmt.Printf("DEBUG: Created new attribute '%v' of type: '%v' & desc: '%v'\n\r", name, t, d)
	}
}

func (m *STManager) AddGlobalEntry(name string) *Entry {
	return m.Global.AddEntry(name)
}
func (m *STManager) AddLocalEntry(name string) *Entry {
	if m.Local == nil {
		if debug {
			fmt.Println("DEBUG: Can't add entry to a 'nil' local table.")
			return nil
		}
	}
	return m.Local.AddEntry(name)
}



func (m *STManager) GetGlobalEntry(name string) (*Entry, bool) {
	v, ok := m.Global.GetEntry(name)
	if !ok {
		if debug {
			fmt.Printf("DEBUG: Global entry not found '%v'\n\r", name)
			return nil, ok
		}
	}
	return v, ok
}
func (m *STManager) RemoveGlobalEntry(name string) {
	m.Global.RemoveEntry(name)
}

func (m *STManager) RemoveLocalEntry(name string) {
	m.Global.RemoveEntry(name)
}

func (m *STManager) SetEntryAttribute(e *Entry, name string, val any) {
	if !m.containsAttribute(name) {
		if debug {
			fmt.Printf("DEBUG: Can't add attribute '%v'. It does not exist\n\r", name)
			return
		}
	}
	if !e.containsAttr(name) {
		v, _ := m.Attributes[name]
		e.AddAtribute(name, v)
		if debug {
			fmt.Printf("DEBUG: Added attribute '%v' to entry '%v'\n\r", name, e.Lexeme)
		}
	}

	e.SetAttributeValue(name, val)
	if debug {
		fmt.Printf("DEBUG: Setted attribute value: '%v=%v' to entry\n\r", name, val)
	}
}

// Create a New local table with a specified name. If there was an existing local table
// it will be removed
func (m *STManager) CreateLocalTable(name string) {
	m.Local = createST(name)
	if debug {
		fmt.Printf("DEBUG: Local table created: [%v]\n\r", name)
	}
}

// Returns if the name already exists in the attribute list
func (m *STManager) containsAttribute(name string) bool {
	_, b := m.Attributes[name]
	return b
}

// Writes ST to the file specified
// @st: symbol table to write to the file
func (m *STManager) Write(st *SymbolTable) {
	st.Write(m.Writer)
}
