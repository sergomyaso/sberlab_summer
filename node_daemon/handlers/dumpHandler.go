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
const MemoryDumpRoute = "dump/memory"

type DumpParams struct {
	Ip      string `json:"node_ip"`
	PodName string `json:"pod_name"`
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

func createMemoryDumpScript(params *DumpParams) {
	file, err := os.Open(TemplateScriptMemDump)
	if err != nil {
		return
	}
	stringTemplate, _ := ioutil.ReadAll(file)
	file.Close()
	script := fmt.Sprintf(string(stringTemplate), params.PodName)
	file, err = os.Create(DumpMemScriptName)
	if err != nil {
		log.Printf("Creating memory script failed, error:%v\n", err)
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

func CreateMemoryDump(params *DumpParams) int {
	log.Println("checking ip address for " + params.PodName)
	if !IsCurrentNodeIp(params.Ip) {
		log.Println("redirect to " + params.Ip)
		respCode := RedirectOnNode(params, MemoryDumpRoute)
		log.Printf("riderect with http code:%d\n", respCode)
		return respCode
	}
	log.Println(params.Ip + " is ip address current node")
	createMemoryDumpScript(params)
	err := ExecScript(DumpMemScriptName)
	if err != nil {
		return KettleHttpCode
	}
	log.Printf("memory dump for %s was succeed\n", params.PodName)
	return OkHttpCode
}
