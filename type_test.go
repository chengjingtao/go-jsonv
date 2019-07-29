package jsonv

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

type KVPair struct {
	Name  string `json:"name"`
	Value JsonV  `json:"value"`
}

func TestJson(t *testing.T) {
	var data = `[
		{
			"name": "the-name",
			"value": 1
		},
		{
			"name": "the-name",
			"value": "1"
		},
	
		{
			"name": "the-name",
			"value": ""
		},
		{
			"name": "the-name",
			"value": "1"
		},
	
		{
			"name": "the-name",
			"value": "false"
		},
		{
			"name": "the-name",
			"value": false
		},
		{
			"name": "the-name",
			"value": true
		},
	
		{
			"name": "the-name",
			"value": {
				"key1": 1,
				"key2": "2",
				"key3": true
			}
		},
	
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
		},
	
		{
			"name": "the-name",
			"value": "{\"key1\":\"v1\"}"
		}
	]`

	var kv = []KVPair{}
	err := json.Unmarshal([]byte(data), &kv)
	if err != nil {
		t.Errorf("err should not be nil, err:%s", err.Error())
		return
	}
	assertResult(t, data, kv)
	actual, err := json.MarshalIndent(kv, "", "	")
	assert.NoError(t, err)
	assert.True(t, compareText(data, string(actual)), "should be equal")
}

func TestYaml(t *testing.T) {
	var data = `
- name: the-name
  value: 1
- name: the-name
  value: "1"
- name: the-name
  value: ""
- name: the-name
  value: "1"
- name: the-name
  value: "false"
- name: the-name
  value: false
- name: the-name
  value: true
- name: the-name
  value: 
    key1: 1
    key2: "2"
    key3: true
- name: the-name
  value: 
    - "1"
    - true
    - key1: "1"
      key2: 2
- name: the-name
  value: '{\"key1\":\"v1\"}'`

	var kv = []KVPair{}
	err := yaml.Unmarshal([]byte(data), &kv)
	if err != nil {
		t.Errorf("err should not be nil, err:%s", err.Error())
		return
	}
	assertResult(t, data, kv)

	actual, err := yaml.Marshal(kv)
	assert.NoError(t, err)
	assert.True(t, compareText(data, string(actual)), "should be equal")
}

func assertResult(t *testing.T, data string, kv []KVPair) {
	// case 0
	caseData := kv[0]
	assert.Equal(t, int64(1), caseData.Value.IntVal, "should be equal when value is int")
	// case 1
	caseData = kv[1]
	assert.Equal(t, String, caseData.Value.Type, "Should be String")
	assert.Equal(t, "1", caseData.Value.StringVal, "should be equal ")

	// case 2
	caseData = kv[2]
	assert.Equal(t, String, caseData.Value.Type, "Should be String")
	assert.Equal(t, "", caseData.Value.StringVal, "should be equal")

	// case 3
	caseData = kv[3]
	assert.Equal(t, String, caseData.Value.Type, "Should be String")
	assert.Equal(t, "1", caseData.Value.StringVal, "should be equal")

	// case 4
	caseData = kv[4]
	assert.Equal(t, String, caseData.Value.Type, "Should be String")
	assert.Equal(t, "false", caseData.Value.StringVal, "should be equal")
	// case 5
	caseData = kv[5]
	assert.Equal(t, Bool, caseData.Value.Type, "Should be bool")
	assert.Equal(t, false, caseData.Value.BoolVal, "should be equal")
	// case 6
	caseData = kv[6]
	assert.Equal(t, Bool, caseData.Value.Type, "Should be bool")
	assert.Equal(t, true, caseData.Value.BoolVal, "should be equal")

	// case 7
	caseData = kv[7]
	assert.Equal(t, StringMap, caseData.Value.Type, "Should be stringMap")
	assert.Equal(t, int64(1), caseData.Value.StringMapVal["key1"].IntVal, "should be equal")
	assert.Equal(t, "2", caseData.Value.StringMapVal["key2"].StringVal, "should be equal")
	assert.Equal(t, true, caseData.Value.StringMapVal["key3"].BoolVal, "should be equal")

	// case 8
	caseData = kv[8]
	assert.Equal(t, Arrary, caseData.Value.Type, "Should be array")
	assert.Equal(t, "1", caseData.Value.ArrayVal[0].StringVal, "Should be equal")
	assert.Equal(t, true, caseData.Value.ArrayVal[1].BoolVal, "Should be equal")
	assert.Equal(t, "1", caseData.Value.ArrayVal[2].StringMapVal["key1"].StringVal, "Should be equal")
	assert.Equal(t, int64(2), caseData.Value.ArrayVal[2].StringMapVal["key2"].IntVal, "Should be equal")

	// case 9
	caseData = kv[9]
	assert.Equal(t, String, caseData.Value.Type, "Should be string")
}

func compareText(txtSource string, txtTarget string) bool {
	var replaceSpace = func(str string) string {
		re, _ := regexp.Compile(`\s+`)
		str = re.ReplaceAllString(str, " ")

		re, _ = regexp.Compile(`\n+`)
		str = re.ReplaceAllString(str, "\n")

		re, _ = regexp.Compile(`\t+`)
		str = re.ReplaceAllString(str, "\t")

		return strings.TrimSpace(str)
	}
	return replaceSpace(txtSource) == replaceSpace(txtTarget)

}
