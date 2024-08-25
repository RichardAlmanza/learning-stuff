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
}
