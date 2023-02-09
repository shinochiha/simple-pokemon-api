package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-pokemon/helpers"
)

type Pokemon struct {
	ID                    int    `json:"id,omitempty" gorm:"primaryKey;type:int"`
	Name                  string `json:"name" gorm:"type:varchar(100)"`
	Type                  string `json:"type" gorm:"type:varchar(255)"`
	ImageBackDefault      string `json:"back_default" gorm:"type:varchar(255)"`
	ImageBackFemale       string `json:"back_female" gorm:"type:varchar(255)"`
	ImageBackShiny        string `json:"back_shiny" gorm:"type:varchar(255)"`
	ImageBackShinyFemale  string `json:"back_shiny_female" gorm:"type:varchar(255)"`
	ImageFrontDefault     string `json:"front_default" gorm:"type:varchar(255)"`
	ImageFrontFemale      string `json:"front_female" gorm:"type:varchar(255)"`
	ImageFrontShiny       string `json:"front_shiny" gorm:"type:varchar(255)"`
	ImageFrontShinyFemale string `json:"front_shiny_female" gorm:"type:varchar(255)"`
}

func (Pokemon) TableName() string {
	return "pokemons"
}

func (o *Pokemon) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "pokemons", "as": "p"},
		"fields": map[string]map[string]string{
			"id":                  {"name": "p.id", "as": "id"},
			"name":                {"name": "p.name", "as": "name", "type": "string"},
			"type":                {"name": "p.type", "as": "type", "type": "string"},
			"image.front_default": {"name": "p.image_front_default", "as": "front_default", "type": "string"},
		},
	}
}

type PokemonDetail struct{}

func (o *PokemonDetail) Schema() map[string]interface{} {
	return map[string]interface{}{
		"table": map[string]string{"name": "pokemons", "as": "p"},
		"fields": map[string]map[string]string{
			"id":                       {"name": "p.id", "as": "id"},
			"name":                     {"name": "p.name", "as": "name", "type": "string"},
			"type":                     {"name": "p.type", "as": "type", "type": "string"},
			"move":                     {"name": "p.move", "as": "move", "type": "string"},
			"image.back_default":       {"name": "p.image_back_default", "as": "back_default", "type": "string"},
			"image.back_female":        {"name": "p.image_back_female", "as": "back_female", "type": "string"},
			"image.back_shiny":         {"name": "p.image_back_shiny", "as": "back_shiny", "back_shiny": "string"},
			"image.back_shiny_female":  {"name": "p.image_back_shiny_female", "as": "back_shiny_female", "type": "string"},
			"image.front_default":      {"name": "p.image_front_default", "as": "front_default", "type": "string"},
			"image.front_female":       {"name": "p.image_front_female", "as": "front_female", "type": "string"},
			"image.front_shiny":        {"name": "p.image_front_shiny", "as": "front_shiny", "front_shiny": "string"},
			"image.front_shiny_female": {"name": "p.image_front_shiny_female", "as": "front_shiny_female", "type": "string"},
		},
	}
}

func (o *Pokemon) GetPaginated(ctx helpers.Context, params map[string][]string) map[string]interface{} {
	var item struct{}
	ra := helpers.GetDB(ctx).Table("pokemons").Where("id = ?", 1).Limit(1).Find(&item).RowsAffected
	if ra < 1 {
		response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=1279")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		data := response.Body
		defer data.Close()
		// Parsing data Pokemon
		var dataPokemon struct {
			Results []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"results"`
		}
		err = json.NewDecoder(response.Body).Decode(&dataPokemon)
		if err != nil {
			return helpers.Map{"message": "Error decoding data:" + err.Error()}
		}

		for i, pokemon := range dataPokemon.Results {
			res, err := http.Get(pokemon.URL)
			if err != nil {
				return helpers.Map{"message": "Error fetching data: " + err.Error()}
			}
			var pokeData struct {
				Sprites struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"sprites"`
				Types []struct {
					Type struct {
						Name string `json:"name"`
					} `json:"type"`
				} `json:"types"`
				Moves []struct {
					Move struct {
						Name string `json:"name"`
					} `json:"move"`
				} `json:"moves"`
			}
			err = json.NewDecoder(res.Body).Decode(&pokeData)
			if err != nil {
				return helpers.Map{"message": "Error decoding data:" + err.Error()}
			}
			var types []string
			for _, t := range pokeData.Types {
				types = append(types, t.Type.Name)
			}

			po := Pokemon{
				ID:                    i + 1,
				Name:                  pokemon.Name,
				ImageBackDefault:      pokeData.Sprites.BackDefault,
				ImageBackFemale:       pokeData.Sprites.BackFemale,
				ImageBackShiny:        pokeData.Sprites.BackShiny,
				ImageBackShinyFemale:  pokeData.Sprites.BackShinyFemale,
				ImageFrontDefault:     pokeData.Sprites.FrontDefault,
				ImageFrontFemale:      pokeData.Sprites.FrontFemale,
				ImageFrontShiny:       pokeData.Sprites.FrontShiny,
				ImageFrontShinyFemale: pokeData.Sprites.FrontShinyFemale,
				Type:                  types[0],
			}
			helpers.GetDB(ctx).Model(Pokemon{}).Create(&po)
		}
	}

	return helpers.GetPaginated(ctx, params, o.Schema(), map[string]interface{}{})
}

func (o *PokemonDetail) GetById(ctx helpers.Context, id string, params map[string][]string) map[string]interface{} {
	return helpers.GetById(ctx, "pokemons", "id", id, params, o.Schema(), map[string]interface{}{})
}
