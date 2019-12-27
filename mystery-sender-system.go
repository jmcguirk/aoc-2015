package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MysterySenderSystem struct {
	PotentialSenders map[string]*MysterySender;
	FileName string;
}

type MysterySender struct {
	Label string;
	KnownInformation map[string]int;
}

func (this *MysterySenderSystem) Query(KnownInformation map[string]int) *MysterySender {
	for _, candidate := range this.PotentialSenders{
		allValid := true;
		for k, v := range KnownInformation{
			value, exists := candidate.KnownInformation[k];
			if(!exists){
				continue;
			}
			if(value != v){
				allValid = false;
				break;
			}
		}
		if(allValid){
			return candidate;
		}
	}
	return nil;
}

func (this *MysterySenderSystem) QueryWithRanges(KnownInformation map[string]int) *MysterySender {
	for _, candidate := range this.PotentialSenders{
		allValid := true;
		for k, v := range KnownInformation{
			value, exists := candidate.KnownInformation[k];
			if(!exists){
				continue;
			}
			if(k == "cats" || k == "trees"){
				if(value <= v){
					allValid = false;
					break;
				}
			} else if(k == "pomeranians" || k == "goldfish"){
				if(value >= v){
					allValid = false;
					break;
				}
			} else if(value != v){
				allValid = false;
				break;
			}
		}
		if(allValid){
			return candidate;
		}
	}
	return nil;
}

func (this *MysterySenderSystem) Init(fileName string) error {

	this.FileName = fileName;
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()

	this.PotentialSenders = make(map[string]*MysterySender);

	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			//Log.Info(line);
			parts := strings.Split(line, " ");
			label := parts[0] + " " + parts[1];

			rest := strings.TrimSpace(strings.Replace(line, label, "", -1));

			candidate := &MysterySender{};
			candidate.KnownInformation = make(map[string]int);
			candidate.Label = label[0:len(label)-1];
			knownParts := strings.Split(rest, ",");
			for _, v := range knownParts{
				infoParts := strings.Split(v,":");
				infoVal, err := strconv.ParseInt(strings.TrimSpace(infoParts[1]), 10, 64);
				if(err != nil){
					return err;
				}
				candidate.KnownInformation[strings.TrimSpace(infoParts[0])] = int(infoVal);
			}
			this.PotentialSenders[candidate.Label] = candidate;
			//Log.Info(candidate.Label);
		}
	}

	Log.Info("Successfully parsed kitchen problem from %s - contains %d potential senders", this.FileName, len(this.PotentialSenders));

	return nil;
}

