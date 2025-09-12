package main

import (
	"fmt"
	"io"
)

type AttributeType int

const (
	T_INTEGER AttributeType = iota
	T_STRING
	T_ARRAY
)

// attributes
const (
	A_TIPO        = "Tipo"
	A_DESPL       = "Despl"
	A_NUMPARAM    = "numParam"
	A_TIPOPARAMXX = "TipoParam"
	A_MODOPARAMXX = "ModoParam"
	A_TIPORETORNO = "TipoRetorno"
	A_ETIQFUNC    = "EtiqFuncion"
	A_PARAM       = "Param"
)

type Attribute struct {
	Name      string        // att name
	Type      AttributeType //type of attribute
	Desc      string        //Description of attribute
	StringVal string        //String value if Type is String
	IntVal    int           //Int value if Type is Int
	ArrayVal  []string      //Array value if Type is Array
	HasValue  bool          //Flag if value has been asigned
}

// Writes the Attribute from an Entry to the specified Writer with
// PDL specified format
func (a *Attribute) Write(w io.Writer) {
	fmt.Fprintf(w, "\t+ %v: ", a.Desc)
	if a.Type == 1 { //TODO
		fmt.Fprintf(w, "'%v'", a.StringVal)
	} else {
		fmt.Fprintf(w, "'%d'", a.IntVal)
	}
	fmt.Fprintln(w)
}

// Loads the default attributes to the STManager
func LoadDefAttributes(st *STManager) {
	st.attributes = append(st.attributes, Attribute{
		Desc: A_TIPO,
		Type: T_STRING,
	})
	st.attributes = append(st.attributes, Attribute{
		Desc: A_DESPL,
		Type: T_INTEGER,
	})
	st.attributes = append(st.attributes, Attribute{
		Desc: A_TIPORETORNO,
		Type: T_STRING,
	})
	st.attributes = append(st.attributes, Attribute{
		Desc: A_ETIQFUNC,
		Type: T_STRING,
	})

}
