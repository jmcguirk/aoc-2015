package main

type Problem16B struct {

}

func (this *Problem16B) Solve() {

	Log.Info("Starting Problem 16B");
	senderSystem := &MysterySenderSystem{};
	err := senderSystem.Init("source-data/input-day-16b.txt");
	if(err != nil){
		Log.FatalError(err);
	}

	data := make(map[string]int);
	data["children"] = 3
	data["cats"] = 7
	data["samoyeds"] = 2
	data["pomeranians"] = 3
	data["akitas"] = 0
	data["vizslas"] = 0
	data["goldfish"] = 5
	data["trees"] = 3
	data["cars"] = 2
	data["perfumes"] = 1

	sender := senderSystem.QueryWithRanges(data);
	Log.Info("Deduced sender - %s", sender.Label)
}