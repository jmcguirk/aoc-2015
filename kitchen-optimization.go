package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type KitchenOptimization struct {
	Ingredients map[string]*KitchenIngredient;
	SimulationStep int;
	FileName string;
	BestScore int;
	BestAllocation map[string]int;
}

type KitchenIngredient struct {
	Label string;
	Capacity int;
	Durability int;
	Flavor int;
	Texture int;
	Calories int;
}

func (this *KitchenIngredient) Describe() string {
	return fmt.Sprintf("%s: capacity %d, durability %d, flavor %d, texture %d, calories %d", this.Label, this.Capacity, this.Durability, this.Flavor, this.Texture, this.Calories);
}

func (this *KitchenOptimization) Optimize(amount int) (int) {
	return this.OptimizeWithCalories(amount, -1);
}

func (this *KitchenOptimization) OptimizeWithCalories(amount int, calorieTarget int) (int) {
	flat := make([]*KitchenIngredient, 0);
	state := make(map[*KitchenIngredient]int);
	this.BestAllocation = make(map[string]int);
	for _, in := range this.Ingredients{
		flat = append(flat, in);
		state[in] = 0;
	}
	this.OptimizeRecur(amount, 0, calorieTarget, flat, state);
	return this.BestScore;
}

// Note - this is a very inefficient approach - as a first pass optimization, I'd probably only consider "remaining" ingredients on the state
func (this *KitchenOptimization) OptimizeRecur(amt int, index int, calorieTarget int, flat []*KitchenIngredient, state map[*KitchenIngredient]int) {
	if(index >= len(flat)){
		totalWeight := 0;
		totalCapacity := 0;
		totalDurability := 0;
		totalFlavor := 0;
		totalTexture := 0;
		totalCalories := 0;
		for ing, v := range state{
			totalWeight += v;
			totalCapacity += ing.Capacity * v;
			totalDurability += ing.Durability * v;
			totalFlavor += ing.Flavor * v;
			totalTexture += ing.Texture * v;
			totalCalories += ing.Calories * v;
		}
		if(totalWeight != 100){
			return;
		}
		if(calorieTarget > 0 && calorieTarget != totalCalories){
			return;
		}
		if(totalCapacity <= 0 || totalDurability <= 0 || totalFlavor <= 0 || totalTexture <= 0){
			return;
		}
		score := totalCapacity * totalDurability * totalFlavor * totalTexture;
		if(score > this.BestScore){
			this.BestScore = score;
			for ing, v := range state{
				this.BestAllocation[ing.Label] = v;
			}
		}
	} else{
		for i:= 0; i <= amt; i++{
			state[flat[index]] = i;
			this.OptimizeRecur(amt, index+1, calorieTarget, flat, state);
		}
	}
}

func (this *KitchenOptimization) Init(fileName string) error {

	this.FileName = fileName;
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()

	this.Ingredients = make(map[string]*KitchenIngredient);

	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			lineParts := strings.Split(line, " ");
			ingredient := &KitchenIngredient{};
			ingredient.Label = strings.Split(lineParts[0],":")[0];

			//Sprinkles: capacity 5, durability -1, flavor 0, texture 0, calories 5

			capacity, err := strconv.ParseInt(strings.Split(lineParts[2],",")[0], 10, 64);
			if(err != nil){
				return err;
			}
			ingredient.Capacity = int(capacity);

			durability, err := strconv.ParseInt(strings.Split(lineParts[4],",")[0], 10, 64);
			if(err != nil){
				return err;
			}
			ingredient.Durability = int(durability);


			flavors, err := strconv.ParseInt(strings.Split(lineParts[6],",")[0], 10, 64);
			if(err != nil){
				return err;
			}
			ingredient.Flavor = int(flavors);

			texture, err := strconv.ParseInt(strings.Split(lineParts[8],",")[0], 10, 64);
			if(err != nil){
				return err;
			}
			ingredient.Texture = int(texture);

			calories, err := strconv.ParseInt(strings.Split(lineParts[10],",")[0], 10, 64);
			if(err != nil){
				return err;
			}
			ingredient.Calories = int(calories);

			this.Ingredients[ingredient.Label] = ingredient;
		}
	}

	Log.Info("Successfully parsed kitchen problem from %s - contains %d ingredients", this.FileName, len(this.Ingredients));

	for _, r := range this.Ingredients{
		Log.Info(r.Describe());
	}

	return nil;
}

