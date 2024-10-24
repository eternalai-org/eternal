package zip_hf_model_to_light_house

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

func DownloadHFFileNew(wg *sync.WaitGroup, hfFile chan HFModelZipFile, modelPath string, dlChan chan DownloadFileChan) {
	DownloadHFFile(wg, <-hfFile, modelPath, dlChan)
}

func downloadZipFileFromLightHouseNewNew(info *HFModelInLightHouse, hfDir string) error {
	files := info.Files
	testData := []string{}
	//index downloaded files for sorting.
	if len(files) == 0 {
		return nil
	}

	limit := 8
	errorNum := 0
	var wg sync.WaitGroup
	//total := info.NumOfFile
	//limit := 8 //8 process at once

	retryDownload := []HFModelZipFile{}
	dlChan := make(chan DownloadFileChan)
	inputChan := make(chan HFModelZipFile)

	//_t := float64(total) / float64(limit)
	//create channels
	for range files {
		go DownloadHFFileNew(&wg, inputChan, hfDir, dlChan)
	}

	var mux sync.Mutex
	processing := 0

	allDone := make(chan bool)
	go func(allDone chan bool) {
		for _, _ = range files {
			dFChan := <-dlChan
			if dFChan.Err != nil {
				errorNum++
				log.Println("[DownloadFile][Error] File: ", dFChan.Data.File, " ,error: ", dFChan.Err)
				retryDownload = append(retryDownload, dFChan.Data)

				//remove file if error
				os.Remove(fmt.Sprintf("%s/%s", hfDir, dFChan.Data.File))

			} else {
				log.Println("[DownloadFile][Success] File: ", dFChan.Data.File, " ,filePath: ", *dFChan.Msg)
			}

			mux.Lock()
			processing--
			mux.Unlock()

			testData = append(testData, dFChan.Data.File)

		}

		allDone <- true
	}(allDone)

	index := 0
	for {
		if processing == limit {
			//fmt.Println("====> wait a minute:  ", processing)
		} else {
			wg.Add(1)
			inputChan <- files[index]
			index++
			mux.Lock()
			processing++
			mux.Unlock()
			fmt.Println("====> push to process: ", processing)
		}

		if index == len(files)-1 {
			break
		}
	}

	wg.Wait()
	<-allDone
	if errorNum != 0 {
		retryInfo := HFModelInLightHouse{
			Model:     info.Model,
			NumOfFile: len(retryDownload),
			Files:     retryDownload,
		}

		return downloadZipFileFromLightHouseNew(&retryInfo, hfDir)
		//return errors.New("error while download file")
	}

	return nil
}
