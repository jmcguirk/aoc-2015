package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type Problem12A struct {
	Count 	int;
}

func (this *Problem12A) Solve() {
	Log.Info("Problem 12A solver beginning!")


	file, err := os.Open("source-data/input-day-12a.txt");
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

func (this *Problem12A) TraverseJSONArray(arr []interface{}) {
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

func (this *Problem12A) TraverseJSONMap(arr map[string]interface{}) {
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