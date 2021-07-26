# GPS CALCULATE

Distance convert to decimal degrees

## Converter

```go
 type ConvertData struct {
    Lat float64
    Lng float64
    X float64
    Y float64
  }

  GetConvert(convertData ConvertData) (float64, float64) {}
```

## Parser

```go
type Article struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetJson(articles []Article) {}

func GetJson() []Article {}
```
