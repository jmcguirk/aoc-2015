package main

import "math"

type RPGSimulation struct {
	Actors []*RPGActor;
	Player *RPGActor;
	Boss *RPGActor;
	ShopContents []*ShopItem;
	Weapons []*ShopItem;
	Armor []*ShopItem;
	Rings []*ShopItem;
	SpellBook map[string]*RPGSpell;
	TurnNumber int;
	CombatLog bool;
	BestManaScore int;
	HPDecay int;
}

const itemTypeWeapon = "Weapon";
const itemTypeArmor = "Armor";
const itemTypeRings = "Rings";

type RPGActor struct {
	IsPlayer bool;
	CurrentHP int;
	Damage int;
	Armor int;
	InitialDamage int;
	InitialArmor int;
	InitialHP int;
	CurrentMana int;
	InitialMana int;
	ActiveEffects []*RPGEffect;
	SpentMana int;
	IsActive bool;
}

func (this *RPGActor) Clone() *RPGActor {
	res := &RPGActor{};
	res.IsPlayer = this.IsPlayer;
	res.CurrentHP = this.CurrentHP;
	res.Damage = this.Damage;
	res.Armor = this.Armor;
	res.InitialDamage = this.InitialDamage;
	res.InitialArmor = this.InitialArmor;
	res.InitialHP = this.InitialHP;
	res.CurrentMana = this.CurrentMana;
	res.InitialMana = this.InitialMana;
	res.SpentMana = this.SpentMana;
	res.ActiveEffects = make([]*RPGEffect, len(this.ActiveEffects));
	for i, effect := range this.ActiveEffects{
		res.ActiveEffects[i] = effect.Clone();
		res.ActiveEffects[i].Actor = this;
	}
	res.IsActive = this.IsActive;
	return res;
}

func (this *RPGEffect) Clone() *RPGEffect {
	res := &RPGEffect{};
	res.RemainingTurns = this.RemainingTurns;
	res.Spell = this.Spell;
	res.ShouldLog = this.ShouldLog;
	//res.Actor = this.Actor;
	return res;
}

type RPGEffect struct {
	Actor *RPGActor;
	RemainingTurns int;
	ShouldLog bool;
	Spell *RPGSpell;
}

type RPGState struct{
	Player *RPGActor;
	Boss *RPGActor;
	TurnNum int;
	CommandQueue []*RPGCommand;
}

type RPGSpell struct {
	SpellName string;
	Magnitude int;
	CastingCost int;
	Duration int;
}

type RPGCommand struct {
	SpellName string;
}

type ShopItem struct {
	Name string;
	Type string;
	Cost int;
	Armor int;
	Damage int;
}

func (this *RPGSimulation) Init() {
	this.Actors = make([]*RPGActor, 0);
	this.ShopContents = make([]*ShopItem, 0);
	this.Weapons = make([]*ShopItem, 0);
	this.Armor = make([]*ShopItem, 0);
	this.Rings = make([]*ShopItem, 0);
	this.SpellBook = make(map[string]*RPGSpell);
	this.PopulateShop();
	this.PopulateSpellBook();
}

const spellMagicMissle = "Magic Missle"
const spellDrain = "Drain"
const spellShield = "Shield"
const spellPoison = "Poison"
const spellRecharge = "Recharge"

func (this *RPGSimulation) PopulateSpellBook() {
	this.AddSpell(spellMagicMissle, 53, 0, 4);
	this.AddSpell(spellDrain, 73, 0, 2);
	this.AddSpell(spellShield, 113, 6, 7);
	this.AddSpell(spellPoison, 173, 6, 3);
	this.AddSpell(spellRecharge, 229, 5, 101);
}

func (this *RPGSimulation) AddSpell(spellName string, cost int, duration int, magnitude int) {
	spell := &RPGSpell{};
	spell.SpellName = spellName;
	spell.CastingCost = cost;
	spell.Duration = duration;
	spell.Magnitude = magnitude;
	this.SpellBook[spell.SpellName] = spell;
}

