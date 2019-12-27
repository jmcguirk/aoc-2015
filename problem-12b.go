package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type Problem12B struct {
	Count 	int;
}

func (this *Problem12B) Solve() {
	Log.Info("Problem 12B solver beginning!")


	file, err := os.Open("source-data/input-day-12b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	rawJSON := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			rawJSON = line;
			break;
		}
	}
	var parsed []interface{}
	err = json.Unmarshal([]byte(rawJSON), &parsed)
	if(err != nil){
		Log.FatalError(err);
	}
	this.TraverseJSONArray(parsed);
	Log.Info("Traversal complete - checksum is %d" , this.Count);
}

func (this *Problem12B) TraverseJSONArray(arr []interface{}) {
	for _, rawV := range arr{
		//valueType := reflect.TypeOf(rawV).String();
		//Log.Info(valueType);
		if floatV, ok := rawV.(float64); ok {
			this.Count += int(floatV);
			continue;
		}
		if subMap, ok := rawV.(map[string]interface{}); ok {
			this.TraverseJSONMap(subMap);
			continue;
		}
		if subArr, ok := rawV.([]interface{}); ok {
			this.TraverseJSONArray(subArr);
			continue;
		}
	}
}

func (this *Problem12B) TraverseJSONMap(arr map[string]interface{}) {
	for _, rawV := range arr{ // preflight check for problem b
		if stringV, ok := rawV.(string); ok {
			if(stringV == "red"){
				return;
			}
		}
	}
	for _, rawV := range arr{
		if floatV, ok := rawV.(float64); ok {
			this.Count += int(floatV);
			continue;
		}
		if subMap, ok := rawV.(map[string]interface{}); ok {
			this.TraverseJSONMap(subMap);
			continue;
		}
		if subArr, ok := rawV.([]interface{}); ok {
			this.TraverseJSONArray(subArr);
			continue;
		}
	}
}