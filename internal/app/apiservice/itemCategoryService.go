package apiservice

import (
	"github.com/gin-gonic/gin"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/itemcategory"
)

type ItemCate struct{}

type ItemCateList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Parent *ItemCateList
}

func NewItemCate() *ItemCate {
	return &ItemCate{}
}

// 查询所有项目一级id
func (s *ItemCate) GetItemCateList(ctx *gin.Context) []*ItemCateList {
	var list []*ItemCateList
	itemQuery := store.WithContext(ctx).ItemCategory.Query().SoftDelete()
	data := itemQuery.Select(itemcategory.FieldID, itemcategory.FieldName, itemcategory.FieldPid).WithParent(func(query *ent.ItemCategoryQuery) {
		query.SoftDelete().Where(itemcategory.PidIsNil())
	}).AllX(ctx)

	for _, v := range data {
		re := ItemCateList{
			Parent: &ItemCateList{},
		}
		re.Id = v.ID
		re.Name = v.Name
		if v.Edges.Parent != nil {
			re.Parent.Id = v.Edges.Parent.ID
			re.Parent.Name = v.Edges.Parent.Name
		}

		list = append(list, &re)
	}
	return list
}
