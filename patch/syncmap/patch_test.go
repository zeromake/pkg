package syncmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/zeromake/pkg/patch"
	"github.com/zeromake/pkg/patch/types"
	"reflect"
	"testing"
)

type Progress struct {
	Status      int32 `json:"status,omitempty"`
	ExpiredTime int64 `json:"expired_time,omitempty"`
	FinishTime  int64 `json:"finish_time,omitempty"`
	Progress    int32 `json:"progress,omitempty"`
}

type SeasonQuestInfo struct {
	Progress     map[string]*Progress `json:"progress,omitempty"`
	ExpiredTime  int64                `json:"expired_time,omitempty"`
	RewardStatus map[string]int32     `json:"reward_status,omitempty"`
}

type DailyQuestInfo struct {
	DailyProgress     map[string]*Progress `json:"daily_progress,omitempty"`
	WeeklyExpiredTime int64                `json:"weekly_expired_time,omitempty"`
	WeeklyQuestStatus map[string]int32     `json:"weekly_quest_status,omitempty"`
	DailyReceiveTime  int64                `json:"daily_receive_time,omitempty"`
}

type GuideQuestInfo struct {
	Progress     int32 `json:"progress,omitempty"`
	RecordStatus int32 `json:"record_status,omitempty"`
}

type FrameModel struct {
	ID          int32 `json:"id,omitempty"`
	ReceiveTime int64 `json:"receive_time,omitempty"`
	IsClick     int32 `json:"is_click,omitempty"` // 0 unClick 1 click
}

type PropertyModel struct {
	Lev                 int32         `json:"lev,omitempty"`                  // 当前等级
	LevelXP             int32         `json:"level_xp,omitempty"`             // 等级经验
	VipLev              int32         `json:"vip_lev,omitempty"`              // vip 等级
	VipLevelXp          int32         `json:"vip_level_xp,omitempty"`         // vip关卡经验
	Balance             *PropsMap64   `json:"balance,omitempty"`              // 货币		key:60000~69999
	Props               *PropsMap32   `json:"props,omitempty"`                // 道具  		key: 1~99
	RegularProps        *PropsMap32   `json:"regular_props,omitempty"`        // 常规道具
	ActivityProps       *PropsMap32   `json:"activity_props,omitempty"`       // 活动道具
	FrameList           []*FrameModel `json:"frame_list,omitempty"`           // 头像框资源
	TipLeftTime         int64         `json:"tip_left_time,omitempty"`        // 提示器时间
	CollectionFragments *PuzzleMap    `json:"collection_fragments,omitempty"` // 拼图资源
}

type Model struct {
	PropertyInfo *PropertyModel             `json:"property_info,omitempty"`
	Season       *SeasonQuestInfo           `json:"season,omitempty"`
	Daily        DailyQuestInfo             `json:"daily,omitempty"`
	Guide        map[string]*GuideQuestInfo `json:"guide,omitempty"`
	Guide2       []*GuideQuestInfo          `json:"guide2,omitempty"`
	Test         [][]int64                  `json:"test,omitempty"`
}

func TestAddPatch(t *testing.T) {
	var err error
	model := &Model{}
	rValue := reflect.ValueOf(model)
	opt := &types.Option{}
	err = patch.ModifyPatch(rValue, "/season/expired_time", 6000, opt)
	err = patch.ModifyPatch(rValue, "/daily/weekly_expired_time", 5000, opt)
	err = patch.ModifyPatch(rValue, "/guide/5/progress", 5, opt)
	err = patch.ModifyPatch(rValue, "/guide/5/progress", 9, opt)
	err = patch.ModifyPatch(rValue, "/guide2/0/progress", 5, opt)
	err = patch.ModifyPatch(rValue, "/guide2/0/progress", 9, opt)
	err = patch.ModifyPatch(rValue, "/guide2/0", map[string]interface{}{"progress": 100}, opt)
	err = patch.ModifyPatch(rValue, "/property_info", map[interface{}]interface{}{"lev": 100}, opt)
	err = patch.ModifyPatch(rValue, "/test/0/0", 9, opt)
	err = patch.ModifyPatch(rValue, "/test/0/0", 6, opt)
	err = patch.ModifyPatch(rValue, "/test/1/0", 6, opt)
	err = patch.ModifyPatch(rValue, "/property_info/balance/60000", 1000, opt)
	err = patch.ModifyPatch(rValue, "/property_info/props/1", 50, opt)
	err = patch.ModifyPatch(rValue, "/property_info/collection_fragments/1", map[string]interface{}{
		"collectionRound": 1,
		"complete_time":   3,
		"collectionPuzzles": []int64{
			1,
		},
		"last_puzzle_play_round": 1,
	}, opt)
	assert.Nil(t, err)
	bb, _ := json.MarshalIndent(model, "", "  ")
	t.Log(string(bb))
}

func TestRemovePatch(t *testing.T) {
	var err error
	model := &Model{}
	rValue := reflect.ValueOf(model)
	opt := &types.Option{}
	err = patch.ModifyPatch(rValue, "/season/expired_time", 6000, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/daily/weekly_expired_time", 5000, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/guide/5/progress", 5, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/guide/6/progress", 9, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/guide2/0/progress", 5, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/guide2/0/progress", 9, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/guide2/0", map[string]interface{}{"progress": 100}, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/property_info", map[interface{}]interface{}{"lev": 100}, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/test/0/0", 9, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/test/0/0", 6, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/test/1/0", 6, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/property_info/balance/60000", 1000, opt)
	assert.Nil(t, err)
	err = patch.ModifyPatch(rValue, "/property_info/props/1", 50, opt)
	assert.Nil(t, err)
	bb, _ := json.MarshalIndent(model, "", "  ")
	t.Log(string(bb))

	err = patch.RemovePatch(rValue, "/season")
	assert.Nil(t, err)
	err = patch.RemovePatch(rValue, "/test/0")
	assert.Nil(t, err)
	err = patch.RemovePatch(rValue, "/test/0/0")
	assert.Nil(t, err)
	err = patch.RemovePatch(rValue, "/property_info/props/1")
	assert.Nil(t, err)
	err = patch.RemovePatch(rValue, "/guide/5")
	assert.Nil(t, err)
	bb, _ = json.MarshalIndent(model, "", "  ")
	t.Log(string(bb))
}
