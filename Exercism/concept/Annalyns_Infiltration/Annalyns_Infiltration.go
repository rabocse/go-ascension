package Annalyns_Infiltration

func CanFastAttack(knightIsAwake bool) bool {

	if knightIsAwake == true {

		// fmt.Println("The knight is sleeping. Attack !!!")
		return true
	} else {
		return false
	}
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {

	if knightIsAwake || archerIsAwake || prisonerIsAwake {

		// fmt.Println("At least one is awake, let's spy upon !")
		return true
	} else {
		return false
	}

}

func CanSignalPrisoner(prisonerIsAwake, archerIsAwake bool) bool {

	if archerIsAwake == false && prisonerIsAwake == true {

		// fmt.Println("The archer is sleeping, we can signal the prisoner !!!")
		return true
	} else {
		return false
	}
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, dogIsWithAnnalyn bool) bool {

	if archerIsAwake == false && dogIsWithAnnalyn == true {

		// fmt.Println("My dog is here and the archer is sleeping. Let's rescue the prisoner!!! ")
		return true

	} else if dogIsWithAnnalyn == false && prisonerIsAwake == true && archerIsAwake == false && knightIsAwake == false {

		// fmt.Println("My dog is not here but archer and knight are sleeping while the prisoner is awake. Let's rescue the prisoner!!!")
		return true

	} else {
		return false
	}

}
