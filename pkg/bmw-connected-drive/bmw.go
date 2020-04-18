package bmw

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	helper "github.com/andig/evcc/pkg/http"
)

const (
	auth       = "https://customer.bmwgroup.com/gcdm/oauth/authenticate"
	apiMe      = "https://www.bmw-connecteddrive.com/api/me"
	apiVehicle = "https://www.bmw-connecteddrive.com/api/vehicle"
)

// BMW is an api.Vehicle implementation for BMW cars
type BMW struct {
	*helper.Helper
	user, password string
	token          string
	tokenValid     time.Time
}

// New creates a new vehicle
func New(user, password string) *BMW {
	return &BMW{
		Helper:   helper.NewHelper(nil),
		user:     user,
		password: password,
	}
}

func (v *BMW) Logger(log *log.Logger) {
	v.Helper.Log = log
}

func (v *BMW) Login() (string, error) {
	data := url.Values{
		"username":      []string{v.user},
		"password":      []string{v.password},
		"client_id":     []string{"dbf0a542-ebd1-4ff0-a9a7-55172fbfce35"},
		"redirect_uri":  []string{"https://www.bmw-connecteddrive.com/app/default/static/external-dispatch.html"},
		"response_type": []string{"token"},
		"scope":         []string{"authenticate_user fupo"},
		"state":         []string{"eyJtYXJrZXQiOiJkZSIsImxhbmd1YWdlIjoiZGUiLCJkZXN0aW5hdGlvbiI6ImxhbmRpbmdQYWdlIn0"},
		"locale":        []string{"DE-de"},
	}

	req, err := http.NewRequest(http.MethodPost, auth, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }, // don't follow redirects
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	query, err := url.ParseQuery(resp.Header.Get("Location"))
	if err != nil {
		return "", err
	}

	token := query.Get("access_token")
	expires, err := strconv.Atoi(query.Get("expires_in"))
	if err != nil || token == "" || expires == 0 {
		return token, errors.New("could not obtain token")
	}

	v.token = token
	v.tokenValid = time.Now().Add(time.Duration(expires) * time.Second)

	return token, nil
}

func (v *BMW) request(uri string) (*http.Request, error) {
	if v.token == "" || time.Since(v.tokenValid) > 0 {
		if _, err := v.Login(); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err == nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", v.token))
	}

	return req, nil
}

// Vehicles api
func (v *BMW) Vehicles() ([]Vehicle, error) {
	uri := fmt.Sprintf("%s/vehicles/v2", apiMe)
	req, err := v.request(uri)
	if err != nil {
		return []Vehicle{}, err
	}

	var res []Vehicle
	_, err = v.RequestJSON(req, &res)

	return res, err
}

// Dynamic api
func (v *BMW) Dynamic(vin string) (DynamicAttributes, error) {
	uri := fmt.Sprintf("%s/dynamic/v1/%s", apiVehicle, vin)
	req, err := v.request(uri)
	if err != nil {
		return DynamicAttributes{}, err
	}

	var res DynamicResponse
	_, err = v.RequestJSON(req, &res)

	return res.AttributesMap, err
}
