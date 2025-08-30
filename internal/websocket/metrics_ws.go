package websocket

import (
	"encoding/json"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/metrics"
	"strconv"
	"time"
)

var metricsUpgrader = websocket.Upgrader{
	CheckOrigin:       func(r *http.Request) bool { return true }, // Fix in production
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
}

type MetricsHub struct {
	*Hub
	interval time.Duration
	stop     chan struct{}
}

func NewMetricsHub(interval time.Duration) *MetricsHub {
	return &MetricsHub{
		Hub:      NewHub(),
		interval: interval,
		stop:     make(chan struct{}),
	}
}

func collectMetrics() MetricsPayload {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	host := hostnameSafe()
	sqlDB, _ := db.DB.DB()
	var open, inUse, idle int
	var waitCount int64
	var waitDur time.Duration
	if sqlDB != nil {
		st := sqlDB.Stats()
		open = st.OpenConnections
		inUse = st.InUse
		idle = st.Idle
		waitCount = st.WaitCount
		waitDur = st.WaitDuration
	}

	var lastGCUnix int64
	if ms.LastGC > 0 {
		lastGCUnix = int64(time.Unix(0, int64(ms.LastGC)).Unix())
	}

	return MetricsPayload{
		Timestamp: time.Now(),
		Host:      host,

		NumGoroutine: runtime.NumGoroutine(),
		NumCPU:       runtime.NumCPU(),
		GoVersion:    runtime.Version(),

		AllocBytes:      ms.Alloc,
		TotalAllocBytes: ms.TotalAlloc,
		SysBytes:        ms.Sys,
		HeapAllocBytes:  ms.HeapAlloc,
		HeapSysBytes:    ms.HeapSys,
		NextGCBytes:     ms.NextGC,
		LastGCUnix:      lastGCUnix,
		GCCPUFraction:   readGCCPUFraction(),

		DBOpenConns:    open,
		DBInUse:        inUse,
		DBIdle:         idle,
		DBWaitCount:    waitCount,
		DBWaitDuration: waitDur.Milliseconds(),
	}
}

func hostnameSafe() string {
	if h, err := os.Hostname(); err == nil {
		return h
	}
	addrs, _ := net.InterfaceAddrs()
	if len(addrs) > 0 {
		return addrs[0].String()
	}
	return "unknown"
}

func readGCCPUFraction() float64 {
	sample := []metrics.Sample{{Name: "/gc/cpu:fraction"}}
	metrics.Read(sample)
	if len(sample) > 0 && sample[0].Value.Kind() == metrics.KindFloat64 {
		return sample[0].Value.Float64()
	}
	return 0
}

func (mh *MetricsHub) RunMetricsPublisher() {
	t := time.NewTicker(mh.interval)
	defer t.Stop()
	for {
		select {
		case <-mh.stop:
			return
		case <-t.C:
			payload := collectMetrics()
			b, _ := json.Marshal(payload)
			mh.Broadcast <- b
		}
	}
}

func (mh *MetricsHub) Stop() { close(mh.stop) }

func ServeMetricsWs(mh *MetricsHub, c *gin.Context) {
	if s := c.Query("interval"); s != "" {
		if d, err := time.ParseDuration(s); err == nil && d >= 200*time.Millisecond {
			mh.interval = d
		}
	}
	uidStr := c.Query("user_id")
	_, _ = strconv.Atoi(uidStr)

	conn, err := metricsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &Client{
		ID:   time.Now().UnixNano(),
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	mh.Register <- client

	go client.writePump()
	go client.readPumpNoEcho(mh.Hub)
}

func (c *Client) readPumpNoEcho(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()
	for {

		if _, _, err := c.Conn.ReadMessage(); err != nil {
			return
		}
	}
}
