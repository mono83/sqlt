package inspect

import (
	"math"
)

// Priority evaluates view priority
func Priority(c Column) (p int) {
	ct := ResolveContentType(c)

	// Initial priority
	p = 100

	switch c.Type {
	case Boolean:
		p += 100
	case Integer:
		p += 50
	case Text:
		if c.Size > 0 {
			p -= int(math.Log(float64(c.Size))) * 10
		}
	}

	switch ct {
	case SecretContent:
		p = -1
	case IdentifierContent:
		p *= 10
	case BooleanContent:
		p *= 9
	case RelationIdentifierContent:
		p *= 8
	case TimestampContext:
		p *= 7
	case NameContent, PostalZipContent:
		p *= 5
	case TitleContent, FirstNameContent, LastNameContent, CountryContent, CityContent, AddressContent, PhoneContent:
		p *= 4
	case URIContent:
		p *= 3
	default:
		if c.Type == Enumeration {
			p *= 6
		} else if c.Type == Set {
			p *= 4
		}
	}

	if c.Type == Integer && c.Size > 0 {
		p -= c.Size
	}

	if p < -1 {
		p = -1
	}

	return
}
