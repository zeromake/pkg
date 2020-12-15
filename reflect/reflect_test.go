package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID   int64
	Name string
}

type UserRespone struct {
	ID   int64
	Name string
}

func TestCopyToStruct(t *testing.T) {
	u := User{
		5,
		"Test",
	}
	var up = &UserRespone{}
	err := CopyToStruct(up, u)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, u.ID, up.ID)
	assert.Equal(t, u.Name, up.Name)
}
