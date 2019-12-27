package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CircuitGateRaw = "RAW";
const CircuitGateAnd = "AND";
const CircuitGateLshift = "LSHIFT";
const CircuitGateRshift = "RSHIFT";
const CircuitGateNOT = "NOT";
const CircuitGateOR = "OR";

type CircuitSystem struct {
	Nodes map[string]*CircuitNode;
	FileName string;
}

type CircuitNode struct {
	Label 	string;
	Operation string;
	Inputs []*CircuitInput;
	Signal int;
	HasSignal bool;
}

type CircuitInput struct {
	Label 	string;
	Literal int;
}

func (this *CircuitNode) SetSignal(val int) int {
	this.Signal  = val;
	this.HasSignal = true;
	return val;
}

func (this *CircuitNode) ToString() string {
	if(this.Operation == CircuitGateRaw){
		return fmt.Sprintf("%s -> %s", this.Inputs[0].ToString(), this.Label);
	}
	if(len(this.Inputs) == 1){
		return fmt.Sprintf("%s %s -> %s", this.Operation, this.Inputs[0].ToString(), this.Label);
	} else if(len(this.Inputs) == 2){
		return fmt.Sprintf("%s %s %s -> %s", this.Inputs[0].ToString(), this.Operation, this.Inputs[1].ToString(), this.Label);
	}
	return "Malformed Node";
}

func (this *CircuitInput) ToString() string {
	if(this.Label != ""){
		return this.Label;
	} else{
		return fmt.Sprintf("%d", this.Literal);
	}
}

func (this *CircuitSystem) PrintCircuitLayout() string {
	var str strings.Builder;
	for _, node := range this.Nodes{
		str.WriteString(node.ToString());
		str.WriteString("\n");
	}
	return str.String();
}

func (this *CircuitSystem) CalculateSignal(label string) int {
	//Log.Info("Calculate %s", label);
	node, exists := this.Nodes[label];
	if(!exists){
		Log.Fatal("Couldn't find a node by label " + label);
		return -1;
	}
	if(node.HasSignal){
		return node.Signal;
	}
	op := node.Operation;
	switch(op){
		case CircuitGateRaw:
			input := node.Inputs[0];
			val := input.Literal;
			if(input.Label != ""){
				val =  this.CalculateSignal(input.Label);
			}
			return node.SetSignal(val);
		case CircuitGateNOT:
			input := node.Inputs[0];
			val := input.Literal;
			if(input.Label != ""){
				val = this.CalculateSignal(input.Label);
			}
			return node.SetSignal(int(^(uint16(val))));
		case CircuitGateAnd:
			inputA := node.Inputs[0];
			valA := inputA.Literal;
			if(inputA.Label != ""){
				valA = this.CalculateSignal(inputA.Label);
			}
			inputB := node.Inputs[1];
			valB := inputB.Literal;
			if(inputB.Label != ""){
				valB = this.CalculateSignal(inputB.Label);
			}
			return node.SetSignal(valA & valB);
		case CircuitGateOR:
			inputA := node.Inputs[0];
			valA := inputA.Literal;
			if(inputA.Label != ""){
				valA = this.CalculateSignal(inputA.Label);
			}
			inputB := node.Inputs[1];
			valB := inputB.Literal;
			if(inputB.Label != ""){
				valB = this.CalculateSignal(inputB.Label);
			}
			return node.SetSignal(valA | valB);
		case CircuitGateLshift:
			inputA := node.Inputs[0];
			valA := inputA.Literal;
			if(inputA.Label != ""){
				valA = this.CalculateSignal(inputA.Label);
			}
			inputB := node.Inputs[1];
			valB := inputB.Literal;
			if(inputB.Label != ""){
				valB = this.CalculateSignal(inputB.Label);
			}
			return node.SetSignal(valA << valB);
		case CircuitGateRshift:
			inputA := node.Inputs[0];
			valA := inputA.Literal;
			if(inputA.Label != ""){
				valA = this.CalculateSignal(inputA.Label);
			}
			inputB := node.Inputs[1];
			valB := inputB.Literal;
			if(inputB.Label != ""){
				valB = this.CalculateSignal(inputB.Label);
			}
			return node.SetSignal(valA >> valB);

	}
	Log.Info("Unhandled operation");
	return -1;
}

func (this *CircuitSystem) Init(fileName string) error {

	this.FileName = fileName;
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()

	this.Nodes = make(map[string]*CircuitNode);

	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			parts := strings.Split(line, "->");
			output := strings.TrimSpace(parts[len(parts) - 1]);
			_, exists := this.Nodes[output];
			if(exists) {
				return errors.New("Duplicate output label " + output);
			}
			node := &CircuitNode{};
			node.Label = output;
			node.Inputs = make([]*CircuitInput, 0);
			inputString := strings.TrimSpace(parts[0]);
			inputStringParts := strings.Split(inputString, " ");

			if(len(inputStringParts) == 1){ // Raw transfer
				node.Operation = CircuitGateRaw;
				input := &CircuitInput{};
				v := inputStringParts[0];
				literalInput, err := strconv.ParseInt(v, 10, 64);
				if(err == nil){
					input.Literal = int(literalInput);
				} else{
					input.Label = v;
				}
				node.Inputs = append(node.Inputs, input);
			} else if(len(inputStringParts) == 2){ // Operation is in first slot. Inputs is in second
				node.Operation = inputStringParts[0];
				input := &CircuitInput{};
				v := inputStringParts[1];
				literalInput, err := strconv.ParseInt(v, 10, 64);
				if(err == nil){
					input.Literal = int(literalInput);
				} else{
					input.Label = v;
				}
				node.Inputs = append(node.Inputs, input);

			} else if(len(inputStringParts) == 3){ // Operation is in second slot. Inputs

				node.Operation = inputStringParts[1];
				input := &CircuitInput{};

				// Arg 1
				v := inputStringParts[0];
				literalInput, err := strconv.ParseInt(v, 10, 64);
				if(err == nil){
					input.Literal = int(literalInput);
				} else{
					input.Label = v;
				}
				node.Inputs = append(node.Inputs, input);

				// Arg 2
				input2 := &CircuitInput{};
				v = inputStringParts[2];
				literalInput, err = strconv.ParseInt(v, 10, 64);
				if(err == nil){
					input2.Literal = int(literalInput);
				} else{
					input2.Label = v;
				}
				node.Inputs = append(node.Inputs, input2);
			}
			this.Nodes[node.Label] = node;
		}
	}

	Log.Info("Successfully parsed circuit system from %s - contains %d nodes", this.FileName, len(this.Nodes));

	return nil;
}

