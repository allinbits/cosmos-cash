terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
    helm = {
      source = "hashicorp/helm"
    }
    null = {
      source = "hashicorp/null"
    }
  }
  required_version = ">= 0.13"
}
