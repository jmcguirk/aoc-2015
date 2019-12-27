package main

import (
	"fmt"
)

type Problem11A struct {

}

func (this *Problem11A) Solve() {

	Log.Info("Starting Problem 11 A");

	maxOdometerReading := 26;

	//vzbxkghb
	/*
	indexArr := make([]int, 8);
	indexArr[0] = 'v' - int('a');
	indexArr[1] = 'z' - int('a');
	indexArr[2] = 'b' - int('a');
	indexArr[3] = 'x' - int('a');
	indexArr[4] = 'k' - int('a');
	indexArr[5] = 'g' - int('a');
	indexArr[6] = 'h' - int('a');
	indexArr[7] = 'b' - int('a');*/

	indexArr := this.LoadOdometer("vzbxkghb");

	for{
		atLim := false;
		for j := len(indexArr) - 1; j >= 0; j--{
			if(indexArr[j] + 1 < maxOdometerReading){
				indexArr[j]++;
				break;
			} else{
				if(j == 0){
					atLim = true;
					break;
				}
				indexArr[j] = 0;
			}
		}
		if(atLim){
			Log.Info("Odomoter hit max lim");
			break;
		}
		newPass := "";
		for _, v := range indexArr{
			newPass += fmt.Sprintf("%c", v + int('a'));
		}
		if(this.TestPassword(newPass)){
			Log.Info("Next valid password is %s", newPass);
			break;
		}

	}

}

func (this *Problem11A) LoadOdometer(password string) []int{
	indexArr := make([]int, len(password));
	for i, c := range password {
		indexArr[i] = int(c) - int('a');
	}
	return indexArr;
}

func (this *Problem11A) TestPassword(password string) bool{
	containsStripe := false;
	counts := make(map[int32]int);
	pwLen := len(password);
	for i, c := range password{
		if(c == 'i' || c == 'o' || c == 'l'){
			return false;
		}
		val := int(c);
		if(i > 0){
			if(int(password[i-1]) == val){ // Previous was a match
				if(i == 1 || int(password[i-2]) != val){ // 2nd Previous is explicitly not a match
					if(i == pwLen - 1 || int(password[i+1]) != val){ // Forward is explicitly not a match
						counts[c] = 1;
					}
				}
			}
		}
		if(!containsStripe && i > 1){

			if(int(password[i-1]) == val - 1 && int(password[i-2]) == val - 2){
				containsStripe = true;
			}
		}
	}
	if(!containsStripe){
		return false;
	}

	if(len(counts) < 2){
		return false;
	}


	return true;
}