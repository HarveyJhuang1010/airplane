-- user Table
CREATE TABLE user
(
    id                 BIGINT PRIMARY KEY,
    email              VARCHAR(255),
    phone_country_code VARCHAR(10),
    phone_number       VARCHAR(20),
    status             VARCHAR(20),
    secret_key         VARCHAR(255),
    created_at         DATETIME,
    updated_at         DATETIME,
    deleted_at         DATETIME DEFAULT NULL,
    UNIQUE KEY idx_email_phone (email, phone_country_code, phone_number)
);

-- flight Table
CREATE TABLE flight
(
    id                BIGINT PRIMARY KEY,
    airline_code      VARCHAR(10) NOT NULL,
    flight_number     VARCHAR(20) NOT NULL,
    departure_airport VARCHAR(10) NOT NULL,
    arrival_airport   VARCHAR(10) NOT NULL,
    departure_time    DATETIME    NOT NULL,
    arrival_time      DATETIME    NOT NULL,
    total_seats       INT         NOT NULL,
    overbooking_limit INT      DEFAULT 5,
    sellable_seats    INT         NOT NULL,
    status            VARCHAR(20) NOT NULL,
    created_at        DATETIME,
    updated_at        DATETIME,
    deleted_at        DATETIME DEFAULT NULL,
    INDEX (departure_airport),
    INDEX (arrival_airport),
    INDEX (departure_time)
);

-- Cabin Classes Table
CREATE TABLE cabin_class
(
    id                BIGINT PRIMARY KEY,
    flight_id         BIGINT         NOT NULL,
    class_code        VARCHAR(20)    NOT NULL,
    price             DECIMAL(10, 2) NOT NULL,
    baggage_allowance INT            NOT NULL DEFAULT 20,
    refundable        BOOLEAN                 DEFAULT FALSE,
    seat_selection    BOOLEAN                 DEFAULT FALSE,
    max_seats         INT            NOT NULL,
    remain_seats      INT            NOT NULL,
    created_at        DATETIME,
    updated_at        DATETIME,
    deleted_at        DATETIME                DEFAULT NULL,
    INDEX (flight_id),
    FOREIGN KEY (flight_id) REFERENCES flight (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

-- seat Table
CREATE TABLE seat
(
    id             BIGINT PRIMARY KEY,
    flight_id      BIGINT      NOT NULL,
    cabin_class_id BIGINT,
    seat_number    VARCHAR(5)  NOT NULL,
    status         VARCHAR(20) NOT NULL,
    created_at     DATETIME,
    updated_at     DATETIME,
    deleted_at     DATETIME DEFAULT NULL,
    INDEX (flight_id),
    INDEX (cabin_class_id),
    FOREIGN KEY (cabin_class_id) REFERENCES cabin_class (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    FOREIGN KEY (flight_id) REFERENCES flight (id)
);

-- booking Table
CREATE TABLE booking
(
    id             BIGINT PRIMARY KEY,
    flight_id      BIGINT         NOT NULL,
    user_id        BIGINT         NOT NULL,
    cabin_class_id BIGINT         NOT NULL,
    seat_id        BIGINT,
    status         VARCHAR(20)    NOT NULL,
    price          DECIMAL(10, 2) NOT NULL,
    expired_at     DATETIME       NOT NULL,
    created_at     DATETIME,
    updated_at     DATETIME,
    deleted_at     DATETIME DEFAULT NULL,
    INDEX (flight_id),
    INDEX (user_id),
    INDEX (seat_id),
    FOREIGN KEY (flight_id) REFERENCES flight (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    FOREIGN KEY (user_id) REFERENCES user (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    FOREIGN KEY (cabin_class_id) REFERENCES cabin_class (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    FOREIGN KEY (seat_id) REFERENCES seat (id) ON UPDATE SET NULL ON DELETE SET NULL
);

-- payment Table
CREATE TABLE payment
(
    id               BIGINT PRIMARY KEY,
    booking_id       BIGINT         NOT NULL UNIQUE,
    user_id          BIGINT         NOT NULL,
    payment_provider VARCHAR(50)    NOT NULL,
    payment_method   VARCHAR(50)    NOT NULL,
    payment_status   VARCHAR(20)    NOT NULL,
    amount           DECIMAL(10, 2) NOT NULL,
    currency         VARCHAR(10) DEFAULT 'TWD',
    transaction_id   VARCHAR(100),
    payment_url      TEXT,
    expired_at       DATETIME       NOT NULL,
    paid_at          DATETIME,
    created_at       DATETIME,
    updated_at       DATETIME,
    deleted_at       DATETIME    DEFAULT NULL,
    INDEX (user_id),
    FOREIGN KEY (booking_id) REFERENCES booking (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);


-- Extra payment Table
CREATE TABLE extra_payment
(
    id               BIGINT PRIMARY KEY,
    booking_id       BIGINT         NOT NULL UNIQUE,
    user_id          BIGINT         NOT NULL,
    payment_provider VARCHAR(50)    NOT NULL,
    payment_method   VARCHAR(50)    NOT NULL,
    payment_status   VARCHAR(20)    NOT NULL,
    amount           DECIMAL(10, 2) NOT NULL,
    currency         VARCHAR(10) DEFAULT 'TWD',
    transaction_id   VARCHAR(100),
    payment_url      TEXT,
    expired_at       DATETIME       NOT NULL,
    paid_at          DATETIME,
    created_at       DATETIME,
    updated_at       DATETIME,
    deleted_at       DATETIME    DEFAULT NULL,
    INDEX (user_id),
    FOREIGN KEY (booking_id) REFERENCES booking (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);
