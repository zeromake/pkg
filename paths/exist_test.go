package paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsExist(t *testing.T) {
	list := []struct {
		name  string
		exist bool
	}{
		// 文件
		{
			name:  "test_data/file0",
			exist: true,
		},
		{
			name:  "test_data/file1",
			exist: false,
		},
		// 软链接
		{
			name:  "test_data/symlink0",
			exist: true,
		},
		{
			name:  "test_data/symlink1",
			exist: true,
		},
		// 目录
		{
			name:  "test_data",
			exist: true,
		},
		{
			name:  "test_data1",
			exist: false,
		},
	}
	for _, item := range list {
		exist := IsExist(item.name)
		assert.Equal(t, exist, item.exist)
	}
}

func TestIsNotExist(t *testing.T) {
	list := []struct {
		name  string
		exist bool
	}{
		// 文件
		{
			name:  "test_data/file0",
			exist: true,
		},
		{
			name:  "test_data/file1",
			exist: false,
		},
		// 软链接
		{
			name:  "test_data/symlink0",
			exist: true,
		},
		{
			name:  "test_data/symlink1",
			exist: true,
		},
		// 目录
		{
			name:  "test_data",
			exist: true,
		},
		{
			name:  "test_data1",
			exist: false,
		},
	}
	for _, item := range list {
		notExist := IsNotExist(item.name)
		assert.Equal(t, notExist, !item.exist)
	}
}

func TestIsTargetExist(t *testing.T) {
	list := []struct {
		name  string
		exist bool
	}{
		// 文件
		{
			name:  "test_data/file0",
			exist: true,
		},
		{
			name:  "test_data/file1",
			exist: false,
		},
		// 软链接
		{
			name:  "test_data/symlink0",
			exist: true,
		},
		{
			name:  "test_data/symlink1",
			exist: false,
		},
		// 目录
		{
			name:  "test_data",
			exist: true,
		},
		{
			name:  "test_data1",
			exist: false,
		},
	}
	for _, item := range list {
		exist := IsLinkTargetExist(item.name)
		assert.Equal(t, exist, item.exist)
	}
}
