# Description
Marshal or Unmarshal json value to nondeterministic go struct

# How to use

```go
func main() {
	var data = `
	{
		"name": "the-name",
		"value": [
			"1",
			true,
			{
				"key1": "1",
				"key2": 2
			}
		]
	}`

	type kv struct {
		Name  string
		Value jsonv.JsonV
	}

	m := kv{}
	json.Unmarshal([]byte(data), &m)
	fmt.Printf("%#v", m)
}
```