package pgstorage

type bucketNum uint16

type game struct {
	ID    				int64 `db:"id"`
	GameProto 			[]byte `db:"game_proto"`
}

const (
	bucketPrefix = "bucket_" // "bucket_n"
	tableName = "games"

	idColumnName   		  = "id"
	gameProtoColumnName   = "game_proto"
	createdAtColumnName   = "created_at"
	updatedAtColumnName   = "updated_at"
)
