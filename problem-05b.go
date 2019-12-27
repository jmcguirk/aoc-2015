package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Problem5B struct {

}

func (this *Problem5B) Solve() {
	Log.Info("Problem 5B solver beginning!")


	file, err := os.Open("source-data/input-day-05b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int64(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			if(IsVeryNiceWord(line)){
				sum++;
			}
		}
	}
	Log.Info("Finished calculation, total nice words is %d", sum);
}

func IsVeryNiceWord(word string) bool {
	dupeSkipCount := 0;
	dupeCount := 0;
	counts := make(map[string]int);
	for i, c := range word{
		if(i > 0){
			pair := fmt.Sprintf("%c%c", word[i-1], word[i])
			index, exists := counts[pair];
			if(exists){
				if(i - index >= 2){
					dupeCount++;
				}
			}
			if(!exists){
				Log.Info("Adding %s", pair);
				counts[pair] = i
			}
		}
		if(i > 1 && int(word[i-2]) == int(c)){
			dupeSkipCount++;
		}

	}


	res := dupeCount > 0 && dupeSkipCount > 0;
	Log.Info("%s is nice(%t) because dupe count %d, and skipCount %d", word, res, dupeCount, dupeSkipCount);
	return res;
}