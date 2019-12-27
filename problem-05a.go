package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem5A struct {

}

func (this *Problem5A) Solve() {
	Log.Info("Problem 5A solver beginning!")


	file, err := os.Open("source-data/input-day-05a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int64(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			if(IsNiceWord(line)){
				sum++;
			}
		}
	}
	Log.Info("Finished calculation, total nice words is %d", sum);
}

func IsNiceWord(word string) bool {
	// A little sloppy since we iterate over it below
	if(strings.Contains(word, "ab")){
		return false;
	}
	if(strings.Contains(word, "cd")){
		return false;
	}
	if(strings.Contains(word, "pq")){
		return false;
	}
	if(strings.Contains(word, "xy")){
		return false;
	}

	vowelCount := 0;
	dupeCount := 0;
	for i, c := range word{
		if(c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'){
			vowelCount++;
		}
		if(i > 0 && int(word[i-1]) == int(c)){
			dupeCount++;
		}
	}
	return vowelCount >= 3 && dupeCount > 0;
}