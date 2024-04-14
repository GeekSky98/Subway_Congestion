import datetime, time, json, threading, os, logging, requests, random, string, psycopg2
from heapq import heappush, heappop

class SimulatorSubway():
    HOST_URL = os.environ.get("POSTGRES_HOST", "localhost")
    DATABASE = os.environ.get("POSTGRES_DB", "subway")
    USER = os.environ.get("POSTGRES_USER", "geeksky")
    PASSWD = os.environ.get("POSTGRES_PASSWORD", "geeksky")
    ID_LENGTH = 24
    THREAD_NUM = 3
    LAST_TIME = 23 * 60 + 30
    def __init__(self):
        self.passenger_list = []
        try:
            with open(os.path.join(os.getcwd(), "Simulation", "line_station.json"), 'r') as f:
                self.station = json.load(f)
        except FileNotFoundError as f_e:
            raise f_e
        except Exception as e:
            raise e

        self.line_id = {}
        self.station_id = {}

        self.connection = self.connect_db()
        self.cursor = self.connection.cursor()

        self.line_num = len(self.station)
        self.station_num = [len(self.station[line]) for line in self.station]
        self.today = datetime.date.today()

        self.insert_init()

        self.thread_list = []
        self.cnt = 0

    def connect_db(self):
        try:
            connection = psycopg2.connect(
                host = self.HOST_URL,
                database = self.DATABASE,
                user = self.USER,
                password = self.PASSWD
            )
            return connection
        except Exception as e:
            logging.error(e)
            raise e

    def insert_init(self):
        try:
            with self.connection.cursor() as cur:
                for line in self.station:
                    # line 정보 없을 시 생성
                    cur.execute("SELECT line_id FROM Lines WHERE line_name = %s",
                                (line,))
                    if cur.fetchone() is None:
                        cur.execute("INSERT INTO Lines (line_name) VALUES (%s) RETURNING line_id",
                                    (line,))
                        self.connection.commit()

                    # line_id 캐싱
                    cur.execute("SELECT line_id FROM Lines WHERE line_name = %s",
                                (line,))
                    line_id = cur.fetchone()[0]
                    self.line_id[line] = line_id

                    # 오늘 날짜의 호선 카운팅 레코드 생성
                    cur.execute("SELECT count_id FROM LinePassengerCount WHERE line_id = %s AND record_date = %s",
                                (line_id, self.today))
                    if cur.fetchone() is None:
                        cur.execute(
                            "INSERT INTO LinePassengerCount (line_id, record_date, total_passengers, alighted_passengers) VALUES (%s, %s, 0, 0)",
                            (line_id, self.today))
                        self.connection.commit()

                    for station in self.station[line]:
                        # Statin 정보 없을 시 생성
                        cur.execute("SELECT station_id FROM Stations WHERE station_name = %s AND line_id = %s",
                                    (station, line_id))
                        if cur.fetchone() is None:
                            cur.execute("INSERT INTO Stations (station_name, train_count, line_id) VALUES (%s, 0, %s) RETURNING station_id",
                                        (station, line_id))
                            self.connection.commit()

                        # station_id 캐싱
                        cur.execute("SELECT station_id FROM Stations WHERE station_name = %s",
                                    (station,))
                        station_id = cur.fetchone()[0]
                        self.station_id[station] = station_id

                        # 오늘 날짜의 역 카운팅 레코드 생성
                        cur.execute("SELECT count_id FROM StationPassengerCount WHERE station_id = %s AND record_date = %s",
                                    (station_id, self.today))
                        if cur.fetchone() is None:
                            cur.execute(
                                "INSERT INTO StationPassengerCount (station_id, record_date, total_passengers, alighted_passengers) VALUES (%s, %s, 0, 0)",
                                (station_id, self.today))
                            self.connection.commit()
        except Exception as e:
            logging.error(e)
            self.connection.rollback()

    def insert_passenger_info(self, encrypted_card_id, line_id, station_id, boarding_time):
        try:
            self.cursor.execute(
                """
                INSERT INTO PassengerInfo (encrypted_card_id, boarding_line, boarding_station, boarding_time)
                VALUES (%s, %s, %s, %s)
                """,
                (encrypted_card_id, line_id, station_id, boarding_time)
            )
            self.connection.commit()
        except Exception as e:
            logging.error(e)
            self.connection.rollback()

    def insert_station_passenger_record(self, record_time, action_type, station_id, line_id):
        try:
            self.cursor.execute(
                """
                INSERT INTO StationPassengerRecords (record_time, action_type, station_id, line_id)
                VALUES (%s, %s, %s, %s)
                """,
                (record_time, action_type, station_id, line_id)
            )
            self.connection.commit()
        except Exception as e:
            logging.error(e)
            self.connection.rollback()

    def update_passenger_alight_info(self, line_id, station_id, al_time, card, bd_time):
        try:
            self.cursor.execute(
                """
                UPDATE PassengerInfo SET alighting_line = %s, alighting_station = %s, alighting_time = %s
                WHERE encrypted_card_id = %s AND boarding_time = %s
                """,
                (line_id, station_id, al_time, card, bd_time)
            )
            self.connection.commit()
        except Exception as e:
            logging.error(e)
            self.connection.rollback()

    def update_line_station_count(self, line_id, station_id, total_value, alighted_value):
        try:
            # 호선 승객 수 업데이트
            self.cursor.execute(
                """
                UPDATE LinePassengerCount
                SET total_passengers = total_passengers + %s, alighted_passengers = alighted_passengers + %s
                WHERE line_id = %s AND record_date = %s
                """,
                (total_value, alighted_value, line_id, self.today)
            )
            self.connection.commit()

            self.cursor.execute(
                """
                UPDATE StationPassengerCount
                SET total_passengers = total_passengers + %s, alighted_passengers = alighted_passengers + %s
                WHERE station_id = %s AND record_date = %s
                """,
                (total_value, alighted_value, station_id, self.today)
            )
            self.connection.commit()

        except Exception as e:
            logging.error(e)
            self.connection.rollback()

    def generate_id(self):
        return ''.join(random.choices(string.ascii_lowercase + string.digits, k=self.ID_LENGTH))

    def calc_least_time(self, st_line, st_station, dt_line, dt_station):
        if st_line == dt_line:
            diff = abs(st_station - dt_station)
            return diff * 2
        else:
            pass  # Line 늘릴 때 추가

    def to_minute(self, t):
        h, m = int(t.strftime("%H")), int(t.strftime("%M"))
        return h * 60 + m

    def start_simulation(self):
        for _ in range(self.THREAD_NUM):
            boarding = threading.Thread(target=self.boarding_thread)
            self.thread_list.append(boarding)
        alighting = threading.Thread(target=self.alighting_thread)
        self.thread_list.append(alighting)

        self.generate_flag = True
        for thread in self.thread_list:
            thread.start()

    def boarding_thread(self):
        while self.generate_flag:
            line_index = random.randint(0, self.line_num-1)
            station_index = random.randint(0, self.station_num[line_index]-1)

            dt_line_index = random.randint(0, self.line_num-1)
            while True:
                dt_station_index = random.randint(0, self.station_num[dt_line_index]-1)
                if (dt_line_index, dt_station_index) != (line_index, station_index):
                    break

            card_id = self.generate_id()
            date = datetime.datetime.now()
            line_list = list(self.station)
            line_id = self.line_id[line_list[line_index]]
            dt_line_id = self.line_id[line_list[dt_line_index]]
            station_id = self.station_id[self.station[line_list[line_index]][station_index]]
            dt_station_id = self.station_id[self.station[line_list[dt_line_index]][dt_station_index]]
            index_alpha = date + datetime.timedelta(minutes=self.calc_least_time(line_index, station_index,
                                                                              dt_line_index, dt_station_index))
            index = self.to_minute(index_alpha)

            self.insert_passenger_info(card_id, line_id, station_id, date)
            self.insert_station_passenger_record(date, "Boarding", station_id, line_id)
            self.update_line_station_count(line_id, station_id, 1, 0)

            queue_data = (index, [card_id, date, dt_line_id, dt_station_id])
            heappush(self.passenger_list, queue_data)
            logging.info(f"Passenger: From {line_id}-{station_id} To {dt_line_id}-{dt_station_id}")

            time.sleep(random.uniform(0, 5))
            if index >= self.LAST_TIME:
                self.generate_flag = False

    def alighting_thread(self):
        while self.passenger_list or self.generate_flag:
            if self.passenger_list and self.passenger_list[0][0] <= self.to_minute(datetime.datetime.now()):
                _, inform = heappop(self.passenger_list)
                card, bd_date, line_id, station_id = inform
                self.update_passenger_alight_info(line_id, station_id, datetime.datetime.now(), card, bd_date)
                self.insert_station_passenger_record(bd_date, "Alighting", station_id, line_id)
                self.update_line_station_count(line_id, station_id, 0, 1)
                self.cnt += 1
                continue
            time.sleep(10)
        self.stop_simulation()

    def stop_simulation(self):
        #self.generate_flag = False
        for thread in self.thread_list:
            #if hasattr(self, thread):
            thread.join()
        logging.info(f'The end of the subway service: {datetime.datetime.now().strftime("%Y-%m-%d %H:%M")}\n'
                     f'Today total passengers: {self.cnt}명')

        self.cnt = 0

if __name__ == '__main__':
    simulator = SimulatorSubway()
    simulator.start_simulation()