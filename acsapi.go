package smartrg

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	//"net"
	"bytes"
	"net/http"
	"strconv"
	//"sync"
	"fmt"
	"strings"
	//	"telmax"
	"time"
)

var (
	APIURL  = flag.String("apiurl", "https://telmax.smartrg.ca/prime-home/", "API Base URL")
	APIUser = flag.String("apiuser", "telmax-api", "API Auth User")
	APIPass = flag.String("apipass", "etrr&OG(^i", "API Password")
)

func getData(request string) (result []byte, err error) {
	Client := http.Client{
		Timeout: time.Second * 4,
	}
	url := *APIURL + request
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("Problem generating HTTP request %v", err)
		return
	}
	req.SetBasicAuth(*APIUser, *APIPass)
	var response *http.Response
	response, err = Client.Do(req)
	if err != nil {
		log.Errorf("Problem with HTTP request execution %v", err)
		return
	}
	defer response.Body.Close()
	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Problem Reading HTTP Response %v", err)
		return
	}
	return
}

func deleteData(request string) (result []byte, err error) {
	Client := http.Client{
		Timeout: time.Second * 4,
	}
	url := *APIURL + request
	var req *http.Request
	req, err = http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Errorf("Problem generating HTTP request %v", err)
		return
	}
	req.SetBasicAuth(*APIUser, *APIPass)
	var response *http.Response
	response, err = Client.Do(req)
	if err != nil {
		log.Errorf("Problem with HTTP request execution %v", err)
		return
	}
	defer response.Body.Close()
	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Problem Reading HTTP Response %v", err)
		return
	}
	return
}

func sendData(method string, request string, data interface{}) (result []byte, err error) {
	Client := http.Client{
		Timeout: time.Second * 4,
	}
	url := *APIURL + request
	var req *http.Request
	var jsonStr []byte
	jsonStr, err = json.Marshal(data)
	log.Debugf("Posted string is %v", string(jsonStr))
	if err != nil {
		log.Errorf("Problem marshalling JSON data", err)
		return
	}
	req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Errorf("Problem generating HTTP request %v", err)
		return
	}
	req.SetBasicAuth(*APIUser, *APIPass)
	var response *http.Response
	response, err = Client.Do(req)
	if err != nil {
		log.Errorf("Problem with HTTP request execution %v", err)
		return
	}
	defer response.Body.Close()
	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Problem Reading HTTP Response %v", err)
		return
	}
	return
}

func NewSubscriber(acct ACSSubscriber) (code int, err error) {

	//	var templateResult []byte
	//	templateResult, err = getData("api/v2/templates/acctcriber")
	//	err = json.Unmarshal(templateResult, &acsAcct)

	var createResult []byte
	createResult, err = sendData(http.MethodPost, "api/v2/subscribers", acct)
	log.Debug(string(createResult))
	var errorMessage []ErrorMessage
	errorError := json.Unmarshal(createResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem creating account %v", errorMessage)
		err = fmt.Errorf("Problem creating account %v", errorMessage[0].Message)
		return
	}
	err = json.Unmarshal(createResult, &acct)
	code = acct.SubscriberCode
	return
}

/*
func ModifySubscriber(acct ACSSubscriber) error {
	fetchResult, err := getData("api/v2/subscribers/" + strconv.Itoa(acct.ACSSubscriber))
	log.Debug(string(fetchResult))

	err = json.Unmarshal(fetchResult, &acsAcct)
	var errorMessage []ErrorMessage
	errorError := json.Unmarshal(fetchResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem fetching account %v", errorMessage)
		err = fmt.Errorf("Problem fetching account %v", errorMessage[0].Message)
		return err
	}

	acsAcct.Labels = []ACSLabel{
		ACSLabel{
			Name:     acct.AccountType,
			FGColour: "#000",
			BGColour: "#fff",
		},
	}

	acsAcct.Accountcode = acct.AccountCode
	acsAcct.Document.Subscriber.Email = acct.Email
	if acct.CompanyName != "" {
		acsAcct.Document.Subscriber.FullName = acct.CompanyName
	} else {
		acsAcct.Document.Subscriber.FullName = acct.FullName
	}
	acsAcct.Credentials.Login = acct.Email
	acsAcct.Credentials.Password = "Telmax@5*"

	var updateResult []byte
	updateResult, err = sendData(http.MethodPut, "api/v2/subscribers/"+strconv.Itoa(acsAcct.SubscriberCode), acsAcct)
	log.Debug(string(updateResult))
	errorError = json.Unmarshal(updateResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem modifying account %v", errorMessage)
		err = fmt.Errorf("Problem modifying account %v", errorMessage[0].Message)
		return err
	}
	return err
}
*/
func RemoveSubscriber(code int) error {
	var errorMessage []ErrorMessage
	deleteResult, err := deleteData("api/v1/subscribers/" + strconv.Itoa(code))
	log.Debug(string(deleteResult))
	errorError := json.Unmarshal(deleteResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem deleting account %v", errorMessage)
		err = fmt.Errorf("Problem deleting account %v", errorMessage[0].Message)
		return err
	}
	return err
}

