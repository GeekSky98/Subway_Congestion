-- 승차 이벤트 트리거 함수
CREATE OR REPLACE FUNCTION update_boarding_count()
RETURNS TRIGGER AS $$
BEGIN
    -- DateStationCount
    IF NOT EXISTS (SELECT 1 FROM DateStationCount WHERE station_id = NEW.boarding_station AND record_day = NEW.boarding_time::date AND record_hour = EXTRACT(HOUR FROM NEW.boarding_time)) THEN
        INSERT INTO DateStationCount (record_day, record_hour, total_passengers, station_id, line_id)
        VALUES (NEW.boarding_time::date, EXTRACT(HOUR FROM NEW.boarding_time), 1, NEW.boarding_station, NEW.boarding_line);
    ELSE
        UPDATE DateStationCount
        SET total_passengers = total_passengers + 1
        WHERE station_id = NEW.boarding_station AND
              record_day = NEW.boarding_time::date AND
              record_hour = EXTRACT(HOUR FROM NEW.boarding_time);
    END IF;

    -- LinePassengerCount
    IF NOT EXISTS (SELECT 1 FROM LinePassengerCount WHERE line_id = NEW.boarding_line AND record_date = NEW.boarding_time::date) THEN
        INSERT INTO LinePassengerCount (line_id, record_date, total_passengers)
        VALUES (NEW.boarding_line, NEW.boarding_time::date, 1);
    ELSE
        UPDATE LinePassengerCount
        SET total_passengers = total_passengers + 1
        WHERE line_id = NEW.boarding_line AND record_date = NEW.boarding_time::date;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 하차 이벤트 트리거 함수
CREATE OR REPLACE FUNCTION update_alighting_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE LinePassengerCount
    SET alighted_passengers = alighted_passengers + 1
    WHERE line_id = NEW.alighting_line AND
          record_date = NEW.alighting_time::date;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;