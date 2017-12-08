package tasks_test

import (
	"testing"

	"github.com/Guazi-inc/machinery/v1/tasks"
	"github.com/stretchr/testify/assert"
)

func TestNewChain(t *testing.T) {
	task1 := tasks.Signature{
		Name: "foo",
		Args: []tasks.Arg{
			{
				Type:  "float64",
				Value: interface{}(1),
			},
			{
				Type:  "float64",
				Value: interface{}(1),
			},
		},
	}

	task2 := tasks.Signature{
		Name: "bar",
		Args: []tasks.Arg{
			{
				Type:  "float64",
				Value: interface{}(5),
			},
			{
				Type:  "float64",
				Value: interface{}(6),
			},
		},
	}

	task3 := tasks.Signature{
		Name: "qux",
		Args: []tasks.Arg{
			{
				Type:  "float64",
				Value: interface{}(4),
			},
		},
	}

	chain := tasks.NewChain(&task1, &task2, &task3)

	firstTask := chain.Tasks[0]

	assert.Equal(t, "foo", firstTask.Name)
	assert.Equal(t, "bar", firstTask.OnSuccess[0].Name)
	assert.Equal(t, "qux", firstTask.OnSuccess[0].OnSuccess[0].Name)
}
