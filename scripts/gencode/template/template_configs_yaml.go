package template

var ConfigsYamlTemplate = ParseTemplate(`
name: lark_{{.PackageName}}_server
server_id: 1
log: "./configs/logger.yaml"
grpc_server:
  name: lark_{{.PackageName}}_server
  server_id: 1
  port: 8000
  max_connection_idle: 0
  max_connection_age: 0
  max_connection_age_grace: 0
  time: 7200000
  timeout: 20000
  connection_limit: 2000
  streams_limit: 2000
  max_recv_msg_size: 4096
  credential:
    cert_file: ./configs/tls/grpc/server.pem
    key_file: ./configs/tls/grpc/server.key
    enabled: true
  jaeger:
    host_port: "lark-jaeger:6831"
    sampler_type: "const"
    param: 1
    log_spans: true
    buffer_flush_interval: 1
    max_packet_size: 0
    enabled: false
etcd:
  endpoints: ["lark-etcd-01:12379","lark-etcd-02:12381","lark-etcd-03:12383"]
  username:
  password:
  schema: lark
  read_timeout: 5000
  write_timeout: 5000
  dial_timeout: 5000
mysql:
  max_open_conns: 20
  max_idle_conns: 10
  max_lifetime: 120000
  max_idle_time: 7200000
  charset: utf8
  address: "lark-mysql-user:13306"
  username: root
  password: lark2022
`)
