package main

type Problem14B struct {

}

func (this *Problem14B) Solve() {
	Log.Info("Problem 14B solver beginning!")

	race := &ReindeerRace{};
	err := race.Init("source-data/input-day-14b.txt");
	if(err != nil){
		Log.FatalError(err);
	}

	race.Simulate(2502);
}