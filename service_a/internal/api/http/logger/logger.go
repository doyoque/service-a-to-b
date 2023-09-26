package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/doyoque/service_a/conf"
	"github.com/doyoque/service_a/internal/api/http/middleware/auth"
	chimware "github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type ResponseWrapWriter struct {
	writer     http.ResponseWriter
	body       bytes.Buffer
	statusCode int
	startTime  time.Time
}

func (i *ResponseWrapWriter) Flush() {
	i.writer.(http.Flusher).Flush()
}

func (w *ResponseWrapWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.writer.(http.Hijacker).Hijack()
}

func (i *ResponseWrapWriter) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	i.writer.WriteHeader(statusCode)
}

func (i *ResponseWrapWriter) Header() http.Header {
	return i.writer.Header()
}

func (i *ResponseWrapWriter) Write(buffer []byte) (int, error) {
	i.body.Write(buffer)
	return i.writer.Write(buffer)
}

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		responseWriterWrapper := &ResponseWrapWriter{
			writer:    w,
			startTime: conf.Now(),
		}

		defer logResponse(responseWriterWrapper, r)

		next.ServeHTTP(responseWriterWrapper, r)
	})
}

func logRequest(r *http.Request) {
	reqId := chimware.GetReqID(r.Context())
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	logrus.Infof("[REQ_LOG] [%s] %s - %s\n%s %s",
		reqId,
		r.Method,
		r.URL.Path,
		getPrettyHeader(r.Header),
		string(bodyBytes),
	)
}

func logResponse(w *ResponseWrapWriter, r *http.Request) {
	ctx := r.Context()
	reqId := chimware.GetReqID(ctx)
	authContext := auth.GetUserAuthorizationContext(ctx)

	authLogStr := ""
	if authContext != nil {
		authLogStr = fmt.Sprintf("[TYPE:%s] [USER_ID:%d]",
			authContext.Type,
			authContext.UserId,
		)
	}

	logLevelFn := levelFn(w.statusCode)

	logLevelFn("[RES_LOG] [%s] %s - %s %d %s %s %s\n%s %s",
		reqId,
		r.Method,
		r.URL.Path,
		w.statusCode,
		http.StatusText(w.statusCode),
		time.Since(w.startTime).String(),
		authLogStr,
		getPrettyHeader(w.Header()),
		w.body.String(),
	)
}

func levelFn(statusCode int) func(format string, args ...interface{}) {
	switch {
	case statusCode >= 500:
		return logrus.Errorf
	case statusCode >= 400:
		return logrus.Warnf
	default:
		return logrus.Infof
	}
}

func getPrettyHeader(header http.Header) string {
	headerStr := ""
	for k, v := range header {
		k = http.CanonicalHeaderKey(k)
		if k == "Authorization" || k == "X-Authorization" {
			v = []string{"***"}
		}
		headerStr += fmt.Sprintf("%s: %v\n", k, strings.Join(v, ", "))
	}

	return headerStr
}
