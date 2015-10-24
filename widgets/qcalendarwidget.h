#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCalendarWidget_DateEditAcceptDelay(QtObjectPtr ptr);
int QCalendarWidget_FirstDayOfWeek(QtObjectPtr ptr);
int QCalendarWidget_HorizontalHeaderFormat(QtObjectPtr ptr);
int QCalendarWidget_IsDateEditEnabled(QtObjectPtr ptr);
int QCalendarWidget_IsGridVisible(QtObjectPtr ptr);
int QCalendarWidget_IsNavigationBarVisible(QtObjectPtr ptr);
int QCalendarWidget_SelectionMode(QtObjectPtr ptr);
void QCalendarWidget_SetDateEditAcceptDelay(QtObjectPtr ptr, int delay);
void QCalendarWidget_SetDateEditEnabled(QtObjectPtr ptr, int enable);
void QCalendarWidget_SetFirstDayOfWeek(QtObjectPtr ptr, int dayOfWeek);
void QCalendarWidget_SetGridVisible(QtObjectPtr ptr, int show);
void QCalendarWidget_SetHorizontalHeaderFormat(QtObjectPtr ptr, int format);
void QCalendarWidget_SetMaximumDate(QtObjectPtr ptr, QtObjectPtr date);
void QCalendarWidget_SetMinimumDate(QtObjectPtr ptr, QtObjectPtr date);
void QCalendarWidget_SetNavigationBarVisible(QtObjectPtr ptr, int visible);
void QCalendarWidget_SetSelectedDate(QtObjectPtr ptr, QtObjectPtr date);
void QCalendarWidget_SetSelectionMode(QtObjectPtr ptr, int mode);
void QCalendarWidget_SetVerticalHeaderFormat(QtObjectPtr ptr, int format);
int QCalendarWidget_VerticalHeaderFormat(QtObjectPtr ptr);
QtObjectPtr QCalendarWidget_NewQCalendarWidget(QtObjectPtr parent);
void QCalendarWidget_ConnectCurrentPageChanged(QtObjectPtr ptr);
void QCalendarWidget_DisconnectCurrentPageChanged(QtObjectPtr ptr);
int QCalendarWidget_MonthShown(QtObjectPtr ptr);
void QCalendarWidget_ConnectSelectionChanged(QtObjectPtr ptr);
void QCalendarWidget_DisconnectSelectionChanged(QtObjectPtr ptr);
void QCalendarWidget_SetCurrentPage(QtObjectPtr ptr, int year, int month);
void QCalendarWidget_SetDateRange(QtObjectPtr ptr, QtObjectPtr min, QtObjectPtr max);
void QCalendarWidget_SetDateTextFormat(QtObjectPtr ptr, QtObjectPtr date, QtObjectPtr format);
void QCalendarWidget_SetHeaderTextFormat(QtObjectPtr ptr, QtObjectPtr format);
void QCalendarWidget_SetWeekdayTextFormat(QtObjectPtr ptr, int dayOfWeek, QtObjectPtr format);
void QCalendarWidget_ShowNextMonth(QtObjectPtr ptr);
void QCalendarWidget_ShowNextYear(QtObjectPtr ptr);
void QCalendarWidget_ShowPreviousMonth(QtObjectPtr ptr);
void QCalendarWidget_ShowPreviousYear(QtObjectPtr ptr);
void QCalendarWidget_ShowSelectedDate(QtObjectPtr ptr);
void QCalendarWidget_ShowToday(QtObjectPtr ptr);
int QCalendarWidget_YearShown(QtObjectPtr ptr);
void QCalendarWidget_DestroyQCalendarWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif