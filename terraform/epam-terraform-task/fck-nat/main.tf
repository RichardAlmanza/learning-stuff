data "aws_ami" "fck_nat" {
  filter {
    name   = "name"
    values = ["fck-nat-al2023-*"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }

  owners      = ["568608671756"]
  most_recent = "true"
}

resource "aws_vpc_security_group_ingress_rule" "allows_ipv4_http" {
  security_group_id = aws_security_group.nat_instance.id

  cidr_ipv4   = var.cidr_block
  ip_protocol = "tcp"
  from_port   = 80
  to_port     = 80
}

resource "aws_vpc_security_group_ingress_rule" "allows_ipv4_https" {
  security_group_id = aws_security_group.nat_instance.id

  cidr_ipv4   = var.cidr_block
  ip_protocol = "tcp"
  from_port   = 443
  to_port     = 443
}

resource "aws_vpc_security_group_egress_rule" "allows_ipv4_tcp" {
  security_group_id = aws_security_group.nat_instance.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "tcp"
  from_port   = 0
  to_port     = 65535
}

resource "aws_security_group" "nat_instance" {
  name   = "${var.project_name}-nat-instance-sg"
  vpc_id = var.vpc_id

  tags = merge({
    Name = "${var.project_name}-nat-instance-sg"
  }, var.base_tags)
}

resource "aws_route" "nat_instance" {
  route_table_id = var.private_route_table

  destination_cidr_block = "0.0.0.0/0"
  network_interface_id   = aws_instance.fck-nat.primary_network_interface_id
}

resource "aws_instance" "fck-nat" {
  ami                    = data.aws_ami.fck_nat.id
  instance_type          = "t2.micro"
  subnet_id              = var.public_subnet_id
  vpc_security_group_ids = [aws_security_group.nat_instance.id]

  associate_public_ip_address = true
  source_dest_check           = false

  tags = merge({
    Name = "${var.project_name}-nat-instance-i"
  }, var.base_tags)
}
