package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Skills struct {
	Health string
	Attack string
	Strength string
	Defence string
	Mage string
	Archery string
	Fishing string
	Cooking string
	Mining string
	Smithing string
}

func getSkills() Skills {
	experienceTable := [99]int{0,83,174,276,388,512,650,801,969,1154,1358,1584,1833,2107,2411,2746,3115,3523,3973,4470,5018,5624,6291,7028,7842,8740,9730,10824,12031,13363,14833,16456,18247,20224,22406,24815,27473,30408,33648,37224,41171,45529,50339,55649,61512,67983,75127,83014,91721,101333,111945,123660,136594,150872,166636,184040,203254,224466,247886,273742,302288,333804,368599,407015,449428,496254,547953,605032,668051,737627,814445,899257,992895,1096278,1210421,1336443,1475581,1629200,1798808,1986068,2192818,2421087,2673114,2951373,3258594,3597792,3972294,4385776,4842295,5346332,5902831,6517253,7195629,7944614,8771558,9684577,10692629,11805606,13034431}
	experience, err := ioutil.ReadFile("data/skills.txt")
	if err != nil {
		fmt.Println("err", err)
	}
	//push strings into a slice
	skills := strings.Split(string(experience), ",")
	for i, v := range skills {
		exp, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		for j,vv := range experienceTable {
			if exp < vv {
				skills[i] = strconv.Itoa(j)
				break;
			}
		}
	}
	fmt.Println("Health",skills[0])
	fmt.Println("Attack",skills[1])
	fmt.Println("Strength",skills[2])
	fmt.Println("Defence",skills[3])
	fmt.Println("Mage",skills[4])
	fmt.Println("Archery",skills[5])
	fmt.Println("Fishing",skills[6])
	fmt.Println("Cooking",skills[7])
	fmt.Println("Mining",skills[8])
	fmt.Println("Smithing",skills[9])
	return Skills{
		Health: skills[0],
		Attack: skills[1],
		Strength: skills[2],
		Defence: skills[3], 
		Mage: skills[4],
		Archery: skills[5],
		Fishing: skills[6],
		Cooking: skills[7], 
		Mining: skills[8],
		Smithing: skills[9]}
}


func main() {
	fmt.Println("Welcome to GoScape, please wait while we initialize your load out...")
	fmt.Println("Fetching skill levels...")
	getSkills()
	//skills := getSkills()
}