-- 1. 호선 정보
CREATE TABLE Lines (
    line_id SERIAL PRIMARY KEY,
    line_name VARCHAR(255) NOT NULL,
    line_detail VARCHAR(255)
);

-- 2. 지하철 역 정보
CREATE TABLE Stations (
    station_id SERIAL PRIMARY KEY,
    station_name VARCHAR(255) NOT NULL,
    train_count INTEGER NOT NULL,
    line_id INTEGER REFERENCES Lines(line_id),
    station_detail VARCHAR(255)
);

-- 3. 승객 정보 데이터
CREATE TABLE PassengerInfo (
    passenger_id SERIAL PRIMARY KEY,
    encrypted_card_id VARCHAR(255) NOT NULL,
    boarding_line INTEGER REFERENCES Lines(line_id),
    boarding_station INTEGER REFERENCES Stations(station_id),
    boarding_time TIMESTAMP NOT NULL,
    alighting_line INTEGER REFERENCES Lines(line_id),
    alighting_station INTEGER REFERENCES Stations(station_id),
    alighting_time TIMESTAMP
);

-- 4. 역의 시간별 승객 수
CREATE TABLE DateStationCount (
    record_id SERIAL PRIMARY KEY,
    record_day DATE NOT NULL,
    record_hour INTEGER NOT NULL,
    total_passengers INTEGER DEFAULT 0,
    station_id INTEGER REFERENCES Stations(station_id),
    line_id INTEGER REFERENCES Lines(line_id),
    holiday_check BOOLEAN NOT NULL DEFAULT FALSE
);

-- 5. 현재 호선의 승객 수(카운팅)
CREATE TABLE LinePassengerCount (
    count_id SERIAL PRIMARY KEY,
    line_id INTEGER REFERENCES Lines(line_id),
    record_date DATE NOT NULL,
    total_passengers INTEGER DEFAULT 0,
    alighted_passengers INTEGER DEFAULT 0
);