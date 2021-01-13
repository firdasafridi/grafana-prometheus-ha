# grafana-prometheus-ha
Sample docker-compose grafana, prometheus, and HA proxy.


## Magic run delete
- docker-compose up && docker-compose rm -fsv

## Inital config
- Add new source with prometheus dashboard port
- Add new panel `rate(haproxy_frontend_connections_total{frontend="http-in"}[1m])`
- Add new pannel `(rate(node_memory_MemAvailable_bytes{server="node-1"}[1m])/1000)-(rate(node_memory_MemFree_bytes{server="node-1"}[1m])/1000)`


