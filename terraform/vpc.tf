resource "aws_vpc" "sock_shop" {
  cidr_block = "172.31.0.0/16"
  tags = {
    "Name" = "sock shop"
  }
}

resource "aws_vpc_dhcp_options" "local" {
  domain_name         = "local"
  domain_name_servers = ["AmazonProvidedDNS"]
}

resource "aws_vpc_dhcp_options_association" "local" {
  vpc_id          = aws_vpc.sock_shop.id
  dhcp_options_id = aws_vpc_dhcp_options.local.id
}

resource "aws_subnet" "public_subnet_1" {
  cidr_block = "172.31.0.0/24"
  vpc_id     = aws_vpc.sock_shop.id
  tags = {
    "Name" = "public subnet 1"
  }
}

resource "aws_subnet" "public_subnet_2" {
  cidr_block = "172.31.1.0/24"
  vpc_id     = aws_vpc.sock_shop.id
  tags = {
    "Name" = "public subnet 2"
  }
}

resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = aws_vpc.sock_shop.id
}

resource "aws_route_table" "route_via_igw" {
  vpc_id = aws_vpc.sock_shop.id
}

resource "aws_route" "default" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id         = aws_route_table.route_via_igw.id
  gateway_id             = aws_internet_gateway.internet_gateway.id
}

resource "aws_route_table_association" "public_subnet_1_via_igw" {
  route_table_id = aws_route_table.route_via_igw.id
  subnet_id      = aws_subnet.public_subnet_1.id
}

resource "aws_route_table_association" "public_subnet_2_via_igw" {
  route_table_id = aws_route_table.route_via_igw.id
  subnet_id      = aws_subnet.public_subnet_2.id
}
