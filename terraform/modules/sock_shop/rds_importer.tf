resource "aws_instance" "rds_importer" {
  ami                         = "ami-56d4ad31"
  instance_type               = "t2.micro"
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.ecs.id]
  subnet_id                   = aws_subnet.public_subnet_1.id

  user_data = <<EOF
#!/bin/bash -xe
yum -y install mysql
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
