package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem13B struct {
	BestSeating []int;
	BestSeatingScore int;
}

func (this *Problem13B) Solve() {
	Log.Info("Problem 13B solver beginning!")


	file, err := os.Open("source-data/input-day-13b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	ruleSet := make(map[int]*map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			//Log.Info(line);
			parts := strings.Split(line, " ");
			from := int(parts[0][0]);
			to := int(parts[len(parts) - 1][0]);

			magnitude, err := strconv.ParseInt(parts[3], 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			if(parts[2] == "lose"){
				magnitude = magnitude * -1;
			}

			_, exists := ruleSet[from];
			if(!exists){
				arr := make(map[int]int);
				ruleSet[from] = &arr;
			}
			rulesP, _ := ruleSet[from];
			rules := *rulesP;
			rules[to] = int(magnitude);

		}
	}
	Log.Info("Finished parsing rule set");

	arr := make(map[int]int);
	for k, v := range ruleSet{
		arr[k] = 0;
		rulesP := *v;
		rulesP['X'] = 0;
	}
	ruleSet['X'] = &arr;

	uniqueSeats := make([]int, 0);
	for k, _ := range ruleSet{
		uniqueSeats = append(uniqueSeats, k);
	}
	PermInt(uniqueSeats, func(ints []int) {
		this.ScoreSeating(ints, ruleSet);
	})
	Log.Info("Best seating determined %s - with total score %d", AsciiArrayToString(this.BestSeating), this.BestSeatingScore)
}
func (this *Problem13B) ScoreSeating(seating []int, rules map[int]*map[int]int) {
	//Log.Info("Considering " + AsciiArrayToString(seating));
	//config := AsciiArrayToString(seating);
	//shouldLog := config == "CDAB";
	sum := 0;
	for i, c := range seating{
		leftIndex := i - 1;
		if i == 0{
			leftIndex = len(seating) - 1;
		}
		left := seating[leftIndex];
		right := seating[(i+1)%len(seating)];

		preferencesP, _ := rules[int(c)];
		preferences := *preferencesP;
		leftScore, _ := preferences[int(left)];
		rightScore, _ := preferences[int(right)];
		sum += leftScore;
		sum += rightScore;
	}
	if(this.BestSeating == nil || sum > this.BestSeatingScore){
		this.BestSeatingScore = sum;
		this.BestSeating = seating;
	}
}