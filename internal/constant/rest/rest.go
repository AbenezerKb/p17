package rest

import (
	"github.com/gin-gonic/gin"
)

func ParsePgn(c *gin.Context) (*FilterParams, error) {
	filterParam := QueryParams{}
	err := c.BindQuery(&filterParam)
	if err != nil {

		return nil, err
	}
	param, err := filterParam.Get()
	if err != nil {
		return nil, err
	}
	return param, nil
}

type Response struct {
	MetaData interface{} `json:"meta_data,omitempty"`
	Data     interface{} `json:"data"`
}

type ErrData struct {
	Error interface{} `json:"error"`
}

//ResponseJson creates new json object
func SuccessResponseJson(c *gin.Context, metaData interface{}, responseData interface{}, statusCode int) {
	c.JSON(statusCode, Response{
		MetaData: metaData,
		Data:     responseData,
	})
}
func Contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
