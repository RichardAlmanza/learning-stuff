terraform {
  required_version = "~> 1.9"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.56"
    }
  }
}

provider "aws" {
  # Uses credentials and configuration in ~/.aws/
  profile = "default"
}

module "vpc" {
  source = "./vpc"

  project_name = var.project_name
  cidr_block   = var.cidr_block
  number_az    = var.number_az
  base_tags    = var.base_tags
}

resource "aws_vpc_security_group_ingress_rule" "allows_ipv4_http" {
  security_group_id = aws_security_group.backend_server.id

  cidr_ipv4   = "0.0.0.0/0" #module.vpc.cidr_block
  ip_protocol = "tcp"
  from_port   = 80
  to_port     = 80
}

resource "aws_vpc_security_group_egress_rule" "allows_all_tcp" {
  security_group_id = aws_instance.backend_server.id

  cidr_ipv4 = "0.0.0.0/0"
  ip_protocol = "tcp"
  from_port = 0
  to_port = 65535
}

resource "aws_security_group" "backend_server" {
  name   = "${var.project_name}-backend-server-sg"
  vpc_id = module.vpc.vpc_id

  tags = merge({
    Name = "${var.project_name}-backend-server-sg"
  }, var.base_tags)
}

resource "aws_instance" "backend_server" {
  subnet_id                   = module.vpc.public_subnets_ids[0]
  ami                         = "ami-066784287e358dad1"
  instance_type               = "t2.micro"
  security_groups             = [aws_security_group.backend_server.id]
  associate_public_ip_address = true #remove after testing!!!
  user_data                   = <<-END
          #! /bin/bash
          dnf -y update
          dnf -y install nginx

          systemctl start nginx.service
          systemctl enable nginx.service

          TOKEN=`curl --request PUT "http://169.254.169.254/latest/api/token" --header "X-aws-ec2-metadata-token-ttl-seconds: 21600"`
          INSTANCE_ID=`curl --header "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/instance-id`
          LOCAL_IP=`curl --header "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/local-ipv4`

          cat << EOF > /usr/share/nginx/html/index.html
          <!DOCTYPE html>
          <html>
          <head>
          <title>Welcome to EPAM TERRAFORM TASK</title>
          <style>
          html { color-scheme: light dark; }
          body { width: 45em; margin: 0 auto;
          font-family: Tahoma, Verdana, Arial, sans-serif; }
          </style>
          </head>
          <body>
          <h1>This instance is "$${INSTANCE_ID}"</h1>
          <p>The local IPV4 is "$${LOCAL_IP}"</p>
          </body>
          </html>
          EOF

          systemctl restart nginx.service
          END

  tags = merge({
    Name = "${var.project_name}-backend-server-i"
  }, var.base_tags)
}