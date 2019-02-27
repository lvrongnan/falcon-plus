package funcs
import (
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/toolkits/file"
	"strings"
	"bytes"
	"os/exec"
)

func OsName() string {
	ospath := "/etc/redhat-release"
	tmp1, _:= file.ToString(ospath)
	tmp2 := strings.Fields(tmp1)
	return tmp2[0] + tmp2[2]
}
func UpTime() string {
	cmd := exec.Command("uptime")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		return ""
	}
	tmp := strings.Fields(stdout.String())
	return tmp[2]+tmp[3]

}
func ExtraMetrics() []*model.MetricValue {
	osname := GaugeValue("extra.osname", OsName())
	uptime := GaugeValue("extra.uptime", UpTime())
	return  []*model.MetricValue{osname,uptime}
}