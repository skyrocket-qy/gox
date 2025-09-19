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
	"github.com/skyrocket-qy/erx"
)

type ErrBinder struct {
	ErrToHTTP map[erx.Code]int // Exported
}

func NewErrBinder(errToHTTP map[erx.Code]int) *ErrBinder {
	return &ErrBinder{ErrToHTTP: errToHTTP} // Updated to use ErrToHTTP
}

type ErrResp struct {
	ReqId string `json:"reqId"`
	Code  string `json:"code"`
}

// for some error, log as error, some log as debug.
func (b *ErrBinder) Bind(c *gin.Context, err error) {
	reqId := c.GetString("reqId")

	var ctxErr *erx.CtxErr
	if !errors.As(err, &ctxErr) {
		c.JSON(http.StatusInternalServerError, ErrResp{ReqId: reqId, Code: erx.ErrUnknown.Str()})

		callerInfos := GetCallStack(2) // Updated to use GetCallStack
		log.Error().Err(err).Str("call",
			fmt.Sprintf("%+v", callerInfos)).Msg("error not wrapped by erx")

		return
	}

	e := log.Error()

	underlyingErr := ctxErr.Unwrap()
	if underlyingErr != nil && underlyingErr.Error() != "" {
		e.Str("cause", underlyingErr.Error())
	}

	e.Str("code", ctxErr.Code.Str())

	// Convert callerInfos to pretty strings
	filtered := FilterCallerInfos(ctxErr.CallerInfos) // Updated to use FilterCallerInfos

	trace := make([]string, 0, len(filtered))
	for _, ci := range filtered {
		trace = append(trace, fmt.Sprintf("%s %d %s",
			TrimToProject(ci.File), // Updated to use TrimToProject
			ci.Line,
			ExtractFuncName(ci.Function), // Updated to use ExtractFuncName
		))
	}

	e.Strs("callerTrace", trace)
	e.Msg("error")

	httpStatus, ok := b.ErrToHTTP[ctxErr.Code] // Updated to use ErrToHTTP
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	c.JSON(httpStatus, ErrResp{ReqId: reqId, Code: ctxErr.Code.Str()})
}

func TrimToProject(path string) string { // Exported
	projectRoot, _ := os.Getwd()
	rel, _ := strings.CutPrefix(path, projectRoot)

	return rel
}

func ExtractFuncName(fullFunc string) string { // Exported
	// e.g., input: srv/internal/logic/inter.(*Logic).Login
	// output: (*Logic).Login
	if idx := strings.LastIndex(fullFunc, "/"); idx >= 0 {
		return fullFunc[idx+1:]
	}

	return fullFunc
}

func FilterCallerInfos(infos []erx.CallerInfo) []erx.CallerInfo { // Exported
	projectPrefix, _ := os.Getwd()

	var filtered []erx.CallerInfo

	for _, ci := range infos {
		if strings.HasPrefix(ci.File, projectPrefix) {
			filtered = append(filtered, ci)
		} else {
			break
		}
	}

	if filtered == nil {
		return []erx.CallerInfo{}
	}

	return filtered
}

type CallerInfo struct {
	Function string
	File     string
	Line     int
}

func GetCallStack(callerSkip ...int) (callerInfos []CallerInfo) { // Exported
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
