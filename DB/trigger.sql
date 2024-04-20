-- 승차 트리거
CREATE TRIGGER trigger_update_boarding_count
AFTER INSERT ON PassengerInfo
FOR EACH ROW
EXECUTE PROCEDURE update_boarding_count();

-- 하차 트리거
CREATE TRIGGER trigger_update_alighting_count
AFTER UPDATE ON PassengerInfo
FOR EACH ROW
WHEN (OLD.alighting_time IS NULL AND NEW.alighting_time IS NOT NULL)
EXECUTE PROCEDURE update_alighting_count();