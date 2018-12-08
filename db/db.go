package db

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gireeshramji/food-ui/config"
)

//Recipe: Stores all metadata pertaining to a specific recipe
type Recipe struct {
	recipeLink      string `toml:recipeLink`
	containsMeat    string `toml:containsMeat`
	containsSeaFood string `toml:containsSeafood`
}

//RecipeExecutions: Stores all metadata pertaining to possible executions
type RecipeExecution struct {
	dateToCook time.Time `toml:dateToCook`
}

//Recipe Entry: Stores linkage between a specific recipe and a specific recipeExecution
type RecipeEntry struct {
	recipe     Recipe          `toml:recipe`
	recipeExec RecipeExecution `toml:recipeExecution`
}

type RecipeEntryDb struct {
	recipeEntries []RecipeEntry `toml:recipeEntries`
}

func ParseRecipeDb(config *config.Config) *RecipeEntryDb {
	recipe := &RecipeEntryDb{}
	toml.DecodeFile(config.StoragePath, recipe)
	return recipe
}
