#ifdef __cplusplus
extern "C" {
#endif

int QDate_QDate_IsLeapYear(int year);
char* QDate_ToString2(void* ptr, int format);
void* QDate_NewQDate();
void* QDate_NewQDate3(int y, int m, int d);
int QDate_Day(void* ptr);
int QDate_DayOfWeek(void* ptr);
int QDate_DayOfYear(void* ptr);
int QDate_DaysInMonth(void* ptr);
int QDate_DaysInYear(void* ptr);
void QDate_GetDate(void* ptr, int year, int month, int day);
int QDate_IsNull(void* ptr);
int QDate_QDate_IsValid2(int year, int month, int day);
int QDate_IsValid(void* ptr);
char* QDate_QDate_LongDayName(int weekday, int ty);
char* QDate_QDate_LongMonthName(int month, int ty);
int QDate_Month(void* ptr);
int QDate_SetDate(void* ptr, int year, int month, int day);
char* QDate_QDate_ShortDayName(int weekday, int ty);
char* QDate_QDate_ShortMonthName(int month, int ty);
char* QDate_ToString(void* ptr, char* format);
int QDate_WeekNumber(void* ptr, int yearNumber);
int QDate_Year(void* ptr);

#ifdef __cplusplus
}
#endif