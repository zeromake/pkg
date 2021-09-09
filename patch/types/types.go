package types

import "strconv"

type Item struct {
	Op    string      `json:"op"`
	From  string      `json:"from"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type Patch interface {
	Patch(patch *Item) error
}


type Option struct {
	Replace bool
}



// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n Number) Float64() float64 {
	nn, _ := strconv.ParseFloat(string(n), 64)
	return nn
}

// Int64 returns the number as an int64.
func (n Number) Int64() int64 {
	nn, _ := strconv.ParseInt(string(n), 10, 64)
	return nn
}

// Uint64 returns the number as an uint64.
func (n Number) Uint64() uint64 {
	nn, _ := strconv.ParseUint(string(n), 10, 64)
	return nn
}

type State struct {
}

func ToInt64(v interface{}) int64 {
	switch vv := v.(type) {
	case int:
		return int64(vv)
	case int8:
		return int64(vv)
	case int16:
		return int64(vv)
	case int32:
		return int64(vv)
	case int64:
		return vv
	case uint8:
		return int64(vv)
	case uint16:
		return int64(vv)
	case uint32:
		return int64(vv)
	case uint64:
		return int64(vv)
	case float32:
		return int64(vv)
	case float64:
		return int64(vv)
	case string:
		return Number(vv).Int64()
	}
	return 0
}

func ToUint64(v interface{}) uint64 {
	switch vv := v.(type) {
	case int:
		return uint64(vv)
	case int8:
		return uint64(vv)
	case int16:
		return uint64(vv)
	case int32:
		return uint64(vv)
	case int64:
		return uint64(vv)
	case uint8:
		return uint64(vv)
	case uint16:
		return uint64(vv)
	case uint32:
		return uint64(vv)
	case uint64:
		return vv
	case float32:
		return uint64(vv)
	case float64:
		return uint64(vv)
	case string:
		return Number(vv).Uint64()
	}
	return 0
}

func ToFloat64(v interface{}) float64 {
	switch vv := v.(type) {
	case int:
		return float64(vv)
	case int8:
		return float64(vv)
	case int16:
		return float64(vv)
	case int32:
		return float64(vv)
	case int64:
		return float64(vv)
	case uint8:
		return float64(vv)
	case uint16:
		return float64(vv)
	case uint32:
		return float64(vv)
	case uint64:
		return float64(vv)
	case float32:
		return float64(vv)
	case float64:
		return vv
	case string:
		return Number(vv).Float64()
	}
	return 0
}


func Split(path string) (key, p string) {
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}
	i := 0
	count := len(path)
	for i < count && path[i] != '/' {
		i++
	}
	if i < count {
		return path[:i], path[i+1:]
	}
	return path[:i], ""
}
