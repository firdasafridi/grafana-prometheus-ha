# grafana-prometheus-ha
Sample docker-compose grafana, prometheus, and HA proxy.


## Magic run delete
- docker-compose up && docker-compose rm -fsv

## Inital config
- Add new source with prometheus dashboard port
- Add new panel `rate(haproxy_frontend_connections_total{frontend="http-in"}[1m])`