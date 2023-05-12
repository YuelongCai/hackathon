package handlerutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var validSortKeyMap = map[string]string{
	"updated_at": "updated_at",
	"created_at": "created_at",
}

var validSortDirectionMap = map[string]string{
	"ASC":  "ASC",
	"DESC": "DESC",
}

// GetOffsetAndLimit from context page_no and per_page
func GetOffsetAndLimit(c *gin.Context) (int, int) {
	paramsInt, _ := LoadIntParams(c, true, true, "page_no", "per_page")
	page := paramsInt[0]
	pageSize := paramsInt[1]
	if pageSize == 0 {
		pageSize = -1
	}
	if page <= 0 {
		page = 1
	}
	limit := pageSize
	offset := (page - 1) * pageSize
	return offset, limit
}

// BuildQueryClause build query clause
func BuildQueryClause(c *gin.Context, table string) (string, map[string]interface{}, string, int, int, error) {
	// offset & limit
	offset, limit := GetOffsetAndLimit(c)

	// order clause
	sortBy := c.Query("sort_by")
	sortDirection := c.Query("sort_direction")
	sortBy, err := parseValidKey(sortBy, "updated_at", "sort key", validSortKeyMap)
	if err != nil {
		return "", nil, "", 0, 0, err
	}
	sortBy = table + "." + sortBy
	sortDirection, err = parseValidKey(sortDirection, "DESC", "sort direction", validSortDirectionMap)
	if err != nil {
		return "", nil, "", 0, 0, err
	}
	orderClause := sortBy + " " + sortDirection

	// filter clause
	filterClause, paramMap := buildFilterClause(c, table)

	return filterClause, paramMap, orderClause, offset, limit, nil
}

// parseValidKey return err if not valid
func parseValidKey(key, defaultValue, fieldName string, validMap map[string]string) (string, error) {
	if key == "" {
		return defaultValue, nil
	}
	key = strings.ToLower(key)
	if _, ok := validMap[key]; !ok {
		validKeys := make([]string, 0, len(validMap))
		for k := range validMap {
			validKeys = append(validKeys, k)
		}
		return "", fmt.Errorf("[%s] should be in {%s}", fieldName, strings.Join(validKeys, ", "))
	}
	return validMap[key], nil
}

// buildFilterClause build filter clause
func buildFilterClause(c *gin.Context, table string) (string, map[string]interface{}) {
	var filterClauses []string
	paramMap := make(map[string]interface{})
	idx := 0
	likeParamsMap := map[string]string{"name": "name", "description": "description", "category": "category", "rarity": "rarity", "status": "status"}
	processLikeOrFilters(c, &filterClauses, paramMap, &idx, likeParamsMap, table)
	equalParamsMap := map[string]string{}
	processCompareOrFilters(c, "=", &filterClauses, paramMap, &idx, equalParamsMap, table)
	filterClauseStr := strings.Join(filterClauses, " and ")
	return filterClauseStr, paramMap
}

// processLikeOrFilters process "like or" type filter, e.g. status=RUNNING,INIT => status like '%RUNNING%' or status like '%INIT%'
func processLikeOrFilters(c *gin.Context, filterClauses *[]string, paramMap map[string]interface{}, idx *int, likeAndFilters map[string]string, table string) {
	for likeParam, likeColumn := range likeAndFilters {
		paramValue, ok := c.GetQuery(likeParam)
		if !ok {
			continue
		}
		if s := buildLikeOrFilterClause(paramMap, paramValue, table+"."+likeColumn, idx); s != "" {
			*filterClauses = append(*filterClauses, s)
		}
	}
}

// processCompareOrFilters process "compare or" type filter, e.g. status=RUNNING,INIT => status = 'RUNNING' or status = 'INIT'
func processCompareOrFilters(c *gin.Context, sign string, filterClauses *[]string, paramMap map[string]interface{}, idx *int, likeAndFilters map[string]string, table string) {
	for likeParam, likeColumn := range likeAndFilters {
		paramValue, ok := c.GetQuery(likeParam)
		if !ok {
			continue
		}
		if s := buildCompareOrFilterClause(paramMap, sign, paramValue, table+"."+likeColumn, idx); s != "" {
			*filterClauses = append(*filterClauses, s)
		}
	}
}

// buildLikeOrFilterClause build sql clause for "like or" type filter
func buildLikeOrFilterClause(paramMap map[string]interface{}, paramValue string, column string, idx *int) string {
	if idx == nil {
		newIdx := 0
		idx = &newIdx
	}
	var filters []string
	// the paramValue is the user input string
	for _, filter := range strings.Split(paramValue, ",") {
		// filter=email like @email1. paramMap["email1"]="%xxx%"
		// will be replaced by Where(clause, paramMap)
		columnParm := column + strconv.Itoa(*idx)
		*idx = *idx + 1
		paramMap[columnParm] = fmt.Sprintf("%%%v%%", filter)
		filters = append(filters, column+" like @"+columnParm)
	}
	if len(filters) > 0 {
		// (email like @email1 or email like @email2 or email like @email3)
		return fmt.Sprintf("(%v)", strings.Join(filters, " or "))
	}
	return ""
}

// buildCompareOrFilterClause build sql clause for "compare or" type filter
func buildCompareOrFilterClause(paramMap map[string]interface{}, sign string, paramValue string, column string, idx *int) string {
	if idx == nil {
		newIdx := 0
		idx = &newIdx
	}
	var filters []string
	for _, filter := range strings.Split(paramValue, ",") {
		columnParm := column + strconv.Itoa(*idx)
		*idx = *idx + 1
		paramMap[columnParm] = filter
		filters = append(filters, column+sign+" @"+columnParm)
	}
	if len(filters) > 0 {
		return fmt.Sprintf("(%v)", strings.Join(filters, " or "))
	}
	return ""
}
