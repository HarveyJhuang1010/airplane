Table users {
  id BIGINT [pk]
  email VARCHAR [unique, note: "聯合唯一"]
  phone_country_code VARCHAR [note: "聯合唯一"]
  phone_number VARCHAR [note: "聯合唯一"]
  status VARCHAR
  secret_key VARCHAR
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
  indexes {
    (email, phone_country_code, phone_number) [name: "idx_email_phone"]
  }
}

Table flights {
  id BIGINT [pk]
  airline_code VARCHAR(10)
  flight_number VARCHAR(20)
  departure_airport VARCHAR(10)
  arrival_airport VARCHAR(10)
  departure_time DATETIME
  arrival_time DATETIME
  total_seats INT
  overbooking_limit INT
  sellable_seats INT
  status VARCHAR(20)
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
}

Table cabin_classes {
  id BIGINT [pk]
  flight_id BIGINT
  class_code VARCHAR(20)
  price DECIMAL(10,2)
  baggage_allowance INT
  refundable BOOLEAN
  seat_selection BOOLEAN
  max_seats INT
  remain_seats INT
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
}

Table seats {
  id BIGINT [pk]
  flight_id BIGINT
  cabin_class_id BIGINT
  seat_number VARCHAR(5)
  status VARCHAR(20)
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
}

Table bookings {
  id BIGINT [pk]
  flight_id BIGINT
  user_id BIGINT
  cabin_class_id BIGINT
  seat_id BIGINT
  status VARCHAR(20)
  price DECIMAL(10,2)
  expired_at DATETIME
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
}

Table payments {
  id BIGINT [pk]
  booking_id BIGINT [unique]
  user_id BIGINT
  provider VARCHAR(50)
  method VARCHAR(50)
  status VARCHAR(20)
  amount DECIMAL(10,2)
  currency VARCHAR(10)
  transaction_id VARCHAR(100)
  payment_url TEXT
  expired_at DATETIME
  paid_at DATETIME
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
}

Table extra_payments {
  id BIGINT [pk]
  booking_id BIGINT [unique]
  user_id BIGINT
  provider VARCHAR(50)
  method VARCHAR(50)
  status VARCHAR(20)
  amount DECIMAL(10,2)
  currency VARCHAR(10)
  transaction_id VARCHAR(100)
  payment_url TEXT
  expired_at DATETIME
  paid_at DATETIME
  created_at DATETIME
  updated_at DATETIME
  deleted_at DATETIME
}

Ref: cabin_classes.flight_id > flights.id
Ref: seats.flight_id > flights.id
Ref: seats.cabin_class_id > cabin_classes.id
Ref: bookings.flight_id > flights.id
Ref: bookings.user_id > users.id
Ref: bookings.cabin_class_id > cabin_classes.id
Ref: bookings.seat_id > seats.id
Ref: payments.booking_id > bookings.id
Ref: payments.user_id > users.id
Ref: extra_payments.booking_id > bookings.id
Ref: extra_payments.user_id > users.id
