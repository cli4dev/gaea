package md

import (
	"testing"

	"github.com/micro-plat/lib4go/ut"
)

func TestTable(t *testing.T) {
	input := []string{
		"### 一、媒资信息",
		"#### 1. 媒资信息 [epg_media_info]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
		"|media_id|number(20)|-|否|PK|编号|",
		"#### 1. 媒资信息 [epg_media_info]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
		"|media_id|number(20)|-|否|PK|编号|",
	}
	tables := markdown2Strings(input)
	ut.ExpectSkip(t, len(tables), 2)
	ut.ExpectSkip(t, len(tables[0]), 4)
	ut.ExpectSkip(t, len(tables[1]), 4)
}

func TestTables(t *testing.T) {
	input := []string{
		"### 一、媒资信息",
		"#### 1. 媒资信息 [epg_media_info]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
		"|media_id|number(20)|-|否|PK|编号|",
		"|media_no|varchar(32)|-|否||唯一标识|",
		"|title|varchar(32)|-|否||标题|",
		"|en_title|varchar(32)|-|否||英文名|",
		"|title_py|varchar(64)|-|否||拼音|",
		"|initials|varchar(32)|-|否||首字母|",
	}
	tables := markdown2Strings(input)
	ut.ExpectSkip(t, len(tables), 1)
	tbs, err := strings2Tables(tables)
	ut.ExpectSkip(t, err, nil)
	ut.ExpectSkip(t, len(tbs), 1)
	ut.Expect(t, tbs[0].Name, "epg_media_info")
	ut.Expect(t, tbs[0].Desc, "媒资信息")
	ut.Expect(t, len(tbs[0].CNames), 6)
}

func TestTables2(t *testing.T) {
	input := []string{

		"#### 1. 媒资图片 [epg_media_picture]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
		"|id|number(20)|-|否|PK|编号|",
		"|media_id|number(20)|-|否||媒资编号|",
		"|show_group|number(2)|-|否||显示分组,轮播图片|",
		"|pic_url|varchar(128)|-|否||地址|",
		"|sortrank|number(10)|-|否||显示顺序|",
		"",
		"",
		"#### 2. 媒资剧集 [epg_media_series]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
		"|id|number(20)|-|否|PK|编号|",
		"|series_id|varchar(32)|-|否||唯一标识|",
		"|media_id|number(20)|-|否||媒资编号|",
		"|serial_num|number(10)|-|否||序号|",
		"|serical_type|number(1)|-|否||剧集类型: 1：正剧 2：宣传片|",
		"|head_duration|number(20)|0|是||片头时长|",
		"|duration|number(20)|0|是||片尾时长|",
		"|serial_name|varchar(32)|-|是||序集名称|",
		"|zx_play_url|varchar(128)|-|是||中兴播放地址|",
		"|hw_play_url|varchar(128)|-|是||华为播放地址|",
		"|fh_play_url|varchar(128)|-|是||烽火播放地址|",
		"",
		"",
		"#### 3. 轮播信息 [epg_media_carousel]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
		"|id|number(20)|-|否|PK|编号|",
		"|media_id|number(20)|-|否||媒资编号|",
		"|type|number(2)|-|否||类型 1. media 2. 广告|",
		"|big_pic_url|varchar(128)|-|否||大图地址|",
		"|small_pic_url|varchar(128)|-|否||小图地址|",
		"|show_big|number(1)|-|否||显示大图|",
		"|show_in_fp|number(1)|-|否||显示到首页|",
		"|show_category_id|number(2)|-|否||显示到分类，电影/电视据/综艺/动漫，0:所有分类都显示|",
		"|view_url|varchar(128)|-|否||地址,media可空为meida详情地址|",
		"|sortrank|number(10)|-|否||显示顺序|",
	}
	tables := markdown2Strings(input)
	ut.ExpectSkip(t, len(tables), 3)
	tbs, err := strings2Tables(tables)
	ut.ExpectSkip(t, err, nil)
	ut.ExpectSkip(t, len(tbs), 3)
	ut.Expect(t, len(tbs[0].CNames), 5)
	ut.Expect(t, len(tbs[1].CNames), 11)
	ut.Expect(t, len(tbs[2].CNames), 10)
}

func TestTableex1(t *testing.T) {
	input := []string{
		"### 一、媒资信息",
		"#### 1. 媒资信息 [epg_media_info]",
		"|字段名|类型|默认值|为空|约束|描述|",
		"|---|---|---|---|---|---|",
	}
	tables := markdown2Strings(input)
	ut.ExpectSkip(t, len(tables), 1)
	tbs, err := strings2Tables(tables)
	ut.ExpectSkip(t, err, nil)
	ut.ExpectSkip(t, len(tbs), 0)
}
func TestTableex2(t *testing.T) {
	input := []string{
		"### 一、媒资信息",
		"#### 1. 媒资信息 [epg_media_info]",
		"|字段名|类型|默认值|为空|约束|描述|",
	}
	tables := markdown2Strings(input)
	ut.ExpectSkip(t, len(tables), 1)
	tbs, err := strings2Tables(tables)
	ut.ExpectSkip(t, err, nil)
	ut.ExpectSkip(t, len(tbs), 0)
}
func TestTableex3(t *testing.T) {
	input := []string{
		"### 一、媒资信息",
	}
	tables := markdown2Strings(input)
	ut.ExpectSkip(t, len(tables), 0)
	tbs, err := strings2Tables(tables)
	ut.ExpectSkip(t, err, nil)
	ut.ExpectSkip(t, len(tbs), 0)
}
