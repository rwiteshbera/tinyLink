# WeMakeDevs <> Napptive : CloudNative Hackathon 2023

## Project Name : <u>tinyLink</u>

## Project Description: 
Our project is a microservices-based URL shortener that is run on the Napptive Kubernetes development platform. When creating the project, we followed the OAM (Open Application Model). The Authentication Service, Email Service, and Shorten Service are the three microservices that we developed. For simple deployment and management, all of these services have been fully dockerized.

Only authenticated users can access the tiny URLs because the Authentication Service handles user authentication and authorisation. One-Time Passwords (OTPs) are sent to users by the Email Service during login, strengthening the authentication process' security. Long URLs are converted into shorter URLs by the Shorten Service, which then stores them in redis for quick access.

To ensure seamless communication between the microservices, we used Kafka as our internal messaging system. This allows for efficient communication and data transfer between the services, ensuring smooth functioning of the entire system.

Overall, our project provides a reliable and secure URL shortening service that can be easily deployed and managed using Napptive.

## Technologies Used:
1. GoLang
2. Apache Kafka
3. Redis
4. MongoDB
5. Docker

## Installation
1.  Clone the repository:
`git clone https://github.com/rwiteshbera/tinyLink`

2.  Build and start the containers:
```bash
cd tinyLink
docker-compose up
```
This will build the services andn run containers in the background.

## Challenges faced:
1.  Communication issues: We had trouble getting different parts of our project to communicate with each other. We were using a system called Apache Kafka to handle messaging between services, but something wasn't working properly. We had to spend some time trying to figure out what was causing the problem and how to fix it.
    
2.  Dockerization errors: We ran into some unexpected errors when we tried to containerize our application using Docker. We weren't sure why these errors were happening, and they seemed to occur randomly but we had to keep testing different solutions to try to get everything working.
    
3.  Deployment errors: When we tried to deploy our application using a web server called Nginx, we kept getting a 503 error. This meant that our application wasn't accessible to users, and we had to figure out what was causing the error and how to fix it. 

When it comes to time management, working in a team of two individuals might be difficult. It can be challenging to split the job and make sure that each team member is contributing equally when there are only two employees.

## Accomplishments:
1. We learned how to use Kafka in distributed systems for messaging.
2. We learned how to send emails using the Go programming language's SMTP package.
3. We learned how to use Redis as a database for faster retrieval.
4. Our services were successfully dockerized, making it simpler to deploy them.
5. We learned how to set up YAML files for the Kubernetes development platform Napptive.

## Contributors:
[Irshit Mukherjee](https://github.com/IRSHIT033)

[Rwitesh Bera](https://github.com/rwiteshbera)
