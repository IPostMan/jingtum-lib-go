/***  读取配置文件
*** config.go
*** 主要用于读取配置文件数据
*** author:              1416205324@qq.com
*** last_modified_time:  2018-6-6 23:13:23
 */

package jingtumLib

import (
	"bufio"
	"io"
	"os"
	"strings"
    "strconv"
)

const middle = "========="

type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) error {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		Error(err)
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			Error(err)
			return err
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
	Info("Init config succ.")
	return nil
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}

func (c Config) ReadInt(node, key string, defaultv int) int {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return defaultv
	}
	return strconv.Atoi(v)
}
