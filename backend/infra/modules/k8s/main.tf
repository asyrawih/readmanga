# Create kubernetes_deployment 
resource "kubernetes_deployment" "nginx" {
  metadata {
    name = "scalable-nginx-example"
    namespace = "terraform"
    labels = {
      App = "ScalableNginxExample"
    }
  }
  spec {
    replicas = 4
    selector {
      match_labels = {
        App = "ScalableNginxExample"
      }
    }
    template {
      metadata {
        labels = {
          App = "ScalableNginxExample"
        }
      }
      spec {
        container {
          image             = "nginx:latest"
          name              = "nginxserer"
          image_pull_policy = "IfNotPresent"
          port {
            container_port = 80
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
      port = 80
      target_port = 80
    }
    type = "NodePort"
  }
}

