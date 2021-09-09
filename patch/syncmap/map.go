package syncmap

import (
	"encoding/json"
	"github.com/zeromake/pkg/patch"
	"github.com/zeromake/pkg/patch/types"
	"reflect"
	"strconv"
	"strings"
)

// go get -u github.com/a8m/syncmap
//go:generate syncmap -pkg syncmap -name PropsMap64 map[string]int64
//go:generate syncmap -pkg syncmap -name PropsMap32 map[string]int32
//go:generate syncmap -pkg syncmap -name PuzzleMap map[string]*PuzzleCollectionStatus

func (m *PropsMap64) MarshalJSON() ([]byte, error) {
	builder := &strings.Builder{}
	builder.WriteByte('{')
	var arr []string
	m.Range(func(key string, value int64) bool {
		arr = append(arr, "\""+key+"\":"+strconv.FormatInt(value, 10))
		return true
	})
	builder.WriteString(strings.Join(arr, ","))
	builder.WriteByte('}')
	return []byte(builder.String()), nil
}

func (m *PropsMap64) CloneToMap() map[string]int64 {
	result := map[string]int64{}
	m.Range(func(key string, value int64) bool {
		result[key] = value
		return true
	})
	return result
}

func (m *PropsMap64) Patch(item *types.Item) (err error) {
	switch item.Op {
	case "add":
		_, has := m.Load(item.Path)
		if !has {
			m.Store(item.Path, types.ToInt64(item.Value))
		}
	case "replace":
		m.Store(item.Path, types.ToInt64(item.Value))
	case "remove":
		m.Delete(item.Path)
	}
	return
}

func (m *PropsMap32) CloneToMap() map[string]int32 {
	result := map[string]int32{}
	m.Range(func(key string, value int32) bool {
		result[key] = value
		return true
	})
	return result
}

func (m *PropsMap32) MarshalJSON() ([]byte, error) {
	builder := &strings.Builder{}
	builder.WriteByte('{')
	var arr []string
	m.Range(func(key string, value int32) bool {
		arr = append(arr, "\""+key+"\":"+strconv.FormatInt(int64(value), 10))
		return true
	})
	builder.WriteString(strings.Join(arr, ","))
	builder.WriteByte('}')
	return []byte(builder.String()), nil
}


func (m *PropsMap32) Patch(item *types.Item) (err error) {
	switch item.Op {
	case "add":
		_, has := m.Load(item.Path)
		if !has {
			m.Store(item.Path, int32(types.ToInt64(item.Value)))
		}
	case "replace":
		m.Store(item.Path, int32(types.ToInt64(item.Value)))
	case "remove":
		m.Delete(item.Path)
	}
	return
}

type PuzzleCollectionStatus struct {
	CompleteRound       int32   `json:"collectionRound,omitempty"`        // 收集完成轮数
	CompleteTime        int64   `json:"complete_time,omitempty"`          // 收集完成时间
	Puzzles             []int32 `json:"collectionPuzzles,omitempty"`      // 收集状态
	LastPuzzlePlayRound int32   `json:"last_puzzle_play_round,omitempty"` // 当收集到最后一个碎片的时候 开始计数，一共玩了多少句 (用于用户收集碎片的补偿值计算 map中的puzzle_game字段)
}

func (p *PuzzleCollectionStatus) GetCollectionPuzzleCount(fragments []int32) int32 {
	if p.CompleteRound >= int32(len(fragments)) {
		return fragments[len(fragments)-1]
	}
	return int32(len(p.Puzzles))
}

func (m *PuzzleMap) CloneToMap() map[string]*PuzzleCollectionStatus {
	result := map[string]*PuzzleCollectionStatus{}
	m.Range(func(key string, value *PuzzleCollectionStatus) bool {
		result[key] = value
		return true
	})
	return result
}

func (m *PuzzleMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.CloneToMap())
}


func (m *PuzzleMap) Patch(item *types.Item) (err error) {
	key, nextPath := types.Split(item.Path)
	if nextPath != "" {
		switch item.Op {
		case "add", "replace":
			opt := &types.Option{}
			if item.Op == "replace" {
				opt.Replace = true
			}
			v, has := m.Load(key)
			if !has {
				v = &PuzzleCollectionStatus{}
			}
			rv := reflect.ValueOf(v)
			return patch.ModifyPatch(rv, nextPath, item.Value, opt)
		case "remove":
			v, has := m.Load(key)
			if !has {
				return
			}
			rv := reflect.ValueOf(v)
			return patch.RemovePatch(rv, nextPath)
		}
		return
	}
	switch item.Op {
	case "add":
		_, has := m.Load(key)
		if has {
			return
		}
		v := &PuzzleCollectionStatus{}
		rv := reflect.ValueOf(v)
		err = patch.SetValue(item.Value, rv)
		if err != nil {
			return
		}
		m.Store(key, v)
	case "replace":
		v, has := m.Load(key)
		if !has {
			v = &PuzzleCollectionStatus{}
			m.Store(key, v)
		}
		rv := reflect.ValueOf(v)
		err = patch.SetValue(item.Value, rv)
		if err != nil {
			return
		}
	case "remove":
		m.Delete(key)
	}
	return
}

