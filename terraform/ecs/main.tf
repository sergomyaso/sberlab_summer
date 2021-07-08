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
    access_key = "Y8L5MUQELXP58RGGDCWF"
    secret_key = "dSJ8CHeNjqSep7jScmiw7janbKtl7aCe24F2BLBv"
    project_name = "ru-moscow-1"
}
