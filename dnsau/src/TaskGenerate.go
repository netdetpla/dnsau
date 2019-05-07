package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

const ConfPath = "/tmp/conf/busi.conf"

type Record struct {
	domain          string
	detectAs        []string
	detectCNames    []string
	authServerNames string
	timeoutFlag     bool
}

type Task struct {
	taskID   string
	taskName string
	uuid     string
	records  []*Record
}

func GetTaskConfig() (task *Task, err error) {
	task = new(Task)
	taskConfigBase64, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		return nil, err
	}
	taskConfigB, err := base64.StdEncoding.DecodeString(string(taskConfigBase64))
	if err != nil {
		return nil, err
	}
	taskConfig := strings.Split(string(taskConfigB), ";")
	fmt.Println(taskConfig)

	task.taskID = taskConfig[0]
	task.taskName = taskConfig[1]
	//组合域名、正确值
	domains := taskConfig[2 : len(taskConfig)-1]
	for _, domain := range domains {
		record := new(Record)
		record.domain = domain
		task.records = append(task.records, record)
	}
	task.uuid = taskConfig[len(taskConfig)-1]
	return
}
