package main

type Problem20A struct {

}

func (this *Problem20A) Solve() {
	Log.Info("Problem 20A solver beginning!")


	targetSum := 29000000;
	for i := 1; i <= targetSum; i++{
		sum := i * 11;
		for j := 1; j <= i/2; j++{
			if(i % j == 0){
				if(j * 50 <= i){
					sum += j * 11;
				}
			}
		}
		if(sum >= targetSum){
			Log.Info("First index occurs at %d", i);
			break;
		} else if(i % 100000 == 0){
			Log.Info("%d - %d", i, sum);
		}
	}
	// [INFO] 1.15h - 1500000 - 19001807
	//[INFO] 1.30h - 1600000 - 19600218
	//[INFO] 1.41h - First index occurs at 1663200
	// 887040 too high
	// 1663200 too high
}