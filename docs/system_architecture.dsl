@startuml
left to right direction
skinparam packageStyle rectangle

actor User

package "Gateway" {
  [nginx (gateway)]
}

package "Core Services" {
  [booking_service]
  [payment_service]
  [flight_service]
  [cron_service]
  [qworker_service]
}

package "Infrastructure" {
  [mysql]
  [redis]
  [kafka]
  [zookeeper]
}

User --> [nginx (gateway)]
[nginx (gateway)] --> [booking_service]
[nginx (gateway)] --> [payment_service]
[nginx (gateway)] --> [flight_service]

[booking_service] --> [mysql]
[booking_service] --> [redis]
[booking_service] --> [kafka]

[payment_service] --> [mysql]
[payment_service] --> [redis]
[payment_service] --> [kafka]

[flight_service] --> [mysql]
[flight_service] --> [redis]

[cron_service] --> [mysql]
[cron_service] --> [redis]

[qworker_service] --> [mysql]
[qworker_service] --> [redis]
[qworker_service] --> [kafka]

[kafka] --> [zookeeper]
@enduml