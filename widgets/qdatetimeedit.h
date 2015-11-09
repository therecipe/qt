#ifdef __cplusplus
extern "C" {
#endif

void* QDateTimeEdit_NewQDateTimeEdit3(void* date, void* parent);
void* QDateTimeEdit_NewQDateTimeEdit4(void* time, void* parent);
int QDateTimeEdit_CalendarPopup(void* ptr);
void QDateTimeEdit_ClearMaximumDate(void* ptr);
void QDateTimeEdit_ClearMaximumDateTime(void* ptr);
void QDateTimeEdit_ClearMaximumTime(void* ptr);
void QDateTimeEdit_ClearMinimumDate(void* ptr);
void QDateTimeEdit_ClearMinimumDateTime(void* ptr);
void QDateTimeEdit_ClearMinimumTime(void* ptr);
int QDateTimeEdit_CurrentSection(void* ptr);
int QDateTimeEdit_CurrentSectionIndex(void* ptr);
void* QDateTimeEdit_DateTime(void* ptr);
char* QDateTimeEdit_DisplayFormat(void* ptr);
int QDateTimeEdit_DisplayedSections(void* ptr);
void* QDateTimeEdit_MaximumDateTime(void* ptr);
void* QDateTimeEdit_MinimumDateTime(void* ptr);
int QDateTimeEdit_SectionCount(void* ptr);
char* QDateTimeEdit_SectionText(void* ptr, int section);
void QDateTimeEdit_SetCalendarPopup(void* ptr, int enable);
void QDateTimeEdit_SetCurrentSection(void* ptr, int section);
void QDateTimeEdit_SetCurrentSectionIndex(void* ptr, int index);
void QDateTimeEdit_SetDate(void* ptr, void* date);
void QDateTimeEdit_SetDateTime(void* ptr, void* dateTime);
void QDateTimeEdit_SetDisplayFormat(void* ptr, char* format);
void QDateTimeEdit_SetMaximumDate(void* ptr, void* max);
void QDateTimeEdit_SetMaximumDateTime(void* ptr, void* dt);
void QDateTimeEdit_SetMaximumTime(void* ptr, void* max);
void QDateTimeEdit_SetMinimumDate(void* ptr, void* min);
void QDateTimeEdit_SetMinimumDateTime(void* ptr, void* dt);
void QDateTimeEdit_SetMinimumTime(void* ptr, void* min);
void QDateTimeEdit_SetTime(void* ptr, void* time);
void QDateTimeEdit_SetTimeSpec(void* ptr, int spec);
int QDateTimeEdit_TimeSpec(void* ptr);
void* QDateTimeEdit_NewQDateTimeEdit(void* parent);
void* QDateTimeEdit_NewQDateTimeEdit2(void* datetime, void* parent);
void* QDateTimeEdit_CalendarWidget(void* ptr);
void QDateTimeEdit_Clear(void* ptr);
void QDateTimeEdit_ConnectDateTimeChanged(void* ptr);
void QDateTimeEdit_DisconnectDateTimeChanged(void* ptr);
int QDateTimeEdit_Event(void* ptr, void* event);
int QDateTimeEdit_SectionAt(void* ptr, int index);
void QDateTimeEdit_SetCalendarWidget(void* ptr, void* calendarWidget);
void QDateTimeEdit_SetDateRange(void* ptr, void* min, void* max);
void QDateTimeEdit_SetDateTimeRange(void* ptr, void* min, void* max);
void QDateTimeEdit_SetSelectedSection(void* ptr, int section);
void QDateTimeEdit_SetTimeRange(void* ptr, void* min, void* max);
void QDateTimeEdit_StepBy(void* ptr, int steps);
void QDateTimeEdit_DestroyQDateTimeEdit(void* ptr);

#ifdef __cplusplus
}
#endif