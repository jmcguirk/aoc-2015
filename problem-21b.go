package main

import (
	"math"
	"strings"
)

type Problem21B struct {

}

func (this *Problem21B) Solve() {
	Log.Info("Problem 21B solver beginning!")

	rpg := &RPGSimulation{};
	rpg.Init();
	rpg.AddBoss(100, 8, 2);
	equipment := make([]*ShopItem, 4);

	weapons := rpg.Weapons;
	armor := rpg.Armor;
	rings := rpg.Rings;



	equipmentArray := make([]int, 4);
	odometerMax := make([]int, 4);
	odometerMax[0] = len(weapons); // 0 is weap, which is required
	odometerMax[1] = len(armor) + 1; // 1 is armor, which is optional
	odometerMax[2] = len(rings) + 1; // 2 is ring 1, which is optional
	odometerMax[3] = len(rings) + 1; // 3 is ring 2, which is optional

	bestCost := int(math.MinInt64);

	for{
		atLim := false;
		for j := len(equipmentArray) - 1; j >= 0; j--{
			if(equipmentArray[j] + 1 < odometerMax[j]){
				equipmentArray[j]++;
				break;
			} else{
				if(j == 0){
					atLim = true;
					break;
				}
				equipmentArray[j] = 0;
			}
		}
		if(atLim){
			Log.Info("Odomoter hit max lim");
			break;
		}
		for i, v := range equipmentArray {
			if(i == 0){
				equipment[0] = weapons[v];
			} else if(i == 1){
				if(v == 0){
					equipment[1] = nil;
				} else {
					equipment[1] = armor[v-1];
				}
			} else{
				if(v == 0){
					equipment[i] = nil;
				} else {
					equipment[i] = rings[v-1];
				}
			}
		}
		if(equipment[2] != nil && equipment[3] != nil && equipment[2] == equipment[3]){
			continue;
		}
		cost := 0;
		for _, eq := range equipment {
			if(eq != nil){
				cost += eq.Cost;
			}
		}
		var buff strings.Builder;
		for _, v := range equipment{
			if(v != nil){
				buff.WriteString(v.Name);
				buff.WriteString( " ");
			}
		}
		if(!rpg.Simulate(equipment)){
			//Log.Info("Win %d %s", cost, buff.String())
			if(cost > bestCost){

				Log.Info("New best cost %d - with equipment %s", cost, buff.String());
				bestCost = cost;
			}
		} else{
			//Log.Info("Lose %d %s", cost, buff.String())
		}
	}

}