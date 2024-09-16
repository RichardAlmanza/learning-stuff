resource "aws_vpc_security_group_ingress_rule" "allows_ipv4_ssh" {
  security_group_id = aws_security_group.ssh_server.id

  cidr_ipv4   = var.external_ip_to_allows_ssh_connections
  ip_protocol = "tcp"
  from_port   = 22
  to_port     = 22
}

resource "aws_vpc_security_group_ingress_rule" "allows_internal_ipv4_ssh" {
  security_group_id = aws_security_group.ssh_server.id

  cidr_ipv4   = var.cidr_block
  ip_protocol = "tcp"
  from_port   = 22
  to_port     = 22
}

resource "aws_vpc_security_group_egress_rule" "allows_all_tcp" {
  security_group_id = aws_security_group.ssh_server.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "tcp"
  from_port   = 0
  to_port     = 65535
}

resource "aws_vpc_security_group_egress_rule" "allows_all_icmp" {
  security_group_id = aws_security_group.ssh_server.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "icmp"
  from_port   = -1
  to_port     = -1
}

resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = var.public_ssh_key
}

resource "aws_security_group" "ssh_server" {
  name   = "${var.project_name}-ssh-server-instance-sg"
  vpc_id = var.vpc_id

  tags = merge({
    Name = "${var.project_name}-ssh-server-instance-sg"
  }, var.base_tags)
}

resource "aws_instance" "ssh_server" {
  subnet_id              = var.public_subnet_id
  ami                    = "ami-066784287e358dad1"
  vpc_security_group_ids = [aws_security_group.ssh_server.id]
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.deployer.key_name

  associate_public_ip_address = true

  tags = merge({
    Name = "${var.project_name}-ssh-server-i"
  }, var.base_tags)
}
