package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

type EquipmentType int

const (
	Weapon EquipmentType = iota
	Armor
	Ring
)

type BattleOutcome int

const (
	InProgress BattleOutcome = iota
	PlayerLoss
	PlayerWin
)

type Equipment struct {
	Name   string
	Type   EquipmentType
	Cost   int
	Damage int
	Armor  int
}

type Entity struct {
	Name           string
	StartHitPoints int
	HitPoints      int
	StartDamage    int
	Damage         int
	StartArmor     int
	Armor          int
}

func (entity *Entity) String() string {
	deltaDamage := entity.Damage - entity.StartDamage
	deltaArmor := entity.Armor - entity.StartArmor
	return fmt.Sprintf("%s: %d hp, %d dmg (+%d), %d armor (+%d)", entity.Name, entity.HitPoints, entity.Damage, deltaDamage, entity.Armor, deltaArmor)
}

func (entity *Entity) Print() {
	fmt.Println(entity.String())
}

func (entity *Entity) ApplyLoadout(loadout *Loadout) {
	entity.Damage += loadout.Delta.Damage
	entity.Armor += loadout.Delta.Armor
}

type Delta struct {
	Damage int
	Armor  int
}

type Loadout struct {
	Weapon *Equipment
	Armor  *Equipment
	Ring1  *Equipment
	Ring2  *Equipment
	Cost   int
	Delta  Delta
}

func (loadout *Loadout) Calculate() {
	cost := 0
	loadout.Delta = Delta{}

	if loadout.Weapon != nil {
		cost += loadout.Weapon.Cost
		loadout.Delta.Damage += loadout.Weapon.Damage
	}

	if loadout.Armor != nil {
		cost += loadout.Armor.Cost
		loadout.Delta.Armor += loadout.Armor.Armor
	}

	if loadout.Ring1 != nil {
		cost += loadout.Ring1.Cost
		loadout.Delta.Damage += loadout.Ring1.Damage
		loadout.Delta.Armor += loadout.Ring1.Armor
	}

	if loadout.Ring2 != nil {
		cost += loadout.Ring2.Cost
		loadout.Delta.Damage += loadout.Ring2.Damage
		loadout.Delta.Armor += loadout.Ring2.Armor
	}

	loadout.Cost = cost
}

func (loadout *Loadout) String() string {
	return fmt.Sprintf("[%s, %s, %s, %s] $%d, +%d dmg, +%d def", loadout.Weapon.Name, loadout.Armor.Name, loadout.Ring1.Name, loadout.Ring2.Name, loadout.Cost, loadout.Delta.Damage, loadout.Delta.Armor)
}

func (loadout *Loadout) Print() {
	fmt.Printf("%s\n", loadout.String())
}

