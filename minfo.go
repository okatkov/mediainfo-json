package media

import (
	"os/exec"
	"log"
	"encoding/json"
)

type container struct {
	Media struct{
		Ref string `json:"@ref"`
		Track []interface{} `json:"track"`

	}
}

type mediainfo struct {
	General general
	Video video
	Audio audio
}


type general struct {
	Format                string `json:"Format"`
	Duration              string `json:"Duration"`
	File_size             string `json:"FileSize"`
	Overall_bit_rate_mode string `json:"OverallBitRateMode"`
	Overall_bit_rate      string `json:"OverallBitRate"`
	Complete_name         string `json:"CompleteName"`
	File_name             string `json:"FileName"`
	File_extension        string `json:"FileExtension"`
	Frame_rate            string `json:"FrameRate"`
	Stream_size           string `json:"StreamSize"`
	Writing_application   string `json:"WritingApplication"`
}

type video struct {
	Width                     string `json:"Width"`
	Height                    string `json:"Height"`
	Format                    string `json:"Format"`
	Bit_rate                  string `json:"Bitrate"`
	Duration                  string `json:"Duration"`
	Format_Info               string `json:"FormatInfo"`
	Format_profile            string `json:"FormatProfile"`
	Frame_rate                string `json:"FrameRate"`
	Bit_depth                 string `json:"BitDepth"`
	Scan_type                 string `json:"ScanType"`
	Interlacement             string `json:"Interlacement"`
	Writing_library           string `json:"WritingLibrary"`
}

type audio struct {
	Format         string `json:"Format"`
	Duration       string `json:"Duration"`
	Bit_rate       string `json:"Bitrate"`
	Channel_s_     string `json:"Channels"`
	Frame_rate     string `json:"FrameRate"`
	Format_Info    string `json:"FormatInfo"`
	Sampling_rate  string `json:"SamplingRate"`
	Format_profile string `json:"FormatProfile"`
}


func GetInfo(url string)(info mediainfo){
	container := container{}
	general := general{}
	video := video{}
	audio := audio{}

	res, err := exec.Command("mediainfo", "--Output=JSON", "-f", url).Output()

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(res,&container)

	raw0,_:=  json.Marshal(container.Media.Track[0])
	json.Unmarshal(raw0,&general)
	raw1,_:=  json.Marshal(container.Media.Track[1])
	json.Unmarshal(raw1,&video)
	raw2,_:=  json.Marshal(container.Media.Track[2])
	json.Unmarshal(raw2,&audio)

	return mediainfo{general,video,audio}
}
