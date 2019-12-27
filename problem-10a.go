package main

type Problem10A struct {

}

func (this *Problem10A) Solve() {

	Log.Info("Starting problem 10 A");

	//iv := make([]int, 0);
	//iv = append(iv, 1);

	iv := IntToDigitArray(3113322113);

	state := iv;
	currIter := 0;
	maxIter := 40;
	for {
		if(currIter >= maxIter){
			// 146784 too low
			// 191596 too low
			Log.Info("Completed %d iterations - len is %d", maxIter, len(state));
			break;
		}
		currIter++;
		new := SpeakAndSayArray(state);
		//Log.Info("%d .) %d becomes %d", currIter, ArrayToInt(state), ArrayToInt(new));
		state = new;
	}
}