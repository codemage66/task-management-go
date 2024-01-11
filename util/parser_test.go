package util

type ExampleStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

/* func TestParseRequestBodyJSON(t *testing.T) {
	requestBody := `{"name": "example", "value": 42}`
	req, err := http.NewRequest("POST", "/test", bytes.NewBufferString(requestBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	var data ExampleStruct
	err = ParseRequestBody(req, &data)
	assert.NoError(t, err)

	expectedData := ExampleStruct{Name: "example", Value: 42}
	assert.Equal(t, expectedData, data)
} */