func NewDevice(mac string, accountcode string, label string) (code int, err error) {
	var acsDevice ACSDevice
	var fetchResult []byte

	fetchResult, err = getData("api/v1/templates/device")
	log.Debug(string(fetchResult))
	err = json.Unmarshal(fetchResult, &acsDevice)
	var errorMessage []ErrorMessage
	errorError := json.Unmarshal(fetchResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem fetching account %v", errorMessage)
		err = fmt.Errorf("Problem fetching account %v", errorMessage[0].Message)
		return
	}

	acsDevice.Accountcode = accountcode
	acsDevice.MAC = mac
	acsDevice.OUI = mac[:6]
	if label != "" {
		acsDevice.Labels = []ACSLabel{
			ACSLabel{
				Name:     label,
				FGColour: "#000",
				BGColour: "#fff",
			},
		}
	}

	var updateResult []byte
	updateResult, err = sendData(http.MethodPost, "api/v1/devices", acsDevice)
	log.Debug(string(updateResult))
	errorError = json.Unmarshal(updateResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem adding device %v", errorMessage)
		err = fmt.Errorf("Problem adding device %v", errorMessage[0].Message)
		return
	}
	err = json.Unmarshal(updateResult, &acsDevice)
	code = acsDevice.DeviceID
	return
}

/*

func ReplaceDevice(mac string, oldid int) (code int, err error) {
	var acsDevice ACSDevice
	var fetchResult []byte

	fetchResult, err = getData("api/v1/devices/" + strconv.Itoa(oldid))
	log.Debug(string(fetchResult))
	err = json.Unmarshal(fetchResult, &acsDevice)
	var errorMessage []ErrorMessage
	errorError := json.Unmarshal(fetchResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem fetching account %v", errorMessage)
		err = fmt.Errorf("Problem fetching account %v", errorMessage[0].Message)
		return
	}

	acsDevice.MAC = mac
	acsDevice.OUI = mac[:6]
	acsDevice.DeviceID = 0
	acsDevice.Disposition = ""

	var updateResult []byte
	updateResult, err = sendData(http.MethodPost, "api/v1/devices", acsDevice)
	log.Debug(string(updateResult))
	errorError = json.Unmarshal(updateResult, &errorMessage)
	if errorError == nil {
		log.Errorf("Problem replacing device %v", errorMessage)
		err = fmt.Errorf("Problem replacing device %v", errorMessage[0].Message)
		return
	}
	err = json.Unmarshal(updateResult, &acsDevice)
	code = acsDevice.DeviceID
	return
}
*/

func RemoveDevice(code int) error {
	var errorMessage []ErrorMessage
	deleteResult, err := deleteData("api/v1/devices/" + strconv.Itoa(code))
	log.Debug(string(deleteResult))
	errorError := json.Unmarshal(deleteResult, &errorMessage)
	if errorError == nil {
		log.Debugf("Problem deleting device %v", errorMessage)
		err = fmt.Errorf("Problem deleting device %v", errorMessage[0].Message)
		return err
	}
	return err
}

func GetDeviceStatus(code int) (status ACSDeviceStatus, err error) {
	var result []byte
	result, err = getData("portal/devices/" + strconv.Itoa(code))
	if err != nil {
		log.Errorf("Problem doing API call for get device status %v", err)
	}
	log.Debug(string(result))
	var deviceResult struct {
		ID          int      `json:"id"`
		Labels      []string `json:"labels"`
		MAC         string   `json:"serialNumber"`
		FirstInform int64    `json:"firstInform"`
		LastInform  int64    `json:"lastInform"`
	}
	err = json.Unmarshal(result, &deviceResult)
	result, err = getData("portal/devices/" + strconv.Itoa(code) + "/attributes")
	if err != nil {
		log.Errorf("Problem doing API call for get device status %v", err)
	}
	log.Debug(string(result))
	var attrResult struct {
		Device struct {
			DeviceInfo struct {
				SoftwareVersion string
			}
			ManagementServer struct {
				ConnectionRequestURL string
			}
		}
	}
	err = json.Unmarshal(result, &attrResult)

	informArr := strings.Split(attrResult.Device.ManagementServer.ConnectionRequestURL, ":")

	status.Firmware = attrResult.Device.DeviceInfo.SoftwareVersion
	if len(informArr) > 1 {
		status.InformURL = informArr[1][2:]
	}
	status.MAC = deviceResult.MAC
	status.FirstInform = time.Unix(0, deviceResult.FirstInform*1000000)
	status.LastInform = time.Unix(0, deviceResult.LastInform*1000000)
	status.SubscriberID = deviceResult.ID

	return
}
