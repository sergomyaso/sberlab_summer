package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

var ecsPage = `{{template "providerConfig" .EcsParams }}

{{template "ecsParams" .EcsParams}}
`

var ecsTemplate = `{{define "ecsParams"}}

# Get the latest Ubuntu image
data "sbercloud_images_image" "ubuntu_image" {
  name        = "{{.ImageTitle}}"
  most_recent = true
}

# Get the subnet where ECS will be created
data "sbercloud_vpc_subnet" "subnet_01" {
  name = "{{.SubnetName}}"
}

# Create ECS
resource "sbercloud_compute_instance" "ecs_01"{
  name              = "{{.Name}}"
  image_id          = {{.ImageId}}
  flavor_id         = "{{.FlavorId}}"
  security_groups   = ["{{.SecGroup}}"]
  availability_zone = "ru-moscow-1a"
  admin_pass        = "adminAdmin123"

  system_disk_type = "SAS"
  system_disk_size = {{.DiskSize}}

 network {
    uuid = data.sbercloud_vpc_subnet.subnet_01.id
  }
}
{{end}}`

type EcsParams struct {
	
	TerraformVersion string `json:"terraformVersion" description:"uuid of the todo"`
	ProviderVersion    string `json:"providerVersion" description:"uuid of the todo"`
	Region   string `json:"region" description:"uuid of the todo"`
	AccessKey   string `json:"accessKey" description:"uuid of the todo"`
	SecretKey   string `json:"secretKey" description:"uuid of the todo"`
	ProjectName string `json:"projectName" description:"uuid of the todo"`

	Name       string `json:"name" description:"uuid of the todo"`
	ImageId    string `json:"imageId" description:"uuid of the todo"`
	FlavorId   string `json:"flavorId" description:"uuid of the todo"`
	SecGroup   string `json:"secGroup" description:"uuid of the todo"`
	DiskSize   int `json:"diskSize" description:"uuid of the todo"`
	SubnetName string `json:"subnetName" description:"uuid of the todo"`
	ImageTitle string `json:"imageTitle" description:"uuid of the todo"`
}

type EcsPage struct {
	EcsParams *EcsParams
}

func GetRenderEcsScript(params *EcsParams) string {
	pageData := &EcsPage{EcsParams: params}
	tmpl := template.New("ecsPage")
	var err error
	if tmpl, err = tmpl.Parse(ecsPage); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(configTemplate); err != nil {
		fmt.Println(err)
	}
	if tmpl, err = tmpl.Parse(ecsTemplate); err != nil {
		fmt.Println(err)
	}


	var buf bytes.Buffer
	tmpl.Execute(&buf, pageData)
	log.Println(buf.String())
	return buf.String()
}
