package paginator

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-es/internal/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int   `json:"current_page"` // 当前页
	PageSize    int   `json:"page_size"`    // 每页条数
	TotalCount  int64 `json:"total_count"`  // 总条数
	TotalPage   int   `json:"total_page"`   // 总页数
}

// Paginator 分页操作类
type Paginator struct {
	CurrentPage int
	PageSize    int
	TotalCount  int64
	TotalPage   int
	Offset      int

	Sort  string // 排序规则
	Order string // 排序顺序

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin context，方便调用
}

// Paginate 分页
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, pageSize int) Paging {
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(pageSize)

	// 查询数据库
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PageSize).
		Offset(p.Offset).
		Find(data).
		Error
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.CurrentPage,
		PageSize:    p.PageSize,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
	}
}

// 初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(pageSize int) {
	p.TotalCount = p.getTotalCount()
	p.PageSize = p.getPageSize(pageSize)
	p.TotalPage = p.getTotalPage()
	p.CurrentPage = p.getCurrentPage()
	p.Offset = (p.CurrentPage - 1) * p.PageSize

	p.Order = "asc"
	p.Sort = "id"
}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}

	return count
}

func (p *Paginator) getPageSize(pageSize int) int {
	queryPeerage := p.ctx.Query("page-size")
	if len(queryPeerage) > 0 {
		pageSize = cast.ToInt(queryPeerage)
	}

	// 没有传参，使用默认
	if pageSize <= 0 {
		pageSize = 10
	}

	return pageSize
}

func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}

	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PageSize)))
	if nums == 0 {
		nums = 1
	}

	return int(nums)
}

// getCurrentPage 返回当前页码
func (p *Paginator) getCurrentPage() int {
	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query("per_page"))
	if page <= 0 {
		// 默认为 1
		page = 1
	}

	// TotalPage 等于 0 ，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}

	// 请求页数大于总页数，返回总页数
	/*if page > p.TotalPage {
		return p.TotalPage
	}*/

	return page
}
