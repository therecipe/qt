#ifdef __cplusplus
extern "C" {
#endif

void* QTimeZone_NewQTimeZone();
void* QTimeZone_NewQTimeZone2(void* ianaId);
void* QTimeZone_NewQTimeZone4(void* ianaId, int offsetSeconds, char* name, char* abbreviation, int country, char* comment);
void* QTimeZone_NewQTimeZone5(void* other);
void* QTimeZone_NewQTimeZone3(int offsetSeconds);
char* QTimeZone_Abbreviation(void* ptr, void* atDateTime);
char* QTimeZone_Comment(void* ptr);
int QTimeZone_Country(void* ptr);
int QTimeZone_DaylightTimeOffset(void* ptr, void* atDateTime);
char* QTimeZone_DisplayName2(void* ptr, int timeType, int nameType, void* locale);
char* QTimeZone_DisplayName(void* ptr, void* atDateTime, int nameType, void* locale);
int QTimeZone_HasDaylightTime(void* ptr);
int QTimeZone_HasTransitions(void* ptr);
void* QTimeZone_QTimeZone_IanaIdToWindowsId(void* ianaId);
void* QTimeZone_Id(void* ptr);
int QTimeZone_IsDaylightTime(void* ptr, void* atDateTime);
int QTimeZone_QTimeZone_IsTimeZoneIdAvailable(void* ianaId);
int QTimeZone_IsValid(void* ptr);
int QTimeZone_OffsetFromUtc(void* ptr, void* atDateTime);
int QTimeZone_StandardTimeOffset(void* ptr, void* atDateTime);
void QTimeZone_Swap(void* ptr, void* other);
void* QTimeZone_QTimeZone_SystemTimeZone();
void* QTimeZone_QTimeZone_SystemTimeZoneId();
void* QTimeZone_QTimeZone_Utc();
void* QTimeZone_QTimeZone_WindowsIdToDefaultIanaId(void* windowsId);
void* QTimeZone_QTimeZone_WindowsIdToDefaultIanaId2(void* windowsId, int country);
void QTimeZone_DestroyQTimeZone(void* ptr);

#ifdef __cplusplus
}
#endif