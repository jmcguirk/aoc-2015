package main

type RPGSimulation struct {
	Actors []*RPGActor;
	Player *RPGActor;
	Boss *RPGActor;
	ShopContents []*ShopItem;
	Weapons []*ShopItem;
	Armor []*ShopItem;
	Rings []*ShopItem;
	TurnNumber int;
	CombatLog bool;
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
	this.PopulateShop();
	this.AddPlayer();
}

func (this *RPGSimulation) AddPlayer() *RPGActor {
	actor := &RPGActor{};
	actor.IsPlayer = true;
	actor.InitialHP = 100;
	actor.CurrentHP = 100;
	actor.Damage = 0;
	actor.Armor = 0;
	actor.InitialArmor = 0;
	actor.InitialDamage = 0;
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