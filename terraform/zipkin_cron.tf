module "ecs_zipkin_cron" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-ZipkinCronService-0TaWGFJtHFUC"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "zipkin-cron"
    image              = "openzipkin/zipkin-dependencies"
    family             = "sock-shop-ZipkinCronTask-9hMhrJVJioTd"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    environment = [
      {
        name  = "STORAGE_TYPE",
        value = "mysql",
      },
      {
        name  = "MYSQL_HOST",
        value = "zipkin-mysql",
      },
      {
        name  = "MYSQL_USER",
        value = "zipkin",
      },
      {
        name  = "MYSQL_PASS",
        value = "zipkin",
      }
    ]
    entryPoint = [
      "crond",
      "-f",
    ]
    memory = 1024
  }
}
