import threading
import requests
from time import sleep


class LoadGenerator:
    __TIMEOUT_LENGTH = 2

    @staticmethod
    def generate_post_load(url, load, time_sec, dict_body=None):
        if dict_body is None:
            dict_body = {"name": "test load", "content": "test body"}
        count_requests = int(load * time_sec)
        sleep_time = int(1 / load)
        for numb_req in range(count_requests):
            try:
                resp = requests.post(url, json=dict_body, timeout=LoadGenerator.__TIMEOUT_LENGTH)
                print(f"\n[INFO] response from POST test load with status {resp.status_code}")
            except requests.exceptions.RequestException:
                print(f"\n[ERROR] bad response from POST test load")
            sleep(sleep_time)

    @staticmethod
    def generate_get_load(url, load, time_sec, dict_body=None):
        count_requests = int(load * time_sec)
        sleep_time = int(1 / load)
        for numb_req in range(count_requests):
            try:
                resp = requests.get(url, json=dict_body, timeout=LoadGenerator.__TIMEOUT_LENGTH)
                print(f"\n[INFO] response from GET test load with status {resp.status_code}")
            except requests.exceptions.RequestException:
                print(f"\n[ERROR] bad response from GET test load")
            sleep(sleep_time)

    @staticmethod
    def __get_thread_count(load):
        if load < 1:
            # if load less than 1, we need only on thread
            return 1
        return int(load)

    @staticmethod
    def __wait_all_threads(threads):
        for thread in threads:
            thread.join()

    @staticmethod
    def __run_test(test_function, url, load, time_sec, body=None):
        threads = list()
        count_threads = LoadGenerator.__get_thread_count(load)
        for thread_number in range(count_threads):
            thread = threading.Thread(target=test_function, args=(url, load, time_sec, body))
            threads.append(thread)
            thread.start()
        LoadGenerator.__wait_all_threads(threads)

    @staticmethod
    def generate_load(request_flag, url, load, time_sec, body=None):
        load_function = LoadGenerator.generate_get_load
        if request_flag == "POST":
            load_function = LoadGenerator.generate_post_load
        LoadGenerator.__run_test(load_function, url, load, time_sec, body)
