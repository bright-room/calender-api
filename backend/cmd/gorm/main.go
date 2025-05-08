package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "infrastructure/datasource/mapper",
		ModelPkgPath:  "infrastructure/datasource/entity",
		Mode:          gen.WithDefaultQuery,
		FieldNullable: false,
	})
}