func (this *RPGSimulation) AddPlayer(currentHp int, damage int, armor int) *RPGActor {
	actor := &RPGActor{};
	actor.IsPlayer = true;
	actor.InitialHP = currentHp;
	actor.CurrentHP = currentHp;
	actor.Damage = damage;
	actor.Armor = armor;
	actor.InitialArmor = armor;
	actor.InitialDamage = damage;
	actor.ActiveEffects = make([]*RPGEffect, 0);
	this.Actors = append(this.Actors, actor);
	this.Player = actor;
	return actor;
}

func (this *RPGSimulation) AddBoss(currentHp int, damage int, armor int) *RPGActor {
	actor := &RPGActor{};
	actor.IsPlayer = false;
	actor.InitialHP = currentHp;
	actor.CurrentHP = currentHp;
	actor.Damage = damage;
	actor.Armor = armor;
	actor.InitialArmor = armor;
	actor.InitialDamage = damage;
	this.Actors = append(this.Actors, actor);
	this.Boss = actor;
	return actor;
}

func (this *RPGActor) Reset() {
	this.CurrentHP = this.InitialHP;
	this.Armor = this.InitialArmor;
	this.Damage = this.InitialDamage;
	this.CurrentMana = this.InitialMana;
}

func (this *RPGActor) Dead() bool {
	return this.CurrentHP <= 0;
}

func (this *RPGActor) ApplyDamage(target *RPGActor) int {
	effectiveDamage := this.Damage - target.Armor;
	if(effectiveDamage < 1){
		effectiveDamage = 1;
	}
	target.CurrentHP -= effectiveDamage;
	return effectiveDamage;
}

func (this *RPGActor) Equip(equipment []*ShopItem) {
	for _, eq := range equipment{
		if(eq == nil){
			continue;
		}
		this.Armor += eq.Armor;
		this.Damage += eq.Damage;
	}
}



func (this *RPGSimulation) SimulateSpellFight() {
	this.SimulateSpellFightWithCommandLog(nil);
}

func (this *RPGSimulation) LogWin(state *RPGState) {
	//Log.Info("win");
	if(state.Player.SpentMana < this.BestManaScore){
		this.BestManaScore = state.Player.SpentMana;
		Log.Info("Found new best mana score %d after turns %d", this.BestManaScore, state.TurnNum)
	}
}

func (this *RPGSimulation) SimulateSpellFightWithCommandLog(commandLog []*RPGCommand) {
	this.BestManaScore = int(math.MaxInt64);

	state := &RPGState{};
	state.Player = this.Player.Clone();
	state.Boss = this.Boss.Clone();
	state.Player.Reset();
	state.Boss.Reset();
	state.CommandQueue = commandLog;
	this.AdvanceState(state);
}

func (this *RPGActor) UpdateEffects() {
	newEffects := make([]*RPGEffect, 0);
	for _, effect := range this.ActiveEffects{
		effect.Actor = this;
		effect.Apply();
		effect.RemainingTurns--;
		if(effect.RemainingTurns > 0){
			newEffects = append(newEffects, effect);
		} else{
			effect.End();
		}
	}
	this.ActiveEffects = newEffects;
}

func (this *RPGActor) AddEffect(spell *RPGSpell, shouldlog bool) {
	effect := &RPGEffect{};
	effect.Spell = spell;
	effect.RemainingTurns = spell.Duration;
	effect.Actor = this;
	effect.ShouldLog = shouldlog;
	this.ActiveEffects = append(this.ActiveEffects, effect);
	effect.Begin();
}

func (this *RPGEffect) Begin() {
	if(this.Spell.SpellName == spellShield){
		this.Actor.Armor += this.Spell.Magnitude;
		//Log.Info("applying shield armor now %d", this.Actor.Armor);
	}
}

func (this *RPGEffect) End() {
	//Log.Info(this.Spell.SpellName + " ended")
	if(this.Spell.SpellName == spellShield){
		this.Actor.Armor -= this.Spell.Magnitude;
	}
}

func (this *RPGEffect) Apply() {
	if(this.Spell.SpellName == spellPoison){
		this.Actor.CurrentHP -= this.Spell.Magnitude;
		if(this.ShouldLog){
			Log.Info("Poison Deals %d Damage. Its Timer Is Now %d - HP is %d", this.Spell.Magnitude, this.RemainingTurns, this.Actor.CurrentHP)
		}
	} else if (this.Spell.SpellName == spellRecharge) {
		this.Actor.CurrentMana += this.Spell.Magnitude;
		if(this.ShouldLog){
			Log.Info("Recharge provides %d mana its timer is now %d", this.Spell.Magnitude, this.RemainingTurns)
		}

	}
}

