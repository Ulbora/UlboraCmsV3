package services

import (
	"fmt"
	"testing"
)

var key string

func TestChallengeService_GetChallenge(t *testing.T) {
	var c ChallengeService
	c.Host = "http://localhost:3003"
	res := c.GetChallenge("en_us")
	fmt.Print("res: ")
	fmt.Println(res)
	key = res.Key
	if res.Question == "" {
		t.Fail()
	}
}

func TestChallengeService_SendChallenge(t *testing.T) {
	var c ChallengeService
	c.Host = "http://localhost:3003"
	var ch Challenge
	ch.Answer = "some answer"
	ch.Key = key
	res := c.SendChallenge(&ch)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}
