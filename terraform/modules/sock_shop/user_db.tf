module "ecs_user_db" {
  source = "../ecs"

  service = {
    name               = "sock-shop-UserDBService-bcNIEWFwyFiJ"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "user-db"
    image              = "weaveworksdemos/user-db"
    family             = "sock-shop-UserDBTask-JH16IOxl33fR"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 27017
        hostPort      = 27017
      }
    ]
    dockerLabels = {
      "agent.signalfx.com.port.27017" = "true",
    }
  }

  registry = {
    name         = "user-db"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
