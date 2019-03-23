package a

type TestType struct {
	UserId    string `json:"user"`       // want "invalid snake case json tag"
	TestField string `json:"test_field"` // OK
}
