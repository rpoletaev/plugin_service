package plugin_service

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func testClient() {
	xmlBts, err := ioutil.ReadFile("/home/roma/fcsAddInfo_0376300004816000001_9595.xml")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:3131/load", bytes.NewBuffer(xmlBts))
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/xml")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
