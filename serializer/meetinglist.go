package serializer

import (
	"facecheckin/model"
)

type Meeting struct{
	MeetingCover string `json:"mcover"`
	MeetingName string `json:"mname"`
	MeetingId string `json:"mid"`
	Timespace TimeSpace `json:"timespace"`
}

func BuildMeetingListResponse(meetings []model.Meeting) Response{
	var meetinglist []Meeting
	for _,m := range meetings{
		newmeeting := Meeting{
			MeetingCover: m.MeetingCover,
			MeetingName:  m.MeetingName,
			MeetingId:    m.Mid,
			Timespace:    TimeSpace{
				TimeStart: m.TimeStart,
				TimeEnd:   m.TimeEnd,
			},
		}
		meetinglist = append(meetinglist, newmeeting)
	}
	return Response{
		Code:  0,
		Data:  meetinglist,
	}
}