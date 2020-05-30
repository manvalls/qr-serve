package main

import (
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/skip2/go-qrcode"
)

const size = 300

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.To4()[0] == 192 {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func getIcon() *walk.Icon {
	q, _ := qrcode.New("QR Serve", qrcode.Medium)
	ic, _ := walk.NewIconFromImage(q.Image(64))
	return ic
}

func showError(text string) {
	MainWindow{
		Title:  "QR Serve",
		Icon:   getIcon(),
		Size:   Size{size + 100, 200},
		Layout: VBox{},
		Children: []Widget{
			TextLabel{
				Text:          text,
				TextAlignment: AlignHCenterVCenter,
			},
		},
	}.Run()
}

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	if len(os.Args) < 2 {
		showError("File not specified. Please open your file with QR Serve.")
		return
	}

	filePath := os.Args[1]
	fileName := filepath.Base(filePath)
	ip := getIP()
	if ip == "" {
		showError("Could not find LAN address")
		return
	}

	ln, err := net.Listen("tcp", ip+":0")
	if err != nil {
		showError(err.Error())
		return
	}

	defer ln.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		defer wg.Done()

		http.ServeFile(w, r, filePath)
	})

	go http.Serve(ln, nil)

	q, _ := qrcode.New("http://"+ln.Addr().String()+"/"+url.QueryEscape(fileName), qrcode.RecoveryLevel(0))
	bm, _ := walk.NewBitmapFromImage(q.Image(size))

	MainWindow{
		Title:  fileName,
		Icon:   getIcon(),
		Size:   Size{size + 100, size + 100},
		Layout: VBox{},
		Children: []Widget{
			ImageView{
				MinSize: Size{size, size},
				MaxSize: Size{size, size},
				Image:   bm,
			},
		},
	}.Run()
}
