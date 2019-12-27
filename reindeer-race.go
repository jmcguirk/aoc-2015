package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ReindeerRace struct {
	Reindeer map[string]*Reindeer;
	SimulationStep int;
	FileName string;
}

type Reindeer struct {
	RestPeriodRequired int;
	RunSpeed int;
	RunEndurance int;
	Name string;

	Position int;
	State int;
	RemainingFrames int;
	PointsEarned int;
}

func (this *Reindeer) Describe() string {
	return fmt.Sprintf("%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", this.Name, this.RunSpeed, this.RunEndurance, this.RestPeriodRequired);
}

func (this *ReindeerRace) Simulate(maxFrames int) {
	currFrame := 0;
	for{
		if(currFrame > maxFrames){
			break;
		}
		furthestOnFrame := int(math.MinInt64);
		for _, r := range this.Reindeer{
			if(r.State == 1){
				r.Position += r.RunSpeed;
			}
			if(r.Position > furthestOnFrame){
				furthestOnFrame = r.Position;
			}
			r.RemainingFrames--;
			if(r.RemainingFrames == 0){
				if(r.State == 0){
					r.RemainingFrames = r.RunEndurance;
					r.State = 1;
				} else{
					r.RemainingFrames = r.RestPeriodRequired;
					r.State = 0;
				}
			}
		}
		for _, r := range this.Reindeer{
			if(r.Position >= furthestOnFrame){
				r.PointsEarned++;
			}
		}
		currFrame++;
	}
	Log.Info("Race complete - after %d frames: ", currFrame)
	for k, v := range this.Reindeer{
		Log.Info("%s - %d km - Points Earned %d", k, v.Position, v.PointsEarned);
	}
}

func (this *ReindeerRace) Init(fileName string) error {

	this.FileName = fileName;
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()

	this.Reindeer = make(map[string]*Reindeer);

	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			lineParts := strings.Split(line, " ");
			reindeer := &Reindeer{};
			reindeer.Name = lineParts[0];

			speed, err := strconv.ParseInt(lineParts[3], 10, 64);
			if(err != nil){
				return err;
			}
			reindeer.RunSpeed = int(speed);

			end, err := strconv.ParseInt(lineParts[6], 10, 64);
			if(err != nil){
				return err;
			}
			reindeer.RunEndurance = int(end);

			rest, err := strconv.ParseInt(lineParts[13], 10, 64);
			if(err != nil){
				return err;
			}
			reindeer.RestPeriodRequired = int(rest);

			reindeer.State = 1;
			reindeer.Position = 0;
			reindeer.RemainingFrames = reindeer.RunEndurance;

			this.Reindeer[reindeer.Name] = reindeer;
		}
	}

	Log.Info("Successfully parsed reindeer from %s - contains %d reindeer", this.FileName, len(this.Reindeer));

	for _, r := range this.Reindeer{
		Log.Info(r.Describe());
	}

	return nil;
}

