package html

import "github.com/gin-gonic/gin"

func IsHtmxReq(c *gin.Context) bool {
    return c.GetHeader("HX-Request") == "true"
}
