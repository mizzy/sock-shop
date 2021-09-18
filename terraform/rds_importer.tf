resource "aws_security_group" "rds_importer_allowed_ports" {
  # This description is from the original cloudformation template, but it seems not appropriate
  description = "ECS Allowed Ports"
  tags = {
    "Name" = "rds-importer-allowed-ports"
  }
}

resource "aws_security_group_rule" "rds_importer_allowed_port_ssh" {
  type              = "ingress"
  security_group_id = aws_security_group.rds_importer_allowed_ports.id
  protocol          = "tcp"
  from_port         = 22
  to_port           = 22
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "rds_importer_allowed_from_elb" {
  type                     = "ingress"
  security_group_id        = aws_security_group.rds_importer_allowed_ports.id
  protocol                 = -1
  from_port                = 0
  to_port                  = 65536
  source_security_group_id = aws_security_group.elb_allowed_ports.id
}

resource "aws_security_group_rule" "rds_importer_allowed_from_self" {
  type                     = "ingress"
  security_group_id        = aws_security_group.rds_importer_allowed_ports.id
  protocol                 = -1
  from_port                = 0
  to_port                  = 65536
  source_security_group_id = aws_security_group.rds_importer_allowed_ports.id
}

resource "aws_instance" "rds_importer" {
  ami                         = "ami-56d4ad31"
  instance_type               = "t2.micro"
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.rds_importer_allowed_ports.id]
  subnet_id                   = aws_subnet.public_subnet_1.id
  key_name                    = "sock-shop-cloudformation"

  user_data = <<EOF
#!/bin/bash -xe
wget https://raw.githubusercontent.com/microservices-demo/catalogue/master/docker/catalogue-db/data/dump.sql
mysql -u catalogue_user --password=default_password -h ${aws_db_instance.catalogue.address} \
  -f -D socksdb < dump.sql
EOF

  tags = {
    "Name" = "RDS Importer - sock-shop"
  }

  # Temporary ignore user_data because it's different from the original cloudformation template
  lifecycle {
    ignore_changes = [user_data]
  }
}
