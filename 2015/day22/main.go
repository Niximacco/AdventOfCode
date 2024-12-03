package main

import (
	"flag"
	"fmt"
	"math"
)

type BattleOutcome int

const (
	InProgress BattleOutcome = iota
	PlayerLoss
	PlayerWin
)

type Turn int

const (
	PlayerTurn Turn = iota
	BossTurn
)

func (t Turn) GetNextTurn() Turn {
	if t == PlayerTurn {
		return BossTurn
	} else {
		return PlayerTurn
	}
}

type Entity struct {
	Name      string
	HitPoints int
	Damage    int
	Mana      int
	Armor     int
	Cooldowns map[string]int
	Spells    []*Spell
}

type Spell struct {
	Name   string
	Cost   int
	Effect Effect
}

type Effect struct {
	Name     string
	Duration int
	Apply    func(player Entity, boss Entity)
	OnExpire func(player Entity, boss Entity)
}

func main() {
	boolPtr := flag.Bool("test", false, "test mode")
	flag.Parse()

	player := Entity{
		Name:      "Player",
		HitPoints: 50,
		Mana:      500,
	}

	boss := Entity{
		Name:      "Boss",
		HitPoints: 51,
		Damage:    9,
	}

	if *boolPtr {
		boss.HitPoints = 13
		boss.Damage = 8

		player.HitPoints = 10
		player.Mana = 250
	}

	// Init Spells
	magicMissile := Spell{
		Name: "Magic Missile",
		Cost: 53,
		Effect: Effect{
			Name:     "Magic Missile Effect",
			Duration: 0,
			Apply: func(player Entity, boss Entity) {
				//fmt.Println("Magic Missile: Boss -4 hp")
				boss.HitPoints -= 4
			},
		},
	}

	drain := Spell{
		Name: "Drain",
		Cost: 73,
		Effect: Effect{
			Name:     "Drain Effect",
			Duration: 0,
			Apply: func(player Entity, boss Entity) {
				//fmt.Println("Drain Effect: +2 player.hp, -2 boss.hp")
				boss.HitPoints -= 2
				player.HitPoints += 2
			},
		},
	}

	shield := Spell{
		Name: "Shield",
		Cost: 113,
		Effect: Effect{
			Name:     "Shield Effect",
			Duration: 6,
			Apply: func(player Entity, boss Entity) {
				//fmt.Println("Shielded: Armor=7")
				player.Armor = 7
			},
			OnExpire: func(player Entity, boss Entity) {
				fmt.Println("No longer Shielded: Armor=0")
				player.Armor = 0
			},
		},
	}

	poison := Spell{
		Name: "Poison",
		Cost: 173,
		Effect: Effect{
			Name:     "Poison Effect",
			Duration: 6,
			Apply: func(player Entity, boss Entity) {
				//fmt.Println("Poison effect: Boss -3 hp")
				boss.HitPoints -= 3
			},
			OnExpire: nil,
		},
	}

	recharge := Spell{
		Name: "Recharge",
		Cost: 229,
		Effect: Effect{
			Name:     "Recharge Effect",
			Duration: 5,
			Apply: func(player Entity, boss Entity) {
				//fmt.Println("Recharging. +101 Mana")
				player.Mana += 101
			},
			OnExpire: nil,
		},
	}

	spells := []*Spell{&magicMissile, &drain, &shield, &poison, &recharge}
	player.Spells = spells
	player.Cooldowns = make(map[string]int)
	part1 := math.MaxInt

	player.printAvailableSpells()

	_, part1 = simulate(player, boss, 0, PlayerTurn, nil)

	// Your Code goes below!
	part1, part2 := 0, 0

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func simulate(player Entity, boss Entity, manaSpent int, turn Turn, effects []Effect) (outcome BattleOutcome, mana int) {
	for _, effect := range effects {
		effect.Apply(player, boss)
		effect.Duration--
	}

	var newEffects []Effect
	for _, effect := range effects {
		if effect.Duration >= 0 {
			newEffects = append(newEffects, effect)
		}
	}

	if turn == PlayerTurn {
		availableSpells := player.getAvailableSpells()
		if len(availableSpells) == 0 {
			fmt.Printf("No Spells Available, player loses")
			outcome = PlayerLoss
			mana = manaSpent
			return
		}

		resultMana := math.MaxInt
		for _, spell := range availableSpells {
			newPlayer := player
			newBoss := boss

			totalMana := manaSpent + spell.Cost
			if spell.Effect.Duration == 0 {
				// spell is instant, apply now
				spell.Effect.Apply(newPlayer, newBoss)
			} else {
				newEffects = append(newEffects, spell.Effect)
				player.Cooldowns[spell.Name] = spell.Effect.Duration
			}

			result, treeMana := simulate(newPlayer, newBoss, totalMana, BossTurn, newEffects)
			if result == PlayerWin {
				outcome = PlayerWin
				fmt.Printf("PLAYER WIN with %d mana", treeMana)
				if treeMana < resultMana {
					resultMana = treeMana
				}
			}
		}

		if outcome == PlayerWin {
			return
		}
	} else {
		boss.attack(&player)
		if player.HitPoints <= 0 {
			outcome = PlayerLoss
			return
		}
		simulate(player, boss, manaSpent, PlayerTurn, newEffects)
	}

	return
}

func (player *Entity) getAvailableSpells() (available []*Spell) {
	for _, spell := range player.Spells {
		if spell.Cost <= player.Mana && player.Cooldowns[spell.Name] == 0 {
			available = append(available, spell)
		}
	}
	return
}

func (player *Entity) printAvailableSpells() {
	available := player.getAvailableSpells()
	fmt.Printf("%s (%d mana) available spells: ", player.Name, player.Mana)
	for _, spell := range available {
		fmt.Printf("\t%s (%d)", spell.Name, spell.Cost)
	}

	fmt.Printf("\n")
}

func (attacker *Entity) attack(defender *Entity) {
	attackDamage := max(1, attacker.Damage)
	defender.HitPoints -= attackDamage
	//fmt.Printf("\t%s deals %d-%d = %d damage; %s goes down to %d hit points.\n", attacker.Name, attacker.Damage, defender.Armor, attackDamage, defender.Name, defender.HitPoints)
}
