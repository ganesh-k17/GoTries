package math

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/testGo/entities"
)

func TestAdd(t *testing.T) {

	// testData := []entities.TestData{
	// 	{[]int{1, 2}, 3, false},
	// 	{[]int{4, 5}, 9, false},
	// 	{[]int{9, 10}, 19, false},
	// }

	testData := []entities.TestData{
		{Inputs: []int{1, 2}, Result: 3, IsError: false},
		{Inputs: []int{4, 5}, Result: 9, IsError: false},
		{Inputs: []int{9, 10}, Result: 19, IsError: false},
	}

	for _, data := range testData {
		result := Add(data.Inputs[0], data.Inputs[1])
		assert.Equal(t, data.Result, result)
	}
	//t.Error() // to indicate test failed
}
