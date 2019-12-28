package main

type Problem20B struct {

}

func (this *Problem20B) Solve() {
	Log.Info("Problem 20B solver beginning!")

	// Submitted, was too high
	targetSum := 29000000;
	for i := 0; i < targetSum; i++{
		sum := i * 11;
		for j := 1; j <= i/2; j++{
			if(i % j == 0){
				if(i/j <= 50){
					sum += j * 11;
				}
			}
		}
		if(sum >= targetSum){
			Log.Info("First index occurs at %d - %d", i, sum);
			break;
		} else if(i % 100000 == 0){
			Log.Info("%d - %d", i, sum);
		}
	}
}