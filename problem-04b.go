package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

type Problem4B struct {

}

func (this *Problem4B) Solve() {
	Log.Info("Problem 4B solver beginning!")


	file, err := os.Open("source-data/input-day-04b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prefix := "";
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			prefix = line;
			break;
		}
	}
	Log.Info("Loaded prefix %s", prefix);
	check := 0;
	for {
		data := []byte(fmt.Sprintf("%s%d", prefix, check));
		hash := fmt.Sprintf("%x", md5.Sum(data));
		if(len(hash) < 6){
			continue;
		}
		zeroes := 0;
		for _, c := range hash{
			if(int(c) != int('0')){
				break;
			} else{
				zeroes++;
			}
		}
		if(zeroes >= 6){
			Log.Info("Found suffix at index %d", check);
			break;
		}
		check++;
	}
}
