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
)

const IpScriptName = "ip_check.sh"
const IpLogName = "ip_req_log.txt"
const IpScriptTemplate = "#!/bin/sh\nifconfig eth0 | grep %s | cat >> " + IpLogName
const JsonRequestTemplate = "{\"node_ip\":\"%s\",\"pod_name\":\"%s\"}"
const NodePort = ":7070"
const KettleHttpCode = 418
const OkHttpCode = 200
const TemplateScriptMemDump = "template_mem_script.sh"
const DumpMemScriptName = "dump_mem.sh"

type DumpParams struct {
	Ip      string `json:"node_ip"`
	PodName string `json:"pod_name"`
}

func execScript(scriptName string) error {
	log.Println("Run script " + scriptName)
	cmd := exec.Command("./" + scriptName)
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}

func CreateCheckIpScript(ip string) (string, error) {
	file, err := os.Create(IpScriptName)
	if err != nil {
		return "", err
	}
	file.Chmod(0777)
	defer file.Close()
	script := fmt.Sprintf(IpScriptTemplate, ip)
	log.Println(script)
	file.WriteString(script)
	return file.Name(), err
}

func isIpScriptLog() bool {
	file, err := os.Open(IpLogName)
	if err != nil {
		log.Println("Open error " + IpLogName)
		return false
	}
	buffer, _ := ioutil.ReadAll(file)
	log.Println(buffer)

	file.Close()
	os.RemoveAll(file.Name())
	if string(buffer) == "" {
		return false
	}
	return true
}

func isCurrentNodeIp(ip string) bool {
	fileName, err := CreateCheckIpScript(ip)
	if err != nil {
		//#TODO error!
		return false
	}
	execScript(fileName)
	return isIpScriptLog()
}

func redirectOnNode(params *DumpParams) int {
	var jsonStr = []byte(fmt.Sprintf(JsonRequestTemplate, params.Ip, params.PodName))
	req, err := http.NewRequest("POST",  "http://"+ params.Ip+NodePort + "/dump/memory", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", restful.MIME_JSON)
	// #TODO add route to request
	client := &http.Client{}
	resp := new(http.Response)
	resp, err = client.Do(req)

	if err != nil {
		log.Println(err)
		return KettleHttpCode
	}
	return resp.StatusCode
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
		return
	}
	file.Chmod(0777)
	defer file.Close()
	file.WriteString(script)

}

func CreateMemoryDump(params *DumpParams) int {
	if !isCurrentNodeIp(params.Ip) {
		respCode := redirectOnNode(params)
		log.Println("redirect to " + params.Ip)
		return respCode
	}
	log.Println("ip is correct")
	createMemoryDumpScript(params)
	execScript(DumpMemScriptName)
	return OkHttpCode
}
