#############
# Master (Public Subnet)
#############
resource "aws_instance" "master" {
  ami                    = data.aws_ami.ubuntu_2204.id
  instance_type          = var.master_instance_type
  subnet_id              = aws_subnet.public_subnet.id
  vpc_security_group_ids = [aws_security_group.allow_ssh.id]

  # Root disk 20GB
  root_block_device {
    volume_size = var.root_volume_size
    volume_type = "gp2"
  }

  # SSH key (if exists)
  key_name = var.key_name != "" ? var.key_name : null

  tags = {
    Name = "master"
  }
}

# Elastic IP attached to master (public)
resource "aws_eip" "master_eip" {
  instance = aws_instance.master.id

  depends_on = [aws_internet_gateway.main_igw]

  tags = {
    Name = "master-eip"
  }
}

#############
# Workers (Private Subnet)
#############
resource "aws_instance" "worker" {
  count                  = var.worker_count
  ami                    = data.aws_ami.ubuntu_2204.id
  instance_type          = var.worker_instance_type
  subnet_id              = aws_subnet.private_subnet.id
  vpc_security_group_ids = [aws_security_group.allow_ssh.id]

  # Root disk 20GB
  root_block_device {
    volume_size = var.root_volume_size
    volume_type = "gp2"
  }

  key_name = var.key_name != "" ? var.key_name : null

  tags = {
    Name = "worker-${count.index}"
  }
}
