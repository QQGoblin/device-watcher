package logs

import (
	"flag"
	"github.com/spf13/pflag"
	. "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"log"
	"time"
)

var logFlushFreq = pflag.Duration("log-flush-frequency", 5*time.Second, "Maximum number of seconds between log flushes")

// TODO(thockin): This is temporary until we agree on log dirs and put those into each cmd.
func init() {
	klog.InitFlags(flag.CommandLine)
	flag.Set("logtostderr", "true")
}

// KlogWriter serves as a bridge between the standard log package and the glog package.
type KlogWriter struct{}

// Write implements the io.Writer interface.
func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.InfoDepth(1, string(data))
	return len(data), nil
}

// InitLogs initializes logs the way we want for kubernetes.
func InitLogs() {
	log.SetOutput(KlogWriter{})
	log.SetFlags(0)
	// The default glog flush interval is 5 seconds.
	go Until(klog.Flush, *logFlushFreq, NeverStop)
}

// FlushLogs flushes logs immediately.
func FlushLogs() {
	klog.Flush()
}

// NewLogger creates a new log.Logger which sends logs to klog.Info.
func NewLogger(prefix string) *log.Logger {
	return log.New(KlogWriter{}, prefix, 0)
}
