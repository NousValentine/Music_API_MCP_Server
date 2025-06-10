provider "aws" {
  region = "us-east-1" # Change to your preferred region
}

resource "aws_vpc" "lambda_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "lambda_subnet" {
  vpc_id                  = aws_vpc.lambda_vpc.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "us-east-1a"
  map_public_ip_on_launch = true
}

resource "aws_internet_gateway" "lambda_igw" {
  vpc_id = aws_vpc.lambda_vpc.id
}

resource "aws_route_table" "lambda_rt" {
  vpc_id = aws_vpc.lambda_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.lambda_igw.id
  }
}

resource "aws_route_table_association" "lambda_rta" {
  subnet_id      = aws_subnet.lambda_subnet.id
  route_table_id = aws_route_table.lambda_rt.id
}

resource "aws_security_group" "lambda_sg" {
  name        = "lambda-sg"
  description = "Allow outbound traffic"
  vpc_id      = aws_vpc.lambda_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Optional: Allow incoming traffic for testing
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
