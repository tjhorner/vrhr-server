package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/tjhorner/vrhr-server/ui"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	writeCsv := flag.Bool("csv", false, "Write data to a csv file")
	flag.Parse()

	now := time.Now()

	context := Context{
		WriteToCsv: *writeCsv,
	}

	if context.WriteToCsv {
		fn := now.Format("vrhr_data_2006-01-02_15-04-05.csv")
		fmt.Printf("Writing data to %s\n", fn)

		file, err := os.Create(fn)
		if err != nil {
			fmt.Printf("Failed to open file for writing: %s\n", err)
			context.WriteToCsv = false
		} else {
			defer file.Close()
			file.WriteString("Date,Heart Rate,Accuracy\n")
			context.CsvFile = file
		}
	}

	router := mux.NewRouter()
	routeApi(router, &context)

	frontendBox := packr.New("frontend", "./frontend/build")

	serveIndex := func(w http.ResponseWriter, r *http.Request) {
		index, err := frontendBox.Find("index.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}

		w.Write(index)
	}

	serveSpa := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		if !frontendBox.Has(path) {
			serveIndex(w, r)
			return
		}

		http.FileServer(frontendBox).ServeHTTP(w, r)
	}

	// http.FileServer will redirect index.html to / and will end up in a redirect loop.
	// So we need to do this to fix that redirect loop.
	router.HandleFunc("/", serveIndex)
	router.PathPrefix("/").HandlerFunc(serveSpa)

	listener, err := net.Listen("tcp", envListenAddr)
	if err != nil {
		ui.Error(fmt.Errorf("error when attempting to listen: %v", err))
		return
	}

	fmt.Printf("vrhr is now listening at %s\n", envListenAddr)

	go http.Serve(listener, handlers.CORS()(router))

	pc, err := getPairingCode(listener.Addr().(*net.TCPAddr).Port)
	if err != nil {
		ui.Error(fmt.Errorf("error generating pairing code: %v", err))
		return
	}

	mdnsServer := advertiseService(listener.Addr().(*net.TCPAddr).Port)
	defer mdnsServer.Shutdown()

	ui.Run(pc)
}
