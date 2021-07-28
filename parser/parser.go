package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/coseo12/gps-calculate/converter"
)

type Rotation struct {
	Yaw float64 `json:"yaw"`
	Pitch float64 `json:"pitch"`
	Roll float64 `json:"roll"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Slam struct {
	Position Position `json:"position"`
	Rotation Rotation `json:"rotation"`
}

type Scene struct {
	UserId string `json:"userId"`
	VendorId string `json:"venderId"`
	TourId string `json:"tourId"`
	Name string `json:"name"`
	OnProcess bool `json:"onProcess"`
	PanoJson map[string]interface{} `json:"panoJson"`
	Obj3d []string `json:"obj3d"`
	Slam Slam `json:"slam"`
	Autoleveled string `json:"autoleveled"`
	Autoblurred string `json:"autoblurred"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
	ImageUrl string `json:"imageUrl"`
	ThumbUrl string `json:"thumbUrl"`
	TilesUrl string `json:"tilesUrl"`
	Id string `json:"id"`
	FileName string `json:"fileName"`
	ThumbName string `json:"thumbName"`
	TilesNmae string `json:"tilesNmae"`
}
type Floorplan struct {
	VendorId string `json:"vendorId"`
	UserId string `json:"userId"`
	Width int `json:"width"`
	Height int `json:"height"`
	Scale float64 `json:"scale"`
	Dx int `json:"dx"`
	Dy int `json:"dy"`
	Opacity int `json:"opacity"`
	Angle int `json:"angle"`
	FileName string `json:"fileName"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
	Id string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	ThumbUrl string `json:"thumbUrl"`
	SmallImageUrl string `json:"smallImageUrl"`
	FpName string `json:"fpName"`
	FpThumbName string `json:"fpThumbName"`
	FpResizedThumbName string `json:"fpResizedThumbName"`
}
type Tour struct {
	UserId string `json:"userId"`
	VendorId string `json:"vendorId"`
	FloorId string `json:"floorId"`
	BuildingId string `json:"buildingId"`
	FpId string `json:"fpId"`
	Name string `json:"name"`
	Sweeps map[string]interface{} `json:"sweeps"`
	Groups []string `json:"groups"`
	Demorooms []string `json:"demorooms"`
	PlayerOptions map[string]interface{} `json:"playerOptions"`
	OnProcess bool `json:"onProcess"`
	Reconstruction map[string]interface{} `json:"reconstruction"`
	CameraHeight int `json:"cameraHeight"`
	CapturedAt int `json:"capturedAt"`
	ScenesOrder []string `json:"scenesOrder"`
	MeshRequested bool `json:"meshRequested"`
	WallHeight int `json:"wallHeight"`
	FolderId string `json:"folderId"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
	ImageUrl string `json:"imageUrl"`
	ThumbUrl string `json:"thumbUrl"`
	SharedThumbUrl string `json:"sharedThumbUrl"`
	Id string `json:"id"`
	Deleted bool `json:"deleted"`
	Scenes []Scene `json:"scenes"`
	Tags []string `json:"tags"`
	Floorplan Floorplan `json:"floorplan"`
	Annotations map[string]interface{} `json:"annotations"`
	Measurement string `json:"Measurement"`
	Tracks []string `json:"tracks"`
}
type Article struct {
	Tours []Tour `json:"tours"`
}

func (d *Article) SetGPS(cLat float64, cLng float64, filename string) {
	for i, tour := range(d.Tours) {
		for j, scene := range(tour.Scenes) {
			x := converter.ConvertDistanceInPixelsToMeter(scene.Slam.Position.X)
			y := converter.ConvertDistanceInPixelsToMeter(scene.Slam.Position.Z)
			convertData := converter.ConvertData{Lat: cLat, Lng: cLng, X: x, Y: y}
			lat, lng := converter.GetConvert(convertData)
			d.Tours[i].Scenes[j].Slam.Position.Lat = lat;
			d.Tours[i].Scenes[j].Slam.Position.Lng = lng;
		}
	}
	SetJson(d, filename)
}
 
func SetJson(articles *Article, filename string) {
	doc, _ := json.Marshal(articles) // data를 JSON 문서로 변환

	err := ioutil.WriteFile("./json/" + filename, doc, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		log.Panic(err)
	}
}

func GetJson(filename string) *Article {
	b, err := ioutil.ReadFile("./archive/" + filename) // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		log.Panic(err)
	}

	var data Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언
	
	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장
	
	return &data
}
