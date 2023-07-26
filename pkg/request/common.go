package request

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"

	"golang.org/x/net/context"
	"strings"
)

// PageParams 分页请求基础对象
type PageParams struct {
	Offset   int32     `json:"offset"`   // 偏移量
	PageSize int32     `json:"pageSize"` // 每页大小
	Sorter   SortField `json:"sorter"`   // 排序字段,默认created_time-descend

}

type SearchParams struct {
	Privacy               string `json:"privacy"`                 // 访问类别
	OrgId                 string `json:"org_id"`                  // org资源id
	PlatformProjectId     string `json:"platform_project_id"`     // platform级的project_id
	OrgProjectId          string `json:"org_project_id"`          // org级的project_id
	ProjectId             string `json:"project_id"`              // project_id
	UserId                string `json:"user_id"`                 // user_id
	Labels                string `json:"labels"`                  // label标签
	Name                  string `json:"name"`                    // name
	WorkflowTemplateName  string `json:"workflow_template_name"`  //workflow_template_name
	WorkflowTemplateId    int32  `json:"workflow_template_id"`    // workflow_template_id
	WorkspaceTemplateName string `json:"workspace_template_name"` //workspace_template_name
	WorkspaceTemplateId   int32  `json:"workspace_template_id"`   // workspace_template_id
}

const (
	DefaultPageNo      = 1   // 页码最小值为 1
	DefaultPageSize    = 15  // 每页条数, 默认15
	DefaultMaxPageSize = 100 // 最大每页条数, 默认100
	DefaultSep         = "-" // 最大每页条数, 默认100
	DefaultSortField   = "created_time"
	DefaultSortOrder   = "descend"
)

// SortField 排序对象
type SortField struct {
	Field string // 排序字段
	Order string // 排序顺序
}

const (
	ORDER_ASC  = "ascend"  // 升序
	ORDER_DESC = "descend" // 降序
)

func PaginateValidate(pageNo, pageSize int32, sorter string) (*PageParams, error) {

	if pageNo <= 0 {
		pageNo = DefaultPageNo
	}

	switch {
	case pageSize > DefaultMaxPageSize:
		pageSize = DefaultPageSize
	case pageSize <= 0:
		pageSize = DefaultPageSize
	}
	var sortFields []string
	sortField := DefaultSortField
	sortOrder := DefaultSortOrder
	if sorter != "" {
		sortFields = strings.Split(sorter, DefaultSep)
		if len(sortFields) != 2 {
			return nil, errors.BadRequest("查询参数异常", "sorter参数格式异常")
		}
		sortField = sortFields[0]
		sortOrder = sortFields[1]
		if sortOrder != ORDER_ASC && sortOrder != ORDER_DESC {
			return nil, errors.BadRequest("查询参数异常", "sorter参数格式异常")
		}
	}
	offset := (pageNo - 1) * pageSize
	return &PageParams{
		Offset:   offset,
		PageSize: pageSize,
		Sorter: SortField{
			Field: sortField,
			Order: sortOrder,
		},
	}, nil
}

type AimpHeader struct {
	AimpToken string `json:"aimp_token"`
}

func ValidateAimpHeader(ctx context.Context) error {

	if tp, ok := transport.FromServerContext(ctx); ok {
		XAimpToken := tp.RequestHeader().Get("X_AIMP_TOKEN")
		if XAimpToken == "" {
			return errors.BadRequest("X_AIMP_TOKEN 缺失", "")
		}
		// 临时密码
		if XAimpToken != "!@#$%^&*()" {
			return errors.BadRequest("X_AIMP_TOKEN 错误", "")
		}
	} else {
		//log.Info("md error:")
		return errors.BadRequest("token 缺失", "")
	}

	return nil
}
