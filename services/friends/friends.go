package friends

import (
	"fmt"
	"go-gin-chat/models"
	"go-gin-chat/services/session"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendController struct{}

// AddFriend 添加好友
func AddFriend(c *gin.Context) {
	// 获取当前用户ID
	userInfo := session.GetSessionUserInfo(c)
	if userInfo == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请先登录",
		})
		return
	}

	// 获取要添加的好友ID
	toUserIdStr := c.PostForm("to_user_id")
	toUserId, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}

	// 检查是否已经是好友
	var existingFriend models.Friends
	result := models.ChatDB.Where("user_id = ? AND to_user_id = ?", userInfo["uid"], toUserId).First(&existingFriend)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "已经是好友关系",
		})
		return
	}

	user_id := int(userInfo["uid"].(uint))

	fmt.Println(user_id, reflect.TypeOf(user_id))
	// 创建好友关系
	friend := models.Friends{
		UserId:   user_id,
		ToUserId: toUserId,
	}

	if err := models.ChatDB.Create(&friend).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "添加好友失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "添加好友成功",
	})
}
