package thegraph

import (
	"os"

	"github.com/dfuse-io/logging"
)

var traceEnabled = logging.IsTraceEnabled("thegraphhose", "github.com/streamingfast/thegraphhose")
var traceMemoryEnabled = os.Getenv("TRACE_MEMORY") == "true"
var zlog = logging.NewSimpleLogger("thegraphhose", "github.com/streamingfast/thegraphhose")
