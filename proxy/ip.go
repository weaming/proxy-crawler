package proxy

type IP struct {
	IP       string `bson:"ip" json:"ip"`
	Port     int    `bson:"port" json:"port"`
	Protocol string `bson:"protocol" json:"protocol"`
	Usable   bool   `bson:"usable" json:"usable"`
}

func NewIP() *IP {
	return &IP{}
}
