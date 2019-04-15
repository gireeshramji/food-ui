package db

import (
	"reflect"
	"testing"
	"time"

	"github.com/gireeshramji/food-ui/config"
)

func TestParseRecipeDb(t *testing.T) {
	type args struct {
		config *config.Config
	}
	type testStruct struct {
		name string
		args args
		want *RecipeEntryDb
	}
	tests := []testStruct{
		// TODO: Add test cases.
		testStruct{
			name: "Test1",
			args: args{config: &config.Config{StoragePath: "data/recipes.toml"}},
			want: &RecipeEntryDb{
				RecipeEntries: []RecipeEntry{
					RecipeEntry{
						Recipe{
							RecipeLink:      "abc",
							ContainsMeat:    true,
							ContainsSeaFood: true,
						},
						RecipeExecution{
							DateToCook: time.Time{
								2019, 04, 20, 0, 0, 0, 0, time.UTC,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseRecipeDb(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRecipeDb() = %v, want %v", got, tt.want)
			}
		})
	}
}
