package form

type State int

const (
	Initial State = iota
	Success

	// Error is *not* a validation error.
	Error
)

type Field struct {
	Val string
	Err string
}

// Val returns a form.Field with just a value set
func Val(val string) Field {
	return Field{
		Val: val,
	}
}
