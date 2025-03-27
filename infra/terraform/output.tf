output "master_public_ip" {
  description = "Public IP of instance master"
  value       = aws_eip.master_eip.public_ip
}

output "worker_public_ips" {
  description = "Public IP(s) of worker"
  value       = [for w in aws_instance.worker : w.public_ip]
}

output "vpc_id" {
  description = "VPC ID"
  value       = aws_vpc.main.id
}

output "nat_gateway_ip" {
  description = "IP of NAT Gateway"
  value       = aws_eip.nat_eip.public_ip
}
