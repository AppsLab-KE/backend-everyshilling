provider "aws"{
  region= "us-east-1"
  version= "~>4.61.0"
}

resource "aws_vpc" "my_vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "twalavpc"
  }
}

resource "aws_subnet" "my_subnet" {
  vpc_id            = aws_vpc.my_vpc.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = "us-east-1"

  tags = {
    Name = "twalasubnet"
  }
}

resource "aws_network_interface" "example" {
  subnet_id   = aws_subnet.my_subnet.id
  private_ips = ["172.16.10.100"]

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_instance" "mockServer" {
  ami= "ami-004811053d831c2c2"
  instance_type = "t3.medium"
  key_name="popo"

  network_interface {
    network_interface_id = aws_network_interface.example.id
    device_index         = 0
  }

  credit_specification {
    cpu_credits = "unlimited"
  }

  security_groups = [aws_security_group.example.name]

}

resource "aws_security_group" "example"{
  name= "Mockserver security-group"
  description="testing to see if my mockserver works"

  ingress {
    description = "Allow SSH"
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks= ["0.0.0.0/0"]

    description = "Allow Mockserver API"
    from_port = 4010
    to_port =4010
    protocol = "tcp"
    cidr_blocks= ["0.0.0.0/0"]
  }
}

output "instance_public_ip" {
  value = aws_instance.mockServer.public_ip
}
