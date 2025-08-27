package entities

type JadwalSholat struct {
	Id     int    `json:"id"`
	Lokasi string `json:"lokasi"`
	Daerah string `json:"daerah"`
	Jadwal struct {
		Imsak   string `json:"imsak"`
		Subuh   string `json:"subuh"`
		Terbit  string `json:"terbit"`
		Dhuha   string `json:"dhuha"`
		Dzuhur  string `json:"dzuhur"`
		Ashar   string `json:"ashar"`
		Maghrib string `json:"maghrib"`
		Isya    string `json:"isya"`
		Date    string `json:"date"`
	}
}

type JadwalSholaCityIdRequest struct {
	Path    string `json:"path"`
	Keyword string `json:"keyword"`
}

type JadwalSholaCityIdResponse struct {
	Status  bool                     `json:"status"`
	Request JadwalSholaCityIdRequest `json:"request"`
	Data    []struct {
		Id     string `json:"id"`
		Lokasi string `json:"lokasi"`
	} `json:"data"`
}

type JadwalSholatRequest struct {
	Path  string `json:"path"`
	Year  string `json:"year"`
	Month string `json:"month"`
	Date  string `json:"date"`
}

type JadwalSholatResponse struct {
	Status  bool                `json:"status"`
	Request JadwalSholatRequest `json:"request"`
	Data    JadwalSholat        `json:"data"`
}
