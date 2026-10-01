# FleetPulse

FleetPulse is a cloud-based fleet management application deployed on AWS using a 3-tier architecture. The infrastructure is provisioned using Terraform, with containerized frontend and backend services running on Amazon EC2 behind an Application Load Balancer.

## Architecture

```text
                         Internet
                            |
                            v
                     Route 53 (DNS)
                            |
                            v
                    CloudFront (CDN)
                            |
                            v
                  Application Load Balancer
                            |
              +-------------+-------------+
              |                           |
              v                           v
       Frontend Target Group       Backend Target Group
              |                           |
              +-------------+-------------+
                            |
                            v
                  EC2 Auto Scaling Group
                   Private App Subnets
                            |
              +-------------+-------------+
              |                           |
              v                           v
        FleetPulse Frontend       FleetPulse Backend
           Docker :80                Docker :8080
                                         |
                     +-------------------+------------------+
                     |                   |                  |
                     v                   v                  v
              Amazon RDS          ElastiCache          OpenSearch
              PostgreSQL             Redis              Search/Logs
```

## AWS 3-Tier Architecture

### Tier 1 - Presentation Layer

The presentation layer handles incoming web traffic.

- Route 53 - DNS management
- CloudFront - CDN and content delivery
- Application Load Balancer - HTTP traffic distribution
- Frontend container - Nginx-based frontend

### Tier 2 - Application Layer

The application layer contains the FleetPulse application services.

- Amazon EC2
- Auto Scaling Group
- Docker
- Backend API
- Frontend application
- Application Load Balancer target groups
- CloudWatch monitoring

### Tier 3 - Data Layer

The data layer contains persistent and supporting services.

- Amazon RDS PostgreSQL - relational database
- Amazon ElastiCache - caching layer
- Amazon OpenSearch - search and log analytics

## Technology Stack

### Application

- Frontend: FleetPulse Frontend
- Backend: FleetPulse Backend
- Docker
- REST API

### Cloud

- AWS EC2
- AWS VPC
- Application Load Balancer
- Auto Scaling
- Amazon RDS PostgreSQL
- Amazon ElastiCache
- Amazon OpenSearch
- Amazon CloudFront
- Amazon Route 53
- Amazon CloudWatch
- Amazon ECR
- AWS Systems Manager

### Infrastructure as Code

- Terraform

## VPC Architecture

The infrastructure uses a VPC with separate subnet groups:

```text
VPC
|
+-- Public Subnets
|   +-- Application Load Balancer
|   +-- Internet Gateway connectivity
|
+-- Private Application Subnets
|   +-- EC2 Auto Scaling Group
|   +-- FleetPulse containers
|
+-- Private Database Subnets
    +-- RDS PostgreSQL
    +-- ElastiCache
    +-- OpenSearch
```

The application servers are isolated from direct public access and receive traffic through the Application Load Balancer.

## Terraform

Terraform is used to provision and manage the AWS infrastructure.

Example resources include:

```text
VPC
Subnets
Internet Gateway
NAT Gateway
Route Tables
Security Groups
Application Load Balancer
Target Groups
EC2 Launch Template
Auto Scaling Group
RDS PostgreSQL
ElastiCache
OpenSearch
CloudWatch
```

Initialize Terraform:

```bash
terraform init
```

Format the configuration:

```bash
terraform fmt
```

Validate the configuration:

```bash
terraform validate
```

Review the infrastructure plan:

```bash
terraform plan
```

Apply the infrastructure:

```bash
terraform apply
```

Destroy the infrastructure when it is no longer required:

```bash
terraform destroy
```

## Container Images

The FleetPulse application images are stored in Amazon ECR.

Example repositories:

```text
fleetpulse-backend
fleetpulse-frontend
```

Example image versions:

```text
fleetpulse-backend:1.1
fleetpulse-frontend:1.0
```

The EC2 instances authenticate with ECR and pull the required images during instance initialization.

## EC2 Auto Scaling

FleetPulse uses an EC2 Auto Scaling Group to maintain application instances.

The configuration includes:

```text
Minimum instances: 2
Desired instances: 2
Maximum instances: 2
```

The instances use an EC2 Launch Template.

The Launch Template configures:

- Amazon Linux 2023
- Docker
- CloudWatch Agent
- ECR authentication
- Backend container
- Frontend container

## Application Ports

| Component | Port |
|---|---:|
| Frontend | 80 |
| Backend | 8080 |
| PostgreSQL | 5432 |

Backend health check:

```text
/api/health
```

Frontend health check:

```text
/
```

## Application Load Balancer

The ALB distributes traffic to the application instances.

Target groups:

```text
fleetpulse-frontend-tg
fleetpulse-backend-tg
```

Frontend:

```text
Port 80
Health Check: /
```

Backend:

```text
Port 8080
Health Check: /api/health
```

## Database

FleetPulse uses Amazon RDS PostgreSQL for persistent application data.

The database is deployed in private database subnets and is not directly exposed to the internet.

```text
Application EC2
       |
       | PostgreSQL :5432
       v
RDS PostgreSQL
```

## ElastiCache

Amazon ElastiCache can be used as the caching layer for frequently accessed application data.

