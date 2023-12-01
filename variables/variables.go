package variables

type Variable struct {
	Name  string
	Value string
}

var Variables = make(map[string]Variable)
