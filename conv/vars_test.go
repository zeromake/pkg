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
		// {
		// 	"u_id",
		// 	"UId",
		// },
		{
			"uid",
			"UID",
		},
		{
			"ut7",
			"Ut7",
		},
		{
			"ip_add_r",
			"IPAddR",
		},
		{
			"",
			"",
		},
	}
	for _, name := range names {
		n := UnderlineToPascalCase(name[0])
		assert.Equal(t, n, name[1])
	}
}

func TestPascalCaseToUnderline(t *testing.T) {
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
		// {
		// 	"u_id",
		// 	"UId",
		// },
		{
			"uid",
			"UID",
		},
		{
			"ut7",
			"Ut7",
		},
		{
			"ip_add_r",
			"IPAddR",
		},
		{
			"",
			"",
		},
	}
	for _, name := range names {
		n := PascalCaseToUnderline(name[1])
		assert.Equal(t, n, name[0])
	}
}
