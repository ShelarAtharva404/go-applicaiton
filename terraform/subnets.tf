# Public subnets
resource "aws_subnet" "public" {
  count = 2

  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.${count.index + 1}.0/24"
  availability_zone       = var.availability_zones[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name    = "${var.project_name}-public-${count.index + 1}"
    Tier    = "public"
    Project = var.project_name
  }
}

# Private application subnets
resource "aws_subnet" "app" {
  count = 2

  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.${count.index + 11}.0/24"
  availability_zone = var.availability_zones[count.index]

  tags = {
    Name    = "${var.project_name}-app-${count.index + 1}"
    Tier    = "application"
    Project = var.project_name
  }
}

# Private database subnets
resource "aws_subnet" "database" {
  count = 2

  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.${count.index + 21}.0/24"
  availability_zone = var.availability_zones[count.index]

  tags = {
    Name    = "${var.project_name}-database-${count.index + 1}"
    Tier    = "database"
    Project = var.project_name
  }
}
