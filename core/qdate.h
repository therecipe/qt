#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDate_QDate_IsLeapYear(int year);
char* QDate_ToString2(QtObjectPtr ptr, int format);
QtObjectPtr QDate_NewQDate();
QtObjectPtr QDate_NewQDate3(int y, int m, int d);
int QDate_Day(QtObjectPtr ptr);
int QDate_DayOfWeek(QtObjectPtr ptr);
int QDate_DayOfYear(QtObjectPtr ptr);
int QDate_DaysInMonth(QtObjectPtr ptr);
int QDate_DaysInYear(QtObjectPtr ptr);
void QDate_GetDate(QtObjectPtr ptr, int year, int month, int day);
int QDate_IsNull(QtObjectPtr ptr);
int QDate_QDate_IsValid2(int year, int month, int day);
int QDate_IsValid(QtObjectPtr ptr);
char* QDate_QDate_LongDayName(int weekday, int ty);
char* QDate_QDate_LongMonthName(int month, int ty);
int QDate_Month(QtObjectPtr ptr);
int QDate_SetDate(QtObjectPtr ptr, int year, int month, int day);
char* QDate_QDate_ShortDayName(int weekday, int ty);
char* QDate_QDate_ShortMonthName(int month, int ty);
char* QDate_ToString(QtObjectPtr ptr, char* format);
int QDate_WeekNumber(QtObjectPtr ptr, int yearNumber);
int QDate_Year(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif