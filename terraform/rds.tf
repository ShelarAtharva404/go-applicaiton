# -------------------------
# RDS Subnet Group
# -------------------------
resource "aws_db_subnet_group" "main" {
  name       = "${var.project_name}-db-subnet-group"
  subnet_ids = aws_subnet.database[*].id

  tags = {
    Name    = "${var.project_name}-db-subnet-group"
    Project = var.project_name
  }
}

# -------------------------
# RDS PostgreSQL
# -------------------------
resource "aws_db_instance" "postgres" {
  identifier = "${var.project_name}-postgres"

  engine         = "postgres"
  engine_version = "16"

  instance_class        = "db.t3.micro"
  allocated_storage     = 20
  max_allocated_storage = 20
  storage_type          = "gp3"

  db_name  = "fleetpulse"
  username = "fleetpulse_user"
  password = "FleetPulseDB123!"

  port = 5432

  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.database.id]

  publicly_accessible = false

  multi_az = false

  backup_retention_period = 0

  skip_final_snapshot = true
  deletion_protection = false

  tags = {
    Name    = "${var.project_name}-postgres"
    Project = var.project_name
  }
}
