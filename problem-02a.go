package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Problem2A struct {

}

func (this *Problem2A) Solve() {
	Log.Info("Problem 2A solver beginning!")


	file, err := os.Open("source-data/input-day-02a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int64(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			parts := strings.Split(line, "x");
			l, _ := strconv.ParseInt(parts[0], 10, 64);
			w, _ := strconv.ParseInt(parts[1], 10, 64);
			h, _ := strconv.ParseInt(parts[2], 10, 64);
			s1 := l * w;
			s2 := w * h;
			s3 := l * h;
			slack := math.Min(float64(s1), float64(s2));
			slack = math.Min(slack, float64(s3));
			sum += (2*s1) + (2*s2) + (2*s3) + int64(slack);
		}
	}
	Log.Info("Finished calculation, total wrapping paper required is %d", sum);
}
