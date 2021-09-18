resource "aws_ecs_task_definition" "carts" {
  container_definitions = jsonencode([
    {
      name                  = "carts",
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
      image                 = "weaveworksdemos/carts:0.4.8",
      logConfiguration = {
        secretOptions = [],
        logDriver     = "awslogs",
        options = {
          awslogs-group         = "sock-shop",
          awslogs-region        = data.aws_region.current.name,
          awslogs-stream-prefix = "carts",
        }
      },
      portMappings = [
        {
          containerPort = 80,
          hostPort      = 80,
          protocol      = "tcp",
        }
      ],
    }
  ])

  cpu    = 256
  memory = 1024

  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  family                   = "sock-shop-CartsTask-eIq5v1xKpl13"
  requires_compatibilities = ["FARGATE"]
}
