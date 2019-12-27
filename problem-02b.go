package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Problem2B struct {

}

func (this *Problem2B) Solve() {
	Log.Info("Problem 2B solver beginning!")


	file, err := os.Open("source-data/input-day-02b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			parts := strings.Split(line, "x");
			l, _ := strconv.ParseInt(parts[0], 10, 64);
			w, _ := strconv.ParseInt(parts[1], 10, 64);
			h, _ := strconv.ParseInt(parts[2], 10, 64);
			s1 := int(l);
			s2 := int(w);
			s3 := int(h);
			faces := make([]int, 3);
			faces[0] = s1;
			faces[1] = s2;
			faces[2] = s3;
			sort.Ints(faces);

			pRibbon := (faces[0] * 2) + (faces[1] * 2);
			vRibbon := int(l * w * h);
			Log.Info("pRib %d, vRib %d", pRibbon, vRibbon);
			sum += pRibbon + vRibbon;
		}
	}
	Log.Info("Finished calculation, total ribbon required is %d", sum);
}
