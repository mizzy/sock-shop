resource "aws_ecs_task_definition" "carts_db" {
  container_definitions = jsonencode([
    {
      name                  = "carts-db",
      command               = [],
      cpu                   = 0,
      dnsSearchDomains      = [],
      dnsServers            = [],
      dockerLabels          = {},
      dockerSecurityOptions = [],
      entryPoint            = [],
      environment           = [],
      environmentFiles      = [],
      extraHosts            = [],
      links                 = [],
      mountPoints           = [],
      secrets               = [],
      systemControls        = [],
      ulimits               = [],
      volumesFrom           = [],
      essential             = true,
      image                 = "mongo",
      logConfiguration = {
        secretOptions = [],
        logDriver     = "awslogs",
        options = {
          awslogs-group         = "sock-shop",
          awslogs-region        = data.aws_region.current.name,
          awslogs-stream-prefix = "carts-db",
        }
      },
      portMappings = [
        {
          containerPort = 27017,
          hostPort      = 27017,
          protocol      = "tcp",
        }
      ],
    }
  ])

  cpu    = 256
  memory = 512

  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  family                   = "sock-shop-CartsDBTask-6LKqO3WQLkdN"
  requires_compatibilities = ["FARGATE"]
}
