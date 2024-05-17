package litedb

import (
	"encoding/json"
	"testing"
)

type user struct {
	First_Name  string  `json:"first_name"`
	Second_Name string  `json:"second_name"`
	Nickname    string  `json:"nickname"`
	Age         float64 `json:"age"`
}

var (
	test_db    = LiteDb{dir_path: "./test_db"}
	test_table = test_db.Table("test")
	test_user  = user{
		First_Name:  "Ben",
		Second_Name: "Ten",
		Nickname:    "Ben10",
		Age:         21,
	}
)

func TestAllTabMethods(t *testing.T) {
	// NewDoc
	test_user_bytes, _ := json.Marshal(test_user)
	new_id := test_table.NewDoc(test_user_bytes)

	// Doc
	opened_test_user_data := OpenUser(new_id)
	if opened_test_user_data.Age != test_user.Age {
		t.Error("Wrong opened data!")
	}

	// UpdateDoc
	test_user.Age++
	test_user_bytes, _ = json.Marshal(test_user)
	test_table.UpdateDoc(new_id, test_user_bytes)

	updated_test_user_data := OpenUser(new_id)
	if updated_test_user_data.Age != test_user.Age {
		t.Error("Wrong updated opened data!")
	}

	// DocPath
	new_doc_path := test_table.DocPath(new_id)
	if !IsFile(new_doc_path) {
		t.Error("Doc doesn't exists")
	}

	// Filter
	criteria := []Criterion{
		{Name: "first_name", Value: "Ben"},
		{Name: "second_name", Value: "Ten"},
		{Name: "age", Value: float64(22)},
	}

	filter_result := test_table.Filter(criteria)

	if len(filter_result) != len(test_table.IDs()) {
		t.Error("Wrong filter result length")
	}

	//DeleteDoc
	test_table.DeleteDoc(new_id)
	if IsFile(new_doc_path) {
		t.Error("Doc exists even after delete")
	}

}

func OpenUser(id string) user {
	return BytesToJson[user](test_table.Doc(id))
}
