package station

type Station struct {
	Id   string `json:"nid"`
	Name string `json:"title"`
}

type Response struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
