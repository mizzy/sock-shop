resource "aws_ecs_service" "carts_db" {
  name            = "sock-shop-CartsDBService-mPiXfiIT9Aky"
  desired_count   = 1
  launch_type     = "FARGATE"
  cluster         = aws_ecs_cluster.sock_shop.id
  task_definition = "${aws_ecs_task_definition.carts_db.id}:${aws_ecs_task_definition.carts_db.revision}"

  network_configuration {
    assign_public_ip = true
    security_groups  = [aws_security_group.rds_importer_allowed_ports.id]
    subnets = [
      aws_subnet.public_subnet_1.id,
      aws_subnet.public_subnet_2.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.carts_db.arn
  }
}

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

resource "aws_service_discovery_service" "carts_db" {
  name = "carts-db"

  dns_config {
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
    dns_records {
      ttl  = 10
      type = "A"
    }
  }

  health_check_custom_config {
    failure_threshold = 1
  }
}
