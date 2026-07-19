package projects

import "github.com/jackc/pgx/v5/pgtype"

func ParseUUID(id string) (pgtype.UUID, error) {
	var uid pgtype.UUID

	err := uid.Scan(id)

	return uid, err
}
