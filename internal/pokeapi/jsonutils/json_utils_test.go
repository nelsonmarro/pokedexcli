package jsonutils

import "testing"

type TestModel struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func TestShouldDeserializeJson(t *testing.T) {
	// Arrange
	jsonString := `{"name":"test","id":1}`
	want := &TestModel{
		Name: "test",
		Id:   1,
	}
	// Act
	model, err := DeserializeJson[TestModel]([]byte(jsonString))
	// Assert
	if err != nil {
		t.Error("Expected to desirealize json without errors")
		return
	}

	if model.Id != want.Id && model.Name != want.Name {
		t.Error("Expected to deserialize json to TestModel")
	}
}
