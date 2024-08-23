terraform {
  required_version = "~> 1.9"

  backend "s3" {
    key    = "terraform/states/0-basics.tfstate"
    bucket = "private.richardalmanza"
    region = "us-east-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.55"
    }
  }
}

# Configure AWS Provider
provider "aws" {
  region = "us-east-1"
}
