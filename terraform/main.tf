terraform {
  required_version = ">= 0.12"
}

provider "sbercloud" {
      region     = "ru-moscow-1"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}

