package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type ChangeMakingSystem struct {
	Denominations []int;
}


type ChangeMakingState struct {
	TargetAmount int;
	UniqueArrangements int;
	MinDesired bool;
	BestScore int;
	UniqueChangeArrangements map[string]int;
}

func (this *ChangeMakingSystem) Init(){
	this.Denominations = make([]int, 0);
}

func (this *ChangeMakingSystem) AddDenomination(val int){
	this.Denominations = append(this.Denominations, val);
}

func (this *ChangeMakingSystem) LoadDenominations(fileName string) error{
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			parsed, err := strconv.ParseInt(line, 10, 64);
			if(err != nil){
				return err;
			}
			this.AddDenomination(int(parsed));
		}
	}
	return nil;
}

func (this *ChangeMakingSystem) CountWaysToMakeChange(amt int) int {
	state := &ChangeMakingState{};
	state.TargetAmount = amt;
	state.UniqueChangeArrangements = make(map[string]int);
	currentWeights := make([]int, len(this.Denominations));
	for i, _ := range this.Denominations{
		currentWeights[i] = 0;
	}
	this.MakeChangeCountRecur(state, 0, currentWeights);
	return state.UniqueArrangements;
}

func (this *ChangeMakingSystem) FindMinWaysToMakeChange(amt int) int {
	state := &ChangeMakingState{};
	state.TargetAmount = amt;
	state.MinDesired = true;
	state.BestScore = int(math.MaxInt64);
	state.UniqueChangeArrangements = make(map[string]int);
	currentWeights := make([]int, len(this.Denominations));
	for i, _ := range this.Denominations{
		currentWeights[i] = 0;
	}
	this.MakeChangeCountRecur(state, 0, currentWeights);
	return state.UniqueArrangements;
}


func (this *ChangeMakingSystem) MakeChangeCountRecur(state *ChangeMakingState, index int, currentWeights []int) {
	if(index >= len(this.Denominations)){
		totalWeight := 0;
		totalCurrencies := 0;
		for i, v := range currentWeights{
			if(v > 0){
				totalCurrencies++;
			}
			totalWeight += v * this.Denominations[i];
		}
		if(totalWeight != state.TargetAmount){
			return;
		}
		if(state.MinDesired){
			if(totalCurrencies > state.BestScore){
				return;
			}
			if(totalCurrencies < state.BestScore){
				state.UniqueArrangements = 0;
				state.BestScore = totalCurrencies;
			}
		}



		signature := "";
		for k, v := range currentWeights{
			if(v > 0){
				signature += strconv.Itoa(k);
				signature += ",";
			}
		}
		_, exists := state.UniqueChangeArrangements[signature];
		if(!exists){
			state.UniqueChangeArrangements[signature] = 1;
			state.UniqueArrangements++;
		}


	} else{
		//max := state.TargetAmount/this.Denominations[index];
		for i := 0; i <= 1; i++ {
			currentWeights[index] = i;
			this.MakeChangeCountRecur(state, index+1, currentWeights);
		}

	}
}
