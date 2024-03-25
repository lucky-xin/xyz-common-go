package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type GinLogger struct {
	Log *logrus.Logger
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r bodyLogWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r bodyLogWriter) WriteString(s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

var Logger = New()

func New() *GinLogger {
	var ginLog = logrus.New()

	ginLog.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式.
	// 同样地,也可以单独为某个logrus实例设置日志级别和hook,这里不详细叙述.
	ginLog.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	ginLog.Level = logrus.InfoLevel

	return &GinLogger{ginLog}
}

// AccessRecordLog 获取返回体的中间件
func AccessRecordLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		//开始时间
		startTime := time.Now()

		c.Next()

		//结束时间
		endTime := time.Now()
		//执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		//请求ip
		clientIP := c.ClientIP()
		//请求参数

		reqParams, _ := RequestInputs(c)

		//请求ua
		reqUa := c.Request.UserAgent()
		var resultBody logrus.Fields
		resultBody = make(map[string]interface{})
		uid := c.GetInt64("uid")
		uname := c.GetString("uname")
		if uid != 0 {
			resultBody["user_id"] = uid
		}
		if uname != "" {
			resultBody["user_account"] = uname
		}

		reqId := c.GetHeader("Req-Id")
		resultBody["req_uri"] = reqUri
		resultBody["http_method"] = reqMethod
		resultBody["client_ip"] = clientIP
		resultBody["req_params"] = reqParams
		resultBody["took"] = latencyTime.Milliseconds()
		resultBody["user_agent"] = reqUa

		var resp map[string]interface{}
		err := json.Unmarshal(blw.body.Bytes(), &resp)
		if err != nil {
			if reqId == "" {
				reqId = resp["reqId"].(string)
			}
			resultBody["biz_code"] = resp["code"]
			resultBody["biz_resp"] = resp
			return
		}
		if reqId != "" {
			resultBody["req_id"] = reqId
		}
		Logger.Log.WithFields(resultBody).Info()
	}
}

// RequestInputs 获取所有参数
func RequestInputs(c *gin.Context) (map[string]interface{}, error) {

	const defaultMemory = 32 << 20
	contentType := c.ContentType()

	var (
		dataMap  = make(map[string]interface{})
		queryMap = make(map[string]interface{})
		postMap  = make(map[string]interface{})
	)

	// @see gin@v1.7.7/binding/query.go ==> func (queryBinding) Bind(req *http.Request, obj interface{})
	for k := range c.Request.URL.Query() {
		queryMap[k] = c.Query(k)
	}

	if "application/json" == contentType {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		// @see gin@v1.7.7/binding/json.go ==> func (jsonBinding) Bind(req *http.Request, obj interface{})
		if c.Request != nil && c.Request.Body != nil {
			if err := json.NewDecoder(c.Request.Body).Decode(&postMap); err != nil {
				return nil, err
			}
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	} else if "multipart/form-data" == contentType {
		// @see gin@v1.7.7/binding/form.go ==> func (formMultipartBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	} else {
		// ParseForm 解析 URL 中的查询字符串，并将解析结果更新到 r.Form 字段
		// 对于 POST 或 PUT 请求，ParseForm 还会将 body 当作表单解析，
		// 并将结果既更新到 r.PostForm 也更新到 r.Form。解析结果中，
		// POST 或 PUT 请求主体要优先于 URL 查询字符串（同名变量，主体的值在查询字符串的值前面）
		// @see gin@v1.7.7/binding/form.go ==> func (formBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseForm(); err != nil {
			return nil, err
		}
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			if err != http.ErrNotMultipart {
				return nil, err
			}
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	}

	var mu sync.RWMutex
	for k, v := range queryMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}
	for k, v := range postMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}

	return dataMap, nil
}

func (logger *GinLogger) Info(c *gin.Context, args ...interface{}) {
	logger.Log.WithFields(logrus.Fields{
		"username": c.GetString("username"),
		"user_id":  c.GetInt64("userId"),
	}).Info(args)
}

func (logger *GinLogger) Error(c *gin.Context, args ...interface{}) {
	logger.Log.WithFields(logrus.Fields{
		"username": c.GetString("username"),
		"user_id":  c.GetInt64("userId"),
	}).Error(args)
}
