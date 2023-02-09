package models

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/api-pokemon/helpers"
)

type MyPokemon struct {
	ID                int    `json:"id,omitempty" gorm:"primaryKey;type:int"`
	NickName          string `json:"nickname" gorm:"type:varchar(100)"`
	Name              string `json:"name" gorm:"type:varchar(100)"`
	Type              string `json:"type" gorm:"type:varchar(255)"`
	IsActive          *bool  `json:"is_active" gorm:"type:bool"`
	ImageBackDefault  string `json:"back_default" gorm:"type:varchar(255)"`
	ImageFrontDefault string `json:"front_default" gorm:"type:varchar(255)"`
	UserID            string `json:"user_id" gorm:"type:bpchar(36)"`
}

func (MyPokemon) TableName() string {
	return "my_pokemons"
}

func (o *MyPokemon) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "my_pokemons", "as": "p"},
		"fields": map[string]map[string]string{
			"id":                  {"name": "p.id", "as": "id"},
			"name":                {"name": "p.name", "as": "name", "type": "string"},
			"nickname":            {"name": "p.nick_name", "as": "nickname", "nickname": "string"},
			"type":                {"name": "p.type", "as": "type", "type": "string"},
			"is_active":           {"name": "p.is_active", "as": "is_active"},
			"image.front_default": {"name": "p.image_front_default", "as": "front_default", "type": "string"},
			"image.back_default":  {"name": "p.image_back_default", "as": "back_default", "type": "string"},
			"user_id":             {"name": "p.user_id", "as": "user_id", "is_hide": "true"},
		},
	}
}

func (o *MyPokemon) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	params["is_active"] = []string{"true"}
	params["user_id"] = []string{ctx.Get("jwt_user_id").(string)}
	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *MyPokemon) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "my_pokemons", "id", id, params, o.Schema(), map[string]interface{}{})
}

func (o *MyPokemon) Create(ctx helpers.Context, id string) map[string]interface{} {

	pokemon := MyPokemon{}
	helpers.GetDB(ctx).Table("pokemons").Where("id = ? ", id).First(&pokemon)

	random := rand.Intn(100)
	if random > 50 {
		mypokemons := MyPokemon{
			ID:                pokemon.ID,
			Name:              pokemon.Name,
			Type:              pokemon.Type,
			ImageBackDefault:  pokemon.ImageBackDefault,
			ImageFrontDefault: pokemon.ImageFrontDefault,
			UserID:            ctx.Get("jwt_user_id").(string),
			IsActive:          helpers.BoolAddr(false),
		}
		helpers.GetDB(ctx).Model(MyPokemon{}).Create(&mypokemons)
		return helpers.Map{
			"message": "Catch pokemon '" + pokemon.Name + "' success with percentage " + fmt.Sprint(random) + "%",
			"code":    200,
		}
	} else {
		return helpers.Map{
			"message": "Catch pokemon " + pokemon.Name + " failed with percentage " + fmt.Sprint(random) + "%",
			"code":    200,
		}
	}
}

func (o *MyPokemon) SavePokemon(ctx helpers.Context) map[string]interface{} {
	o.IsActive = helpers.BoolAddr(true)
	helpers.GetDB(ctx).Model(MyPokemon{}).Where("id = ?", o.ID).Updates(o)
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

func (o *MyPokemon) RealeasePokemon(ctx helpers.Context) map[string]interface{} {
	n := rand.Intn(21)
	if helpers.IsPrime(n) {
		helpers.GetDB(ctx).Model(MyPokemon{}).Where("id = ?", o.ID).Delete(&MyPokemon{})
		return helpers.Map{
			"message": "Pokemon telah dibebaskan menggunakan bilangan prima " + fmt.Sprint(n),
			"code":    200,
		}
	} else {
		return helpers.Map{
			"message": "Gagal membebaskan Pokemon, bilangan " + fmt.Sprint(n) + " bukan bilangan prima",
			"code":    200,
		}
	}
}

func (o *MyPokemon) ChangeNamePokemon(ctx helpers.Context) map[string]interface{} {
	helpers.GetDB(ctx).Model(MyPokemon{}).Where("id = ?", o.ID).Update(MyPokemon{NickName: o.NickName + "-" + strconv.Itoa(fibonacci())})
	return o.GetById(ctx, helpers.Convert(o.ID).String(), map[string][]string{})
}

var count int

func fibonacci() int {
	count++
	if count == 1 {
		return 0
	}
	if count == 2 {
		return 1
	}
	if count == 3 {
		return 1
	}
	if count == 4 {
		return 2
	}
	if count == 5 {
		return 3
	}
	if count == 6 {
		return 5
	}
	if count == 7 {
		return 8
	}
	return fibonacci() + fibonacci()
}
