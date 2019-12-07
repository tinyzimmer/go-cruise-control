package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/tinyzimmer/go-cruise-control/pkg/types"
)

// ServerError implements error and represents a server-side exception from
// CruiseControl.
type ServerError struct {
	ErrorMessage string `json:"errorMessage"`
	StackTrace   string `json:"stackTrace"`
	Version      int    `json:"version"`
}

// Error returns the main error message from a ServerError.
func (s *ServerError) Error() string {
	return s.ErrorMessage
}

// basicRequest performs a request with the given method to the given endpoint.
// Opts are optionally validated and then transformed to URL arguments. The
// out interface (or error) will hold the server response.
func (c *Client) basicRequest(method, endpoint string, opts interface{}, out interface{}) error {
	if validator, ok := opts.(interface{ Validate() error }); ok {
		if err := validator.Validate(); err != nil {
			return err
		}
	}
	u := appendURLOpts(c.newURL(endpoint), opts)
	return c.do(method, u, out)
}

// newURL returns the full URL to the given endpoint.
func (c *Client) newURL(endpoint string) *url.URL {
	raw := fmt.Sprintf("%s/%s", c.baseURL.String(), endpoint)
	url, _ := url.Parse(raw)
	return url
}

// do handles the actual request to the server. If the response is non-200,
// it attempts to unmarshal the response into a ServerError. Otherwise, the
// supplied out interface is used.
func (c *Client) do(method string, u *url.URL, out interface{}) error {
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}
	if c.credentials != nil {
		req.SetBasicAuth(c.credentials.Username, c.credentials.Password)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		serr := &ServerError{}
		if err := json.Unmarshal(body, serr); err != nil {
			return fmt.Errorf("%s returned status: %d %s\n%s", u.String(), res.StatusCode, res.Status, string(body))
		}
		return serr
	}
	return json.Unmarshal(body, out)
}

// appendURLOpts transforms the given struct into URL parameters and returns
// a complete URL. It uses reflection and the `param` struct field to determine
// the correct arguments.
func appendURLOpts(u *url.URL, opts interface{}) *url.URL {
	q := u.Query()
	q.Set("json", "true")

	if opts == nil {
		u.RawQuery = q.Encode()
		return u
	}

	params := reflect.TypeOf(opts).Elem()
	vals := reflect.ValueOf(opts).Elem()
	for i := 0; i < params.NumField(); i++ {
		field := params.Field(i)
		param := field.Tag.Get("param")
		if val := getParamStrValue(vals.Field(i).Interface()); val != "" {
			q.Set(param, val)
		}
	}
	u.RawQuery = q.Encode()
	return u
}

// getParamStrValue will determine the type of the provided value, and return
// the URL query string equivalent.
func getParamStrValue(in interface{}) string {
	switch reflect.TypeOf(in).Name() {

	case "string":
		val := in.(string)
		return val

	case "bool":
		val := in.(bool)
		if val {
			return "true"
		}
		return "false"

	case "Time":
		val := in.(time.Time)
		if val.Year() == 0001 {
			// the time object is actually empty/nil
			return ""
		}
		return fmt.Sprintf("%d", val.UnixNano()/1000000)

	case "int":
		val := in.(int)
		if val == 0 {
			return ""
		}
		return fmt.Sprintf("%d", val)

	case "float64":
		val := in.(float64)
		if val == 0 {
			return ""
		}
		return fmt.Sprintf("%f", val)

	case "LoadType":
		val := in.(types.LoadType)
		return string(val)

	case "DataSource":
		val := in.(types.DataSource)
		return string(val)

	case "ReplicaMovementStrategy":
		val := in.(types.ReplicaMovementStrategy)
		return string(val)
	}

	if reflect.TypeOf(in).Kind() == reflect.Slice {
		switch reflect.TypeOf(in).Elem().Name() {
		case "string":
			val := in.([]string)
			if len(val) == 0 {
				return ""
			}
			return strings.Join(val, ",")

		case "int":
			val := in.([]int)
			out := make([]string, 0)
			for _, x := range val {
				out = append(out, fmt.Sprintf("%d", x))
			}
			return strings.Join(out, ",")

		case "Substate":
			val := in.([]types.Substate)
			if len(val) == 0 {
				return ""
			}
			out := make([]string, 0)
			for _, x := range val {
				out = append(out, string(x))
			}
			return strings.Join(out, ",")
		}

		panic(fmt.Sprintf("Unsupported param type - please add handler for: []%s", reflect.TypeOf(in).Elem().Name()))
	}

	panic(fmt.Sprintf("Unsupported param type - please add handler for: %s", reflect.TypeOf(in).Name()))
}
