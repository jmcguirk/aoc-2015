package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type WeightOptimizationProblem struct {
	CanonicalWeights []int;
	FileName string;
	InitialState string;
	FewestPackagesEncountered int;
	SmallestEntanglementEncountered int;
}

type WeightOptimizationState struct {
	LeftContainer []int;
	RightContainer []int;
	BottomContainer []int;
	QE int;
	Passenger []int;
	Index int;
	IncludeBottom bool;
}

func (this *WeightOptimizationState) Clone() *WeightOptimizationState{
	res := &WeightOptimizationState{};
	res.LeftContainer = make([]int, 0);
	res.LeftContainer = append(res.LeftContainer, this.LeftContainer...);
	res.RightContainer = make([]int, 0);
	res.RightContainer = append(res.RightContainer, this.RightContainer...);
	res.Passenger = make([]int, 0);
	res.Passenger = append(res.Passenger, this.Passenger...);
	res.BottomContainer = make([]int, 0);
	res.BottomContainer = append(res.BottomContainer, this.BottomContainer...);
	res.QE = this.QE;
	res.Index = this.Index+1;
	res.IncludeBottom = this.IncludeBottom;
	return res;
}


func (this *WeightOptimizationProblem) Init(fileName string) error{
	this.FileName = fileName;
	this.CanonicalWeights = make([]int, 0);
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
			weight, err := strconv.ParseInt(line, 10, 64);
			if(err != nil){
				return err;
			}
			this.CanonicalWeights = append(this.CanonicalWeights, int(weight));
		}
	}

	Log.Info("Parsed %d weights from %s", len(this.CanonicalWeights), fileName);
	return nil;
}


func (this *WeightOptimizationProblem) Optimize(includeBottom bool) (int, int) {
	return this.OptimizeWithKnownGood(includeBottom, int(math.MaxInt64), int(math.MaxInt64));
}

func (this *WeightOptimizationProblem) OptimizeWithKnownGood(includeBottom bool, fewestPackages int, smallestEntanglement int) (int, int) {
	this.FewestPackagesEncountered = fewestPackages;
	this.SmallestEntanglementEncountered = smallestEntanglement;
	state := &WeightOptimizationState{};
	state.IncludeBottom = includeBottom;
	state.QE = 1;
	this.SolveRecursive(state);
	return this.FewestPackagesEncountered, this.SmallestEntanglementEncountered;
}

func (this *WeightOptimizationProblem) SolveRecursive(state *WeightOptimizationState) {
	if(len(state.Passenger) > this.FewestPackagesEncountered){
		return;
	} else if(len(state.Passenger) == this.FewestPackagesEncountered && state.QE >= this.SmallestEntanglementEncountered){
		return;
	}
	if(state.Index >= len(this.CanonicalWeights)){
		sumL := SumArray(state.LeftContainer);
		sumR := SumArray(state.RightContainer);
		if(sumL != sumR){
			return;
		}
		sumC := SumArray(state.Passenger);
		if(sumL != sumC){
			return;
		}
		if(state.IncludeBottom){
			sumB := SumArray(state.BottomContainer);
			if(sumB != sumC){
				return;
			}
		}
		if(len(state.Passenger) <= this.FewestPackagesEncountered){
			quant := MulArray(state.Passenger);
			if(quant < this.SmallestEntanglementEncountered){
				Log.Info("Found new best arrangement %d packages, entanglement is %d", len(state.Passenger), quant);
				this.SmallestEntanglementEncountered = quant;
				this.FewestPackagesEncountered = len(state.Passenger);
			}
		}
		return;
	}

	pack := this.CanonicalWeights[state.Index];
	left := state.Clone();
	left.LeftContainer = append(left.LeftContainer, pack);
	right := state.Clone();
	right.RightContainer = append(right.RightContainer, pack);
	center := state.Clone();
	center.Passenger = append(center.Passenger, pack);
	center.QE = center.QE * pack;
	var bottom *WeightOptimizationState;

	if(state.IncludeBottom){
		bottom = state.Clone();
		bottom.BottomContainer = append(bottom.BottomContainer, pack);
	}

	if(state.Index < 2){
		operationDone := make(chan bool)
		go this.SolveRecursive(left);
		go this.SolveRecursive(right);
		go this.SolveRecursive(center);
		if(bottom != nil){
			go this.SolveRecursive(bottom);
		}
		<-operationDone
	} else{
		this.SolveRecursive(left);
		this.SolveRecursive(right);
		this.SolveRecursive(center);
		if(bottom != nil){
			go this.SolveRecursive(bottom);
		}
	}

}