package util

import (
	"encoding/json"
	api "gas-detect/api/gas"
	"github.com/go-kratos/kratos/v2/errors"
	"sort"
	"strconv"
)

func MetricsToMap(apiKeyValues []*api.Metric) []map[string]string {
	result := make([]map[string]string, 0)
	if apiKeyValues == nil {
		return result
	}
	for _, entry := range apiKeyValues {
		item := make(map[string]string)
		item["typeName"] = entry.TypeName
		item["concentration"] = entry.Concentration
		result = append(result, item)
	}

	return result
}

func MapToMetrics(maps []map[string]string) []*api.Metric {
	metrics := make([]*api.Metric, 0)

	for _, item := range maps {
		metrics = append(metrics, &api.Metric{
			TypeName:      item["typeName"],
			Concentration: item["concentration"],
		})
	}

	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].TypeName < metrics[j].TypeName
	})

	return metrics
}

func JSONStringToMetrics(jString string) (metrics []*api.Metric, err error) {
	var res []map[string]string
	err = json.Unmarshal([]byte(jString), &res)
	if err != nil {
		return nil, err
	}

	for _, item := range res {
		metrics = append(metrics, &api.Metric{
			TypeName:      item["TypeName"],
			Concentration: item["Concentration"],
		})
	}
	return metrics, nil
}

func MetricsToJSONString(metrics []*api.Metric) (string, error) {
	result := make([]map[string]string, 0)
	if metrics == nil {
		return "", nil
	}
	for _, entry := range metrics {
		item := make(map[string]string)
		item[entry.TypeName] = entry.TypeName
		item[entry.Concentration] = entry.Concentration
		result = append(result, item)
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(resultBytes), nil
}

func StrToInt64(strV string) (intV int64, err error) {
	intV, err = strconv.ParseInt(strV, 10, 64)
	if err != nil {
		return 0, errors.BadRequest("string to int64 类型转换失败", err.Error())
	}
	return
}

func StrToFloat64(strV string) (intV float64, err error) {
	intV, err = strconv.ParseFloat(strV, 10)
	if err != nil {
		return 0, errors.BadRequest("string to float64 类型转换失败", err.Error())
	}
	return
}

func StrToInt32(strV string) (int32, error) {
	intV, err := strconv.Atoi(strV)
	if err != nil {
		return 0, errors.BadRequest("string to int32 类型转换失败", err.Error())
	}
	return int32(intV), nil
}

func Int64ToStr(intV int64) (strV string, err error) {
	strV = strconv.FormatInt(intV, 10)
	if err != nil {
		return "", errors.BadRequest("string to int64 类型转换失败", "")
	}
	return
}
