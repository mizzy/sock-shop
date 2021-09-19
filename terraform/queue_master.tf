module "ecs_queue_master" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-QueueMasterService-aJBpJZ5ztMnu"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "queue-master"
    volume             = "VarRunDocker_Sock"
    image              = "weaveworksdemos/queue-master"
    family             = "sock-shop-QueueMasterTask-WrLvLSJ3b4Sz"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 80
    memory             = 1024
    task_role_arn      = null
    mountPoints = [
      {
        containerPath = "/var/run/docker.sock"
        sourceVolume  = "VarRunDocker_Sock"
      },
    ]
  }

  registry = {
    name         = "queue-master"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
