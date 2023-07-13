package mysqlt

import "database/sql"

// ConnectSimpleTCP performs simple connection over
// TCP to requested address and database
func ConnectSimpleTCP(login, password, addr, database string) (*sql.DB, error) {
	return sql.Open(
		"mysql",
		login+`:`+password+`@tcp(`+addr+`)/`+database+`?maxAllowedPacket=0`,
	)
}
