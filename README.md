# Smart Flight Management System

## 📌 Target:

Designed and scoped a comprehensive flight management system covering ticket booking, flight scheduling, crew allocation, and passenger information. Focused on delivering a scalable MVP with microservices architecture, secure APIs, and real-time data handling.

---------------------------------------

### The names of the packages used along with the licenses for each are as follows:

1-mysql    License  ==  [License](https://github.com/go-sql-driver/mysql?tab=MPL-2.0-1-ov-file "License mysql")
Mozilla Public License Version 2.0

---------------------------------------

2-gin      License  ==  [License](https://github.com/gin-gonic/gin?tab=MIT-1-ov-file "License gin") 
Copyright (c) 2014 Manuel Martínez-Almeida

---------------------------------------

3-uuid      License  ==  [License](https://github.com/google/uuid?tab=License-1-ov-file "License uuid") 
Copyright (c) 2009,2014 Google Inc. All rights reserved.

---------------------------------------

4-Redis client for Go      License  ==  [License](https://github.com/redis/go-redis?tab=BSD-2-Clause-1-ov-file "License Redis client for Go")
Copyright (c) 2013 The github.com/redis/go-redis Authors.
All rights reserved.

---------------------------------------

5-Gomail      License  ==  [License](https://github.com/go-gomail/gomail?tab=MIT-1-ov-file "Gomail")
Copyright (c) 2014 Alexandre Cesaro

---------------------------------------

5-gRPC-Go      License  ==  [License](https://github.com/grpc/grpc-go?tab=Apache-2.0-1-ov-file "gRPC-Go")
Apache License Version 2.0, January 2004

---------------------------------------
> [!IMPORTANT]
>Built as a realistic prototype to simulate the challenges and solutions of a full-scale airline management platform.

## 🎯 Business Goals
- Provide fast and secure ticket booking for passengers
- Enable intelligent seat management and prevent overbooking
- Support flight scheduling and efficient crew allocation
- Deliver real-time notifications to passengers and staff
- Implement dynamic pricing and smart flight recommendation

------------------------------------------------------------------

## 🛠️ Scope — MVP
The MVP scope covers the following core functionalities:
- Flight search by origin, destination, date, and travel class
- Seat reservation and inventory control with overbooking prevention
- Simulated online payment and digital ticket generation
- Real-time seat availability via WebSocket or polling mechanisms
- Email and push notifications for flight status updates
- Lightweight admin panel for flight scheduling and crew assignment

> [!IMPORTANT]
>These features deliver a complete flight booking experience while maintaining scalability, responsiveness, and operational clarity.

------------------------------------------------------------------

## 🔧 Functional Requirements
- Passengers: Registration, login, flight booking, ticket and flight status viewing, reservation cancellation
- Operators: Flight management, seat configuration, crew assignment
- System: Dynamic pricing calculation , prevention of simultaneous seat booking
- Notifications: Real-time alerts for passengers and operators regarding flight status and updates

> [!IMPORTANT]
> These functional requirements ensure a seamless experience for both passengers and operators, while maintaining system integrity and responsiveness.

------------------------------------------------------------------

## 🛡️ Non-Functional Requirements
- Availability: Target SLA of 99.9% uptime
- Performance: API response time under 300ms for critical endpoints
- Scalability: Horizontally scalable microservices architecture
- Security: HTTPS/TLS encryption, secure password storage, protection against CSRF, SQL injection, and XSS
- Data Privacy: Compliance with GDPR and local data protection regulations
- Monitoring: Integrated Prometheus/Grafana for metrics and ELK Stack for centralized logging

> [!IMPORTANT]
>  These non-functional requirements ensure the system remains reliable, secure, and maintainable under real-world conditions.

------------------------------------------------------------------

## 🔐 Security & Authentication
- Authentication: JWT-based access and refresh tokens
- User Roles: Guest, Passenger, Operator, Admin
- Authorization: Role-Based Access Control (RBAC) enforced at each endpoint
- Data Protection: Encryption of sensitive data such as payment tokens and API keys

> [!IMPORTANT]
>  These mechanisms ensure secure access control, data confidentiality, and compliance with modern security standards.


---------------------------
#### If you have any feedback or suggestions, I’d be happy to hear from you via email :
mreza_4040@outlook.com


