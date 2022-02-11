# c-conn-watcher

An application for watch active channel connections from specifc channels in Rapidpro and give metrics from this to serve for Prometheus.

___

## environment variables


| name                   | description                                    |default|
|------------------------|------------------------------------------------|-|
| CCONN_WATCHER_PORT     |server port                                     |8080|
| CCONN_WATCHER_DB_URL   |database url                                    |postgres://temba:temba@localhost/temba?sslmode=disable&Timezone=UTC|
| CCONN_WATCHER_INTERVAL |watch interval in seconds                       |5|
| CCONN_WATCHER_CHANNELS |channels ids list separated by "," eg: (1, 2, 3)|1|

## how to run

```
$ docker-compose -f docker/docker-compose.yml up --build
```

## more info
dashboard can be imported into grafana from ./docker/grafana/dashboard.json