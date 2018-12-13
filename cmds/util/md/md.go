package md

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/micro-plat/gaea/cmds/util/conf"
)

type Line struct {
	Text   string
	LineID int
}

//Markdown2Table 读取markdown文件并转换为table对象
func Markdown2Table(fn string) ([]*conf.Table, error) {
	lines, err := readMarkdown(fn)
	if err != nil {
		return nil, err
	}
	tlines := markdown2Strings(lines)
	return strings2Tables(tlines)
}

//readMarkdown 读取md文件
func readMarkdown(name string) ([]*Line, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines := make([]*Line, 0, 64)
	rd := bufio.NewReader(f)
	num := 0
	for {
		num++
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Trim(line, "\n")
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines = append(lines, &Line{Text: line, LineID: num})
	}
	return lines, nil
}

//markdown2Strings
func markdown2Strings(lines []*Line) [][]*Line {
	tables := make([][]*Line, 0, 10)
	index := make([]int, 0, 10)
	for i, line := range lines {
		nline := strings.TrimSpace(strings.Replace(line.Text, " ", "", -1))
		if nline == "|字段名|类型|默认值|为空|约束|描述|" {
			index = append(index, i-1)
		}
		if len(index)%2 == 1 && strings.Count(nline, "|") != 7 {
			index = append(index, i-1)
		}
	}
	if len(index)%2 == 1 {
		index = append(index, len(lines)-1)
	}

	for i := 0; i < len(index); i = i + 2 {
		tables = append(tables, lines[index[i]:index[i+1]+1])
	}
	return tables
}
func strings2Tables(tbs [][]*Line) ([]*conf.Table, error) {
	tables := make([]*conf.Table, 0, len(tbs))
	for _, tb := range tbs {
		if len(tb) <= 3 {
			continue
		}
		var name string
		var des string
		var err error
		var table *conf.Table
		for i, line := range tb {
			if i == 0 {
				//获取表名，描述名称
				if name, err = getTableName(line); err != nil {
					return nil, err
				}
				if des, err = getTableDesc(line); err != nil {
					return nil, err
				}
				table = conf.NewTable(name, des)
				continue
			}
			if i < 3 {
				continue
			}
			if strings.Count(line.Text, "|") != 7 {
				return nil, fmt.Errorf("表结构有误(行:%d)", line.LineID)
			}
			colums := strings.Split(strings.Trim(line.Text, "|"), "|")
			if colums[0] == "" {
				return nil, fmt.Errorf("字段名称不能为空 %s(行:%d)", line.Text, line.LineID)
			}
			tp, l, err := getType(line)
			if err != nil {
				return nil, err
			}
			if err := table.AppendColumn(strings.TrimSpace(colums[0]), tp, l, strings.TrimSpace(colums[2]), strings.TrimSpace(colums[3]) == "否", strings.TrimSpace(colums[4]), strings.TrimSpace(colums[5])); err != nil {
				return nil, err
			}
		}
		if table != nil {
			tables = append(tables, table)
		}
	}
	return tables, nil
}
func getTableDesc(line *Line) (string, error) {
	reg := regexp.MustCompile(`[^\d^\.]+[\p{Han}]+[^\[]+`)
	names := reg.FindAllString(line.Text, -1)
	if len(names) == 0 {
		return "", fmt.Errorf("未设置表描述:%s(行:%d)", line.Text, line.LineID)
	}
	return strings.TrimSpace(names[0]), nil
}
func getTableName(line *Line) (string, error) {
	reg := regexp.MustCompile(`\[[\w]+\]`)
	names := reg.FindAllString(line.Text, -1)
	if len(names) == 0 {
		return "", fmt.Errorf("未设置表名称:%s(行:%d)", line.Text, line.LineID)
	}
	return strings.TrimRight(strings.TrimLeft(names[0], "["), "]"), nil
}
func getType(line *Line) (string, string, error) {
	colums := strings.Split(strings.Trim(line.Text, "|"), "|")
	if colums[0] == "" {
		return "", "", fmt.Errorf("字段名称不能为空 %s(行:%d)", line.Text, line.LineID)
	}
	reg := regexp.MustCompile(`[\w]+`)
	names := reg.FindAllString(colums[1], -1)
	if len(names) == 0 || len(names) > 3 {
		return "", "", fmt.Errorf("未设置字段类型:%v(行:%d)", names, line.LineID)
	}
	if len(names) == 1 {
		return colums[1], "", nil
	}
	return colums[1], strings.Join(names[1:], ","), nil
}
