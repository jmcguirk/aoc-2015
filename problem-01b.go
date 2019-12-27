package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem1B struct {

}

func (this *Problem1B) Solve() {
	Log.Info("Problem 1B solver beginning!")


	file, err := os.Open("source-data/input-day-01b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currVal int = 0;
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		for i, c := range line{
			if(int(c)) == int('('){
				currVal++
			}else if(int(c)) == int(')'){

				if(currVal == 0){
					Log.Info("Santa went into the basement on step %d", i + 1);
				}
				currVal--
			}
		}
	}
	Log.Info("Finished simulation, santa should be on floor %d", currVal);
}
