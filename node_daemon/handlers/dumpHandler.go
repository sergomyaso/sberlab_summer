package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const NodePort = ":7070"
const KettleHttpCode = 418
const OkHttpCode = 200
const TemplateScriptMemDump = "template_mem_script.sh"
const DumpMemScriptName = "dump_mem.sh"
const MemoryDumpRoute = "/dump/memory"
const CpuDumpRoute = "/dump/cpu"
const TemplateScriptCpuDump = "template_cpu_script.sh"
const DumpCpuScriptName = "dump_cpu.sh"

type DumpParams struct {
	TestName string `json:"test_name"`
	TestType string `json:"test_type"`
	Ip       string `json:"node_ip"`
	PodUid  string `json:"pod_uid"`
}

func CreateCheckIpScript(ip string) error {
	file, err := os.Create(IpScriptName)
	if err != nil {
		log.Printf("Creating script %s was failed, error:%v\n", IpScriptName, err)
		return err
	}
	err = file.Chmod(0777)
	if err != nil {
		log.Printf("Chmod for %s was failed, error:%v\n", IpScriptName, err)
		return err
	}
	defer file.Close()
	script := fmt.Sprintf(IpScriptTemplate, ip)
	file.WriteString(script)
	return err
}

func createDumpScript(params *DumpParams, templateScript string, scriptName string) {
	file, err := os.Open(templateScript)
	if err != nil {
		return
	}
	stringTemplate, _ := ioutil.ReadAll(file)
	file.Close()
	script := fmt.Sprintf(string(stringTemplate), params.PodUid)
	file, err = os.Create(scriptName)
	if err != nil {
		log.Printf("Creating %s script failed, error:%v\n", params.TestType, err)
		return
	}
	err = file.Chmod(0777)
	if err != nil {
		log.Printf("Chmod for %s was failed, error:%v\n", IpScriptName, err)
		return
	}
	defer file.Close()
	file.WriteString(script)

}

func RunDumpScript(params *DumpParams, templateScript string, dumpRoute string, scriptName string) int {
	log.Println("checking ip address for " + params.PodUid)
	//if !IsCurrentNodeIp(params.Ip) {
	//	log.Println("redirect to " + params.Ip)
	//	respCode := RedirectOnNode(params, dumpRoute)
	//	log.Printf("riderect with http code:%d\n", respCode)
	//	ClearPageCash()
	//	return respCode
	//}
	log.Println(params.Ip + " is ip address current node")
	createDumpScript(params, templateScript, scriptName)
	err := ExecScript(scriptName)
	if err != nil {
		ClearPageCash()
		return KettleHttpCode
	}
	log.Printf("%s dump for %s was succeed\n", params.TestType, params.PodUid)
	SaveDumpedData(params)
	//ClearPageCash()
	return OkHttpCode
}

func CreateDump(params *DumpParams) int {
	if params.TestType == "memory" {
		 return RunDumpScript(params, TemplateScriptMemDump, MemoryDumpRoute, DumpMemScriptName)
	}
	if params.TestType == "cpu" {
		return RunDumpScript(params, TemplateScriptCpuDump, CpuDumpRoute, DumpCpuScriptName)
	}
	log.Printf("Ivalid test type %s\n", params.TestType)
	return KettleHttpCode
}
