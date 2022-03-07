package process

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenRndString(t *testing.T) {
	s := GenerateRandomString()
	fmt.Println(s)
	if !assert.Equal(t, 10, len(s)) {
		t.Error("Epected length of random string 10, but go : ", len(s))
	}
}

func TestUpdateNumOfCases(t *testing.T) {
	s, err := UpdateNumOfUses("1")
	fmt.Println(s)
	if !assert.Equal(t, "2", s) {
		t.Error("Epected result '2', but go : ", s)
	}
	if !assert.Nil(t, err) {
		t.Error("There is an error: ", err)
	}
}
