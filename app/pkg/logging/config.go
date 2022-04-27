package logging

import (
	"github.com/labstack/echo/v4/middleware"
)

var Config = middleware.LoggerConfig{
	Skipper: middleware.DefaultSkipper,
	Format: `{
	"time":"${time_rfc3339_nano}",
	"id":"${id}",
	"remote_ip":"${remote_ip}",
	"host":"${host}",
	"method":"${method}",
	"uri":"${uri}",
	"status":${status},
	"error":"${error}",
	"latency":${latency},
	"latency_human":"${latency_human}",
	"bytes_in":${bytes_in},
	"bytes_out":${bytes_out},
	"user_agent":"${user_agent}"
}` + "\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
}
