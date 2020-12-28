package logging

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getInstance(ch chan *Logger) {
	l := GetLogger()
	l.Info("Create Instance")
	ch <- &l
}

func TestLogger_Logging(t *testing.T) {
	t.Run("Is singleton", func(t *testing.T) {

		ch1 := make(chan *Logger)
		ch2 := make(chan *Logger)
		go getInstance(ch1)
		go getInstance(ch2)
		l1 := <-ch1
		l2 := <-ch2

		res := reflect.DeepEqual(&l1, &l2)
		fmt.Println(fmt.Sprintf("Is same object: %v", res))
		assert.Equal(t, res, true)
	})
}
