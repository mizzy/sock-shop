### Database

resource "aws_security_group" "db_ec2" {
  description = "Open database for access"

  ingress {
    protocol        = "tcp"
    from_port       = 3306
    to_port         = 3306
    security_groups = [aws_security_group.ecs.id]
  }

  tags = {
    "Name" = "db-ecs"
  }
}

resource "aws_db_subnet_group" "my_db_subnet_group" {
  description = "description"
  subnet_ids = [
    aws_subnet.public_subnet_1.id,
    aws_subnet.public_subnet_2.id,
  ]
}

resource "aws_db_instance" "catalogue" {
  instance_class         = "db.t2.medium"
  skip_final_snapshot    = true
  vpc_security_group_ids = [aws_security_group.db_ec2.id]
  db_subnet_group_name   = aws_db_subnet_group.my_db_subnet_group.name
  allocated_storage      = 100

  name     = "socksdb"
  engine   = "MySQL"
  username = "catalogue_user"
  password = "default_password"
}


### ECS

resource "aws_ecs_service" "catalogue" {
  name            = "sock-shop-CatalogueService-mVp9BfkdbXVD"
  desired_count   = 1
  launch_type     = "FARGATE"
  cluster         = aws_ecs_cluster.sock_shop.id
  task_definition = "${aws_ecs_task_definition.catalogue.id}:${aws_ecs_task_definition.catalogue.revision}"

  network_configuration {
    assign_public_ip = true
    security_groups  = [aws_security_group.ecs.id]
    subnets = [
      aws_subnet.public_subnet_1.id,
      aws_subnet.public_subnet_2.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.catalogue.arn
  }
}

resource "aws_ecs_task_definition" "catalogue" {
  container_definitions = jsonencode([
    {
      name = "catalogue",
      environment = [
        {
          name  = "ZIPKIN",
          value = "http://zipkin:9411/api/v1/spans",
        },
      ],
      command = [
        "/app",
        "-port=80",
        "-DSN=catalogue_user:default_password@tcp(${aws_db_instance.catalogue.address}:${aws_db_instance.catalogue.port})/socksdb",
      ],
      cpu                   = 0,
      dnsSearchDomains      = [],
      dnsServers            = [],
      dockerLabels          = {},
      dockerSecurityOptions = [],
      entryPoint            = [],
      environmentFiles      = [],
      extraHosts            = [],
      links                 = [],
      mountPoints           = [],
      secrets               = [],
      systemControls        = [],
      ulimits               = [],
      volumesFrom           = [],
      essential             = true,
      image                 = "weaveworksdemos/catalogue:0.3.5",
      logConfiguration = {
        secretOptions = [],
        logDriver     = "awslogs",
        options = {
          awslogs-group         = "sock-shop",
          awslogs-region        = data.aws_region.current.name,
          awslogs-stream-prefix = "catalogue",
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
  memory = 512

  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  family                   = "sock-shop-CatalogueTask-DfaZsRAivWhG"
  requires_compatibilities = ["FARGATE"]
}

resource "aws_service_discovery_service" "catalogue" {
  name = "catalogue"

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
