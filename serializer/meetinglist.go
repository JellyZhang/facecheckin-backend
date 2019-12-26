package serializer

import (
	"facecheckin/model"
)

type Meeting struct{
	MeetingCover string `json:"mcover"`
	MeetingName string `json:"mname"`
	MeetingId string `json:"mid"`
	Timespace TimeSpace `json:"timespace"`
	Sum int `json:"sum"`
}

func BuildMeetingListResponse(meetings []model.Meeting) Response{
	var meetinglist []Meeting
	for _,m := range meetings{
		sum := 0
		model.DB.Model(&model.Relation{}).Where("meeting_id = ?",m.Mid).Count(&sum)
		newmeeting := Meeting{
			MeetingCover: m.MeetingCover,
			MeetingName:  m.MeetingName,
			MeetingId:    m.Mid,
			Timespace:    TimeSpace{
				TimeStart: m.TimeStart,
				TimeEnd:   m.TimeEnd,
			},
			Sum:sum,
		}
		meetinglist = append(meetinglist, newmeeting)
	}
	return Response{
		Code:  0,
		Data:  meetinglist,
	}
}