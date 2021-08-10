import requests

from deploy_create import ResourceHandler
import numpy as np
import time


class TestEntity:
    def __init__(self, test_name, test_type, config_path, deploy_entity, test_url, daemon_url, mem_values=(20, 25, 1),
                 time_pod_load_sec=15,
                 request_flag="GET", load=1,
                 time_load_sec=3):
        self.test_name = test_name
        self.test_type = test_type
        self.config_path = config_path
        self.deploy_entity = deploy_entity
        self.test_url = test_url
        self.daemon_url = daemon_url
        self.test_values = mem_values
        self.time_pod_load_sec = time_pod_load_sec

        self.request_flag = request_flag
        self.load = load
        self.time_load_sec = time_load_sec


class DeployTester:
    def __init__(self, load_generator):
        self.__load_generator = load_generator

    @staticmethod
    def __create_resource_dump(test_entity: TestEntity, handler):
        pod_list = handler.get_pod_information(test_entity.deploy_entity.namespace)
        url = test_entity.daemon_url + "/dump/" + test_entity.test_type
        try:
            requests.post(url, json=pod_list[0].get_dict_view(test_entity.test_name, test_entity.test_type))
            print(f"\n[ERROR] dump on node " + pod_list[0].node_ip + " was created")
        except requests.exceptions.RequestException:
            print(f"\n[ERROR] dump on node " + pod_list[0].node_ip + " not created")

    @staticmethod
    def __clean_page_cash(test_entity: TestEntity):
        url = test_entity.daemon_url + "/dump/clspg"
        print(f"\n[INFO] try to clean page cash in " + url)
        try:
            requests.post(url, json={})
            print(f"\n[INFO] page cash " + " was cleaned")
        except requests.exceptions.RequestException:
            print(f"\n[ERROR] page cash " + " wasn't cleaned")

    def run_deploy_mem_test(self, test_entity):
        handler = ResourceHandler(test_entity.config_path)
        for mem_step in range(test_entity.test_values[0], test_entity.test_values[1], test_entity.test_values[2]):
            #self.__clean_page_cash(test_entity)
            test_entity.deploy_entity.mem_limit = str(mem_step) + "Mi"
            print(f"\n[INFO] init create deployment with {mem_step} mem step")
            handler.apply_deployment(test_entity.deploy_entity)
            time.sleep(test_entity.time_pod_load_sec)
            self.__load_generator.generate_load(test_entity.request_flag, test_entity.test_url, test_entity.load,
                                                test_entity.time_load_sec)
            print("TRY TO CREATE DUMP")
            self.__create_resource_dump(test_entity, handler)
            print("AFTER TO CREATE DUMP")
            handler.remove_deployment(test_entity.deploy_entity)
            time.sleep(10)

    def run_deploy_cpu_test(self, test_entity):
        handler = ResourceHandler(test_entity.config_path)

        for cpu_step in np.arange(test_entity.test_values[0], test_entity.test_values[1], test_entity.test_values[2]):
            test_entity.deploy_entity.cpu_limit = str(cpu_step)
            print(f"\n[INFO] init create deployment with {cpu_step} cpu step")
            handler.apply_deployment(test_entity.deploy_entity)
            time.sleep(test_entity.time_pod_load_sec)
            self.__load_generator.generate_load(test_entity.request_flag, test_entity.test_url, test_entity.load,
                                                test_entity.time_load_sec)
            self.__create_resource_dump(test_entity, handler)
            handler.remove_deployment(test_entity.deploy_entity)
