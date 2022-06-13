//  A CLI rpg

package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type player struct {
	name string
	hp   int // player health
	mp   int // player energy
}

// value receiver function
func (p player) status() {
	fmt.Println("Hero Name:", p.name)
	fmt.Println("HP:", p.hp)
	fmt.Println("MP:", p.mp)
}

func (p *player) takeDamage(damage int) {
	if damage <= p.hp {
		p.hp -= damage
		fmt.Println(p.name, "took a hit of", damage)
		fmt.Println("The hero have", p.hp, "hp left!")
	} else {
		p.hp -= damage
		fmt.Println(p.name, "took a severe hit! The hero falls!")
	}
	time.Sleep(6000 * time.Millisecond)
}

// the player knows two default magic
// fireBall throws a fireball of damage 20 to the opponent , mp usage 20
func (p *player) fireBall() (int, error) {
	const mpUsage = 20
	if p.mp >= mpUsage {
		p.mp -= mpUsage
		fmt.Println(p.name, "successfully performed FireBall.")
		fmt.Println("Made 20 damage!")
		return 20, error(nil)
	} else {

		return 0, errors.New("enough mp, nothing happened")
	}
}

// iceFrost freezes the opponent with damage 10, mp usage 10
func (p *player) iceFrost() (int, error) {
	const mpUsage = 10
	if p.mp >= mpUsage {
		p.mp -= mpUsage
		fmt.Println(p.name, "successfully performed IceFrost.")
		fmt.Println("Made 10 damage!")
		return 10, error(nil)
	} else {

		return 0, errors.New("not enough mp, nothing happened")
	}
}

// the player knows three default One-Handed sword skill
// slant damages the opponent with 30, mp usage 30
func (p *player) slant() (int, error) {
	const mpUsage = 30
	if p.mp >= mpUsage {
		p.mp -= mpUsage
		fmt.Println(p.name, "successfully performed Slant.")
		fmt.Println("Made 30 damage!")
		return 30, error(nil)
	} else {

		return 0, errors.New("not enough mp, nothing happened")
	}
}

// horizontal damages the opponent with 40, mp usage 35
func (p *player) horizontal() (int, error) {
	const mpUsage = 35
	if p.mp >= mpUsage {
		p.mp -= mpUsage
		fmt.Println(p.name, "successfully performed Horizontal Slant.")
		fmt.Println("Made 40 damage!")
		return 40, error(nil)
	} else {

		return 0, errors.New("not enough mp, nothing happened")
	}
}

// horizontal damages the opponent with 20, mp usage 25
func (p *player) vertical() (int, error) {
	const mpUsage = 25
	if p.mp >= mpUsage {
		p.mp -= mpUsage
		fmt.Println(p.name, "successfully performed Vertical Slant.")
		fmt.Println("Made 20 damage!")
		return 20, error(nil)
	} else {

		return 0, errors.New("not enough mp, nothing happened")
	}
}

func (p *player) runAndCheckSkill() (int, error) {
	fmt.Println("f: FireBall(mp: 20, damage: 20)")
	fmt.Println("i: IceFrost(mp: 10, damage: 10)")
	fmt.Println("s: Slant(mp: 30, damage: 30)")
	fmt.Println("v: Vertical Slant(mp: 20, damage: 25)")
	fmt.Println("h: Horizontal Slant(mp: 40, damage: 35)")

	return 0, errors.New("the forgetting hero who knows nothing about his skills runs around to check the dictionary")

}

type tyrant struct {
	name   string
	hp, mp int
}

func (t tyrant) status() {
	fmt.Println("Tyrant Name:", t.name)
	fmt.Println("HP:", t.hp)
	fmt.Println("MP:", t.mp)
}

func (t *tyrant) takeDamage(damage int) {

	if damage <= t.hp {
		t.hp -= damage
		fmt.Println(t.name, "took a hit of", damage)
		fmt.Println("The tyrant have", t.hp, "hp left!")
	} else {
		t.hp -= damage
		fmt.Println(t.name, "took a severe hit! The tyrant is defeated!")
	}
	time.Sleep(6000 * time.Millisecond)

}

