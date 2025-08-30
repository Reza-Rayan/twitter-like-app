package websocket

import "time"

type MetricsPayload struct {
	Timestamp time.Time `json:"timestamp"`
	Host      string    `json:"host"`

	// Go runtime
	NumGoroutine int    `json:"num_goroutine"`
	NumCPU       int    `json:"num_cpu"`
	GoVersion    string `json:"go_version"`

	// Memory
	AllocBytes      uint64  `json:"alloc_bytes"`
	TotalAllocBytes uint64  `json:"total_alloc_bytes"`
	SysBytes        uint64  `json:"sys_bytes"`
	HeapAllocBytes  uint64  `json:"heap_alloc_bytes"`
	HeapSysBytes    uint64  `json:"heap_sys_bytes"`
	NextGCBytes     uint64  `json:"next_gc_bytes"`
	LastGCUnix      int64   `json:"last_gc_unix"`
	GCCPUFraction   float64 `json:"gc_cpu_fraction"`

	// DB (GORM/sql.DB pool)
	DBOpenConns    int   `json:"db_open_conns"`
	DBInUse        int   `json:"db_in_use"`
	DBIdle         int   `json:"db_idle"`
	DBWaitCount    int64 `json:"db_wait_count"`
	DBWaitDuration int64 `json:"db_wait_duration_ms"`
}
