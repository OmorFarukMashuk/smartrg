package smartrg

type DTORequest struct {
	DTO DTO `json:"dto"`
}

type DTO struct {
	Subscriber    Subscriber      `json:"Subscriber"`
	Revision      string          `json:"revision,omitempty"`
	Subscriptions []interface{}   `json:"subscriptions"`
	Labels        []Label         `json:"labels"`
	Credentials   [][]interface{} `json:"credentials"`
	AccountCode   string          `json:"code"`
}

type Subscriber struct {
	FullName     string
	EmailAddress string
}

type Label struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	FGColour string `json:"fgColor"`
	BGColour string `json:"bgColor"`
}

type Device struct {
	MAC           string        `json:"sn"`
	OUI           string        `json:oui`
	ID            int           `json:deviceId`
	Disposition   string        `json:"disposition"`
	ActionLog     []interface{} `json:"actionLog"`
	Applications  interface{}   `json:"applications"`
	Labels        []string      `json:"labels"`
	QueuedActions ActionList    `json:"queuedActions"`
}

type ActionList struct {
	Applications interface{} `json:"applications"`
	Scripts      []string    `json:"scripts"`
	Services     interface{} `json:"services"`
}
