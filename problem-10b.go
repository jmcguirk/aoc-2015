package main

type Problem10B struct {

}

func (this *Problem10B) Solve() {

	Log.Info("Starting problem 10B");

	//iv := make([]int, 0);
	//iv = append(iv, 1);

	iv := IntToDigitArray(3113322113);

	state := iv;
	currIter := 0;
	maxIter := 50;
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