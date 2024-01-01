package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PanicJson struct {
	Code int
	Msg  interface{}
	Data interface{}
}

func (pj PanicJson) Handle(c *gin.Context) {
	defer func() {
		rc := recover()
		if rc != nil {
			//log.Printf("panic: %v\n", r)
			//debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			switch rcv := rc.(type) {
			case error:
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  rcv.Error(),
				})
			case PanicJson:
				c.JSON(http.StatusOK, gin.H{
					"code": rcv.Code,
					"msg":  rcv.Msg,
					"data": rcv.Data,
				})
			default:
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  rc.(string),
				})
			}
			c.Abort()
		}
	}()
	c.Next()
}
