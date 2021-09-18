resource "aws_security_group" "db_ec2" {
  description = "Open database for access"

  ingress {
    protocol        = "tcp"
    from_port       = 3306
    to_port         = 3306
    security_groups = [aws_security_group.rds_importer_allowed_ports.id]
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
