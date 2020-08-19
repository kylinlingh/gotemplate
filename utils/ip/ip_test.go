package ip

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsPrivateIP(t *testing.T) {
	convey.Convey("Test private ip.\n", t, func() {
		ip1 := "192.168.1.1"
		convey.So(IsPrivateIP(ip1), convey.ShouldBeTrue)
		ip2 := "192.168.1.256"
		convey.So(IsPrivateIP(ip2), convey.ShouldBeFalse)
		ip3 := ""
		convey.So(IsPrivateIP(ip3), convey.ShouldBeFalse)
	})
}
