package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/gox/erx"
)

type ErrBinder struct {
	errToHTTP map[erx.Code]int
}

func NewErrBinder(errToHTTP map[erx.Code]int) *ErrBinder {
	return &ErrBinder{errToHTTP: errToHTTP}
}

type ErrResp struct {
	ReqId string `json:"reqId"`
	Code  string `json:"code"`
}

// for some error, log as error, some log as debug
func (b *ErrBinder) Bind(c *gin.Context, err error) {
	reqId := c.GetString("reqId")
	var ctxErr *erx.CtxErr
	if !errors.As(err, &ctxErr) {
		c.JSON(http.StatusInternalServerError, ErrResp{ReqId: reqId, Code: erx.ErrUnknown.Str()})

		callerInfos := getCallStack(2)
		log.Error().Err(err).Str("call", fmt.Sprintf("%+v", callerInfos)).Msg("error not wrapped by erx")
		return
	}

	e := log.Error()
	if ctxErr.Cause != "" {
		e.Str("cause", ctxErr.Cause)
	}
	e.Str("code", ctxErr.Code.Str())

	// Convert callerInfos to pretty strings
	filtered := filterCallerInfos(ctxErr.CallerInfos)
	trace := make([]string, 0, len(filtered))
	for _, ci := range filtered {
		trace = append(trace, fmt.Sprintf("%s %d %s",
			trimToProject(ci.File),
			ci.Line,
			extractFuncName(ci.Function),
		))
	}
	e.Strs("callerTrace", trace)
	e.Msg("error")

	httpStatus, ok := b.errToHTTP[ctxErr.Code]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	c.JSON(httpStatus, ErrResp{ReqId: reqId, Code: ctxErr.Code.Str()})

}

func trimToProject(path string) string {
	projectRoot, _ := os.Getwd()
	rel, _ := strings.CutPrefix(path, projectRoot)
	return rel
}

func extractFuncName(fullFunc string) string {
	// e.g., input: srv/internal/logic/inter.(*Logic).Login
	// output: (*Logic).Login
	if idx := strings.LastIndex(fullFunc, "/"); idx >= 0 {
		return fullFunc[idx+1:]
	}
	return fullFunc
}

func filterCallerInfos(infos []erx.CallerInfo) []erx.CallerInfo {
	projectPrefix, _ := os.Getwd()
	var filtered []erx.CallerInfo
	for _, ci := range infos {
		if strings.HasPrefix(ci.File, projectPrefix) {
			filtered = append(filtered, ci)
		} else {
			break
		}
	}
	return filtered
}

type CallerInfo struct {
	Function string
	File     string
	Line     int
}

func getCallStack(callerSkip ...int) (callerInfos []CallerInfo) {
	pc := make([]uintptr, 5)

	skip := 2
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}

	n := runtime.Callers(skip, pc)

	frames := runtime.CallersFrames(pc[:n])

	for {
		frame, more := frames.Next()
		callerInfos = append(callerInfos, CallerInfo{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		})

		if !more {
			break
		}
	}

	return callerInfos
}
