newmigrate:
	migrate create -ext sql -dir internal/pkg/db/migrations/ -seq $(name)

migrateup:
	migrate -path internal/pkg/db/migrations/ -database "postgresql://root:secret@localhost:5432/hackernews?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/pkg/db/migrations/ -database "postgresql://root:secret@localhost:5432/hackernews?sslmode=disable" -verbose down

.PHONY:migrateup migratedown newmigrate