Typical use cases include:

- Session caching
- Frequently accessed data
- API response caching
- Reducing database load

Architecture:

```text
FleetPulse Backend
       |
       v
ElastiCache
       |
       v
RDS PostgreSQL
```

## OpenSearch

Amazon OpenSearch can be used for centralized search, log analysis, and application observability.

Example flow:

```text
Application
     |
     v
Logs / Events
     |
     v
OpenSearch
     |
     v
Search & Visualization
```

## CloudWatch Monitoring

Amazon CloudWatch is used to monitor EC2 infrastructure.

The CloudWatch Agent collects:

- CPU utilization
- Memory utilization
- Disk utilization

Metrics are published under the custom namespace:

```text
FleetPulse
```

The CloudWatch Agent configuration is installed on EC2 instances during bootstrapping.

## Route 53

Route 53 provides DNS management for the application.

Example flow:

```text
fleetpulse.example.com
        |
        v
Route 53
        |
        v
CloudFront
        |
        v
Application Load Balancer
```

## CloudFront

CloudFront provides CDN capabilities and can distribute frontend content closer to users.

Benefits include:

- Reduced latency
- Edge caching
- HTTPS support
- Integration with Route 53
- Improved content delivery

## Security

The infrastructure follows a layered security model.

### Public

Only the required entry points are publicly accessible:

```text
Route 53
CloudFront
ALB
```

### Private

Application and database resources are placed in private subnets:

```text
EC2
RDS
ElastiCache
OpenSearch
```

Security groups restrict communication between the layers.

Example:

```text
Internet
   |
   v
ALB :80
   |
   v
EC2 :80 / :8080
   |
   v
RDS :5432
```

## Deployment Flow

The deployment process is:

```text
Developer
    |
    v
GitHub
    |
    v
Docker Build
    |
    v
Amazon ECR
    |
    v
Terraform Infrastructure
    |
    v
EC2 Auto Scaling Group
    |
    v
Docker Containers
    |
    v
Application Load Balancer
    |
    v
Users
```

## Health Checks

The ALB verifies application availability using health checks.

Frontend:

```text
GET /
```

Backend:

```text
GET /api/health
```

If an EC2 instance fails its health checks, the Auto Scaling Group can replace the unhealthy instance.

## Instance Refresh

When the EC2 Launch Template is modified, an Auto Scaling instance refresh can be used to replace existing instances with instances using the updated configuration.

Example:

```bash
aws autoscaling start-instance-refresh \
  --auto-scaling-group-name fleetpulse-asg \
  --preferences MinHealthyPercentage=50,InstanceWarmup=300 \
  --region ap-south-1
```

Check the refresh:

```bash
aws autoscaling describe-instance-refreshes \
  --auto-scaling-group-name fleetpulse-asg \
  --region ap-south-1
```

## Troubleshooting

Check Auto Scaling instances:

```bash
aws autoscaling describe-auto-scaling-groups \
  --auto-scaling-group-names fleetpulse-asg \
  --region ap-south-1
```

Check ALB target health:

```bash
aws elbv2 describe-target-health \
  --target-group-arn <TARGET_GROUP_ARN> \
  --region ap-south-1
```

Check EC2 instances through Systems Manager:

```bash
aws ssm send-command \
  --instance-ids <INSTANCE_ID> \
  --region ap-south-1 \
  --document-name "AWS-RunShellScript"
```

Check Docker:

```bash
docker ps -a
```

Check application ports:

```bash
ss -lntp | grep -E ':80|:8080'
```

Check backend health:

```bash
curl http://localhost:8080/api/health
```

Check frontend:

```bash
curl -I http://localhost:80/
```

## Project Structure

```text
fleetpulse/
|
+-- frontend/
|
+-- backend/
|
+-- terraform/
|   +-- ec2.tf
|   +-- vpc.tf
|   +-- alb.tf
|   +-- rds.tf
|   +-- security-groups.tf
|   +-- variables.tf
|   +-- outputs.tf
|   +-- providers.tf
|   +-- ...
|
+-- .gitignore
+-- README.md
```

## Important Security Note

Do not commit passwords, AWS access keys, private keys, Terraform state files, or other secrets to GitHub.

Sensitive values should be stored using services such as:

- AWS Secrets Manager
- AWS Systems Manager Parameter Store
- Environment variables
- GitHub Actions Secrets

Terraform state files should also not be committed to the repository.

## Current Infrastructure

AWS Region:

```text
ap-south-1
```

Project:

```text
FleetPulse
```

Infrastructure is managed using Terraform and the application is deployed using Docker containers on EC2 Auto Scaling instances.

## Future Improvements

Possible improvements include:

- HTTPS using ACM
- Route 53 custom domain
- CloudFront distribution
- ElastiCache Redis integration
- OpenSearch log aggregation
- AWS Secrets Manager
- HTTPS-only ALB listener
- CI/CD using GitHub Actions
- Blue/Green deployment
- Automated CloudWatch alarms
- Centralized application logging
- Container orchestration using Amazon ECS or EKS

## Author

Atharva Shelar

FleetPulse - AWS Cloud & DevOps Deployment Project
