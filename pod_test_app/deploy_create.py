import tempfile

import yaml
from kubernetes import client, config
import os

from kubernetes.client import ApiException


class DeployEntity:
    def __init__(self, template_path, name, image, mem_limit, cpu_limit, replicas=1, namespace="default", port=8080):
        self.template_path = template_path
        self.name = name
        self.image = image
        self.mem_limit = mem_limit
        self.cpu_limit = cpu_limit
        self.replicas = replicas
        self.namespace = namespace
        self.port = port

    def get_render_template(self):
        with open(self.template_path) as template_file:
            string_template = template_file.read()
            return eval(f'f"""{string_template}"""')


class PodEntity:
    def __init__(self, node_ip, pod_uid):
        self.node_ip = node_ip
        self.pod_uid = pod_uid

    def get_dict_view(self, test_name, test_type):
        return {
            "test_name": test_name,
            "test_type": test_type,
            "node_ip": self.node_ip,
            "pod_uid": self.pod_uid
        }


class ResourceHandler:
    __TEMP_PATH = "./temp"

    def __init__(self, config_path):
        self.__config_path = config_path

    def __config_client(self):
        config.load_kube_config(config_file=self.__config_path)

    @staticmethod
    def __create_resource_yaml(resource_entity):
        file = tempfile.TemporaryFile(suffix=".yaml", dir=ResourceHandler.__TEMP_PATH, delete=False)
        script = resource_entity.get_render_template()
        file.write(bytes(script.encode()))
        file.close()
        return file.name

    def apply_deployment(self, resource_entity):
        self.__config_client()
        resource_name = self.__create_resource_yaml(resource_entity)
        with open(os.path.join(os.path.dirname(__file__), resource_name)) as f:
            dep = yaml.safe_load(f)
            k8s_apps_v1 = client.AppsV1Api()
            try:
                k8s_apps_v1.create_namespaced_deployment(
                    body=dep, namespace=resource_entity.namespace)

                print(f"\n[INFO] deployment {resource_entity.name} created.")
            except ApiException:
                print(f"\n[ERROR] deployment {resource_entity.name} NOT created.")

        os.remove(resource_name)

    def remove_deployment(self, resource_entity):
        self.__config_client()
        k8s_apps_v1 = client.AppsV1Api()
        k8s_apps_v1.delete_namespaced_deployment(name=resource_entity.name, namespace=resource_entity.namespace)
        print(f"\n[INFO] deployment {resource_entity.name} deleted.")

    def get_pod_information(self, namespace):
        self.__config_client()
        ret = client.CoreV1Api().list_namespaced_pod(namespace=namespace)
        pod_list = list()
        for item in ret.items:
            pod_list.append(PodEntity(item.status.host_ip, item.metadata.uid.replace("-", "_")))
        return pod_list
