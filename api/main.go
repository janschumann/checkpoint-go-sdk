package main

//go:generate go run ./main.go -in=https://sc1.checkpoint.com/documents/latest/APIs/data/v1.1/dynamic/apis.json -out=../
//go:generate go run ./main.go -in=https://sc1.checkpoint.com/documents/latest/APIs/data/v1.2/dynamic/apis.json -out=../
//go:generate go run ./main.go -in=https://sc1.checkpoint.com/documents/latest/APIs/data/v1.3/dynamic/apis.json -out=../
//go:generate go run ./main.go -in=https://sc1.checkpoint.com/documents/latest/APIs/data/v1.4/dynamic/apis.json -out=../
//go:generate go run ./main.go -in=https://sc1.checkpoint.com/documents/latest/APIs/data/v1.5/dynamic/apis.json -out=../

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/api/model"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	// parse flags
	var inUrl, outPath string
	flag.StringVar(&inUrl, "in", "api.json",
		"URL to the model file.",
	)
	flag.StringVar(&outPath, "out", "",
		"The `file` to generate the swagger json to..",
	)
	flag.Parse()

	// prepare and validate flags
	pUrl, err := url.Parse(inUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	outPath, err = filepath.Abs(filepath.FromSlash(outPath))
	if err != nil {
		fmt.Println(err)
		return
	}

	// load model
	apiModel, err := loadApi(pUrl.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert
	swaggerJson, err := apiModel.Convert().MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	// dump to disk
	writeSwagger(filepath.Join(outPath, fmt.Sprintf("%s-swagger.json", apiModel.Version)), swaggerJson)
}

func loadApi(url string) (*model.Api, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data bytes.Buffer
	_, err = io.Copy(&data, resp.Body)
	if err != nil {
		return nil, err
	}

	err, api := model.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return api, err
}

func writeSwagger(file string, swaggerJson []byte) {
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, swaggerJson, "", "\t")
	if _, err := os.Stat(file); os.IsExist(err) {
		_ = os.Remove(file)
	}

	f, _ := os.Create(file)
	defer f.Close()

	_, _ = f.Write(prettyJSON.Bytes())

	f.Sync()
}
