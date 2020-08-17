package _401BinaryWatch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	ans []string
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 1,
			ans: []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"},
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, readBinaryWatchOne(q.one))
	}

}