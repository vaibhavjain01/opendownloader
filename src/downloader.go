package main

import (
	"os/exec"
	"strings"
	"container/list"
)

type Downloader struct {
	LogObj Logger
	UtilityObj Utility
}

type AudioInfoJson struct {
	FormatCode	string	`json: formatCode`
	Extension	string	`json: extension`
	Type	string	`json: type`
	AudioQuality	string	`json: audioQuality`
	MediaSize	string	`json: mediaSize`
}

type VideoInfoJson struct {
	FormatCode	string	`json: formatCode`
	Extension	string	`json: extension`
	Type	string	`json: type`
	PixelInfo	string	`json: pixelInfo`
	AudioQuality	string	`json: audioQuality`
	MediaSize	string	`json: mediaSize`
}

type AvailableFormatsJson struct {
	URL	string	`json: url`
	AvailableAudioFormats	[]*AudioInfoJson	`json: availableAudioFormats`
	AvailableVideoFormats	[]*VideoInfoJson	`json: availableVideoFormats`
}

type DownloadedMediaJson struct {
	URL	string	`json: url`
	DownloadedURL	string	`json: downloadedUrl`
}

func (downloader Downloader) CheckAvailableFormats(targetUrl string) (*AvailableFormatsJson) {
	downloader.LogObj.LogToConsole("Checking Available Formats on link: " + targetUrl)
	out, err := exec.Command("sh", "-c", "./udlbin/youtube-dl -F " + targetUrl).Output()
	if err != nil {
		downloader.LogObj.LogToConsole("Could not find information. Error Occurred")
		downloader.LogObj.LogToConsole(err.Error())
		return nil
	} else {
		info := strings.Split(string(out), "\n")
		listAudioJson := list.New()
		listVideoJson := list.New()

		for i := 0; i < len(info); i++ {
			if i > 3 {
				audioInfoJson, videoInfoJson := downloader._convertMediaInfoToJson(info[i])
				if audioInfoJson == nil && videoInfoJson == nil {
					// do nothing
				} else if audioInfoJson == nil {
					listVideoJson.PushBack(videoInfoJson)
				} else {
					listAudioJson.PushBack(audioInfoJson)
				}
			}
		}
		
		var availableFormatsJson AvailableFormatsJson
		availableFormatsJson.URL = targetUrl
		availableFormatsJson.AvailableAudioFormats =  make([]*AudioInfoJson, listAudioJson.Len())
		availableFormatsJson.AvailableVideoFormats =  make([]*VideoInfoJson, listVideoJson.Len())

		for element,i := listAudioJson.Front(), 0; element != nil; element,i = element.Next(), i + 1 {
			availableFormatsJson.AvailableAudioFormats[i] = element.Value.(*AudioInfoJson)
		}
		for element,i := listVideoJson.Front(), 0; element != nil; element,i = element.Next(), i + 1 {
			availableFormatsJson.AvailableVideoFormats[i] = element.Value.(*VideoInfoJson)
		}

		return &availableFormatsJson
	}	
}

func (downloader Downloader) _convertMediaInfoToJson(mediaInfo string) (*AudioInfoJson, *VideoInfoJson) {
	
	mediaInfo = downloader.UtilityObj._removeExtraSpaces(mediaInfo)
	mediaInfoArr := strings.Split(mediaInfo, " ")
	
	if(len(mediaInfoArr) < 5) {
		return nil, nil
	}

	if strings.Contains(mediaInfoArr[2], "audio") {
		audioInfoJson := new(AudioInfoJson)
		audioInfoJson.Type = "audio"
		audioInfoJson.AudioQuality = mediaInfoArr[5]
		audioInfoJson.FormatCode = mediaInfoArr[0]
		audioInfoJson.Extension = mediaInfoArr[1]

		remaining := strings.Join(mediaInfoArr[6:]," ")
		remainingArr := strings.Split(remaining, ",")
		remaining = remainingArr[len(remainingArr) - 1]
		remaining = strings.Split(remaining, " ")[1]
		audioInfoJson.MediaSize = remaining

		return audioInfoJson, nil

	} else {
		videoInfoJson := new(VideoInfoJson)
		videoInfoJson.Type = "video"
		videoInfoJson.PixelInfo = mediaInfoArr[3]
		videoInfoJson.AudioQuality = mediaInfoArr[4]
		videoInfoJson.FormatCode = mediaInfoArr[0]
		videoInfoJson.Extension = mediaInfoArr[1]

		remaining := strings.Join(mediaInfoArr[5:]," ")
		remainingArr := strings.Split(remaining, ",")
		remaining = remainingArr[len(remainingArr) - 1]
		remaining = strings.Split(remaining, " ")[1]
		videoInfoJson.MediaSize = remaining

		return nil, videoInfoJson
	}
}

func (downloader Downloader) BeginDownload (targetUrl string, targetFormatCode string) (*DownloadedMediaJson) {
	downloader.LogObj.LogToConsole("Downloading media from link: " + targetUrl + ", with extension: " + targetFormatCode)
	out, err := exec.Command("sh", "-c", "./udlbin/youtube-dl -f " + targetFormatCode + " " + targetUrl).Output()
	if err != nil {
		downloader.LogObj.LogToConsole("Could not find information. Error Occurred")
		downloader.LogObj.LogToConsole(err.Error())
		return nil
	} else {
		var downloadedMediaJson DownloadedMediaJson
		downloadedMediaJson.URL = targetUrl
		downloadedMediaJson.DownloadedURL = "downloadlink.com"
		downloader.LogObj.LogToConsole("File Downloaded")
		downloader.LogObj.LogToConsole(string(out))
		return &downloadedMediaJson
	}
}