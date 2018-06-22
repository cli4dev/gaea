package md

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/micro-plat/gaea/cmds/new/sql"
)

//Markdown2Table 读取markdown文件并转换为table对象
func Markdown2Table(fn string) ([]*sql.Table, error) {
	lines, err := readMarkdown(fn)
	if err != nil {
		return nil, err
	}
	tlines := markdown2Strings(lines)
	return strings2Tables(tlines)
}

//readMarkdown 读取md文件
func readMarkdown(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines := make([]string, 0, 64)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}

//markdown2Strings
func markdown2Strings(lines []string) [][]string {
	tables := make([][]string, 0, 10)
	index := make([]int, 0, 10)
	for i, line := range lines {
		if line == "|字段名|类型|默认值|为空|约束|描述|" {
			index = append(index, i-1)
		}
		if len(index)%2 == 1 && strings.Count(line, "|") != 7 {
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
func strings2Tables(tbs [][]string) ([]*sql.Table, error) {
	tables := make([]*sql.Table, 0, len(tbs))
	for _, tb := range tbs {
		if len(tb) <= 3 {
			continue
		}
		var name string
		var des string
		var err error
		var table *sql.Table
		for i, line := range tb {
			if i == 0 {
				//获取表名，描述名称
				if name, err = getTableName(line); err != nil {
					return nil, err
				}
				if des, err = getTableDesc(line); err != nil {
					return nil, err
				}
				table = sql.NewTable(name, des)
				continue
			}
			if i < 3 {
				continue
			}
			if strings.Count(line, "|") != 7 {
				return nil, fmt.Errorf("表结构有误")
			}

			colums := strings.Split(strings.Trim(line, "|"), "|")
			if colums[0] == "" {
				return nil, fmt.Errorf("表%s的字段名称不能为空 %v", name, colums)
			}
			tp, l, err := getType(colums[1])
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
func getTableDesc(line string) (string, error) {
	reg := regexp.MustCompile(`[^\d^\.]+[\p{Han}]+[^\[]+`)
	names := reg.FindAllString(line, -1)
	if len(names) == 0 {
		return "", fmt.Errorf("未设置表描述:%s", line)
	}
	return strings.TrimSpace(names[0]), nil
}
func getTableName(line string) (string, error) {
	reg := regexp.MustCompile(`\[[\w]+\]`)
	names := reg.FindAllString(line, -1)
	if len(names) == 0 {
		return "", fmt.Errorf("未设置表名称:%s", line)
	}
	return strings.TrimRight(strings.TrimLeft(names[0], "["), "]"), nil
}
func getType(line string) (string, int, error) {
	reg := regexp.MustCompile(`[\w]+`)
	names := reg.FindAllString(line, -1)
	if len(names) == 0 || len(names) > 2 {
		return "", 0, fmt.Errorf("未设置字段类型:%v", names)
	}
	if len(names) == 1 {
		return names[0], 0, nil
	}
	i, err := strconv.ParseInt(names[1], 10, 32)
	return names[0], int(i), err
}
