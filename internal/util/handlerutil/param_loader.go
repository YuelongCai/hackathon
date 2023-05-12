package handlerutil

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// LoadStringParams load string array parameter from gin context
func LoadStringParams(c *gin.Context, allowMissing, allowEmpty bool, paramNames ...string) ([]string, error) {
	params := make([]string, 0)
	for _, name := range paramNames {
		param, ok := c.GetQuery(name)
		if !allowMissing && !ok {
			return nil, errors.Errorf("The parameter [%s] is missing", name)
		}
		if !allowEmpty && param == "" {
			return nil, errors.Errorf("The parameter [%s] is empty", name)
		}
		params = append(params, param)
	}
	return params, nil
}

// LoadInt64Params load int64 parameter from gin context
func LoadInt64Params(c *gin.Context, allowMissing, allowEmpty bool, paramNames ...string) ([]int64, error) {
	params := make([]int64, 0)
	strParams, err := LoadStringParams(c, allowMissing, allowEmpty, paramNames...)
	if err != nil {
		return nil, err
	}
	for i, strParam := range strParams {
		var param int64
		if (allowMissing || allowEmpty) && strParam == "" {
			param = 0
		} else {
			param, err = strconv.ParseInt(strParam, 10, 64)
			if err != nil {
				return nil, errors.Errorf("The parameter [%s] is not an integer", paramNames[i])
			}
		}
		params = append(params, param)
	}
	return params, nil
}

// LoadStringArrayParams load string array parameter from gin context
func LoadStringArrayParams(c *gin.Context, allowMissing, allowEmpty bool, paramNames ...string) ([][]string, error) {
	params := make([][]string, 0)
	strParams, err := LoadStringParams(c, allowMissing, allowEmpty, paramNames...)
	if err != nil {
		return nil, err
	}
	for _, strParam := range strParams {
		arrayParam := strings.Split(strParam, ",")
		params = append(params, arrayParam)
	}
	return params, nil
}

// LoadIntParams load int parameter from gin context
func LoadIntParams(c *gin.Context, allowMissing, allowEmpty bool, paramNames ...string) ([]int, error) {
	params := make([]int, 0)
	strParams, err := LoadStringParams(c, allowMissing, allowEmpty, paramNames...)
	if err != nil {
		return nil, err
	}
	for i, strParam := range strParams {
		var param int
		if (allowMissing || allowEmpty) && strParam == "" {
			param = 0
		} else {
			param, err = strconv.Atoi(strParam)
			if err != nil {
				return nil, errors.Errorf("The parameter [%s] is not an integer", paramNames[i])
			}
		}
		params = append(params, param)
	}
	return params, nil
}
