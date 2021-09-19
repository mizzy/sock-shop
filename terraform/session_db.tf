module "ecs_session_db" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-SessionDBService-QvCqf3gtThsA"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
    load_balancer      = null
  }

  task = {
    name               = "session-db"
    image              = "redis:alpine"
    family             = "sock-shop-SessionDBTask-lmz8Mx5Cmgth"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 6379
    memory             = 512
    command            = []
    environment        = []
    dockerLabels = {
      "agent.signalfx.com.port.6379" = "true"
    }
  }

  registry = {
    name         = "session-db"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
