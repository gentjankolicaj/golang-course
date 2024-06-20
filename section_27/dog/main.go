//package provides functions for converting human years to dogs and vice-versa
package dog

const(
	conversionRate:=10
)

//Convert human years to dog years
func FromHumanToDog(year int) int{
	return year/conversionRate
}


//Convert dog years to human years
func FromDogToHuman(year int) int{
	return year*conversionRate
}




