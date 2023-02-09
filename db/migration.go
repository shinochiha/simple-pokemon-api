package db

import (
	"sort"

	"github.com/api-pokemon/models"
)

func Migrate() {
	hasNewMigration := false
	setting := models.Setting{Key: "db.migration.version"}
	db.AutoMigrate(&setting)
	db.Where(models.Setting{Key: setting.Key}).FirstOrCreate(&setting)

	index := make([]string, 0)
	for i := range migration {
		index = append(index, i)
	}
	sort.Strings(index)
	for _, i := range index {
		if setting.Value == "" || setting.Value < i {
			migration[i]()
			setting.Value = i
			hasNewMigration = true
		}
	}
	if hasNewMigration {
		db.Where(models.Setting{Key: setting.Key}).Assign(setting).FirstOrCreate(&setting)
	}
}

var migration = map[string]func(){
	"0002": func() { db.AutoMigrate(&models.User{}) },
	"0008": func() { db.AutoMigrate(&models.Pokemon{}) },
	"0011": func() { db.AutoMigrate(&models.MyPokemon{}) },
}
