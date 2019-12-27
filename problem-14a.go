package main

type Problem14A struct {

}

func (this *Problem14A) Solve() {
	Log.Info("Problem 14A solver beginning!")

	race := &ReindeerRace{};
	err := race.Init("source-data/input-day-14a.txt");
	if(err != nil){
		Log.FatalError(err);
	}

	race.Simulate(2503);
}