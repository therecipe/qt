#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QDateTime_ToString2(QtObjectPtr ptr, int format);
QtObjectPtr QDateTime_NewQDateTime();
QtObjectPtr QDateTime_NewQDateTime2(QtObjectPtr date);
QtObjectPtr QDateTime_NewQDateTime3(QtObjectPtr date, QtObjectPtr time, int spec);
QtObjectPtr QDateTime_NewQDateTime4(QtObjectPtr date, QtObjectPtr time, int spec, int offsetSeconds);
QtObjectPtr QDateTime_NewQDateTime5(QtObjectPtr date, QtObjectPtr time, QtObjectPtr timeZone);
QtObjectPtr QDateTime_NewQDateTime6(QtObjectPtr other);
int QDateTime_IsDaylightTime(QtObjectPtr ptr);
int QDateTime_IsNull(QtObjectPtr ptr);
int QDateTime_IsValid(QtObjectPtr ptr);
int QDateTime_OffsetFromUtc(QtObjectPtr ptr);
void QDateTime_SetDate(QtObjectPtr ptr, QtObjectPtr date);
void QDateTime_SetOffsetFromUtc(QtObjectPtr ptr, int offsetSeconds);
void QDateTime_SetTime(QtObjectPtr ptr, QtObjectPtr time);
void QDateTime_SetTimeSpec(QtObjectPtr ptr, int spec);
void QDateTime_SetTimeZone(QtObjectPtr ptr, QtObjectPtr toZone);
void QDateTime_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QDateTime_TimeSpec(QtObjectPtr ptr);
char* QDateTime_TimeZoneAbbreviation(QtObjectPtr ptr);
char* QDateTime_ToString(QtObjectPtr ptr, char* format);
void QDateTime_DestroyQDateTime(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif