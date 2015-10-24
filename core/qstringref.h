#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStringRef_Begin(QtObjectPtr ptr);
QtObjectPtr QStringRef_Cbegin(QtObjectPtr ptr);
QtObjectPtr QStringRef_Cend(QtObjectPtr ptr);
void QStringRef_Clear(QtObjectPtr ptr);
int QStringRef_QStringRef_Compare3(QtObjectPtr s1, QtObjectPtr s2, int cs);
int QStringRef_QStringRef_Compare(QtObjectPtr s1, char* s2, int cs);
int QStringRef_QStringRef_Compare2(QtObjectPtr s1, QtObjectPtr s2, int cs);
int QStringRef_Compare6(QtObjectPtr ptr, QtObjectPtr other, int cs);
int QStringRef_Compare4(QtObjectPtr ptr, char* other, int cs);
int QStringRef_Compare5(QtObjectPtr ptr, QtObjectPtr other, int cs);
QtObjectPtr QStringRef_ConstData(QtObjectPtr ptr);
int QStringRef_Contains2(QtObjectPtr ptr, QtObjectPtr ch, int cs);
int QStringRef_Contains4(QtObjectPtr ptr, QtObjectPtr str, int cs);
int QStringRef_Contains(QtObjectPtr ptr, char* str, int cs);
int QStringRef_Contains3(QtObjectPtr ptr, QtObjectPtr str, int cs);
int QStringRef_Count(QtObjectPtr ptr);
int QStringRef_Count3(QtObjectPtr ptr, QtObjectPtr ch, int cs);
int QStringRef_Count2(QtObjectPtr ptr, char* str, int cs);
int QStringRef_Count4(QtObjectPtr ptr, QtObjectPtr str, int cs);
QtObjectPtr QStringRef_Data(QtObjectPtr ptr);
QtObjectPtr QStringRef_End(QtObjectPtr ptr);
int QStringRef_EndsWith2(QtObjectPtr ptr, QtObjectPtr ch, int cs);
int QStringRef_EndsWith3(QtObjectPtr ptr, QtObjectPtr str, int cs);
int QStringRef_EndsWith(QtObjectPtr ptr, char* str, int cs);
int QStringRef_EndsWith4(QtObjectPtr ptr, QtObjectPtr str, int cs);
int QStringRef_IndexOf3(QtObjectPtr ptr, QtObjectPtr ch, int from, int cs);
int QStringRef_IndexOf2(QtObjectPtr ptr, QtObjectPtr str, int from, int cs);
int QStringRef_IndexOf(QtObjectPtr ptr, char* str, int from, int cs);
int QStringRef_IndexOf4(QtObjectPtr ptr, QtObjectPtr str, int from, int cs);
int QStringRef_IsEmpty(QtObjectPtr ptr);
int QStringRef_IsNull(QtObjectPtr ptr);
int QStringRef_LastIndexOf2(QtObjectPtr ptr, QtObjectPtr ch, int from, int cs);
int QStringRef_LastIndexOf3(QtObjectPtr ptr, QtObjectPtr str, int from, int cs);
int QStringRef_LastIndexOf(QtObjectPtr ptr, char* str, int from, int cs);
int QStringRef_LastIndexOf4(QtObjectPtr ptr, QtObjectPtr str, int from, int cs);
int QStringRef_Length(QtObjectPtr ptr);
int QStringRef_QStringRef_LocaleAwareCompare(QtObjectPtr s1, char* s2);
int QStringRef_QStringRef_LocaleAwareCompare2(QtObjectPtr s1, QtObjectPtr s2);
int QStringRef_LocaleAwareCompare3(QtObjectPtr ptr, char* other);
int QStringRef_LocaleAwareCompare4(QtObjectPtr ptr, QtObjectPtr other);
int QStringRef_Position(QtObjectPtr ptr);
int QStringRef_Size(QtObjectPtr ptr);
int QStringRef_StartsWith4(QtObjectPtr ptr, QtObjectPtr ch, int cs);
int QStringRef_StartsWith2(QtObjectPtr ptr, QtObjectPtr str, int cs);
int QStringRef_StartsWith(QtObjectPtr ptr, char* str, int cs);
int QStringRef_StartsWith3(QtObjectPtr ptr, QtObjectPtr str, int cs);
char* QStringRef_String(QtObjectPtr ptr);
int QStringRef_ToInt(QtObjectPtr ptr, int ok, int base);
char* QStringRef_ToString(QtObjectPtr ptr);
QtObjectPtr QStringRef_Unicode(QtObjectPtr ptr);
void QStringRef_DestroyQStringRef(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif