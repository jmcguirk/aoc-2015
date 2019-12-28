package main

type Problem22A struct {

}

func (this *Problem22A) Solve() {
	Log.Info("Problem 22A solver beginning!")

	rpg := &RPGSimulation{};
	rpg.Init();

	/*
	player := rpg.AddPlayer(10, 0, 0);
	player.InitialMana = 250;
	rpg.AddBoss(13, 8, 0);

	spells := make([]string, 0);
	spells = append(spells, spellPoison);
	spells = append(spells, spellMagicMissle);
	//spells = append(spells, spellMagicMissle);

	commands := make([]*RPGCommand, 0);
	for _, spell := range spells{
		cmd := &RPGCommand{};
		cmd.SpellName = spell;
		commands = append(commands, cmd);
	}
	rpg.CombatLog = false;
	rpg.SimulateSpellFightWithCommandLog(commands);*/


	player := rpg.AddPlayer(50, 0, 0);
	player.InitialMana = 500;
	rpg.AddBoss(71, 10, 0);

	//rpg.CombatLog = true;
	//rpg.Player.

	rpg.SimulateSpellFight();

}