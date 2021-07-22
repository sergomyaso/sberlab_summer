import requests

from deploy_create import DeployEntity, ResourceHandler
from test_pod import TestEntity, DeployTester
from load_generator import LoadGenerator

if __name__ == '__main__':
    generator = LoadGenerator()
    dp = DeployEntity("template/DeployTemplate.yaml", "java-http", "sergomyaso/java-http", "50Mi", 1)
    entity = TestEntity("configs/myaso-config.yaml", dp, "http://178.170.195.224")
    tester = DeployTester(generator)
    tester.run_deploy_mem_test(entity)
