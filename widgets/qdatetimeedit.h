#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDateTimeEdit_NewQDateTimeEdit3(QtObjectPtr date, QtObjectPtr parent);
QtObjectPtr QDateTimeEdit_NewQDateTimeEdit4(QtObjectPtr time, QtObjectPtr parent);
int QDateTimeEdit_CalendarPopup(QtObjectPtr ptr);
void QDateTimeEdit_ClearMaximumDate(QtObjectPtr ptr);
void QDateTimeEdit_ClearMaximumDateTime(QtObjectPtr ptr);
void QDateTimeEdit_ClearMaximumTime(QtObjectPtr ptr);
void QDateTimeEdit_ClearMinimumDate(QtObjectPtr ptr);
void QDateTimeEdit_ClearMinimumDateTime(QtObjectPtr ptr);
void QDateTimeEdit_ClearMinimumTime(QtObjectPtr ptr);
int QDateTimeEdit_CurrentSection(QtObjectPtr ptr);
int QDateTimeEdit_CurrentSectionIndex(QtObjectPtr ptr);
char* QDateTimeEdit_DisplayFormat(QtObjectPtr ptr);
int QDateTimeEdit_DisplayedSections(QtObjectPtr ptr);
int QDateTimeEdit_SectionCount(QtObjectPtr ptr);
char* QDateTimeEdit_SectionText(QtObjectPtr ptr, int section);
void QDateTimeEdit_SetCalendarPopup(QtObjectPtr ptr, int enable);
void QDateTimeEdit_SetCurrentSection(QtObjectPtr ptr, int section);
void QDateTimeEdit_SetCurrentSectionIndex(QtObjectPtr ptr, int index);
void QDateTimeEdit_SetDate(QtObjectPtr ptr, QtObjectPtr date);
void QDateTimeEdit_SetDateTime(QtObjectPtr ptr, QtObjectPtr dateTime);
void QDateTimeEdit_SetDisplayFormat(QtObjectPtr ptr, char* format);
void QDateTimeEdit_SetMaximumDate(QtObjectPtr ptr, QtObjectPtr max);
void QDateTimeEdit_SetMaximumDateTime(QtObjectPtr ptr, QtObjectPtr dt);
void QDateTimeEdit_SetMaximumTime(QtObjectPtr ptr, QtObjectPtr max);
void QDateTimeEdit_SetMinimumDate(QtObjectPtr ptr, QtObjectPtr min);
void QDateTimeEdit_SetMinimumDateTime(QtObjectPtr ptr, QtObjectPtr dt);
void QDateTimeEdit_SetMinimumTime(QtObjectPtr ptr, QtObjectPtr min);
void QDateTimeEdit_SetTime(QtObjectPtr ptr, QtObjectPtr time);
void QDateTimeEdit_SetTimeSpec(QtObjectPtr ptr, int spec);
int QDateTimeEdit_TimeSpec(QtObjectPtr ptr);
QtObjectPtr QDateTimeEdit_NewQDateTimeEdit(QtObjectPtr parent);
QtObjectPtr QDateTimeEdit_NewQDateTimeEdit2(QtObjectPtr datetime, QtObjectPtr parent);
QtObjectPtr QDateTimeEdit_CalendarWidget(QtObjectPtr ptr);
void QDateTimeEdit_Clear(QtObjectPtr ptr);
int QDateTimeEdit_Event(QtObjectPtr ptr, QtObjectPtr event);
int QDateTimeEdit_SectionAt(QtObjectPtr ptr, int index);
void QDateTimeEdit_SetCalendarWidget(QtObjectPtr ptr, QtObjectPtr calendarWidget);
void QDateTimeEdit_SetDateRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max);
void QDateTimeEdit_SetDateTimeRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max);
void QDateTimeEdit_SetSelectedSection(QtObjectPtr ptr, int section);
void QDateTimeEdit_SetTimeRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max);
void QDateTimeEdit_StepBy(QtObjectPtr ptr, int steps);
void QDateTimeEdit_DestroyQDateTimeEdit(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif