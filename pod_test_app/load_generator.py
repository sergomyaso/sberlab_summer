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
                resp = requests.get(url, json=dict_body, timeout=LoadGenerator.__TIMEOUT_LENGTH)
                print(f"\n[INFO] response from POST test load with status {resp.status_code}")
            except requests.exceptions.RequestException:
                print(f"\n[ERROR] bad response from POST test load")
            sleep(sleep_time)

    @staticmethod
    def generate_get_load(url, load, time_sec):
        count_requests = int(load * time_sec)
        sleep_time = int(1 / load)
        for numb_req in range(count_requests):
            try:
                resp = requests.get(url, timeout=LoadGenerator.__TIMEOUT_LENGTH)
                print(f"\n[INFO] response from GET test load with status {resp.status_code}")
            except requests.exceptions.RequestException:
                print(f"\n[ERROR] bad response from GET test load")
            sleep(sleep_time)

    def generate_load(self, request_flag, url, load, time_sec, body=None):
        if request_flag == "POST":
            LoadGenerator.generate_post_load(url, load, time_sec, body)
        LoadGenerator.generate_get_load(url, load, time_sec)
