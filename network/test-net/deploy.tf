provider "kubernetes" {
  config_path    = "~/.kube/config"
}

module "cash" {
  source          = "../deployment"
  docker_image    = "paddymctm/cash"
  deployment_name = "cash"
  namespace       = "ghost"
}

module "gaia" {
  source          = "../deployment"
  docker_image    = "paddymctm/gaia"
  deployment_name = "gaia"
  namespace       = "ghost"
}

module "rly" {
  source          = "../deployment"
  docker_image    = "paddymctm/rly"
  deployment_name = "rly"
  namespace       = "ghost"
}
