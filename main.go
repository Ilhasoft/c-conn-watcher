package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var dbURL = "postgres://temba:temba@localhost/temba?sslmode=disable&Timezone=UTC"
var observeInterval = 5
var observedChannelsIds []int64
var observedChannels []Channel
var serverPort = 8080
var db *sqlx.DB
var err error

func init() {
	initVars()
}

func main() {
	db = NewDB()
	defer db.Close()

	prometheus.MustRegister(gauge)

	searchAndSetupChannels()

	go taskObserveChannelsConnections()

	setupRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil))

	select {}
}

type Channel struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

func initVars() {
	observedChannelsIds = []int64{}
	observedChannels = []Channel{}
	if interval, ok := os.LookupEnv("CCONN_WATCHER_INTERVAL"); ok {
		observeInterval, _ = strconv.Atoi(interval)
	}
	if dburl, ok := os.LookupEnv("CCONN_WATCHER_DB_URL"); ok {
		dbURL = dburl
	}
	if observedCIDS, ok := os.LookupEnv("CCONN_WATCHER_CHANNELS"); ok {
		for _, id := range strings.Split(observedCIDS, ",") {
			ID, err := strconv.Atoi(strings.TrimSpace(id))
			if err != nil {
				log.Fatal(err)
			}
			observedChannelsIds = append(observedChannelsIds, int64(ID))
		}
	} else {
		observedChannelsIds = append(observedChannelsIds, 1)
	}
	if port, ok := os.LookupEnv("CCONN_WATCHER_PORT"); ok {
		serverPort, err = strconv.Atoi(port)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func searchAndSetupChannels() {
	for _, chID := range observedChannelsIds {
		channel, err := selectChannelByID(chID)
		if err != nil {
			log.Println(fmt.Errorf("error on select channel by ID: %s", err))
		} else {
			observedChannels = append(observedChannels, channel)
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Channels Conections Watcher"))
	}))

	http.Handle("/metrics", promhttp.Handler())
}
