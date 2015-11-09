#ifdef __cplusplus
extern "C" {
#endif

int QCalendarWidget_DateEditAcceptDelay(void* ptr);
int QCalendarWidget_FirstDayOfWeek(void* ptr);
int QCalendarWidget_HorizontalHeaderFormat(void* ptr);
int QCalendarWidget_IsDateEditEnabled(void* ptr);
int QCalendarWidget_IsGridVisible(void* ptr);
int QCalendarWidget_IsNavigationBarVisible(void* ptr);
int QCalendarWidget_SelectionMode(void* ptr);
void QCalendarWidget_SetDateEditAcceptDelay(void* ptr, int delay);
void QCalendarWidget_SetDateEditEnabled(void* ptr, int enable);
void QCalendarWidget_SetFirstDayOfWeek(void* ptr, int dayOfWeek);
void QCalendarWidget_SetGridVisible(void* ptr, int show);
void QCalendarWidget_SetHorizontalHeaderFormat(void* ptr, int format);
void QCalendarWidget_SetMaximumDate(void* ptr, void* date);
void QCalendarWidget_SetMinimumDate(void* ptr, void* date);
void QCalendarWidget_SetNavigationBarVisible(void* ptr, int visible);
void QCalendarWidget_SetSelectedDate(void* ptr, void* date);
void QCalendarWidget_SetSelectionMode(void* ptr, int mode);
void QCalendarWidget_SetVerticalHeaderFormat(void* ptr, int format);
int QCalendarWidget_VerticalHeaderFormat(void* ptr);
void* QCalendarWidget_NewQCalendarWidget(void* parent);
void QCalendarWidget_ConnectCurrentPageChanged(void* ptr);
void QCalendarWidget_DisconnectCurrentPageChanged(void* ptr);
int QCalendarWidget_MonthShown(void* ptr);
void QCalendarWidget_ConnectSelectionChanged(void* ptr);
void QCalendarWidget_DisconnectSelectionChanged(void* ptr);
void QCalendarWidget_SetCurrentPage(void* ptr, int year, int month);
void QCalendarWidget_SetDateRange(void* ptr, void* min, void* max);
void QCalendarWidget_SetDateTextFormat(void* ptr, void* date, void* format);
void QCalendarWidget_SetHeaderTextFormat(void* ptr, void* format);
void QCalendarWidget_SetWeekdayTextFormat(void* ptr, int dayOfWeek, void* format);
void QCalendarWidget_ShowNextMonth(void* ptr);
void QCalendarWidget_ShowNextYear(void* ptr);
void QCalendarWidget_ShowPreviousMonth(void* ptr);
void QCalendarWidget_ShowPreviousYear(void* ptr);
void QCalendarWidget_ShowSelectedDate(void* ptr);
void QCalendarWidget_ShowToday(void* ptr);
int QCalendarWidget_YearShown(void* ptr);
void QCalendarWidget_DestroyQCalendarWidget(void* ptr);

#ifdef __cplusplus
}
#endif