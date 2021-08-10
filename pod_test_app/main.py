from config_parser import ConfigTestRunner
from test_pod import DeployTester
from load_generator import LoadGenerator


def main():
    tester = DeployTester(LoadGenerator())
    test_runner = ConfigTestRunner(tester)
    test_runner.run_test_from_config("./java-test-config.yaml")


if __name__ == '__main__':
    main()
