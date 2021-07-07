terraform {
  required_version = ">= 0.12"
  required_providers {
    sbercloud = {
        source = "sbercloud-terraform/sbercloud"
        version = "1.3.0"
        }
  }
}

# Configure the SberCloud Provider
provider "sbercloud" {
    region = "ru-moscow-1"
    access_key = ""
    secret_key = ""
    project_name = "summer_school"
}
