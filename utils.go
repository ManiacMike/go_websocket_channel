package main

import (
	"github.com/bitly/go-simplejson" // for json get
	"strconv"
	"strings"
	"time"
)

const LINE_SEPARATOR = "#LINE_SEPARATOR#"

func JsonStrToStruct(jsonStr string) map[string]interface{} {
	jsonStr = strings.Replace(jsonStr, "\n", LINE_SEPARATOR, -1)
	json, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err.Error())
	}
	var nodes = make(map[string]interface{})
	nodes, _ = json.Map()
	return nodes
}

func GenerateId() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
