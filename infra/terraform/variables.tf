variable "aws_region" {
  type        = string
  default     = "ap-southeast-1"
  description = "Region of AWS"
}

variable "vpc_cidr" {
  type        = string
  description = "CIDR of VPC"
  default     = "10.0.0.0/16"
}

variable "subnet_cidr" {
  type        = string
  description = "CIDR cho Subnet"
  default     = "10.0.1.0/24"
}

variable "master_instance_type" {
  type        = string
  description = "Instance type of master"
  default     = "t3.micro"
}

variable "worker_instance_type" {
  type        = string
  description = "Instance type of worker"
  default     = "t3.medium"
}

variable "worker_count" {
  type        = number
  description = "Number of worker"
  default     = 2 # Sorry i'm a poor man 😭😭😭
}

variable "root_volume_size" {
  type        = number
  description = "Volume size of master & worker"
  default     = 20
}

variable "public_subnet_cidr" {
  type        = string
  description = "CIDR cho Subnet Public"
  default     = "10.0.1.0/24"
}

variable "private_subnet_cidr" {
  type        = string
  description = "CIDR cho Subnet Private"
  default     = "10.0.2.0/24"
}

variable "key_name" {
  type        = string
  description = "Key name"
  // TODO: Change this to your own key name
  default = "your-key-name-here"
}

