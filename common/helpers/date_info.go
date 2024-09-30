package helpers

import "github.com/ortizdavid/go-nopain/datetime"

type DateInfo struct {
	CurrentDate           string
	CurrentDateTime       string
	LastDayOfCurrentWeek	string
	LastDayOfCurrentMonth string
	LastDayOfCurrentYear  string
	EighteenYearsOld        string
}

func GetDateInfo() DateInfo {
	return DateInfo{
		CurrentDate:           datetime.CurrentDate(),
		CurrentDateTime:       datetime.CurrentDateTime(),
		LastDayOfCurrentWeek:  datetime.LastDayOfCurrentWeekStr(),
		LastDayOfCurrentMonth: datetime.LastDayOfCurrentMonthStr(),
		LastDayOfCurrentYear:  datetime.LastDateOfYearStr(),
		EighteenYearsOld:      datetime.SubtractYearsStr(datetime.CurrentDate(), 18),
	}
}
