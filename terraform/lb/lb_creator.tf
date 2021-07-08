resource "sbercloud_lb_loadbalancer" "lb_1" {
  vip_subnet_id = "52ef7848-e8b4-49ca-b36e-3d9f5a79e4b4"
  name = "myaso_lb"
}

resource "sbercloud_networking_eip_associate" "eip_1" {
  public_ip = "46.243.201.73"
  port_id   = sbercloud_lb_loadbalancer.lb_1.vip_port_id
}