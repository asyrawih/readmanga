# check your cluster information  by running this commnad 
#  kubectl config view --minify --flatten
variable "host" {
  type = string
}

# You can get the value from this command
# kubectl config view --minify --flatten -o jsonpath='{.users[0].user}'
variable "client_certificate" {
  type = string
}

# kubectl config view --minify --flatten -o jsonpath='{.users[0].user}'
variable "client_key" {
  type = string
}

# cluster_ca_certificate
# kubectl config view --minify --flatten -o jsonpath='{.clusters}'
variable "cluster_ca_certificate" {
  type = string
}
