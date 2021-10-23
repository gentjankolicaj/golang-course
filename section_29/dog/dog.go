package dog

const convertionRate=10

func HumanToDogYears(i int)int{
return i/convertionRate
}


func DogToHumanYears(i int) int{
	return i*convertionRate
}