func (this *RPGState) Clone() *RPGState {
	res := &RPGState{};
	res.CommandQueue = make([]*RPGCommand, 0);
	res.CommandQueue = append(res.CommandQueue, this.CommandQueue...);
	res.Player = this.Player.Clone();
	res.Boss = this.Boss.Clone();
	res.TurnNum = this.TurnNum;
	return res;
}

func (this *RPGSimulation) GeneratePossibleCommands(state *RPGState) []*RPGCommand {
	res := make([]*RPGCommand, 0);
	for k, v := range this.SpellBook{
		if(v.CanBeCasted(state)){
			cmd := &RPGCommand{};
			cmd.SpellName = k;
			res = append(res, cmd);
		}
	}
	return res;
}

func (this *RPGSpell) CanBeCasted(state *RPGState) bool {
	if(state.Player.CurrentMana < this.CastingCost){
		return false;
	}
	switch(this.SpellName){
		case spellRecharge:
			if(state.Player.HasActiveEffect(this)){
				return false;
			}
			break;
		case spellPoison:
			if(state.Boss.HasActiveEffect(this)){
				return false;
			}
			break;
		case spellShield:
			if(state.Player.HasActiveEffect(this)){
				return false;
			}
			break;
	}
	return true;
}

func (this *RPGActor) HasActiveEffect(spell *RPGSpell) bool {
	for _, effect := range this.ActiveEffects{
		if(effect.Spell == spell){
			return true;
		}
	}
	return false;
}

func (this *RPGSimulation) AdvanceState(state *RPGState) bool {
	state.Player.CurrentHP -= this.HPDecay;
	if(state.Player.Dead()){
		return false;
	}
	if(state.Boss.Dead()){
		this.LogWin(state);
		return true;
	}
	if(this.CombatLog){
		Log.Info("-- Player Turn --");
	}
	state.Player.UpdateEffects();
	state.Boss.UpdateEffects();
	if(this.CombatLog){
		Log.Info("- Player has %d hit points, %d armor, %d mana", state.Player.CurrentHP, state.Player.Armor, state.Player.CurrentMana);
		Log.Info("- Boss has %d hit points", state.Boss.CurrentHP);

	}


	if(state.Player.Dead()){
		return false;
	}
	if(state.Boss.Dead()){
		this.LogWin(state);
		return true;
	}


	possibleCommands := make([]*RPGCommand, 0);
	if(len(state.CommandQueue) > 0){
		possibleCommands = append(possibleCommands, state.CommandQueue[0]);
		state.CommandQueue = state.CommandQueue[1:];
	} else{
		possibleCommands = this.GeneratePossibleCommands(state);
	}
	for _, c := range possibleCommands{
		newState := state.Clone();
		spell, _ := this.SpellBook[c.SpellName];
		if(this.CombatLog){
			Log.Info("Player casts %s.", c.SpellName);
		}
		newState.Player.CurrentMana -= spell.CastingCost;
		newState.Player.SpentMana += spell.CastingCost;
		if(spell.Duration > 0){
			switch(spell.SpellName){
				case spellShield:
					newState.Player.AddEffect(spell, this.CombatLog);
					break;
				case spellRecharge:
					newState.Player.AddEffect(spell, this.CombatLog);
					break;
				case spellPoison:
					newState.Boss.AddEffect(spell, this.CombatLog);
					break;
			}
		} else{
			newState.Boss.CurrentHP -= spell.Magnitude;
			if(spell.SpellName == spellDrain){
				newState.Player.CurrentHP += spell.Magnitude;
			}
		}

		if(newState.Player.Dead()){
			return false;
		}
		if(newState.Boss.Dead()){
			this.LogWin(newState);
			return true;
		}
		if(this.CombatLog){
			Log.Info("-- Boss Turn --");
		}
		if(this.CombatLog){
			Log.Info("- Player has %d hit points, %d armor, %d mana", newState.Player.CurrentHP, newState.Player.Armor, newState.Player.CurrentMana);
			Log.Info("- Boss has %d hit points", newState.Boss.CurrentHP);
		}
		newState.Player.UpdateEffects();
		//Log.Info("- Applying boss effects %d", newState.Boss.CurrentHP);
		newState.Boss.UpdateEffects();
		//Log.Info("- Done boss effects %d", newState.Boss.CurrentHP);
		if(newState.Player.Dead()){
			return false;
		}
		if(newState.Boss.Dead()){
			this.LogWin(newState);
			return true;
		}

		newState.TurnNum++;
		dealt := newState.Boss.ApplyDamage(newState.Player);
		if(this.CombatLog){
			//Log.Info("- Player has %d hit points, %d armor, %d mana", newState.Player.CurrentHP, newState.Player.Armor, newState.Player.CurrentMana);
			Log.Info("Boss Attacks for %d Damage!", dealt);
		}
		this.AdvanceState(newState);
	}
	return false;
}


