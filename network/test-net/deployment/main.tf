resource "kubernetes_deployment" "main" {
  metadata {
    name = var.deployment_name
    namespace = var.namespace
    labels = {
      app = var.deployment_name
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = var.deployment_name
      }
    }

    template {
      metadata {
        labels = {
          app = var.deployment_name
        }
      }

      spec {
        container {
          image = var.docker_image
          name  = var.deployment_name
        }
      } 
    }
  }
}

resource "kubernetes_service" "main" {
  metadata {
    name = var.deployment_name
    namespace = var.namespace
  }
  spec {
    selector = {
      app = var.deployment_name
    }
    session_affinity = "ClientIP"
    port {
      name = "1"
      port        = 12345
      target_port = 12345
    }
    port {
      name = "rest-server"
      port        = 1317
      target_port = 1317
    }
    port {
      name = "tendermint"
      port        = 26657
      target_port = 26657
    }
    port {
      name = "gossip"
      port        = 26656
      target_port = 26656
    }
  }
}
