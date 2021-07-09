package handlers

import (
	"bytes"
	"fmt"
	"html/template"
)

var configPage = `{{template "providerConfig" .ProviderConfig }}`

var configTemplate = `{{define "providerConfig"}}
terraform {
  required_version = ">= {{.TerraformVersion}}"
  required_providers {
    sbercloud = {
        source = "sbercloud-terraform/sbercloud"
        version = "{{.ProviderVersion}}"
        }
  }
}

# Configure the SberCloud Provider
provider "sbercloud" {
    region = "{{.Region}}"
    access_key = "{{.AccessKey}}"
    secret_key = "{{.SecretKey}}"
    project_name = "{{.ProjectName}}"
}

{{end}}`

type ProviderConfig struct {
	TerraformVersion string `json:"terraformVersion" description:"uuid of the todo"`
	ProviderVersion    string `json:"providerVersion" description:"uuid of the todo"`
	Region   string `json:"region" description:"uuid of the todo"`
	AccessKey   string `json:"accessKey" description:"uuid of the todo"`
	SecretKey   string `json:"secretKey" description:"uuid of the todo"`
	ProjectName string `json:"projectName" description:"uuid of the todo"`
}

type ConfigPage struct {
	ProviderConfig *ProviderConfig
}

func GetRenderConfigScript( configs *ProviderConfig) string {
	pageData := &ConfigPage{ProviderConfig:configs}
	tmpl := template.New("configPage")
	var err error
	if tmpl, err = tmpl.Parse(configPage); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(configTemplate); err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, pageData)
	return buf.String()
}
