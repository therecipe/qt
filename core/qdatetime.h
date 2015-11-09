#ifdef __cplusplus
extern "C" {
#endif

void* QDateTime_QDateTime_CurrentDateTime();
void* QDateTime_QDateTime_CurrentDateTimeUtc();
void* QDateTime_QDateTime_FromString(char* stri, int format);
void* QDateTime_QDateTime_FromString2(char* stri, char* format);
void* QDateTime_ToOffsetFromUtc(void* ptr, int offsetSeconds);
char* QDateTime_ToString2(void* ptr, int format);
void* QDateTime_ToTimeSpec(void* ptr, int spec);
void* QDateTime_NewQDateTime();
void* QDateTime_NewQDateTime2(void* date);
void* QDateTime_NewQDateTime3(void* date, void* time, int spec);
void* QDateTime_NewQDateTime4(void* date, void* time, int spec, int offsetSeconds);
void* QDateTime_NewQDateTime5(void* date, void* time, void* timeZone);
void* QDateTime_NewQDateTime6(void* other);
void* QDateTime_AddMonths(void* ptr, int nmonths);
void* QDateTime_AddYears(void* ptr, int nyears);
int QDateTime_IsDaylightTime(void* ptr);
int QDateTime_IsNull(void* ptr);
int QDateTime_IsValid(void* ptr);
int QDateTime_OffsetFromUtc(void* ptr);
void QDateTime_SetDate(void* ptr, void* date);
void QDateTime_SetOffsetFromUtc(void* ptr, int offsetSeconds);
void QDateTime_SetTime(void* ptr, void* time);
void QDateTime_SetTimeSpec(void* ptr, int spec);
void QDateTime_SetTimeZone(void* ptr, void* toZone);
void QDateTime_Swap(void* ptr, void* other);
int QDateTime_TimeSpec(void* ptr);
void* QDateTime_TimeZone(void* ptr);
char* QDateTime_TimeZoneAbbreviation(void* ptr);
void* QDateTime_ToLocalTime(void* ptr);
char* QDateTime_ToString(void* ptr, char* format);
void* QDateTime_ToTimeZone(void* ptr, void* timeZone);
void* QDateTime_ToUTC(void* ptr);
void QDateTime_DestroyQDateTime(void* ptr);

#ifdef __cplusplus
}
#endif