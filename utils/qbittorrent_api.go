package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Torrent 表示单个种子的详细信息
type Torrent struct {
	AddedOn      *int64   `json:"added_on"`      //添加时间
	AmountLeft   *int64   `json:"amount_left"`   //剩余下载量
	Completed    *int64   `json:"completed"`     //已完成下载量
	CompletionOn *int64   `json:"completion_on"` //完成时间
	ContentPath  *string  `json:"content_path"`  //下载路径
	Downloaded   *int64   `json:"downloaded"`    //总计下载量
	Name         *string  `json:"name"`          //资源名称
	Popularity   *float64 `json:"popularity"`    //资源热度
	Progress     *float64 `json:"progress"`      //进度 0 -1
	SavePath     *string  `json:"save_path"`     //保存路径
	Size         *int64   `json:"size"`          //总大小
	State        *string  `json:"state"`         //当前状态
	Uploaded     *int64   `json:"uploaded"`      //总上传量
	UpSpeed      *int     `json:"upspeed"`       //上传速度
}

// TorrentList 是种子的映射，键是Infohash
type TorrentList map[string]Torrent

type QBittorrentResponse struct {
	Rid      int          `json:"rid"`
	Torrents *TorrentList `json:"torrents"`
}

func (q *QBittorrentResponse) String() string {
	data, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(data)
}

////////////////////////////////////////////////////////

type QbittorrentApi struct {
	cookiesName string
	host        string
}

func (q *QbittorrentApi) SetHost(host string) *QbittorrentApi {
	q.host = host
	return q
}

func (q *QbittorrentApi) Login(user string, pwd string) error {
	q.cookiesName = "qbittorrent"
	data := map[string]string{}
	data["username"] = user
	data["password"] = pwd
	rsp, err := PostForm(fmt.Sprintf("http://%s/api/v2/auth/login", q.host), nil, data, q.cookiesName)
	if err != nil {
		return err
	}

	if string(rsp) != "Ok." {
		return errors.New("user or pwd error")
	}

	return nil
}

func (q *QbittorrentApi) GetAllInfo() (*QBittorrentResponse, error) {
	rsp, err := Get(fmt.Sprintf("http://%s/api/v2/sync/maindata?rid=0", q.host), nil, nil, q.cookiesName)
	if err != nil {
		return nil, err
	}

	qbRsp := &QBittorrentResponse{}
	err = json.Unmarshal(rsp, qbRsp)
	if err != nil {
		return nil, err
	}

	return qbRsp, nil
}

func (q *QbittorrentApi) Start(hashVal string) error {
	data := map[string]string{}
	data["hashes"] = hashVal
	rsp, err := PostForm(fmt.Sprintf("http://%s/api/v2/torrents/start", q.host), nil, data, q.cookiesName)
	if err != nil {
		return err
	}

	if len(rsp) != 0 {
		return errors.New(string(rsp))
	}

	return nil
}

func (q *QbittorrentApi) Stop(hashVal string) error {
	data := map[string]string{}
	data["hashes"] = hashVal
	rsp, err := PostForm(fmt.Sprintf("http://%s/api/v2/torrents/stop", q.host), nil, data, q.cookiesName)
	if err != nil {
		return err
	}

	if len(rsp) != 0 {
		return errors.New(string(rsp))
	}

	return nil
}

func (q *QbittorrentApi) Add(urls string, savepath string) error {
	data := map[string]string{}
	data["urls"] = urls
	data["savepath"] = savepath
	rsp, err := PostForm(fmt.Sprintf("http://%s/api/v2/torrents/add", q.host), nil, data, q.cookiesName)
	if err != nil {
		return err
	}

	if len(rsp) != 0 {
		return errors.New(string(rsp))
	}

	return nil
}

func (q *QbittorrentApi) Delete(hashVal string, deleteFiles bool) error {
	data := map[string]string{}
	data["hashes"] = hashVal
	data["deleteFiles"] = strconv.FormatBool(deleteFiles)
	rsp, err := PostForm(fmt.Sprintf("http://%s/api/v2/torrents/delete", q.host), nil, data, q.cookiesName)
	if err != nil {
		return err
	}

	if len(rsp) != 0 {
		return errors.New(string(rsp))
	}

	return nil
}