func (this *RPGSimulation) Simulate(equipment []*ShopItem) bool {
	for _, actor := range this.Actors{
		actor.Reset();
	}
	this.TurnNumber = 1;
	this.Player.Equip(equipment);
	for{
		if(this.Player.Dead()){
			if(this.CombatLog){
				Log.Info("Player determined dead on turn %d", this.TurnNumber);
			}
			return false;
		}
		dealt := this.Player.ApplyDamage(this.Boss);
		if(this.CombatLog){
			Log.Info("%d Player dealt %d-%d = %d damage; the boss goes down to %d hit points.", this.TurnNumber, this.Player.Damage, this.Boss.Armor, dealt, this.Boss.CurrentHP);
		}
		if(this.Boss.Dead()){
			if(this.CombatLog){
				Log.Info("Boss determined dead on turn %d", this.TurnNumber);
			}
			return true;
		}
		dealt = this.Boss.ApplyDamage(this.Player);
		if(this.CombatLog){
			Log.Info("%d Boss dealt %d-%d = %d damage; the player goes down to %d hit points.", this.TurnNumber, this.Boss.Damage, this.Player.Armor, dealt, this.Player.CurrentHP);
		}
		this.TurnNumber++;
	}
}


func (this *RPGSimulation) PopulateShop() {

	// Weapons
	this.AddShopItem("Dagger", itemTypeWeapon, 8, 4, 0);
	this.AddShopItem("Short Sword", itemTypeWeapon, 10, 5, 0);
	this.AddShopItem("Warhammer", itemTypeWeapon, 25, 6, 0);
	this.AddShopItem("Longsword", itemTypeWeapon, 40, 7, 0);
	this.AddShopItem("Greataxe", itemTypeWeapon, 74, 8, 0);

	// Armor
	this.AddShopItem("Leather", itemTypeArmor, 13, 0, 1);
	this.AddShopItem("Chainmail", itemTypeArmor, 31, 0, 2);
	this.AddShopItem("Splintmail", itemTypeArmor, 53, 0, 3);
	this.AddShopItem("Bandedmail", itemTypeArmor, 75, 0, 4);
	this.AddShopItem("Platemail", itemTypeArmor, 102, 0, 5);

	// Rings
	this.AddShopItem("Damage +1", itemTypeRings, 25, 1, 0);
	this.AddShopItem("Damage +2", itemTypeRings, 50, 2, 0);
	this.AddShopItem("Damage +3", itemTypeRings, 100, 3, 0);

	this.AddShopItem("Defense +1", itemTypeRings, 20, 0, 1);
	this.AddShopItem("Defense +2", itemTypeRings, 40, 0, 2);
	this.AddShopItem("Defense +3", itemTypeRings, 80, 0, 3);
}

func (this *RPGSimulation) AddShopItem(name string, itemType string, cost int, damage int, armor int) *ShopItem {
	item := &ShopItem{};
	item.Name = name;
	item.Type = itemType;
	item.Cost = cost;
	item.Armor = armor;
	item.Damage = damage;
	this.ShopContents = append(this.ShopContents, item);
	if(item.Type == itemTypeArmor){
		this.Armor = append(this.Armor, item);
	} else if(item.Type == itemTypeWeapon){
		this.Weapons = append(this.Weapons, item);
	} else if(item.Type == itemTypeRings){
		this.Rings = append(this.Rings, item);
	}
	return item;
}