package apivideosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Client type handles communicating with the api.video API
type Client struct {
	BaseURL            *url.URL
	APIKey             string
	httpClient         Doer
	chunkSize          int64
	Token              *Token
	applicationName    string
	applicationVersion string
	sdkName            string
	sdkVersion         string

	AdvancedAuthentication AdvancedAuthenticationServiceI
	Analytics              AnalyticsServiceI
	Captions               CaptionsServiceI
	Chapters               ChaptersServiceI
	LiveStreams            LiveStreamsServiceI
	PlayerThemes           PlayerThemesServiceI
	Summaries              SummariesServiceI
	Tags                   TagsServiceI
	UploadTokens           UploadTokensServiceI
	Videos                 VideosServiceI
	Watermarks             WatermarksServiceI
	Webhooks               WebhooksServiceI
}

// Token contains token for connecting to the api.video API
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	ExpireTime   time.Time
}

// ErrorResponse contains an error from the api.video API
type ErrorResponse struct {
	Response *http.Response
	Type     string `json:"type"`
	Title    string `json:"title"`
	Name     string `json:"name"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf(
		"[%d]: %v %v\nType: %v\nTitle: %v\nName: %v",
		r.Response.StatusCode,
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Type,
		r.Title,
		r.Name,
	)
}

const (
	defaultBaseURL        = "https://ws.api.video/"
	defaultSandboxBaseURL = "https://sandbox.api.video/"
	defaultChunkSize      = 50 * 1024 * 1024
	minChunkSize          = 5 * 1024 * 1024
	maxChunkSize          = 128 * 1024 * 1024
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Builder struct {
	apiKey             string
	baseURL            string
	uploadChunkSize    int64
	httpClient         Doer
	applicationName    string
	applicationVersion string
	sdkName            string
	sdkVersion         string
}

func (cb *Builder) BaseURL(url string) *Builder {
	cb.baseURL = url
	return cb
}

func (cb *Builder) APIKey(key string) *Builder {
	cb.apiKey = key
	return cb
}

func (cb *Builder) ApplicationName(applicationName string) *Builder {
	cb.applicationName = applicationName
	return cb
}

func (cb *Builder) ApplicationVersion(applicationVersion string) *Builder {
	cb.applicationVersion = applicationVersion
	return cb
}

func (cb *Builder) SdkName(sdkName string) *Builder {
	cb.sdkName = sdkName
	return cb
}

func (cb *Builder) SdkVersion(sdkVersion string) *Builder {
	cb.sdkVersion = sdkVersion
	return cb
}

func (cb *Builder) UploadChunkSize(size int64) *Builder {
	if size < minChunkSize {
		cb.uploadChunkSize = minChunkSize
	} else if size > maxChunkSize {
		cb.uploadChunkSize = maxChunkSize
	} else {
		cb.uploadChunkSize = size
	}
	return cb
}

func (cb *Builder) HTTPClient(httpClient Doer) *Builder {
	cb.httpClient = httpClient
	return cb
}

// ClientBuilder returns a new api.video API client builder for production
func ClientBuilder(apiKey string) *Builder {
	return &Builder{
		baseURL:         defaultBaseURL,
		uploadChunkSize: defaultChunkSize,
		apiKey:          apiKey,
	}
}

// SandboxClientBuilder returns a new api.video API client builder for sandbox environment
func SandboxClientBuilder(apiKey string) *Builder {
	return ClientBuilder(apiKey).BaseURL(defaultSandboxBaseURL)
}

func (cb *Builder) Build() *Client {

	baseURL, _ := url.Parse(cb.baseURL)

	var httpClient Doer

	if cb.httpClient == nil {
		httpClient = http.DefaultClient
	} else {
		httpClient = cb.httpClient
	}

	c := &Client{
		BaseURL:            baseURL,
		APIKey:             cb.apiKey,
		httpClient:         httpClient,
		chunkSize:          cb.uploadChunkSize,
		applicationName:    cb.applicationName,
		applicationVersion: cb.applicationVersion,
		sdkName:            cb.sdkName,
		sdkVersion:         cb.sdkVersion,
	}

	c.AdvancedAuthentication = &AdvancedAuthenticationService{client: c}
	c.Analytics = &AnalyticsService{client: c}
	c.Captions = &CaptionsService{client: c}
	c.Chapters = &ChaptersService{client: c}
	c.LiveStreams = &LiveStreamsService{client: c}
	c.PlayerThemes = &PlayerThemesService{client: c}
	c.Summaries = &SummariesService{client: c}
	c.Tags = &TagsService{client: c}
	c.UploadTokens = &UploadTokensService{client: c}
	c.Videos = &VideosService{client: c}
	c.Watermarks = &WatermarksService{client: c}
	c.Webhooks = &WebhooksService{client: c}

	return c
}

// ChunkSize changes chunk size for video upload, by default its 128MB
func (c *Client) ChunkSize(size int64) {
	c.chunkSize = size
}

func (c *Client) prepareRequest(
	ctx context.Context,
	method,
	urlStr string,
	body interface{},
	headerParams map[string]string,
	queryParams url.Values) (*http.Request, error) {

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	u.RawQuery = query.Encode()

	buf := new(bytes.Buffer)
	if body != nil {

		switch v := body.(type) {
		case *bytes.Buffer:
			buf = v
		default:
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	originAppHeaderValue, err := getOriginHeaderValue(c.applicationName, c.applicationVersion)
	if err != nil {
		return nil, err
	}
	if originAppHeaderValue != "" {
		req.Header.Set("AV-Origin-App", originAppHeaderValue)
	}

	originSdkHeaderValue, err := getOriginHeaderValue(c.sdkName, c.sdkVersion)
	if err != nil {
		return nil, err
	}
	if originSdkHeaderValue != "" {
		req.Header.Set("AV-Origin-Sdk", originSdkHeaderValue)
	}

	req.Header.Set("AV-Origin-Client", "go:1.4.7")

	for headerName := range headerParams {
		req.Header.Set(headerName, headerParams[headerName])
	}

	if c.APIKey != "" {
		req, err = c.auth(req)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}

func getOriginHeaderValue(name string, version string) (string, error) {
	if name == "" && version == "" {
		return "", nil
	}
	var re = regexp.MustCompile(`(?m)^[\w-]{1,50}$`)
	if !re.MatchString(name) {
		return "", errors.New("Invalid name value. Allowed characters: A-Z, a-z, 0-9, '-', '_'. Max length: 50.")
	}

	var reVersion = regexp.MustCompile(`(?m)^\d{1,3}(\.\d{1,3}(\.\d{1,3})?)?$`)
	if !reVersion.MatchString(version) {
		return "", errors.New("Invalid version value. The version should match the xxx[.yyy][.zzz] pattern.")
	}

	return name + ":" + version, nil
}

func (c *Client) prepareProgressiveUploadRequest(
	ctx context.Context,
	urlStr string,
	fileName string,
	fileReader io.Reader,
	fileSize int64,
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	part int32,
	isLast bool) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	partWriter, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	io.Copy(partWriter, fileReader)
	if err != nil {
		return nil, err
	}

	for key, val := range formParams {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	var ranges string
	if isLast {
		ranges = fmt.Sprintf("part %d/%d", part, part)
	} else {
		ranges = fmt.Sprintf("part %d/*", part)
	}

	req.Header.Set("Content-Range", ranges)

	return req, nil
}

func (c *Client) prepareRangeRequests(
	ctx context.Context,
	urlStr string,
	fileName string,
	fileReader io.Reader,
	fileSize int64,
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string) ([]*http.Request, error) {

	var bufSize int64
	if fileSize > c.chunkSize && c.chunkSize != 0 {
		bufSize = c.chunkSize
	} else {
		bufSize = fileSize
	}

	var requests []*http.Request
	startByte := int64(0)
	part := 1
	partCount := int64(math.Ceil(float64(fileSize) / float64(c.chunkSize)))
	for {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		partWriter, err := writer.CreateFormFile("file", fileName)
		if err != nil {
			return nil, err
		}

		for key, val := range formParams {
			err = writer.WriteField(key, val)
			if err != nil {
				return nil, err
			}
		}

		var bytesToRead int64
		if startByte+bufSize > fileSize {
			bytesToRead = fileSize - startByte
		} else {
			bytesToRead = bufSize
		}

		bytesread, err := io.CopyN(partWriter, fileReader, bytesToRead)
		if err != nil {
			return nil, err
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}

		req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		if fileSize > c.chunkSize && c.chunkSize != 0 {
			ranges := fmt.Sprintf("part %d/%d", part, partCount)
			req.Header.Set("Content-Range", ranges)
			startByte = startByte + bytesread
		}

		requests = append(requests, req)

		if startByte == fileSize || fileSize < c.chunkSize {
			break
		}

		part = part + 1
	}
	return requests, nil
}

func (c *Client) prepareUploadRequest(
	ctx context.Context,
	urlStr string,
	fileName string,
	fileReader io.Reader,
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	partWriter, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(partWriter, fileReader)
	if err != nil {
		return nil, err
	}

	for key, val := range formParams {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (c *Client) auth(req *http.Request) (*http.Request, error) {

	if c.Token == nil || time.Now().After(c.Token.ExpireTime) {
		u, err := c.BaseURL.Parse("/auth/api-key")
		if err != nil {
			return nil, err
		}

		payload := map[string]string{"apiKey": c.APIKey}

		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)

		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", u.String(), buf)

		if err != nil {
			return nil, err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		originAppHeaderValue, err := getOriginHeaderValue(c.applicationName, c.applicationVersion)
		if err != nil {
			return nil, err
		}
		if originAppHeaderValue != "" {
			req.Header.Set("AV-Origin-App", originAppHeaderValue)
		}
		req.Header.Set("AV-Origin-Client", "go:1.2.4")

		originSdkHeaderValue, err := getOriginHeaderValue(c.sdkName, c.sdkVersion)
		if err != nil {
			return nil, err
		}
		if originSdkHeaderValue != "" {
			req.Header.Set("AV-Origin-Sdk", originSdkHeaderValue)
		}

		req.Header.Set("AV-Origin-Client", "go:1.4.7")

		resp, err := c.httpClient.Do(req)

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		err = checkResponse(resp)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(resp.Body).Decode(&c.Token)
		if err != nil {
			return nil, err
		}

		c.Token.ExpireTime = time.Now().Add(time.Duration(c.Token.ExpiresIn) * time.Second)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)
	return req, nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Title = string(data)
		}
	}

	return errorResponse
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

func addDeepQueryParams(filter interface{}, prefix string, queryParams url.Values) {
	v := reflect.ValueOf(filter)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)

			if field.Kind() == reflect.Ptr && !field.IsNil() {
				field = field.Elem()
			}

			lowercaseFirstChar := strings.ToLower(string(fieldType.Name[0]))
			restOfString := fieldType.Name[1:]
			lowercaseName := lowercaseFirstChar + restOfString

			if field.Kind() == reflect.Slice {
				for j := 0; j < field.Len(); j++ {
					item := field.Index(j)
					queryParams.Add(fmt.Sprintf("%s[%s][%d]", prefix, lowercaseName, j), item.String())
				}
			} else if field.Kind() == reflect.String {
				queryParams.Add(fmt.Sprintf("%s[%s]", prefix, lowercaseName), field.String())
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			queryParams.Add(fmt.Sprintf("%s[%s]", prefix, key.String()), fmt.Sprintf("%v", val))
		}
	default:
		fmt.Println("Unsupported type")
	}
}
