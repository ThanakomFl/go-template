package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/F1ukez/sample-go/model"
	"gitlab.com/F1ukez/sample-go/utils"
)

func CountInfectedHandler(c *gin.Context) {
	var response map[string]interface{}
	resp, err := http.Get("https://covid19.th-stat.com/api/open/cases")
	if err != nil {
		c.JSON(500, err)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(500, err)
		return
	}

	data := response["Data"].([]interface{})

	length := len(data)
	ch0 := make(chan model.InfectedInfo)
	data0 := data[0 : length/4]

	ch1 := make(chan model.InfectedInfo)
	data1 := data[length/4 : (length*2)/4]

	ch2 := make(chan model.InfectedInfo)
	data2 := data[(length*2)/4 : (length*3)/4]

	ch3 := make(chan model.InfectedInfo)
	data3 := data[(length*3)/4 : length]

	go utils.CalculateProvinceAndAgeGroup(data0, ch0)
	go utils.CalculateProvinceAndAgeGroup(data1, ch1)
	go utils.CalculateProvinceAndAgeGroup(data2, ch2)
	go utils.CalculateProvinceAndAgeGroup(data3, ch3)

	infectedInfo0 := <-ch0
	infectedInfo1 := <-ch1
	infectedInfo2 := <-ch2
	infectedInfo3 := <-ch3

	ageGroup := map[string]int{
		"0-30":  infectedInfo0.AgeGroup["0-30"] + infectedInfo1.AgeGroup["0-30"] + infectedInfo2.AgeGroup["0-30"] + infectedInfo3.AgeGroup["0-30"],
		"31-60": infectedInfo0.AgeGroup["31-60"] + infectedInfo1.AgeGroup["31-60"] + infectedInfo2.AgeGroup["31-60"] + infectedInfo3.AgeGroup["31-60"],
		"61+":   infectedInfo0.AgeGroup["61+"] + infectedInfo1.AgeGroup["61+"] + infectedInfo2.AgeGroup["61+"] + infectedInfo3.AgeGroup["61+"],
		"N/A":   infectedInfo0.AgeGroup["N/A"] + infectedInfo1.AgeGroup["N/A"] + infectedInfo2.AgeGroup["N/A"] + infectedInfo3.AgeGroup["N/A"],
	}
	province := make(map[string]int)
	for key, value := range infectedInfo0.Province {
		province[key] += value
	}
	for key, value := range infectedInfo1.Province {
		province[key] += value
	}
	for key, value := range infectedInfo2.Province {
		province[key] += value
	}
	for key, value := range infectedInfo3.Province {
		province[key] += value
	}

	c.JSON(200, model.InfectedInfo{AgeGroup: ageGroup, Province: province})
}
