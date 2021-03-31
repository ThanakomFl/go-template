package model

type InfectedInfo struct {
	AgeGroup map[string]int `json:"AgeGroup"`
	Province map[string]int `json:"Province"`
}