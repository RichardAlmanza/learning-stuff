variable "number_az" {
  type        = number
  default     = 2
  description = "The number AZs in the region where the subnets will be deployed"
}

variable "base_tags" {
  type = map(string)
  default = {
    "Enviroment" = "Temporal"
  }
  description = "Tags that will be added to all resource"
}

variable "cidr_block" {
  type    = string
  default = "10.0.0.0/16"
}

variable "project_name" {
  type    = string
  default = "epam-task-tf"
}
