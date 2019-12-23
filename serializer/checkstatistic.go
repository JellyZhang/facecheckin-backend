package serializer

import "facecheckin/model"

func BuildCheckStatisticResponse(statistic model.CheckStatistic) Response {
	return Response{
		Code:  0,
		Data:  statistic,
		Msg:   "",
		Error: "",
	}
}
