package services

import (
	"fmt"
	"testing"
)

func TestChallengeService_GetChallenge(t *testing.T) {
	var c ChallengeService
	c.Host = "http://localhost:3003"
	res := c.GetChallenge("en_us")
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Question == "" {
		t.Fail()
	}
}
