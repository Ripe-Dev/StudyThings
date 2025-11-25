package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    productionFloat := float64(productionRate)
    multiply := productionFloat * successRate
	if multiply == 100{
        return 0.0
    }
    return multiply / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRate /= 60
    productionFloat := float64(productionRate)
    multiply := productionFloat * successRate
    multInt := int(multiply)
    return multInt / 100    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	decimal := carsCount / 10
    unity := carsCount % 10

	if carsCount >= 10 {
        decimal *= 95000 
        unity *= 10000
        return uint(unity + decimal)
    }
    return uint (carsCount * 10000) 
} 

