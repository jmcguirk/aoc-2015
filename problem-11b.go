package main

import (
	"fmt"
)

type Problem11B struct {

}

func (this *Problem11B) Solve() {

	Log.Info("Starting Problem 11B");

	maxOdometerReading := 26;

	indexArr := this.LoadOdometer("vzbxkghb");

	remaining := 2;
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
			remaining--;
			Log.Info("Next valid password is %s", newPass);
			if(remaining <= 0){
				break;
			}
		}

	}

}

func (this *Problem11B) LoadOdometer(password string) []int{
	indexArr := make([]int, len(password));
	for i, c := range password {
		indexArr[i] = int(c) - int('a');
	}
	return indexArr;
}

func (this *Problem11B) TestPassword(password string) bool{
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