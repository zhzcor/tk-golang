package apiservice

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gserver/internal/store"
	"gserver/internal/store/ent/city"

)

type City struct{}

type cityList struct {
	Id  int    `json:"id"`
	Name string `json:"name"`
}

func CityService() *City {
	return &City{}
}

// 查询城市id
func (s *City) GetCityList(ctx *gin.Context) ([]*cityList,error) {
	var list []*cityList
	cityQuery := store.WithContext(ctx).City.Query().SoftDelete()
	err := cityQuery.Select(city.FieldID,city.FieldName).Scan(ctx,&list)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	return list,err
}

