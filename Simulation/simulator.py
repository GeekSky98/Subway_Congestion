import datetime, time, json, threading, os, logging, requests, random, string
from heapq import heappush, heappop

class SimulatorSubway():
    HOST_URL = 'localhost'
    ID_LENGTH = 24
    def __init__(self):
        self.passenger_list = []
        try:
            with open(os.path.join(os.getcwd(), "line_station.json"), 'r') as f:
                self.station = json.load(f)
        except FileNotFoundError as f_e:
            raise f_e
        except Exception as e:
            raise e
        self.thread_list = []
        self.cnt = 0

    def generate_id(self):
        return ''.join(random.choices(string.ascii_lowercase + string.digits, k=self.ID_LENGTH))

    def start_simulation(self):
        boarding = threading.Thread(target=self.boarding_thread, args=)
        alighting = threading.Thread(target=self.alighting_thread, args=)
        self.thread_list.append(boarding)
        self.thread_list.append(alighting)

        self.generate_flag = True
        for thread in self.thread_list:
            thread.start()

    def boarding_thread(self):
        while self.generate_flag:
            for _ in range(random.randint(10, 300)):
                id = self.generate_id()

            least_second = 60 - int(datetime.datetime.now().strftime('%S'))
            time.sleep(least_second)

    def alighting_thread(self):
        while self.passenger_list:
            pass

    def stop_simulation(self):
        self.generate_flag = False
        for thread in self.thread_list:
            if hasattr(self, thread):
                thread.join()
        logging.info(f'The end of the subway service: {datetime.datetime.now().strftime("%Y-%m-%d %H:%M")}\n'
                     f'Today total passengers: {self.cnt}명')

        self.cnt = 0


if __name__ == '__main__':
    pass
