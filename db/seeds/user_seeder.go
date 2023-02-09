package seeds

import (
	"github.com/api-pokemon/helpers"
	"github.com/api-pokemon/models"
	"github.com/jinzhu/gorm"
)

func SeedUser(db *gorm.DB) {
	for _, o := range User {
		db.Where(models.User{ID: o.ID}).Assign(o).FirstOrCreate(&o)
	}
}

var User = []models.User{
	{ID: "7bff3968-617b-4347-b59f-f1afccfca5a0", Username: "test", Email: "test@gmail.com", Password: "$2a$10$fj6mI37VxHcPbHxecpHuoOo4liqhh6Ylz4d6Jyefd38uCHL1CC1KS", IsActive: helpers.BoolAddr(true)},
}
