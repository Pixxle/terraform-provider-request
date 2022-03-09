terraform {
  required_providers {
    terraform-provider-request = {
      source = "github.com/Pixxle/terraform-provider-request"
      version = "0.0.1"
    }
  }
}

provider "terraform-provider-request" {}
