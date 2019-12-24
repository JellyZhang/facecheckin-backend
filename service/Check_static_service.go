package service

import (
	"facecheckin/model"
	"facecheckin/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type CheckstatisticService struct {
	UserId    string `form:"uid" json:"uid" binding:"required"`
	MeetingId string `form:"mid" json:"mid" binding:"required"`
}

func (service CheckstatisticService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("phone_number = ?", service.UserId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到用户", nil)
		return &rtn
	}

	count = 0
	model.DB.Model(&model.Meeting{}).Where("mid = ?", service.MeetingId).Count(&count)
	if count == 0 {
		rtn := serializer.ParamErr("未找到meeting", nil)
		return &rtn
	}

	return nil
}

const Duration = 2 * 7

func (service CheckstatisticService) GetStatic(c *gin.Context) serializer.Response {
	// valid
	if err := service.valid(); err != nil {
		return *err
	}
	var relations []model.Relation
	var personal []model.DayStatistic
	var group []model.DayStatistic
	if err := model.DB.Where("meeting_id = ? AND type = 0", service.MeetingId).Find(&relations).Error; err != nil {
		return serializer.ParamErr("查询用户出错", err)
	}
	for _, v := range relations {
		temp := getonesstatic(v.UserId, v.MeetingId, v.CreatedAt)
		group = append(group, temp...)
		if v.UserId == service.UserId {
			personal = temp
		}
	}
	checkstatic := model.CheckStatistic{
		Personal: personal,
		Group:    group,
	}
	return serializer.BuildCheckStatisticResponse(checkstatic)
}


func getonesstatic(userid string, meetingid string, jointime time.Time) []model.DayStatistic {
	var mainUser model.User
	var mainMeeting model.Meeting
	model.DB.Where("phone_number = ?", userid).First(&mainUser)
	model.DB.Where("mid = ?", meetingid).First(&mainMeeting)

	var personal []model.DayStatistic

	now := time.Now().AddDate(0, 0, 1)
	weeks := strings.Split(mainMeeting.CheckRule, "_")
	w := make(map[int]int)
	for _, v := range weeks {
		temp, _ := strconv.Atoi(v)
		w[temp] = 1
	}
	for i := 0; i < Duration; i++ {
		now = now.AddDate(0, 0, -1)
		fmt.Println(now)
		if now.Before(jointime) {
			break
		}
		if _, ok := w[int(now.Weekday())]; !ok {
			continue
		}
		var checkin model.Check
		var checkout model.Check
		count := 0
		var status string
		model.DB.Where("user_id = ? AND meeting_id = ? AND check_type = 1 AND created_at >= ? AND created_at < ?", userid, meetingid, now.Format("2006-01-02"), now.AddDate(0, 0, 1).Format("2006-01-02")).First(&checkin).Count(&count)
		if count == 0 {
			status = "00"
		} else {
			h, m, _ := checkin.CreatedAt.Clock()
			stamp := (60 * h) + m
			if stamp <= mainMeeting.TimeStart {
				status += "1"
			} else {
				status += "2"
			}
			model.DB.Where("user_id = ? AND meeting_id = ? AND check_type = 2 AND created_at >= ? AND created_at < ?", userid, meetingid, now.Format("2006-01-02"), now.AddDate(0, 0, 1).Format("2006-01-02")).First(&checkout).Count(&count)
			if count == 0 {
				status += "0"
			} else {
				h, m, _ := checkout.CreatedAt.Clock()
				stamp := (60 * h) + m
				if stamp <= mainMeeting.TimeEnd {
					status += "1"
				} else {
					status += "2"
				}
			}
		}
		newday := model.DayStatistic{
			UserId: userid,
			Date:   now.Format("2006-01-02"),
			Status: status,
		}
		personal = append(personal, newday)
	}
	return personal

}
