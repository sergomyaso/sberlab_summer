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

func ExecScript(scriptName string) error {
	log.Println("Run script " + scriptName)
	cmd := exec.Command("./" + scriptName)
	err := cmd.Start()
	if err != nil {
		log.Printf("exec script %s, error:%v\n", scriptName, err)
		return err
	}
	err = cmd.Wait()
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
	var jsonStr = []byte(fmt.Sprintf(JsonRequestTemplate, params.Ip, params.PodName))
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
