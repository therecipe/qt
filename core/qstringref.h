#ifdef __cplusplus
extern "C" {
#endif

void* QStringRef_Left(void* ptr, int n);
void* QStringRef_Mid(void* ptr, int position, int n);
void* QStringRef_Right(void* ptr, int n);
void* QStringRef_AppendTo(void* ptr, char* stri);
void* QStringRef_Begin(void* ptr);
void* QStringRef_Cbegin(void* ptr);
void* QStringRef_Cend(void* ptr);
void QStringRef_Clear(void* ptr);
int QStringRef_QStringRef_Compare3(void* s1, void* s2, int cs);
int QStringRef_QStringRef_Compare(void* s1, char* s2, int cs);
int QStringRef_QStringRef_Compare2(void* s1, void* s2, int cs);
int QStringRef_Compare6(void* ptr, void* other, int cs);
int QStringRef_Compare4(void* ptr, char* other, int cs);
int QStringRef_Compare5(void* ptr, void* other, int cs);
void* QStringRef_ConstData(void* ptr);
int QStringRef_Contains2(void* ptr, void* ch, int cs);
int QStringRef_Contains4(void* ptr, void* str, int cs);
int QStringRef_Contains(void* ptr, char* str, int cs);
int QStringRef_Contains3(void* ptr, void* str, int cs);
int QStringRef_Count(void* ptr);
int QStringRef_Count3(void* ptr, void* ch, int cs);
int QStringRef_Count2(void* ptr, char* str, int cs);
int QStringRef_Count4(void* ptr, void* str, int cs);
void* QStringRef_Data(void* ptr);
void* QStringRef_End(void* ptr);
int QStringRef_EndsWith2(void* ptr, void* ch, int cs);
int QStringRef_EndsWith3(void* ptr, void* str, int cs);
int QStringRef_EndsWith(void* ptr, char* str, int cs);
int QStringRef_EndsWith4(void* ptr, void* str, int cs);
int QStringRef_IndexOf3(void* ptr, void* ch, int from, int cs);
int QStringRef_IndexOf2(void* ptr, void* str, int from, int cs);
int QStringRef_IndexOf(void* ptr, char* str, int from, int cs);
int QStringRef_IndexOf4(void* ptr, void* str, int from, int cs);
int QStringRef_IsEmpty(void* ptr);
int QStringRef_IsNull(void* ptr);
int QStringRef_LastIndexOf2(void* ptr, void* ch, int from, int cs);
int QStringRef_LastIndexOf3(void* ptr, void* str, int from, int cs);
int QStringRef_LastIndexOf(void* ptr, char* str, int from, int cs);
int QStringRef_LastIndexOf4(void* ptr, void* str, int from, int cs);
int QStringRef_Length(void* ptr);
int QStringRef_QStringRef_LocaleAwareCompare(void* s1, char* s2);
int QStringRef_QStringRef_LocaleAwareCompare2(void* s1, void* s2);
int QStringRef_LocaleAwareCompare3(void* ptr, char* other);
int QStringRef_LocaleAwareCompare4(void* ptr, void* other);
int QStringRef_Position(void* ptr);
int QStringRef_Size(void* ptr);
int QStringRef_StartsWith4(void* ptr, void* ch, int cs);
int QStringRef_StartsWith2(void* ptr, void* str, int cs);
int QStringRef_StartsWith(void* ptr, char* str, int cs);
int QStringRef_StartsWith3(void* ptr, void* str, int cs);
char* QStringRef_String(void* ptr);
int QStringRef_ToInt(void* ptr, int ok, int base);
void* QStringRef_ToLatin1(void* ptr);
void* QStringRef_ToLocal8Bit(void* ptr);
char* QStringRef_ToString(void* ptr);
void* QStringRef_ToUtf8(void* ptr);
void* QStringRef_Trimmed(void* ptr);
void* QStringRef_Unicode(void* ptr);
void QStringRef_DestroyQStringRef(void* ptr);

#ifdef __cplusplus
}
#endif