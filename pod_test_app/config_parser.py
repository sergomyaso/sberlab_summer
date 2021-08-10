import yaml
from test_pod import DeployTester, TestEntity
from deploy_create import DeployEntity


class ConfigTestRunner:

    def __init__(self, tester: DeployTester):
        self.__test_func_dict = dict()
        self.__tester = tester
        self.__init_func_dict()

    def __init_func_dict(self):
        self.__test_func_dict["memory"] = self.__tester.run_deploy_mem_test
        self.__test_func_dict["cpu"] = self.__tester.run_deploy_cpu_test

    @staticmethod
    def __parse_config( config_path):
        with open(config_path) as f:
            config = yaml.safe_load(f)
        deploy_entity = DeployEntity(
            config["deploy"]["template"],
            config["deploy"]["name"],
            config["deploy"]["image"],
            config["deploy"]["memory-limit"],
            config["deploy"]["cpu-limit"],
            config["deploy"]["replicas"],
            config["deploy"]["namespace"],
            config["deploy"]["container-port"]
        )
        return TestEntity(
            config["test"]["name"],
            config["type"],
            config["deploy"]["cluster-configuration"],
            deploy_entity,
            config["test"]["server-url"],
            config["test"]["daemon-url"],
            (config["test"]["values"]["start"], config["test"]["values"]["end"], config["test"]["values"]["step"]),
            config["test"]["pod-prepare"],
            config["test"]["test-method"],
            config["test"]["load"],
            config["test"]["time-load"]
        )

    def run_test_from_config(self, config_path):
        test_entity = self.__parse_config(config_path)
        self.__test_func_dict[test_entity.test_type](test_entity)
