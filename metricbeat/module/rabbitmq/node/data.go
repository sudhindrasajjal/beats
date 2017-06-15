package node

import (
	"encoding/json"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"github.com/elastic/beats/libbeat/logp"
)

var (
	schema = s.Schema{
		"disk": s.Object{
			"free": s.Object{
				"bytes": c.Int("disk_free"),
				"limit": s.Object{
					"bytes": c.Int("disk_free_limit"),
				},
			},
		},
		"fd": s.Object{
			"total": c.Int("fd_total"),
			"used":  c.Int("fd_used"),
		},
		"gc": s.Object{
			"reclaimed": s.Object{
				"bytes": c.Int("gc_bytes_reclaimed"),
			},
			"num": s.Object{
				"count": c.Int("gc_num"),
			},
		},
		"io": s.Object{
			"file_handle": s.Object{
				"open_attempt": s.Object{
					"avg": s.Object{
						"ms": c.Int("io_file_handle_open_attempt_avg_time"),
					},
					"count": c.Int("io_file_handle_open_attempt_count"),
				},
			},
			"read": s.Object{
				"avg": s.Object{
					"ms": c.Int("io_read_avg_time"),
				},
				"bytes": c.Int("io_read_bytes"),
				"count": c.Int("io_read_count"),
			},
			"reopen": s.Object{
				"count": c.Int("io_read_count"),
			},
			"seek": s.Object{
				"avg": s.Object{
					"ms": c.Int("io_seek_avg_time"),
				},
				"count": c.Int("io_seek_count"),
			},
			"sync": s.Object{
				"avg": s.Object{
					"ms": c.Int("io_sync_avg_time"),
				},
				"count": c.Int("io_sync_count"),
			},
			"write": s.Object{
				"avg": s.Object{
					"ms": c.Int("io_write_avg_time"),
				},
				"bytes": c.Int("io_write_bytes"),
				"count": c.Int("io_write_count"),
			},
		},
		"mem": s.Object{
			"limit": s.Object{
				"bytes": c.Int("mem_limit"),
			},
			"used": s.Object{
				"bytes": c.Int("mem_used"),
			},
		},
		"mnesia": s.Object{
			"disk": s.Object{
				"tx": s.Object{
					"count": c.Int("mnesia_disk_tx_count"),
				},
			},
			"ram": s.Object{
				"tx": s.Object{
					"count": c.Int("mnesia_ram_tx_count"),
				},
			},
		},
		"msg": s.Object{
			"store_read": s.Object{
				"count": c.Int("msg_store_read_count"),
			},
			"store_write": s.Object{
				"count": c.Int("msg_store_write_count"),
			},
		},
		"name": c.Str("name"),
		"proc": s.Object{
			"total": c.Int("proc_total"),
			"used":  c.Int("proc_used"),
		},
		"processors": c.Int("processors"),
		"queue": s.Object{
			"index": s.Object{
				"journal_write": s.Object{
					"count": c.Int("queue_index_journal_write_count"),
				},
				"read": s.Object{
					"count": c.Int("queue_index_read_count"),
				},
				"write": s.Object{
					"count": c.Int("queue_index_write_count"),
				},
			},
		},
		"run": s.Object{
			"queue": c.Int("run_queue"),
		},
		"socket": s.Object{
			"total": c.Int("sockets_total"),
			"used":  c.Int("sockets_used"),
		},
		"type":   c.Str("type"),
		"uptime": c.Int("uptime"),
	}
)

func eventsMapping(content []byte) ([]common.MapStr, error) {

	var nodes []map[string]interface{}
	err := json.Unmarshal(content, &nodes)
	if err != nil {
		logp.Err("Error: ", err)
	}

	events := []common.MapStr{}
	errors := s.NewErrors()

	for _, node := range nodes {
		event, errs := eventMapping(node)
		events = append(events, event)
		errors.AddErrors(errs)

	}

	return events, errors
}

func eventMapping(node map[string]interface{}) (common.MapStr, *s.Errors) {
	return schema.Apply(node)
}
