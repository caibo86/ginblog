package middleware

import (
	"fmt"
	"github.com/caibo86/ginblog/api/base"
	"github.com/caibo86/ginblog/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Request.Header.Get("Authorization")
		if a == "" {
			base.RenderError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		ss := strings.SplitN(a, " ", 2)
		if len(ss) != 2 && ss[0] != "Bearer" {
			base.RenderError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		claims, err := utils.CheckToken(ss[1])
		if err != nil {
			base.RenderError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	wlog := logrus.New()
	linkName := "log/latest_log.log"
	//logFile := "log/ginblog.log"
	//file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	//if err != nil {
	//	log.Fatalln("Open log file err:", err)
	//}
	//logger.Out = file
	wlog.SetFormatter(&logrus.TextFormatter{
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
	wlog.AddHook(hook)

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
		//userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		method := c.Request.Method
		path := c.Request.RequestURI
		entry := wlog.WithFields(logrus.Fields{
			"Hostname":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"IP":        clientIP,
			"Method":    method,
			"Path":      path,
			//"Agent":     userAgent,
			"DataSize": dataSize,
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
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

}
