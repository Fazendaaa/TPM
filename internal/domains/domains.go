package domains

type domain struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
}

type subdomain struct {
	Port   int    `json:"port"`
	Name   string `json:"name"`
	Domain domain `json:"domain"`
}

var DomainTest = []domain{
	{
		IP:   "192.168.0.1",
		Name: "example.com",
	},
	{
		IP:   "192.168.0.2",
		Name: "example.org",
	},
	{
		IP:   "192.168.0.3",
		Name: "example.net",
	},
}

var SubdomainTest = []subdomain{
	{
		Port: 80,
		Name: "about",
		Domain: domain{
			IP:   "192.168.0.1",
			Name: "example.com",
		},
	},
	{
		Port: 443,
		Name: "mail",
		Domain: domain{
			IP:   "192.168.0.1",
			Name: "example.com",
		},
	},
	{
		Port: 8080,
		Name: "api",
		Domain: domain{
			IP:   "192.168.0.2",
			Name: "example.org",
		},
	},
}
