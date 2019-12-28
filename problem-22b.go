package main

type Problem22B struct {

}

func (this *Problem22B) Solve() {
	Log.Info("Problem 22B solver beginning!")

	rpg := &RPGSimulation{};
	rpg.Init();
	player := rpg.AddPlayer(50, 0, 0);
	player.InitialMana = 500;
	rpg.HPDecay = 1;
	rpg.AddBoss(71, 10, 0);
	rpg.SimulateSpellFight();

}