package handlers

import (
	"bytes"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const IpScriptName = "ip_check.sh"
const IpLogName = "ip_req_log.txt"
const IpScriptTemplate = "#!/bin/sh\nifconfig eth0 | grep %s | cat >> " + IpLogName
const JsonRequestTemplate = "{\"test_name\":\"%s\",\"test_type\":\"%s\",\"node_ip\":\"%s\",\"pod_uid\":\"%s\"}"
const ClearPageCashScriptName = "clear_page_cash.sh"

const TemplateMemoryDataPath = "mem_pods_dump/kubepods-pod%s.slice/"
const MemMaxUsageFile = "memory.max_usage_in_bytes"
const MemLimitsFile = "memory.limit_in_bytes"

const TemplateCpuDataPath = "cpu_pods_dump/kubepods-pod%s.slice/"
const CpuUsageFile = "cpuacct.usage"
const CpuSharesFile = "cpu.shares"
const CpuPeriodFile = "cpu.cfs_period_us"
const CpuQuotaFile = "cpu.cfs_quota_us"
const CpuThrottlingDataFile = "cpu.stat"
const KeyPeriods = "nr_periods"
const KeyThrottlingPeriods = "nr_throttled"
const KeyThrottlingTime = "throttled_time"

func ExecScript(scriptName string) error {
	log.Println("Run script " + scriptName)
	cmd := exec.Command("./" + scriptName)
	err := cmd.Start()
	if err != nil {
		log.Printf("exec script %s, error:%v\n", scriptName, err)
		return err
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("exec script %s, error:%v\n", scriptName, err)
		return err
	}
	return err
}

func IsIpScriptLog() bool {
	file, err := os.Open(IpLogName)
	if err != nil {
		log.Println("Open error " + IpLogName)
		return false
	}
	buffer, _ := ioutil.ReadAll(file)
	file.Close()
	os.RemoveAll(file.Name())
	if string(buffer) == "" {
		return false
	}
	return true
}

func IsCurrentNodeIp(ip string) bool {
	err := CreateCheckIpScript(ip)
	if err != nil {
		return false
	}
	err = ExecScript(IpScriptName)
	if err != nil {
		return false
	}
	return IsIpScriptLog()
}

func RedirectOnNode(params *DumpParams, route string) int {
	var jsonStr = []byte(fmt.Sprintf(JsonRequestTemplate, params.TestName, params.TestType, params.Ip, params.PodUid))
	req, err := http.NewRequest("POST", "http://"+params.Ip+NodePort+route, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)
	client := &http.Client{}
	resp := new(http.Response)
	resp, err = client.Do(req)

	if err != nil {
		log.Printf("redirecting on %s failed, error:%v\n", params.Ip, err)
		return KettleHttpCode
	}
	return resp.StatusCode
}

func ClearPageCash() {
	err := ExecScript(ClearPageCashScriptName)
	if err != nil {
		return
	}
}

func GetDataFromDumpFile(templatePath string, podUid string, fileName string) string {
	filePath := fmt.Sprintf(templatePath, podUid) + fileName
	fContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Read dumped file %s was failed, error%v\n:", filePath, err)
		panic(err)
	}
	return strings.Split(string(fContent), "\n")[0]
}

func GetDataFromThrottlingFile(templatePath string, podUid string, fileName string) map[string]int64 {
	filePath := fmt.Sprintf(templatePath, podUid) + fileName
	fContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Read dumped file %s was failed, error%v\n:", filePath, err)
		panic(err)
	}
	log.Println(string(fContent))
	arrParams := strings.Split(string(fContent), "\n")
	arrParams = arrParams[:len(arrParams)-1]
	log.Println(arrParams)
	paramsMap := make(map[string]int64)
	for _, param := range arrParams {
		nameValueArr := strings.Split(param, " ")
		log.Println(nameValueArr)
		paramsMap[nameValueArr[0]], _ = strconv.ParseInt(nameValueArr[1], 0, 64)
	}

	return paramsMap
}

func SaveDumpedMemoryData(params *DumpParams) {
	log.Println("Try to parse memory files")
	limits, _ := strconv.ParseInt(GetDataFromDumpFile(TemplateMemoryDataPath, params.PodUid, MemLimitsFile), 0, 64)
	maxUsage, _ := strconv.ParseInt(GetDataFromDumpFile(TemplateMemoryDataPath, params.PodUid, MemMaxUsageFile), 0, 64)
	memoryModel := MemoryModel{TestName: params.TestName, Limits: limits, MaxUsage: maxUsage}
	log.Println(limits)
	log.Println(memoryModel.Limits)
	InsertMemoryLogsToDb(&memoryModel)
}

func SaveDumpedCpuData(params *DumpParams) {
	usage, _ := strconv.ParseInt(GetDataFromDumpFile(TemplateCpuDataPath, params.PodUid, CpuUsageFile), 0, 64)
	shares, _ := strconv.ParseInt(GetDataFromDumpFile(TemplateCpuDataPath, params.PodUid, CpuSharesFile), 0, 64)
	cpuPeriod, _ := strconv.ParseInt(GetDataFromDumpFile(TemplateCpuDataPath, params.PodUid, CpuPeriodFile), 0, 64)
	cpuQuota, _ := strconv.ParseInt(GetDataFromDumpFile(TemplateCpuDataPath, params.PodUid, CpuQuotaFile), 0, 64)
	throttlingParamMap := GetDataFromThrottlingFile(TemplateCpuDataPath, params.PodUid, CpuThrottlingDataFile)
	cpuModel := CpuModel{
		TestName:          params.TestName,
		Usage:             usage,
		Shares:            shares,
		CpuPeriod:         cpuPeriod,
		CpuQuota:          cpuQuota,
		Periods:           throttlingParamMap[KeyPeriods],
		ThrottlingPeriods: throttlingParamMap[KeyThrottlingPeriods],
		ThrottlingTime:    throttlingParamMap[KeyThrottlingTime],
	}
	log.Println(cpuModel.CpuPeriod, cpuModel.Periods, cpuModel.ThrottlingPeriods, cpuModel.Shares, cpuModel.Usage, cpuModel.TestName)
	InsertCpuLogsToDb(&cpuModel)
}

func SaveDumpedData(params *DumpParams) {
	if params.TestType == "memory" {
		SaveDumpedMemoryData(params)
	}
	if params.TestType == "cpu" {
		SaveDumpedCpuData(params)
	}

}
