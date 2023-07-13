package inspect

type Type byte

const (
	Unknown     Type = iota // Unknown or unidentified column type
	Text                    // String column type of any size
	Integer                 // Integer column type
	Decimal                 // Floating point column type
	Binary                  // Binary column type
	Boolean                 // Boolean column type
	Enumeration             // Enumeration column type
	Set                     // Set column type
	TimeStamp               // Timestamp column type
	DateTime                // Datetime column type
	Date                    // Date column type (without time)
	Time                    // Time column type (without date)
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
	case TimeStamp:
		return "timestamp"
	case DateTime:
		return "datetime"
	case Date:
		return "date"
	case Time:
		return "time"
	default:
		return "unknown"
	}
}
