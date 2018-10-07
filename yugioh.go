package main

import "fmt"

type yugioh struct {
	id string
	name string
	cardtype string
	level string
	attribute string
	race string
	atk string
	def string
	effect string
}

func (m yugioh) getCard() string {
	return fmt.Sprintf("No. %s -- %s\nCard Type : %s\nLevel : %s\nAttribute : %s\nRace : %s\n",m.id,m.name,m.cardtype,m.level,m.attribute,m.race)
} 

func main() {
	gradius := yugioh{
		id:"001",
		name:"Gradius",
		cardtype:"Normal Monster",
		level:"4",
		attribute:"Light",
		race:"Machine"}
	fmt.Println(gradius.getCard())
}

