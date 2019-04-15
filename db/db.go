package db

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gireeshramji/food-ui/config"
)

//Recipe: Stores all metadata pertaining to a specific recipe
type Recipe struct {
	RecipeLink      string `toml:recipeLink`
	ContainsMeat    bool   `toml:containsMeat`
	ContainsSeaFood bool   `toml:containsSeafood`
}

//RecipeExecutions: Stores all metadata pertaining to possible executions
type RecipeExecution struct {
	DateToCook time.Time `toml:dateToCook`
}

//Recipe Entry: Stores linkage between a specific recipe and a specific recipeExecution
type RecipeEntry struct {
	Recipe     Recipe          `toml:recipe`
	RecipeExec RecipeExecution `toml:recipeExecution`
}

type RecipeEntryDb struct {
	RecipeEntries []RecipeEntry `toml:recipeEntries`
}

//Loads RecipeDb from toml file
func ParseRecipeDb(config *config.Config) *RecipeEntryDb {
	recipe := &RecipeEntryDb{}
	toml.DecodeFile(config.StoragePath, recipe)
	return recipe
}
