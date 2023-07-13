package inspect

type Type byte

const (
	Unknown Type = iota
	Text
	Integer
	Decimal
	Binary
	Boolean
	Enumeration
	Set
)

func (t Type) String() string {
	switch t {
	case Text:
		return "text"
	case Integer:
		return "integer"
	case Decimal:
		return "decimal"
	case Binary:
		return "binary"
	case Boolean:
		return "boolean"
	case Enumeration:
		return "enum"
	case Set:
		return "set"
	default:
		return "unknown"
	}
}
