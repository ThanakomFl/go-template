package utils

import (
	"gitlab.com/F1ukez/sample-go/model"
)

func CalculateProvinceAndAgeGroup(people []interface{}, channel chan<- model.InfectedInfo) {
	ageGroup := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}
	province := make(map[string]int)

	for _, person := range people {
		person := person.(map[string]interface{})
		age, canGetAge := person["Age"].(float64)
		if !canGetAge {
			ageGroup["N/A"] += 1
		} else if age >= 0 && age <= 30 {
			ageGroup["0-30"] += 1
		} else if age >= 31 && age <= 60 {
			ageGroup["31-60"] += 1
		} else {
			ageGroup["61+"] += 1
		}

		provinceEN, canGetProvinceEN := person["ProvinceEn"].(string)
		if !canGetProvinceEN {
			province["N/A"] += 1
		} else {
			province[provinceEN] += 1
		}
	}

	infectedInfo := model.InfectedInfo{
		AgeGroup: ageGroup,
		Province: province,
	}

	channel <- infectedInfo
}
