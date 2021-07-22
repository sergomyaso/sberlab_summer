from deploy_create import ResourceHandler, DeployEntity
import time


class TestEntity:
    def __init__(self, config_path, deploy_entity, test_url, mem_values=(20, 25, 1), time_life_sec=15,
                 request_flag="GET", load=1,
                 time_load_sec=3):
        self.config_path = config_path
        self.deploy_entity = deploy_entity
        self.test_url = test_url
        self.test_values = mem_values
        self.time_life_sec = time_life_sec

        self.request_flag = request_flag
        self.load = load
        self.time_load_sec = time_load_sec


class DeployTester:
    def __init__(self, load_generator):
        self.__load_generator = load_generator

    def run_deploy_mem_test(self, test_entity):
        handler = ResourceHandler(test_entity.config_path)
        for mem_step in range(test_entity.test_values[0], test_entity.test_values[1], test_entity.test_values[2]):
            test_entity.deploy_entity.mem_limit = str(mem_step) + "Mi"
            print(f"\n[INFO] init create deployment with {mem_step} mem step")
            handler.apply_deployment(test_entity.deploy_entity)
            time.sleep(test_entity.time_life_sec)
            self.__load_generator.generate_load(test_entity.request_flag, test_entity.test_url, test_entity.load,
                                                test_entity.time_load_sec)
            handler.remove_deployment(test_entity.deploy_entity)

    def run_deploy_cpu_test(self, test_entity):
        handler = ResourceHandler(test_entity.config_path)
        for cpu_step in range(test_entity.test_values[0], test_entity.test_values[1], test_entity.test_values[2]):
            test_entity.deploy_entity.cpu_limit = str(cpu_step)
            print(f"\n[INFO] init create deployment with {cpu_step} cpu step")
            handler.apply_deployment(test_entity.deploy_entity)
            self.__load_generator.generate_load(test_entity.request_flag, test_entity.test_url, test_entity.load,
                                                test_entity.time_load_sec)
            time.sleep(test_entity.time_life_sec)
            handler.remove_deployment(test_entity.deploy_entity)
