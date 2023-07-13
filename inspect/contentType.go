package inspect

import (
	"strconv"
	"strings"
)

// ContentType defines column content
type ContentType int

const (
	RawContent ContentType = iota
	IdentifierContent
	RelationIdentifierContent
	UnixSecondsContent
	BooleanContent
	SecretContent
	NameContent
	TitleContent
	EmailContent
	URIContent
)

func (c ContentType) String() string {
	switch c {
	case RawContent:
		return "raw"
	case IdentifierContent:
		return "id"
	case RelationIdentifierContent:
		return "relation id"
	case UnixSecondsContent:
		return "unix seconds"
	case BooleanContent:
		return "boolean"
	case SecretContent:
		return "secret"
	case NameContent:
		return "name"
	case TitleContent:
		return "title"
	case EmailContent:
		return "email"
	case URIContent:
		return "uri"
	default:
		return "unknown " + strconv.Itoa(int(c))
	}
}

// ResolveContentType resolves column content
func ResolveContentType(c Column) ContentType {
	name := strings.ToLower(c.Name)
	if name == "id" {
		return IdentifierContent
	}
	if strings.HasSuffix(name, "password") || strings.HasPrefix(name, "password") {
		return SecretContent
	}
	if (c.Type == Integer || c.Type == Text) && strings.HasSuffix(name, "id") {
		return RelationIdentifierContent
	}
	if (c.Type == Integer || c.Type == Enumeration) && (name == "enabled" || isBooleanValues(c.Values)) {
		return BooleanContent
	}
	if c.Type == Text {
		switch name {
		case "name":
			return NameContent
		case "email":
			return EmailContent
		case "title":
			return TitleContent
		}
		if strings.HasSuffix(name, "uri") || strings.HasPrefix(name, "uri") || strings.HasSuffix(name, "url") || strings.HasPrefix(name, "url") {
			return URIContent
		}
	}
	if c.Type == Integer && (name == "createdat" || name == "updatedat" || name == "publishedat" || name == "expiryat") {
		return UnixSecondsContent
	}

	return RawContent
}

func isBooleanValues(v []string) bool {
	if len(v) != 2 {
		return false
	}

	if (v[0] == "true" && v[1] == "false") || (v[0] == "false" && v[1] == "true") {
		return true
	}
	return false
}
