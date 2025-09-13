package st

import (
	"fmt"
	"io"
)

// Defines the types that an attribute can be
type AttributeType int

const (
	T_INTEGER AttributeType = iota
	T_STRING
	T_ARRAY
	T_NONE
)

// PREDEFINED ATTRIBUTES
const (
	DESC_DESPL        = "Despl"       // relative offset
	DESC_NUM_PARAM    = "numParam"    //num of params
	DESC_TIPO_PARAM   = "TipoParam"   // type of params
	DESC_MODO_PARAM   = "ModoParam"   // param mode
	DESC_TIPO_RETORNO = "TipoRetorno" // return type
	DESC_ETIQ_FUNCION = "EtiqFuncion" // function label
	DESC_PARAM        = "Param"       // param
	DESC_OTRO         = "Otro"        // type of identifier
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

// Creates new attribute
func NewAttribute(name string, tp AttributeType, ad string) Attribute {

	return Attribute{
		Name:     name,
		Type:     tp,
		Desc:     ad,
		HasValue: false,
	}
}

// Writes the Attribute from an Entry to the specified Writer with
// PDL specified format
func (a *Attribute) Write(w io.Writer) {
	switch a.Desc {
	case DESC_DESPL:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_DESPL, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_DESPL, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_DESPL)
				}
			case T_ARRAY:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_DESPL, a.ArrayVal[0])
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_DESPL)
				}
			}
		}
	case DESC_NUM_PARAM:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_NUM_PARAM, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_NUM_PARAM, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_NUM_PARAM)
				}
			case T_ARRAY:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_NUM_PARAM, a.ArrayVal[0])
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_NUM_PARAM)
				}
			}
		}
	case DESC_TIPO_PARAM:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_TIPO_PARAM, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_TIPO_PARAM, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_TIPO_PARAM)
				}
			case T_ARRAY:
				if a.HasValue {
					for i := range a.ArrayVal {
						fmt.Fprintf(w, "    +%v%v: '%v'", DESC_TIPO_PARAM, i, a.ArrayVal)
					}
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_TIPO_PARAM)
				}
			}
		}
	case DESC_MODO_PARAM:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_MODO_PARAM, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_MODO_PARAM, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_MODO_PARAM)
				}
			case T_ARRAY:
				if a.HasValue {
					for i := range a.ArrayVal {
						fmt.Fprintf(w, "    +%v%v: '%v'", DESC_MODO_PARAM, i, a.ArrayVal)
					}
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_MODO_PARAM)
				}
			}

		}
	case DESC_TIPO_RETORNO:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_TIPO_RETORNO, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_TIPO_RETORNO, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_TIPO_RETORNO)
				}
			case T_ARRAY:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_TIPO_RETORNO, a.ArrayVal[0])
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_TIPO_RETORNO)
				}
			}

		}
	case DESC_ETIQ_FUNCION:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_ETIQ_FUNCION, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_ETIQ_FUNCION, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_ETIQ_FUNCION)
				}
			case T_ARRAY:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_ETIQ_FUNCION, a.ArrayVal[0])
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_ETIQ_FUNCION)
				}
			}
		}
	case DESC_PARAM:
		{
			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", DESC_PARAM, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_PARAM, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_PARAM)
				}
			case T_ARRAY:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", DESC_PARAM, a.ArrayVal[0])
				} else {
					fmt.Fprintf(w, "    +%v: '-'", DESC_PARAM)
				}
			}
		}
	default:
		{

			switch a.Type {
			case T_INTEGER:
				fmt.Fprintf(w, "    +%v: %v", a.Desc, a.IntVal)
			case T_STRING:
				if a.HasValue {
					fmt.Fprintf(w, "    +%v: '%v'", a.Desc, a.StringVal)
				} else {
					fmt.Fprintf(w, "    +%v: '-'", a.Desc)
				}
			case T_ARRAY:
				if a.HasValue {
					for i := range a.ArrayVal {
						fmt.Fprintf(w, "    +%v%v: '%v'", a.Desc, i, a.ArrayVal)
					}
				} else {
					fmt.Fprintf(w, "    +%v: '-'", a.Desc)
				}
			}
		}
	}
	fmt.Fprintln(w)
}
