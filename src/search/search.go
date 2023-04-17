package search

import (
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"github.com/gin-gonic/gin"
	"net/http"
	"search/src/common"
	"search/src/service"
)

type RecordResp struct {
	Title string `json:"title"`
}
type RecordListResp struct {
	common.Response
	RecordList []RecordResp `json:"record_list"`
}

func Search(c *gin.Context) {

	keyword := c.DefaultQuery("Keyword", "请输入要查询的内容")
	recordList, err := service.GetRecordList(keyword)
	// 拼音匹配
	dict := pinyin.Dict{}
	keyword_pinyin := dict.Name(keyword, " ").None()
	recordList_pinyin, _ := service.GetRecordList(keyword_pinyin)

	recordList = append(recordList, recordList_pinyin...)

	// save keyword
	//service.SaveKeyword(keyword)

	// record not found
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 403,
			StatusMsg:  "Failed to get recordList",
		})
		c.Abort()
		return
	}

	//found record
	var RespRecordList []RecordResp
	for i := 0; i < len(recordList); i++ {
		RespRecord := RecordResp{
			Title: recordList[i].Title,
		}
		RespRecordList = append(RespRecordList, RespRecord)

	}

	//resp
	if RespRecordList == nil {
		service.SaveKeyword(keyword)
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 403,
			StatusMsg:  "Failed to get recordList",
		})
		c.Abort()
		return

	} else {
		c.JSON(http.StatusOK, RecordListResp{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "Successfully found the record list.",
			},
			RecordList: RespRecordList,
		})
	}

}
