package log

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLog(t *testing.T) {
	convey.Convey("test uber zap log\n", t, func() {
		GlobalLogger.Debugf("This is a test.")
	})
}
