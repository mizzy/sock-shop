locals {
  task = defaults(var.task, {
    cpu    = 256
    memory = 512
  })
}

data "aws_region" "current" {}

resource "aws_ecs_task_definition" "main" {
  container_definitions = jsonencode([
    {
      name                  = local.task.name,
      command               = local.task.command,
      cpu                   = 0,
      dnsSearchDomains      = [],
      dnsServers            = [],
      dockerLabels          = local.task.dockerLabels,
      dockerSecurityOptions = [],
      entryPoint            = local.task.entryPoint,
      environment           = local.task.environment,
      environmentFiles      = [],
      extraHosts            = [],
      links                 = [],
      mountPoints           = local.task.mountPoints,
      secrets               = [],
      systemControls        = [],
      ulimits               = [],
      volumesFrom           = [],
      essential             = true,
      image                 = local.task.image,
      logConfiguration = {
        secretOptions = [],
        logDriver     = "awslogs",
        options = {
          awslogs-group         = "sock-shop",
          awslogs-region        = data.aws_region.current.name,
          awslogs-stream-prefix = local.task.name,
        }
      },
      portMappings = local.task.portMappings,
    }
  ])

  cpu    = local.task.cpu
  memory = local.task.memory

  execution_role_arn       = local.task.execution_role_arn
  task_role_arn            = local.task.task_role_arn
  family                   = local.task.family
  requires_compatibilities = ["FARGATE"]

  dynamic "volume" {
    for_each = local.task.volume != null ? [local.task.volume] : []
    content {
      name = volume.value
    }
  }
}
