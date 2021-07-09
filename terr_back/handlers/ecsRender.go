package handlers

import (
	"bytes"
	"fmt"
	"html/template"
)

var ecsPage = `{{template "ecsParams" .EcsParams}}`

var ecsTemplate = `{{define "ecsParams"}}
# Declare all required input variables
variable "root_password" {
  description = "Root password for ECS"
  sensitive   = true
}

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
  admin_pass        = "admin"

  system_disk_type = "SAS"
  system_disk_size = {{.DiskSize}}

 network {
    uuid = data.sbercloud_vpc_subnet.subnet_01.id
  }
}
{{end}}`


type EcsParams struct {
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
	if tmpl, err = tmpl.Parse(ecsTemplate); err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, pageData)
	return buf.String()
}

/*func m() {
	/*pagedata :=EcsParams: &EcsParams{Name: name, ImageId: imageId, FlavorId: flavorId,
	SecGroup: secGroup, DiskSize: diskSize, SubnetName: subnetName, ImageTitle: imageTitle}*/
	/*params := &EcsParams{Name: "name", ImageId: "imageId", FlavorId: "flavorId",
		SecGroup: "secGroup", DiskSize: 12, SubnetName: "subnetName", ImageTitle: "imageTitle"}
	println(GetRenderEcsScript(configPage, ecsTemplate, params))
}*/