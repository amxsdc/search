package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"search/src/common"
	"search/src/dao"
	"search/src/model"
)

var record model.Record

type Record struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func GetRecordList(keyword string) ([]model.Record, error) {
	var recordlist []model.Record
	if err := dao.SqlSession.Table("record").Select("title").Where("title LIKE ?", keyword+"%").Find(&recordlist).Error; err != nil {
		return recordlist, err
	}
	return recordlist, nil
}

func SaveKeyword(Keyword string) {
	var c *gin.Context
	id, err := getlastId()
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 403,
			StatusMsg:  "Failed to get recordList",
		})
		c.Abort()
		return
	}
	record := Record{Title: Keyword,
		Id: id}
	dao.SqlSession.Select("Title").Create(&record)
}

func getlastId() (uint, error) {
	var lastId uint
	if err := dao.SqlSession.Table("record").Select("id").Last(&record).Error; err != nil {
		return lastId, err
	}
	return lastId, nil
}
