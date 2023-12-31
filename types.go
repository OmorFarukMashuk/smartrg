package smartrg

import (
	"time"
)

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

type ACSDeviceStatus struct {
	InformURL    string    `json:"ConneectionRequestURL"`
	FirstInform  time.Time `json:"firstInform"`
	LastInform   time.Time `json:"lastInform"`
	MAC          string    `json:"serialNumber"`
	Firmware     string    `json:"SoftwareVersion"`
	SubscriberID int       `json:subscriberid`
}

type ACSSubscriber struct {
	DTO           struct{}      `json:"dto"`
	Revision      interface{}   `json:"revision,omitempty"`
	Subscriptions []interface{} `json:"subscriptions,nilasempty"`
	Labels        []ACSLabel    `json:"labels,nilasempty"`
	Credentials   struct {
		Login    string `json:"login,omitempty"`
		Password string `json:"password,omitempty"`
		Locked   bool   `json:"lockUser,omitempty"`
	} `json:"credentials"`
	Accountcode  string `json:"code"`
	SubscriberID int    `json:"subscriberID,omitempty"`
	Attributes   struct {
		Email string `json:"Subscriber.EmailAddress"`
		Name  string `json:"Subscriber.FullName"`
	} `json:"attributes"`
}

/*
type ACSSubscriber struct {
	Document struct {
		Subscriber struct {
			FullName string `json:"FullName,omitempty"`
			Email    string `json:"EmailAddress,omitempty"`
		}
	} `json:"dto,omitempty"`
	Revision       interface{} `json:"revision,omitempty"`
	Subscriptions  []string    `json:"subscriptions"`
	Labels         []ACSLabel  `json:"labels"`
	SubscriberCode int         `json:"subscriberId,omitempty"`
	Credentials    struct {
		Login    string `json:"login,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"credentials"`
	Accountcode string `json:"code"`
}
*/
type ACSDevice struct {
	Accountcode   string                 `json:"subscriberCode"`
	MAC           string                 `json:"sn"`
	OUI           string                 `json:"oui"`
	ActionLog     []string               `json:"actionLog"`
	Labels        []ACSLabel             `json:"labels"`
	Applications  map[string]interface{} `json:"applications"`
	QueuedActions map[string]interface{} `json:"queuedActions"`
	DeviceID      int                    `json:"deviceid"`
	Disposition   string                 `json:"disposition"`
}

type ACSLabel struct {
	Name     string `json:"name"`
	FGColour string `json:"fgcolor"`
	BGColour string `json:"bgcolor"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

type ACSResponse struct {
}

// Device record return components

type ACSDeviceRecord struct {
	ID     string       `json:"id"`
	DocID  string       `json:"docid"`
	Type   string       `json:"type"`
	Fields DeviceFields `json:"fields"`
}

type DeviceFields struct {
	WanType        []string  `json:"wanType"`
	Model          string    `json:"model"`
	WanMAC         []string  `json:"wanMacAddress"`
	Email          string    `json:"emailAddress"`
	FullName       string    `json:"fullName"`
	Serial         string    `json:"serialNumber"`
	WanIPv4        []string  `json:"wanIPv4Address"`
	LastInform     time.Time `json:"lastInform"`
	SubscriberCode string    `json:"subscriberCode"`
	OUI            string    `json:"oui"`
	SubscriberID   string    `json:"subscriberId"`
	DeviceID       string    `json:"deviceId"`
	Manufacturer   string    `json:"manufacturer"`
	Disposition    string    `json:"disposition"`
}
