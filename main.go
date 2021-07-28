package main

import "github.com/coseo12/gps-calculate/parser"

type files struct {
	name string
	lat float64
	lng float64
}

func main() {

	fileList := []files{}

	fileList = append(fileList,files{"210621_complete_210629.437a8ccf-f083-4cb5-84a3-544d074c57be.main.json", 35.881493, 128.573618});
	fileList = append(fileList,files{"210621_complete_210702.0ad76f69-e0ae-4f89-9ebe-84f273c401a5.main.json", 35.878925, 128.575591});
	fileList = append(fileList,files{"210621_complete_210704.12eea293-9429-4989-82d9-d54e68312920.main.json", 35.880059, 128.574067});
	fileList = append(fileList,files{"210622_complete_210702.70aef337-d823-4a53-8474-a12f20050674.main.json", 35.880063, 128.57245});
	fileList = append(fileList,files{"210622_complete_210704.aee41662-776c-45cb-8f7a-0d585051b91f.main.json", 35.878925, 128.575591});
	fileList = append(fileList,files{"210622_complete_210704.e3611664-98b3-4bec-b106-ee494e37788f.main.json", 35.878925, 128.575591});
	fileList = append(fileList,files{"210622-1_complete_210629.adb7811d-ef78-4e7e-873b-be6b8f941bb6.main.json", 35.880063, 128.57245});

	for _, file := range(fileList) {
		data := parser.GetJson(file.name)
		data.SetGPS(file.lat, file.lng, file.name)
	}
}