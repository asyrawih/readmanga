# Create kubernetes_deployment 
resource "kubernetes_deployment" "nginx" {
  metadata {
    name = "nginx-server-deployment"
    namespace = "terraform"
    labels = {
      App = "nginx-server"
    }
  }
  spec {
    replicas = var.replica_set
    selector {
      match_labels = {
        App = "nginx-server"
      }
    }
    template {
      metadata {
        labels = {
          App = "nginx-server"
        }
      }
      spec {
        container {
          image             = "nginx:latest"
          name              = "nginx-server"
          image_pull_policy = "IfNotPresent"
          port {
            container_port = var.nginx_port
          }
        }
      }
    }
  }
}

# Create Service
resource "kubernetes_service" "nginx_terraform" {
  metadata {
    namespace = "terraform"
    name = "nginx-server"
  }
  spec {
    selector = {
      App = kubernetes_deployment.nginx.spec[0].template[0].metadata[0].labels.App
    }
    port {
      node_port = 30000 
      port = var.nginx_port
      target_port = var.nginx_port
    }
    type = var.service_type
  }
}