// strike is the one and only the tyrant have, mp usage 40
// damage is 60 when hit(successful strike), else damages 5(the aftereffects of striking)
func (t *tyrant) strike() (int, error) {
	const mpUsage = 40
	rand.Seed(time.Now().Unix())
	if t.mp >= mpUsage {
		if rand.Intn(2) == 1 {
			damage := 60
			fmt.Println(t.name, "strikes, made 60 damage.")
			return damage, error(nil)
		} else {
			damage := 5
			fmt.Println(t.name, "strikes misses, the ground rumbling, made 5 damage.")
			return damage, error(nil)
		}
	} else {
		return 0, errors.New("not enough mp, nothing happened")
	}
}

func (t *tyrant) restAndCheckSkill() (int, error) {
	fmt.Println("Strike(mp: 20, damage: 5 or 60 depends on how focus you are)")

	return 0, errors.New("the yelling hero is annoying, let me rest and check what skills i have")

}

func main() {
	rand.Seed(time.Now().Unix())
	minHpHero := 100
	maxHpHero := 500
	minMpHero := 200
	maxMpHero := 400

	theHero := player{
		name: "The Hero",
		hp:   rand.Intn(maxHpHero-minHpHero) + minHpHero,
		mp:   rand.Intn(maxMpHero-minMpHero) + minMpHero,
	}

	minHpBoss := 300
	maxHpBoss := 1000
	minMpBoss := 400
	maxMpBoss := 2000

	bossAlpha := tyrant{
		name: "Alpha",
		hp:   rand.Intn(maxHpBoss-minHpBoss) + minHpBoss,
		mp:   rand.Intn(maxMpBoss-minMpBoss) + minMpBoss,
	}

	fmt.Println("The hero met the tyrant!")
	var char string
	fmt.Println("Are you the hero or the tyrant? Press h for hero, t for tyrant: ")
	fmt.Scanln(&char)
	if char == "t" {
		for bossAlpha.hp > 0 {
			bossAlpha.status()
			fmt.Println("You're turn to strike! Decide your move: (Press on the following key: s, r)")
			var attack string
			fmt.Scanln(&attack)
			var damage int
			var err error
			switch {
			case attack == "s":
				damage, err = bossAlpha.strike()
			case attack == "r":
				damage, err = bossAlpha.restAndCheckSkill()
			default:
				damage, err = 0, fmt.Errorf("what a boring fight...nothing happened")
			}

			if err == nil {
				theHero.takeDamage(damage)
				fmt.Println()
				if theHero.hp <= 0 {
					break
				} else {
					fmt.Println(err)
					fmt.Println()
				}

			}

			theHero.status()
			rand.Seed(time.Now().Unix())
			attackIdx := rand.Intn(5)
			attacks := []string{"f", "i", "s", "v", "h"}

			switch {
			case attacks[attackIdx] == "f":
				damage, err = theHero.fireBall()
			case attacks[attackIdx] == "i":
				damage, err = theHero.iceFrost()
			case attacks[attackIdx] == "s":
				damage, err = theHero.slant()
			case attacks[attackIdx] == "v":
				damage, err = theHero.vertical()
			case attacks[attackIdx] == "h":
				damage, err = theHero.horizontal()
			}

			if err == nil {
				bossAlpha.takeDamage(damage)
				fmt.Println()

			} else {
				bossAlpha.takeDamage(damage)
				fmt.Println(err)
				fmt.Println()

			}

		}
	} else if char == "h" {

		for theHero.hp > 0 {
			theHero.status()
			fmt.Println("You're turn to strike! Decide your move: (Press on the following key: f, i, s, v, h, c)")
			var attack string
			fmt.Scanln(&attack)
			var damage int
			var err error
			switch {
			case attack == "f":
				damage, err = theHero.fireBall()
			case attack == "i":
				damage, err = theHero.iceFrost()
			case attack == "s":
				damage, err = theHero.slant()
			case attack == "v":
				damage, err = theHero.vertical()
			case attack == "h":
				damage, err = theHero.horizontal()
			case attack == "c":
				damage, err = theHero.runAndCheckSkill()
			default:
				damage, err = 0, fmt.Errorf("what are you dreaming of when in a tensive fight...nothing happened")
			}
			if err == nil {
				bossAlpha.takeDamage(damage)
				fmt.Println()

				if bossAlpha.hp <= 0 {
					break
				}
			} else {
				fmt.Println(err)
				fmt.Println()
			}

			bossAlpha.status()
			damage, err = bossAlpha.strike()
			if err == nil {
				theHero.takeDamage(damage)
				fmt.Println()

			} else {
				theHero.takeDamage(damage)
				fmt.Println(err)
				fmt.Println()

			}
		}
	} else {
		log.Fatal("Try to pick a character!")
	}
}
