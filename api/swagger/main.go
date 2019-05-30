package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/api/swagger/model"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	// parse flags
	var version, outPath string
	flag.StringVar(&version, "version", "v1.5",
		"The version to load.",
	)
	flag.StringVar(&outPath, "out", "",
		"The `file` to generate the swagger json to..",
	)
	flag.Parse()

	inUrl := fmt.Sprintf("https://sc1.checkpoint.com/documents/latest/APIs/data/%s/dynamic/apis.json", version)
	mUrl, err := url.Parse(inUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	inUrl = fmt.Sprintf("https://sc1.checkpoint.com/documents/latest/APIs/data/%s/dynamic/content.json", version)
	tUrl, err := url.Parse(inUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	outPath, err = filepath.Abs(filepath.FromSlash(outPath))
	if err != nil {
		fmt.Println(err)
		return
	}

	// load tags
	buffer, err := loadFromUrl(tUrl.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	tags, err := model.UnmarshalTags(*buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// load model
	buffer, err = loadFromUrl(mUrl.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	apiModel, err := model.UnmarshalApi(*buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert
	swaggerJson, err := apiModel.Convert(tags).MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	// dump to disk
	writeJson(filepath.Join(outPath, fmt.Sprintf("%s-swagger.json", apiModel.Version)), swaggerJson)
}

func loadApi(data *bytes.Buffer) (error, *model.Api) {
	err, api := model.UnmarshalApi(*data)
	if err != nil {
		return nil, err
	}

	return api, err
}

func loadTags(data *bytes.Buffer) (error, *model.Tags) {
	err, api := model.UnmarshalTags(*data)
	if err != nil {
		return nil, err
	}

	return api, err
}

func loadFromUrl(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, resp.Body)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func writeJson(file string, jsonData []byte) {
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, jsonData, "", "\t")
	if _, err := os.Stat(file); os.IsExist(err) {
		_ = os.Remove(file)
	}

	f, _ := os.Create(file)
	defer f.Close()

	_, _ = f.Write(prettyJSON.Bytes())

	f.Sync()
}
