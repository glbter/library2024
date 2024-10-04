package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"library/internal/config"
	"os"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/store/query",
		OutFile: "query.go",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	cfg := config.MustLoadConfig()
	gormdb := mustOpen(cfg.DSN)
	g.UseDB(gormdb)

	g.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		"uuid": func(columnType gorm.ColumnType) (dataType string) {
			//pgtype.UUID{}
			return "uuid.UUID"
		},
	})
	g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(
		g.GenerateModel("users", gen.FieldJSONTag("password_hash", "-")),
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
