#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTimeZone_NewQTimeZone();
QtObjectPtr QTimeZone_NewQTimeZone2(QtObjectPtr ianaId);
QtObjectPtr QTimeZone_NewQTimeZone4(QtObjectPtr ianaId, int offsetSeconds, char* name, char* abbreviation, int country, char* comment);
QtObjectPtr QTimeZone_NewQTimeZone5(QtObjectPtr other);
QtObjectPtr QTimeZone_NewQTimeZone3(int offsetSeconds);
char* QTimeZone_Abbreviation(QtObjectPtr ptr, QtObjectPtr atDateTime);
char* QTimeZone_Comment(QtObjectPtr ptr);
int QTimeZone_Country(QtObjectPtr ptr);
int QTimeZone_DaylightTimeOffset(QtObjectPtr ptr, QtObjectPtr atDateTime);
char* QTimeZone_DisplayName2(QtObjectPtr ptr, int timeType, int nameType, QtObjectPtr locale);
char* QTimeZone_DisplayName(QtObjectPtr ptr, QtObjectPtr atDateTime, int nameType, QtObjectPtr locale);
int QTimeZone_HasDaylightTime(QtObjectPtr ptr);
int QTimeZone_HasTransitions(QtObjectPtr ptr);
int QTimeZone_IsDaylightTime(QtObjectPtr ptr, QtObjectPtr atDateTime);
int QTimeZone_QTimeZone_IsTimeZoneIdAvailable(QtObjectPtr ianaId);
int QTimeZone_IsValid(QtObjectPtr ptr);
int QTimeZone_OffsetFromUtc(QtObjectPtr ptr, QtObjectPtr atDateTime);
int QTimeZone_StandardTimeOffset(QtObjectPtr ptr, QtObjectPtr atDateTime);
void QTimeZone_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QTimeZone_DestroyQTimeZone(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif