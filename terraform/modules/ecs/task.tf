data "aws_region" "current" {}

resource "aws_ecs_task_definition" "main" {
  container_definitions = jsonencode([
    {
      name                  = var.task.name,
      command               = var.task.command,
      cpu                   = 0,
      dnsSearchDomains      = [],
      dnsServers            = [],
      dockerLabels          = var.task.dockerLabels,
      dockerSecurityOptions = [],
      entryPoint            = [],
      environment           = var.task.environment,
      environmentFiles      = [],
      extraHosts            = [],
      links                 = [],
      mountPoints           = [],
      secrets               = [],
      systemControls        = [],
      ulimits               = [],
      volumesFrom           = [],
      essential             = true,
      image                 = var.task.image,
      logConfiguration = {
        secretOptions = [],
        logDriver     = "awslogs",
        options = {
          awslogs-group         = "sock-shop",
          awslogs-region        = data.aws_region.current.name,
          awslogs-stream-prefix = var.task.name,
        }
      },
      portMappings = [
        {
          containerPort = var.task.port,
          hostPort      = var.task.port,
          protocol      = "tcp",
        }
      ],
    }
  ])

  cpu    = 256
  memory = var.task.memory

  execution_role_arn       = var.task.execution_role_arn
  task_role_arn            = var.task.task_role_arn
  family                   = var.task.family
  requires_compatibilities = ["FARGATE"]
}
