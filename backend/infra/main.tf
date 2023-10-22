# Setup Provider kubernates
# Tell terraform for use kubernates as provider
provider "kubernetes" {
  host                   = var.host
  client_key             = base64decode(var.client_key)
  client_certificate     = base64decode(var.client_certificate)
  cluster_ca_certificate = base64decode(var.cluster_ca_certificate)
}

module "nginx_workloads" {
  source = "./modules/k8s"
}
