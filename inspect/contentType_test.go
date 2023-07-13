package inspect

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var resolveContentTypeData = []struct {
	Expected ContentType
	Given    Column
}{
	{Expected: IdentifierContent, Given: Column{Name: "id"}},
	{Expected: IdentifierContent, Given: Column{Name: "ID"}},

	{Expected: SecretContent, Given: Column{Name: "passWord"}},
	{Expected: SecretContent, Given: Column{Name: "userPassword"}},
	{Expected: SecretContent, Given: Column{Name: "passwordForStaff"}},

	{Expected: RelationIdentifierContent, Given: Column{Type: Integer, Name: "userId"}},
	{Expected: RelationIdentifierContent, Given: Column{Type: Integer, Name: "USER_ID"}},
	{Expected: RelationIdentifierContent, Given: Column{Type: Text, Name: "userId"}},
	{Expected: RelationIdentifierContent, Given: Column{Type: Text, Name: "USER_ID"}},
	{Expected: RawContent, Given: Column{Type: Decimal, Name: "userId"}},

	{Expected: BooleanContent, Given: Column{Type: Boolean}},
	{Expected: BooleanContent, Given: Column{Type: Integer, Name: "Enabled"}},
	{Expected: BooleanContent, Given: Column{Type: Enumeration, Name: "Enabled"}},
	{Expected: BooleanContent, Given: Column{Type: Enumeration, Name: "Whatever", Values: []string{"true", "false"}}},

	{Expected: NameContent, Given: Column{Type: Text, Name: "Name"}},
	{Expected: TitleContent, Given: Column{Type: Text, Name: "Title"}},
	{Expected: EmailContent, Given: Column{Type: Text, Name: "eMail"}},
	{Expected: URIContent, Given: Column{Type: Text, Name: "URL"}},
	{Expected: URIContent, Given: Column{Type: Text, Name: "URI"}},
	{Expected: FirstNameContent, Given: Column{Type: Text, Name: "FirstName"}},
	{Expected: FirstNameContent, Given: Column{Type: Text, Name: "First_Name"}},
	{Expected: LastNameContent, Given: Column{Type: Text, Name: "LastName"}},
	{Expected: LastNameContent, Given: Column{Type: Text, Name: "Last_Name"}},
	{Expected: AddressContent, Given: Column{Type: Text, Name: "aDDress"}},
	{Expected: CountryContent, Given: Column{Type: Text, Name: "country"}},
	{Expected: CityContent, Given: Column{Type: Text, Name: "city"}},
	{Expected: PostalZipContent, Given: Column{Type: Text, Name: "zip"}},
	{Expected: PhoneContent, Given: Column{Type: Text, Name: "phone"}},

	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "Time"}},
	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "CreatedAt"}},
	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "UpdatedAt"}},
	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "ExpiryAt"}},
	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "PublishedAt"}},
	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "NextAt"}},
	{Expected: TimestampContext, Given: Column{Type: Integer, Name: "ScheduledAt"}},
}

func TestResolveContentType(t *testing.T) {
	for _, datum := range resolveContentTypeData {
		t.Run(fmt.Sprint(datum.Given), func(t *testing.T) {
			assert.Equal(t, datum.Expected, ResolveContentType(datum.Given))
		})
	}
}
