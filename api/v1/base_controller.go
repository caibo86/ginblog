package v1

import (
	"fmt"
	"github.com/caibo86/ginblog/utils"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func RenderResult(c *gin.Context, code error, data interface{}) {
	if code != errmsg.OK {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
			"data":    data,
		})
	}
}

func RenderResultWithTotal(c *gin.Context, code error, data interface{}, total int64) {
	if code != errmsg.OK {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
			"data":    data,
			"total":   total,
		})
	}
}

func RenderError(c *gin.Context, status int, err error, msg ...string) {
	var code errmsg.Code
	switch status {
	case http.StatusBadRequest:
		code = errmsg.ErrBadRequest
	case http.StatusUnauthorized:
		code = errmsg.ErrUnauthorized
	default:
		code = errmsg.ErrServer
	}
	message := code.With(err)
	if len(msg) > 0 {
		for _, m := range msg {
			message += ":" + m
		}
	}
	c.JSON(status, gin.H{
		"status":  code,
		"message": message,
	})
}

func GetPaginate(c *gin.Context) (int, int) {
	perPage, _ := strconv.Atoi(c.Query("per_page"))
	page, _ := strconv.Atoi(c.Query("page"))
	if perPage < 1 {
		perPage = 20
	}
	if page < 1 {
		page = 1
	}
	return perPage, page
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Request.Header.Get("Authorization")
		if a == "" {
			RenderError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		ss := strings.SplitN(a, " ", 2)
		if len(ss) != 2 && ss[0] != "Bearer" {
			RenderError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		claims, code := utils.CheckToken(ss[1])
		if code != errmsg.OK {
			RenderError(c, http.StatusUnauthorized, code)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	logger := logrus.New()
	linkName := "log/latest_log.log"
	//logFile := "log/ginblog.log"
	//file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	//if err != nil {
	//	log.Fatalln("Open log file err:", err)
	//}
	//logger.Out = file
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05.999",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})

	logWriter, _ := rotatelogs.New(
		"log/ginblog_%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05.999",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime).Nanoseconds()
		spendTime := fmt.Sprintf("%f ms", float64(duration)/1000000)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		method := c.Request.Method
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"Hostname":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"IP":        clientIP,
			"Method":    method,
			"Path":      path,
			"Agent":     userAgent,
			"DataSize":  dataSize,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			//AllowAllOrigins: true,
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			//AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		})
	}
}
