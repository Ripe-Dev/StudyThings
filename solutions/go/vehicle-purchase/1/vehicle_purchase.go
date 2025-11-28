package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		option2 += " is clearly the better choice."
		return option2
	}
	option1 += " is clearly the better choice."
	return option1
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		originalPrice *= 80
		return originalPrice / 100
		//return 80%

	case age >= 10:
		originalPrice *= 50
		return originalPrice / 100
		//return 70%

	case age >= 3 && age < 10:
		originalPrice *= 70
		return originalPrice / 100
		// return 50%
	}
	return 3.14
	/*
		if less then 3 years old == 80 % of it original price
		if betwen 3y (>= 3y) and 10y == 70% of it original price
		if at least 10 years old == 50% of it original price
	*/
}