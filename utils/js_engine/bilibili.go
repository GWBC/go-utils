package jsengine

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type BlibiliDashInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Dash              struct {
			Duration      int     `json:"duration"`
			MinBufferTime float64 `json:"min_buffer_time"`
			Video         []struct {
				ID           int      `json:"id"`
				BaseURL      string   `json:"base_url"`
				BackupURL    []string `json:"backup_url"`
				Bandwidth    int      `json:"bandwidth"`
				MimeType     string   `json:"mime_type"`
				Codecs       string   `json:"codecs"`
				Width        int      `json:"width"`
				Height       int      `json:"height"`
				FrameRate    string   `json:"frame_rate"`
				Sar          string   `json:"sar"`
				StartWithSap int      `json:"start_with_sap"`
				SegmentBase  struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				ID           int      `json:"id"`
				BaseURL      string   `json:"base_url"`
				BackupURL    []string `json:"backup_url"`
				Bandwidth    int      `json:"bandwidth"`
				MimeType     string   `json:"mime_type"`
				Codecs       string   `json:"codecs"`
				Width        int      `json:"width"`
				Height       int      `json:"height"`
				FrameRate    string   `json:"frame_rate"`
				Sar          string   `json:"sar"`
				StartWithSap int      `json:"start_with_sap"`
				SegmentBase  struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"audio"`
			Dolby struct {
				Type  int         `json:"type"`
				Audio interface{} `json:"audio"`
			} `json:"dolby"`
			Flac interface{} `json:"flac"`
		} `json:"dash"`
		SupportFormats []struct {
			Quality        int      `json:"quality"`
			Format         string   `json:"format"`
			NewDescription string   `json:"new_description"`
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Codecs         []string `json:"codecs"`
		} `json:"support_formats"`
		HighFormat   interface{} `json:"high_format"`
		LastPlayTime int         `json:"last_play_time"`
		LastPlayCid  int64       `json:"last_play_cid"`
		ViewInfo     interface{} `json:"view_info"`
		PlayConf     struct {
			IsNewDescription bool `json:"is_new_description"`
		} `json:"play_conf"`
	} `json:"data"`
}

// 主要结构体定义
type MPD struct {
	XMLName                   xml.Name `xml:"MPD"`
	MediaPresentationDuration string   `xml:"mediaPresentationDuration,attr"`
	MinBufferTime             string   `xml:"minBufferTime,attr"`
	Profiles                  string   `xml:"profiles,attr"`
	Type                      string   `xml:"type,attr"`
	Xmlns                     string   `xml:"xmlns,attr"`
	Xsi                       string   `xml:"xmlns:xsi,attr"`
	SchemaLocation            string   `xml:"xsi:schemaLocation,attr"`
	BaseURL                   string   `xml:"BaseURL"`
	Period                    Period   `xml:"Period"`
}

type Period struct {
	AdaptationSets []AdaptationSet `xml:"AdaptationSet"`
}

type AdaptationSet struct {
	MimeType                string           `xml:"mimeType,attr"`
	ContentType             string           `xml:"contentType,attr"`
	SubsegmentAlignment     string           `xml:"subsegmentAlignment,attr"`
	SubsegmentStartsWithSAP string           `xml:"subsegmentStartsWithSAP,attr"`
	Par                     string           `xml:"par,attr"`
	Representations         []Representation `xml:"Representation"`
}

type Representation struct {
	ID          string        `xml:"id,attr"`
	Bandwidth   string        `xml:"bandwidth,attr"`
	Width       string        `xml:"width,attr"`
	Height      string        `xml:"height,attr"`
	Codecs      string        `xml:"codecs,attr"`
	MimeType    string        `xml:"mimeType,attr"`
	BaseURL     string        `xml:"BaseURL"`
	SegmentBase []SegmentBase `xml:"SegmentBase"`
}

type SegmentBase struct {
	IndexRange     string         `xml:"indexRange,attr,omitempty"`
	Initialization Initialization `xml:"Initialization"`
}

type Initialization struct {
	Range string `xml:"range,attr"`
}

func BlibiliData2MPD(biliData string, proxyPath string) string {
	dashInfo := BlibiliDashInfo{}
	err := json.Unmarshal([]byte(biliData), &dashInfo)
	if err != nil {
		return ""
	}

	mpd := MPD{}
	mpd.Profiles = "urn:hbbtv:dash:profile:isoff-live:2012,urn:mpeg:dash:profile:isoff-live:2011"
	mpd.Type = "static"
	mpd.Xmlns = "urn:mpeg:dash:schema:mpd:2011"
	mpd.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	mpd.SchemaLocation = "urn:mpeg:DASH:schema:MPD:2011 DASH-MPD.xsd"
	mpd.MediaPresentationDuration = fmt.Sprintf("PT%vS", dashInfo.Data.Dash.Duration)
	mpd.MinBufferTime = fmt.Sprintf("PT%vS", dashInfo.Data.Dash.MinBufferTime)

	proxyParam := ""

	if len(proxyPath) != 0 {
		proxyRes := strings.Split(proxyPath, "?")
		if len(proxyRes) == 2 {
			mpd.BaseURL = proxyRes[0]
			proxyParam = proxyRes[1]
			if !strings.HasSuffix(mpd.BaseURL, "/") {
				mpd.BaseURL += "/"
			}
		}
	}

	var videoSet *AdaptationSet

	for _, v := range dashInfo.Data.Dash.Video {
		if videoSet == nil {
			videoSet = &AdaptationSet{}
			videoSet.MimeType = v.MimeType
			videoSet.ContentType = "video"
			videoSet.SubsegmentAlignment = "true"
			videoSet.SubsegmentStartsWithSAP = strconv.Itoa(v.StartWithSap)
			videoSet.Par = "16:9"
		}

		represent := Representation{}
		represent.ID = strconv.Itoa(v.ID)
		represent.Bandwidth = strconv.Itoa(v.Bandwidth)
		represent.Width = strconv.Itoa(v.Width)
		represent.Height = strconv.Itoa(v.Height)
		represent.Codecs = v.Codecs
		represent.MimeType = v.MimeType

		if len(proxyParam) != 0 {
			represent.BaseURL += "?" + proxyParam
			represent.BaseURL += base64.URLEncoding.EncodeToString([]byte(v.BaseURL))
		} else {
			represent.BaseURL += v.BaseURL
		}

		baseUrl := SegmentBase{}
		baseUrl.IndexRange = v.SegmentBase.IndexRange
		baseUrl.Initialization.Range = v.SegmentBase.Initialization
		represent.SegmentBase = append(represent.SegmentBase, baseUrl)

		videoSet.Representations = append(videoSet.Representations, represent)
	}

	if videoSet != nil {
		mpd.Period.AdaptationSets = append(mpd.Period.AdaptationSets, *videoSet)
	}

	data, err := xml.MarshalIndent(&mpd, "", "  ")
	if err != nil {
		return ""
	}

	return string(data)
}