func main() {
	boolPtr := flag.Bool("test", false, "test mode")
	flag.Parse()

	player := Entity{
		Name:           "Player",
		StartHitPoints: 100,
		StartDamage:    0,
		StartArmor:     0,
	}

	boss := Entity{
		Name:           "Boss",
		StartHitPoints: 109,
		StartDamage:    8,
		StartArmor:     2,
	}

	if *boolPtr {
		boss.StartHitPoints = 12
		boss.StartDamage = 7
		boss.StartArmor = 2

		player.StartHitPoints = 8
		player.StartDamage = 5
		player.StartArmor = 5
	}

	player.reset()
	boss.reset()

	// Init Equipment
	dagger := Equipment{
		Name:   "Dagger",
		Type:   Weapon,
		Cost:   8,
		Damage: 4,
		Armor:  0,
	}
	shortsword := Equipment{
		Name:   "Shortsword",
		Type:   Weapon,
		Cost:   10,
		Damage: 5,
		Armor:  0,
	}

	warhammer := Equipment{
		Name:   "Warhammer",
		Type:   Weapon,
		Cost:   25,
		Damage: 6,
		Armor:  0,
	}

	longsword := Equipment{
		Name:   "Longsword",
		Type:   Weapon,
		Cost:   40,
		Damage: 7,
		Armor:  0,
	}

	greataxe := Equipment{
		Name:   "Greataxe",
		Type:   Weapon,
		Cost:   74,
		Damage: 8,
		Armor:  0,
	}

	weapons := []*Equipment{&dagger, &shortsword, &warhammer, &longsword, &greataxe}
	fmt.Printf("Weapons: %v\n", weapons)

	// Armor
	noArmor := Equipment{
		Name:   "NoArmor",
		Type:   Armor,
		Cost:   0,
		Damage: 0,
		Armor:  0,
	}

	leather := Equipment{
		Name:   "Leather",
		Type:   Armor,
		Cost:   13,
		Damage: 0,
		Armor:  1,
	}

	chainmail := Equipment{
		Name:   "Chainmail",
		Type:   Armor,
		Cost:   31,
		Damage: 0,
		Armor:  2,
	}

	splintmail := Equipment{
		Name:   "Splintmail",
		Type:   Armor,
		Cost:   53,
		Damage: 0,
		Armor:  3,
	}

	bandedmail := Equipment{
		Name:   "Bandedmail",
		Type:   Armor,
		Cost:   75,
		Damage: 0,
		Armor:  4,
	}

	platedmail := Equipment{
		Name:   "Platedmail",
		Type:   Armor,
		Cost:   102,
		Damage: 0,
		Armor:  5,
	}

	armors := []*Equipment{&noArmor, &leather, &chainmail, &splintmail, &bandedmail, &platedmail}
	fmt.Printf("Armors: %v\n", armors)

	// Rings
	noRing1 := Equipment{
		Name:   "NoRing1",
		Type:   Ring,
		Cost:   0,
		Damage: 0,
		Armor:  0,
	}

	noRing2 := Equipment{
		Name:   "NoRing2",
		Type:   Ring,
		Cost:   0,
		Damage: 0,
		Armor:  0,
	}

	damage1 := Equipment{
		Name:   "Damage+1",
		Type:   Ring,
		Cost:   25,
		Damage: 1,
		Armor:  0,
	}

	damage2 := Equipment{
		Name:   "Damage+2",
		Type:   Ring,
		Cost:   50,
		Damage: 2,
		Armor:  0,
	}

	damage3 := Equipment{
		Name:   "Damage+3",
		Type:   Ring,
		Cost:   100,
		Damage: 3,
		Armor:  0,
	}

	defense1 := Equipment{
		Name:   "Defense+1",
		Type:   Ring,
		Cost:   20,
		Damage: 0,
		Armor:  1,
	}

	defense2 := Equipment{
		Name:   "Defense+2",
		Type:   Ring,
		Cost:   40,
		Damage: 0,
		Armor:  2,
	}

	defense3 := Equipment{
		Name:   "Defense+3",
		Type:   Ring,
		Cost:   80,
		Damage: 0,
		Armor:  3,
	}

	rings := []*Equipment{&noRing1, &noRing2, &damage1, &damage2, &damage3, &defense1, &defense2, &defense3}
	fmt.Printf("Rings: %v\n", rings)

	simulate(player, boss)

	// Your Code goes below!
	loadouts := generateAllLoadouts(weapons, armors, rings)
	part1 := math.MaxInt
	part2 := 0
	for _, loadout := range loadouts {
		loadout.Print()
		player.ApplyLoadout(loadout)

		result := simulate(player, boss)
		if result == PlayerWin && loadout.Cost < part1 {
			part1 = loadout.Cost
		}

		if result == PlayerLoss && loadout.Cost > part2 {
			part2 = loadout.Cost
		}

		player.reset()
		boss.reset()
	}

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func generateAllLoadouts(weapons []*Equipment, armors []*Equipment, rings []*Equipment) (loadouts []*Loadout) {
	// 1 weapon required
	for _, weapon := range weapons {
		// armor is optional
		for _, armor := range armors {
			// choose 0-2 rings
			for i, ring1 := range rings {
				for j, ring2 := range rings {
					if i != j {
						loadout := &Loadout{
							Weapon: weapon,
							Armor:  armor,
							Ring1:  ring1,
							Ring2:  ring2,
							Cost:   0,
						}
						loadout.Calculate()
						loadouts = append(loadouts, loadout)
					}
				}
			}
		}
	}
	return
}

func simulate(player Entity, boss Entity) (outcome BattleOutcome) {
	round := 0

	for outcome == InProgress {
		round++
		fmt.Printf("Round %d\n", round)
		// player turn
		player.attack(&boss)
		if boss.HitPoints <= 0 {
			outcome = PlayerWin
			return
		}

		// boss turn
		boss.attack(&player)
		if player.HitPoints <= 0 {
			outcome = PlayerLoss
			return
		}
	}

	return
}

func (attacker *Entity) attack(defender *Entity) {
	attackDamage := max(1, attacker.Damage-defender.Armor)
	defender.HitPoints -= attackDamage
	fmt.Printf("\t%s deals %d-%d = %d damage; %s goes down to %d hit points.\n", attacker.Name, attacker.Damage, defender.Armor, attackDamage, defender.Name, defender.HitPoints)
}

func (entity *Entity) reset() {
	entity.HitPoints = entity.StartHitPoints
	entity.Damage = entity.StartDamage
	entity.Armor = entity.StartArmor
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
