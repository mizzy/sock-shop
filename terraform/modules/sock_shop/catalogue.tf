### Database

resource "aws_security_group" "db_ec2" {
  vpc_id      = aws_vpc.sock_shop.id
  description = "Open database for access"

  tags = {
    "Name" = "db-ecs"
  }
}

resource "aws_security_group_rule" "db_ec2_allow_from_sg" {
  type                     = "ingress"
  security_group_id        = aws_security_group.db_ec2.id
  from_port                = 3306
  to_port                  = 3306
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.ecs.id
}

resource "aws_security_group_rule" "db_ec2_allow_to_all" {
  type              = "egress"
  security_group_id = aws_security_group.db_ec2.id
  from_port         = 0
  to_port           = 65536
  protocol          = -1
  cidr_blocks       = ["0.0.0.0/0"]
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
module "ecs_catalogue" {
  source = "../ecs"

  service = {
    name               = "sock-shop-CatalogueService-mVp9BfkdbXVD"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "catalogue"
    image              = "weaveworksdemos/catalogue:0.3.5"
    family             = "sock-shop-CatalogueTask-DfaZsRAivWhG"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 80
        hostPort      = 80
      }
    ]
    command = [
      "/app",
      "-port=80",
      "-DSN=catalogue_user:default_password@tcp(${aws_db_instance.catalogue.address}:${aws_db_instance.catalogue.port})/socksdb",
    ]
    environment = [
      {
        name  = "ZIPKIN"
        value = "http://zipkin:9411/api/v1/spans"
      }
    ]
  }

  registry = {
    name         = "catalogue"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
