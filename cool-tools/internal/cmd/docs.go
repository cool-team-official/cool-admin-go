package cmd

import (
	"context"
	"net"

	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/utility/mlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Docs = gcmd.Command{
		Name:        "docs",
		Usage:       "cool-tools docs",
		Brief:       "查看帮助文档",
		Description: "查看帮助文档",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			s := g.Server("docs")
			// 获取本机未占用的端口
			port, err := getfreeport()
			if err != nil {
				mlog.Fatal(err)
				return err
			}
			// 获取本机ip
			// ip, err := getlocalip()
			// if err != nil {
			// 	mlog.Fatal(err)
			// 	return err
			// }
			s.SetServerRoot("docs")
			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.RedirectTo("/cool-admin-go/")
			})
			// 设置端口
			s.SetPort(gconv.Int(port))
			mlog.Printf("CoolAdminGo docs server is running at %s", "http://"+"127.0.0.1"+":"+gconv.String(port)+"/cool-admin-go/")
			s.Run()
			return nil
		},
	}
)

func init() {
	Main.AddCommand(&Docs)
}

// getfreeport 获取本机未占用的端口
func getfreeport() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// getlocalip 获取本机ip
// func getlocalip() (string, error) {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return "", err
// 	}
// 	for _, address := range addrs {
// 		// 检查ip地址判断是否回环地址
// 		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				return ipnet.IP.String(), nil
// 			}
// 		}
// 	}
// 	return "", nil
// }
