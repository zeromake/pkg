package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnderlineToPascalCase(t *testing.T) {
	names := [][2]string{
		{
			"name",
			"Name",
		},
		{
			"api_endpoind",
			"APIEndpoind",
		},
		{
			"id",
			"ID",
		},
		{
			"u_id",
			"UID",
		},
		{
			"uid",
			"UID",
		},
	}
	for _, name := range names {
		assert.True(t, UnderlineToPascalCase(name[0]) == name[1])
	}
}
