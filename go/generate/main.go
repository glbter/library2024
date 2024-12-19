package main

import (
	"library/internal/config"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/store/query",
		OutFile:           "query.go",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithTypeTag:  true,
		FieldWithIndexTag: true,
	})

	cfg := config.MustLoadConfig()
	db := mustOpen(cfg.DSN)
	g.UseDB(db)

	g.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		"uuid": func(columnType gorm.ColumnType) (dataType string) {
			_, hasDefaultValue := columnType.DefaultValue()
			if hasDefaultValue {
				return "pgtype.UUID"
			}
			return "uuid.UUID"
		},
		"text": func(columnType gorm.ColumnType) (dataType string) {
			isNullable, ok := columnType.Nullable()
			if ok && !isNullable {
				return "string"
			}
			return "pgtype.Text"
		},
		"date": func(columnType gorm.ColumnType) (dataType string) {
			return "pgtype.Date"
		},
		"timestamp": func(columnType gorm.ColumnType) (dataType string) {
			return "pgtype.Timestamp"
		},
	})
	g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(
		g.GenerateModel(
			"users",
			gen.FieldJSONTag("password_hash", "-"),
			gen.FieldComment("password_hash", "Must not be exposed"),
			gen.FieldType("password_hash", "string"),
			gen.FieldGORMTag("email", func(tag field.GormTag) field.GormTag {
				return tag.Set("uniqueIndex")
			}),
		),
		g.GenerateModel(
			"sessions",
			gen.FieldGORMTag("user_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:CASCADE")
			}),
		),
		g.GenerateModel("genres", gen.FieldGORMTag("name", func(tag field.GormTag) field.GormTag {
			return tag.Set("uniqueIndex")
		})),
		g.GenerateModel(
			"book_to_authors",
			gen.FieldGORMTag("book_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:CASCADE")
			}),
			gen.FieldGORMTag("author_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:CASCADE")
			}),
		),
		g.GenerateModel(
			"book_to_genres",
			gen.FieldGORMTag("book_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:CASCADE")
			}),
			gen.FieldGORMTag("genre_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:CASCADE")
			}),
		),
		g.GenerateModel(
			"book_lend_requests",
			gen.FieldGORMTag("book_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:RESTRICT")
			}),
			gen.FieldGORMTag("user_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:RESTRICT")
			}),
		),
		g.GenerateModel(
			"book_lend_transactions",
			gen.FieldComment("request_id", "RequestID is a foreign key to `book_lend_requests`"),
			gen.FieldGORMTag("request_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:RESTRICT")
			}),
		),
		g.GenerateModel(
			"book_return_transactions",
			gen.FieldComment("request_id", "RequestID is a foreign key to `book_lend_requests` as well as `book_lend_transactions`"),
			gen.FieldGORMTag("request_id", func(tag field.GormTag) field.GormTag {
				return tag.Set("constraint", "OnUpdate:CASCADE,OnDelete:RESTRICT")
			}),
		),
	)

	g.Execute()
}

func open(dsn string) (*gorm.DB, error) {
	// make the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
}

func mustOpen(dsn string) *gorm.DB {
	db, err := open(dsn)
	if err != nil {
		panic(err)
	}

	return db
}
