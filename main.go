package main

import (
	"bufio"
	"fmt"
	"os"
)

type hero struct {
	heroName   string
	playerName string
	race       string
	class      string
	background string
	stats      statsList
	hits       int
}

type statsList struct {
	strength     int
	dexterity    int
	constitution int
	intelligence int
	wisdom       int
	charisma     int
}

type weapon struct {
	name   string
	damage int
	typeOf string
}

type armor struct {
	name       string
	armorClass int
	typeOf     string
}

func main() {
	myHero := createHero()
	printing(myHero)

	/* notes
	1. СДЕЛАТЬ ПРОВЕРКУ СУЩЕСТВУЕТ ЛИ ТАКОЙ КЛАСС/РАСА/ПРЕДЫСТОРИЯ
	2. ЛЮБОЙ РЕГИСТР БУКВ
	3. НАПИСАТЬ СИМУЛЯЦИИ БРОСКОВ КУБИКА
	4. УМНОЕ РАСПРЕДЕЛЕНИЕ ХАРАКТЕРИСТИК
	ОТ 27 ОЧКОВ В ЗАВИСИМОСТИ ОТ ТИПА ИГРЫ
	5. СИМУЛЯЦИЯ ХИТОВ И ЯЧЕЕК ЗАКЛИНАНИЙ
	*/
}

func initiationLists() ([]string, []string, []string, []int) {
	startStats := []int{15, 14, 13, 12, 10, 8}

	class := []string{
		"paladin",
		"warrior",
		"bard",
	}

	race := []string{
		"human",
		"lizardfolk",
	}

	background := []string{
		"guardian",
	}

	return class, race, background, startStats
}

func printing(myHero hero) {
	fmt.Println("Your hero: ", myHero.heroName)
	fmt.Println("Race:", myHero.race)
	fmt.Println("Class: ", myHero.class)
	fmt.Println("Stats: ", myHero.stats)
	fmt.Println("Background: ", myHero.background)
}

func heroName() {
	/*
		enter hero class
		if there is no
		heroClass() again

	*/
}

func scanner() (string, string, string, string) {

	var heroName, race, class, background string

	fmt.Println("Enter Hero name: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	heroName = sc.Text()

	fmt.Println("Enter Hero race: ")
	sc.Scan()
	race = sc.Text()

	fmt.Println("Enter Hero class: ")
	sc.Scan()
	class = sc.Text()

	fmt.Println("Enter Hero background: ")
	sc.Scan()
	background = sc.Text()

	fmt.Println()
	fmt.Println()
	return heroName, race, class, background

}

func createHero() hero {
	_, _, _, startStats := initiationLists()

	// For manual hero creating.
	//heroName, race, class, background := scanner()

	/*
		If you want to try to create your character,
		then activate the "scanner", see the list
		of available races and classes in L#59 - initiationLists().
	*/

	heroName, race, class, background := "Krestar", "human", "paladin", "warrior"

	myHero := hero{
		heroName:   heroName,
		playerName: "Kaplan",
		race:       race,
		class:      class,
		background: background,
		stats: statsList{
			strength:     8,
			dexterity:    8,
			constitution: 8,
			intelligence: 8,
			wisdom:       8,
			charisma:     8,
		},
	}

	_, _, myHero = placeStats(myHero, startStats)

	myHero = raceInit(myHero)
	myHero = classInit(myHero)

	return myHero
}

func placeStats(myHero hero, startStats []int) (int, int, hero) {

	var firstMainCharacteristic *int
	var secondMainCharacteristic *int
	var otherCharacteristic []*int

	strength := &(myHero.stats.strength)
	dexterity := &(myHero.stats.dexterity)
	constitution := &(myHero.stats.constitution)
	intelligence := &(myHero.stats.intelligence)
	wisdom := &(myHero.stats.wisdom)
	charisma := &(myHero.stats.charisma)

	switch myHero.class {
	case "paladin":
		firstMainCharacteristic = strength
		secondMainCharacteristic = charisma
		otherCharacteristic = append(otherCharacteristic, dexterity, constitution, intelligence, wisdom)

	case "warrior":
		firstMainCharacteristic = strength
		secondMainCharacteristic = constitution
		otherCharacteristic = append(otherCharacteristic, dexterity, intelligence, wisdom, charisma)

	case "bard":
		firstMainCharacteristic = charisma
		secondMainCharacteristic = dexterity
		otherCharacteristic = append(otherCharacteristic, constitution, intelligence, wisdom, strength)
	}

	*firstMainCharacteristic, startStats = distributeStats(startStats)
	*secondMainCharacteristic, startStats = distributeStats(startStats)

	for i := range otherCharacteristic {
		*otherCharacteristic[i], startStats = distributeStats(startStats)
	}

	return *firstMainCharacteristic, *secondMainCharacteristic, myHero
}

func distributeStats(startStats []int) (int, []int) {
	current := startStats[0]
	startStats = startStats[1:]
	return current, startStats
}

func classInit(myHero hero) hero {
	switch myHero.class {
	case "paladin":
		myHero.hits = 10

	case "warrior":
		myHero.hits = 10

	case "bard":
		myHero.hits = 8
	}
	return myHero
}

func raceInit(myHero hero) hero {

	switch myHero.race {
	case "human":
		myHero.stats.strength = myHero.stats.strength + 1
		myHero.stats.dexterity = myHero.stats.dexterity + 1
		myHero.stats.constitution = myHero.stats.constitution + 1
		myHero.stats.intelligence = myHero.stats.intelligence + 1
		myHero.stats.wisdom = myHero.stats.wisdom + 1
		myHero.stats.charisma = myHero.stats.charisma + 1

	case "lizardfolk":
		myHero.stats.strength = myHero.stats.strength + 2
		myHero.stats.constitution = myHero.stats.constitution + 2
	}
	return myHero

}
