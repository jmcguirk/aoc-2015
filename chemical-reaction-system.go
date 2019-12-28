package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type ChemicalReactionSystem struct {
	Rules []*ChemicalReactionRule;
	FileName string;
	InitialState string;
}


type ChemicalReactionRule struct {
	Input string;
	Output string;
}

type ChemicalReactionSearchState struct {
	ShortestPath int;
	MaxDepth int;
}


type ChemicalReactionSubState struct {
	NewString string;
	Depth int;
}

func (this *ChemicalReactionRule) Describe() string {
	return fmt.Sprintf("%s => %s", this.Input, this.Output);
}

func (this *ChemicalReactionSystem) Init(fileName string) error{
	this.Rules = make([]*ChemicalReactionRule, 0);
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			if(strings.Contains(line, "=>")){ // Rule
				parts := strings.Split(line, "=>");
				rule := &ChemicalReactionRule{};
				rule.Input = strings.TrimSpace(parts[0]);
				rule.Output = strings.TrimSpace(parts[1]);
				this.Rules = append(this.Rules, rule);
			} else{
				this.InitialState = line;
			}
		}
	}
	return nil;
}



func (this *ChemicalReactionSystem) CalculateMinStepsToGenerateDesiredCompound() int{
	search := &ChemicalReactionSearchState{};
	search.ShortestPath = math.MaxInt64;
	search.MaxDepth = 10;
	this.CalculateMinStepsIter(this.InitialState, 0, search);
	return search.ShortestPath;
}


func (this *ChemicalReactionSystem) CalculateMinStepsIter(startString string, startDepth int, state *ChemicalReactionSearchState){


	frontier := make(map[string]*ChemicalReactionSubState);
	frontierArr := make([]*ChemicalReactionSubState, 0);
	seen := make(map[string]int);

	initial := this.GeneratePossibleReplacements(startString, true);
	for k, _ := range initial{
		sub := &ChemicalReactionSubState{};
		frontier[k] = sub;
		sub.Depth = startDepth+1;
		sub.NewString = k;
		frontierArr = append(frontierArr, sub);
	}


	for{
		if(len(frontier) <= 0) {
			Log.Info("Failed to find a valid replacement path");
			break;
		}


		sort.SliceStable(frontierArr, func(i, j int) bool {
			lenI := len(frontierArr[i].NewString);
			lenJ := len(frontierArr[j].NewString);
			if(lenI != lenJ){
				return lenI < lenJ;
			}
			return frontierArr[i].Depth < frontierArr[j].Depth;
		});

		next := frontierArr[0];
		frontierArr = frontierArr[1:];
		delete(frontier, next.NewString);

		if(next.NewString == "e"){
			Log.Info("Found path at depth %d!", next.Depth)
			state.ShortestPath = next.Depth;
			return;
		}

		sub := this.GeneratePossibleReplacements(next.NewString, true);
		for k, _ := range sub{

			shortest, exists := seen[k];
			if(exists && next.Depth >= shortest){
				continue;
			}
			seen[k] = next.Depth;
			existing, exists := frontier[k];
			if(exists){
				if(next.Depth + 1 < existing.Depth){
					existing.Depth = next.Depth;
				}
			} else{
				subState := &ChemicalReactionSubState{};
				frontier[k] = subState;
				subState.Depth = next.Depth+1;
				subState.NewString = k;
				frontierArr = append(frontierArr, subState);
			}

		}
	}
}

func (this *ChemicalReactionSystem) GeneratePossibleReplacements(sourceString string, reverse bool) map[string]int{
	successorStates := make(map[string]int);
	for _, rule := range this.Rules{
		substr := sourceString;
		index := 0;
		for{
			target := rule.Input;
			replace := rule.Output;
			if(reverse){
				target = rule.Output;
				replace = rule.Input;
			}
			next := strings.Index(substr, target);
			if(next == -1){
				break;
			}
			actualIndex := index + next;
			rewrite := sourceString[0:actualIndex];
			rewrite += replace;
			pivot := actualIndex + len(target);
			rewrite += sourceString[pivot:len(sourceString)];
			_, exists := successorStates[rewrite];
			if(!exists){
				successorStates[rewrite] = 1;
			}
			substr = substr[next+1:];
			index += next+1;
		}
	}
	return successorStates;
}


func (this *ChemicalReactionSystem) CalculateMinStepsRecur(currentString string, depth int, search *ChemicalReactionSearchState){
	if(currentString == this.InitialState){
		if(depth < search.ShortestPath){
			Log.Info("New shortest path found %s %d", currentString, depth);
			search.ShortestPath = depth;
		}
		return;
	}
	if(depth > search.MaxDepth){
		return;
	}
	successorStates := make(map[string]int);
	for _, rule := range this.Rules{
		substr := currentString;
		index := 0;
		for{
			next := strings.Index(substr, rule.Input);
			if(next == -1){
				break;
			}
			actualIndex := index + next;
			rewrite := currentString[0:actualIndex];
			rewrite += rule.Output;
			pivot := actualIndex + len(rule.Input);
			rewrite += currentString[pivot:len(currentString)];
			_, exists := successorStates[rewrite];
			if(!exists){
				successorStates[rewrite] = 1;
			}
			substr = substr[next+1:];
			index += next+1;
		}
	}
	for k, _ := range successorStates{
		this.CalculateMinStepsRecur(k, depth+1, search);
	}
}


func (this *ChemicalReactionSystem) CountUniqueOutputs() int{
	sum := 0;
	seen := make(map[string]int);
	for _, rule := range this.Rules{
		substr := this.InitialState;
		index := 0;
		for{
			//Log.Info("Remaining %s",substr);
			next := strings.Index(substr, rule.Input);
			if(next == -1){
				break;
			}
			actualIndex := index + next;
			rewrite := this.InitialState[0:actualIndex];
			//Log.Info("Left %s", rewrite);
			rewrite += rule.Output;
			//Log.Info("Middle %s", rewrite);
			pivot := actualIndex + len(rule.Input);
			rewrite += this.InitialState[pivot:len(this.InitialState)];
			//Log.Info("Right %s", rewrite);
			//Log.Info("Rewrote %s to %s using %s", this.InitialState, rewrite, rule.Describe())
			_, exists := seen[rewrite];
			if(!exists){
				seen[rewrite] = 1;
				sum++;
			}
			substr = substr[next+1:];
			index += next+1;
		}
	}

	return sum;
}
