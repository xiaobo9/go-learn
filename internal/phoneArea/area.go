/**
 * 电话归属地信息
 */
package phoneArea

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/xiaobo9/go-learn/config"
)

type Area struct {
	AreaCode string
	AreaName string
	Province string
}

// 缓存一下
var areas struct {
	values map[string]Area
}

func GetArea(phone string) Area {
	if areas.values == nil {
		areas.values = BuildArea()
	}
	return areas.values[phone]
}

type Areas []Area

func (a Areas) Len() int {
	return len(a)
}

func (a Areas) Less(i, j int) bool {
	ai, _ := strconv.Atoi(a[i].AreaCode)
	aj, _ := strconv.Atoi(a[j].AreaCode)
	return ai < aj
}

func (a Areas) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func ToFile(areas map[string]Area, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	var as Areas
	for _, v := range areas {
		as = append(as, v)
	}
	sort.Sort(as)
	for _, v1 := range as {
		var line = fmt.Sprintf("%s,%s,%s\n", v1.AreaCode, v1.AreaName, v1.Province)
		file.WriteString(line)
	}

}

func BuildArea() map[string]Area {
	var result = make(map[string]Area)
	
	file, err := os.OpenFile(config.CC.CsvFilePath, os.O_RDONLY, 0)
	if err != nil {
		log.Printf("file: [%s], error: %s", config.CC.CsvFilePath, err)
		return result
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				if len(lineData) > 0 {
					log.Println(lineData)
				}
				break
			}
		} else {
			var ss = strings.Split(strings.TrimSuffix(lineData, "\n"), ",")
			result[ss[0]] = Area{AreaCode: ss[0], AreaName: ss[1], Province: ss[2]}
		}
	}
	return result
}
