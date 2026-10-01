# -------------------------
# EC2 Launch Template
# -------------------------
resource "aws_launch_template" "app" {
  name = "${var.project_name}-app-template"

  image_id      = "ami-0351a972f17e4d123"
  instance_type = var.app_instance_type

  vpc_security_group_ids = [
    aws_security_group.app.id
  ]

  iam_instance_profile {
    name = aws_iam_instance_profile.ec2.name
  }

  user_data = base64encode(<<-EOF
    #!/bin/bash

    set -e

    # Update packages
    dnf update -y

    # Install Docker
    dnf install -y docker

    systemctl enable docker
    systemctl start docker

    # Install CloudWatch Agent
curl -fsSL \
  https://amazoncloudwatch-agent.s3.amazonaws.com/amazon_linux/amd64/latest/amazon-cloudwatch-agent.rpm \
  -o /tmp/amazon-cloudwatch-agent.rpm

rpm -U /tmp/amazon-cloudwatch-agent.rpm

# Configure CloudWatch Agent
cat > /opt/aws/amazon-cloudwatch-agent/etc/amazon-cloudwatch-agent.json <<'CWEOF'
{
  "metrics": {
    "namespace": "FleetPulse",
    "metrics_collected": {
      "cpu": {
        "measurement": [
          "cpu_usage_idle",
          "cpu_usage_user",
          "cpu_usage_system"
        ],
        "metrics_collection_interval": 60
      },
      "mem": {
        "measurement": [
          "mem_used_percent"
        ],
        "metrics_collection_interval": 60
      },
      "disk": {
        "measurement": [
          "used_percent"
        ],
        "resources": [
          "*"
        ],
        "metrics_collection_interval": 60
      }
    }
  }
}

CWEOF

# Start CloudWatch Agent
systemctl enable amazon-cloudwatch-agent
systemctl start amazon-cloudwatch-agent

# Login to ECR
    aws ecr get-login-password --region ap-south-1 | \
      docker login --username AWS --password-stdin \
      670518341147.dkr.ecr.ap-south-1.amazonaws.com

    # Pull images
    docker pull 670518341147.dkr.ecr.ap-south-1.amazonaws.com/fleetpulse-backend:1.1
    docker pull 670518341147.dkr.ecr.ap-south-1.amazonaws.com/fleetpulse-frontend:1.0

    # Backend
    docker run -d \
      --name fleetpulse-backend \
      --restart unless-stopped \
      -p 8080:8080 \
      -e DB_HOST="fleetpulse-postgres.czygysaks7l0.ap-south-1.rds.amazonaws.com" \
      -e DB_PORT="5432" \
      -e DB_USER="fleetpulse_user" \
      -e DB_PASSWORD="FleetPulseDB123!" \
      -e DB_NAME="fleetpulse" \
      670518341147.dkr.ecr.ap-south-1.amazonaws.com/fleetpulse-backend:1.1

    # Frontend
    docker run -d \
      --name fleetpulse-frontend \
      --restart unless-stopped \
      -p 80:80 \
      670518341147.dkr.ecr.ap-south-1.amazonaws.com/fleetpulse-frontend:1.0
  EOF
  )

  tag_specifications {
    resource_type = "instance"

    tags = {
      Name    = "${var.project_name}-app"
      Project = var.project_name
    }
  }

  tags = {
    Name    = "${var.project_name}-launch-template"
    Project = var.project_name
  }
}
