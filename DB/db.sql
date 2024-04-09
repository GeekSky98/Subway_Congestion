-- 1. 호선 정보
CREATE TABLE Lines (
    line_id SERIAL PRIMARY KEY,
    line_name VARCHAR(255) NOT NULL
);

-- 2. 지하철 역 정보
CREATE TABLE Stations (
    station_id SERIAL PRIMARY KEY,
    station_name VARCHAR(255) NOT NULL,
    train_count INTEGER NOT NULL,
    line_id INTEGER REFERENCES Lines(line_id)
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

-- 4. 지하철 역 승객 기록
CREATE TABLE StationPassengerRecords (
    record_id SERIAL PRIMARY KEY,
    record_time TIMESTAMP NOT NULL,
    action_type VARCHAR(50) CHECK (action_type IN ('Boarding', 'Alighting')),
    station_id INTEGER REFERENCES Stations(station_id),
    line_id INTEGER REFERENCES Lines(line_id)
);

-- 5. 현재 호선의 승객 수(카운팅)
CREATE TABLE LinePassengerCount (
    count_id SERIAL PRIMARY KEY,
    line_id INTEGER REFERENCES Lines(line_id),
    record_date DATE NOT NULL,
    total_passengers INTEGER DEFAULT 0,
    alighted_passengers INTEGER DEFAULT 0
);

-- 6. 현재 역의 승객 수(카운팅)
CREATE TABLE StationPassengerCount (
    count_id SERIAL PRIMARY KEY,
    station_id INTEGER REFERENCES Stations(station_id),
    record_date DATE NOT NULL,
    total_passengers INTEGER DEFAULT 0,
    alighted_passengers INTEGER DEFAULT 0
);
