package litedb

import (
	"encoding/json"
	"fmt"
	"testing"
)

type user struct {
	First_Name  string `json:"first_name"`
	Second_Name string `json:"second_name"`
	Nickname    string `json:"nickname"`
}

var (
	test_db    = LiteDb{dir_path: "./test_db"}
	test_table = test_db.Table("test")
	test_user  = user{
		First_Name:  "Ben",
		Second_Name: "Ten",
		Nickname:    "Ben10",
	}
)

func TestMain(t *testing.T) {
	t.Run("Saving", Saving)
	t.Run("Reading", Reading)
	t.Run("Filtering", Filter)
}

func Saving(t *testing.T) {
	json_data, _ := json.Marshal(test_user)
	test_table.NewDoc(json_data)
}

func Reading(t *testing.T) {
	last_index := test_table.LastIndex()

	user_data := test_table.Doc(last_index)

	var user_json user
	json.Unmarshal(user_data, &user_json)

	if user_json.Nickname != test_user.Nickname {
		t.Error("Wrong data!")
	}
}

func Filter(t *testing.T) {
	criteria := []Criterion{
		{Name: "first_name", Value: "Ben"},
		{Name: "age", Value: float64(21)},
	}

	fmt.Println(test_table.Filter(criteria))
}
