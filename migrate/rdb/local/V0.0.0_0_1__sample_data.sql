-- Insert sample users
INSERT INTO users (id, email, phone_country_code, phone_number, status, secret_key, created_at, updated_at)
VALUES 
(1001, 'alice@example.com', '+886', '912345678', 'ACTIVE', 'secret123', NOW(), NOW()),
(1002, 'bob@example.com', '+886', '922222222', 'ACTIVE', 'secret456', NOW(), NOW());

-- Insert sample flights
INSERT INTO flights (id, airline_code, flight_number, departure_airport, arrival_airport, departure_time, arrival_time, total_seats, overbooking_limit, sellable_seats, status, created_at, updated_at)
VALUES
(2001, 'CI', 'CI102', 'TPE', 'NRT', '2025-06-01 08:00:00', '2025-06-01 12:00:00', 180, 5, 185, 'Scheduled', NOW(), NOW());

-- Insert sample cabin classes
INSERT INTO cabin_classes (id, flight_id, class_code, price, baggage_allowance, refundable, seat_selection, max_seats, remain_seats, created_at, updated_at)
VALUES
(3001, 2001, 'Economy', 7800.00, 20, TRUE, TRUE, 160, 160, NOW(), NOW()),
(3002, 2001, 'Business', 18800.00, 35, TRUE, TRUE, 20, 20, NOW(), NOW());

-- Insert sample seats
INSERT INTO seats (id, flight_id, cabin_class_id, seat_number, status, created_at, updated_at)
VALUES
(4001, 2001, 3001, '12A', 'Available', NOW(), NOW()),
(4002, 2001, 3001, '12B', 'Available', NOW(), NOW()),
(4003, 2001, 3002, '2A', 'Available', NOW(), NOW());

-- Insert sample booking
INSERT INTO bookings (id, flight_id, user_id, cabin_class_id, seat_id, status, price, expired_at, created_at, updated_at)
VALUES
(5001, 2001, 1001, 3001, 4001, 'Confirmed', 7800.00, DATE_ADD(NOW(), INTERVAL 15 MINUTE), NOW(), NOW());

-- Insert sample payment
INSERT INTO payments (id, booking_id, user_id, payment_provider, payment_method, payment_status, amount, currency, transaction_id, payment_url, expired_at, created_at, updated_at)
VALUES
(6001, 5001, 1001, 'TapPay', 'credit_card', 'Success', 7800.00, 'TWD', 'TX123456789', 'https://pay.example.com/tx/123', DATE_ADD(NOW(), INTERVAL 15 MINUTE), NOW(), NOW());

-- Insert another booking that fills the seat
INSERT INTO bookings (id, flight_id, user_id, cabin_class_id, seat_id, status, price, expired_at, created_at, updated_at)
VALUES
(5002, 2001, 1002, 3001, 4002, 'Confirmed', 7800.00, DATE_ADD(NOW(), INTERVAL 15 MINUTE), NOW(), NOW());

-- Simulate overbooking (no seats left)
INSERT INTO bookings (id, flight_id, user_id, cabin_class_id, seat_id, status, price, expired_at, created_at, updated_at)
VALUES
(5003, 2001, 1002, 3001, NULL, 'Overbooked', 7800.00, DATE_ADD(NOW(), INTERVAL 15 MINUTE), NOW(), NOW());

-- Later someone cancels, seat is released (assume ID 4004 exists)
-- We reassign the seat to the overbooked booking to simulate waitlist processing
-- (simulate manually here, system would normally do this automatically)
UPDATE bookings
SET seat_id = 4001, status = 'Confirmed', updated_at = NOW()
WHERE id = 5003;

-- Mark seat 4001 as Booked
UPDATE seats
SET status = 'Booked', updated_at = NOW()
WHERE id = 4001;