#ifdef __cplusplus
extern "C" {
#endif

void QByteArray_Clear(void* ptr);
int QByteArray_IndexOf3(void* ptr, char* str, int from);
int QByteArray_IsNull(void* ptr);
int QByteArray_LastIndexOf(void* ptr, void* ba, int from);
int QByteArray_LastIndexOf3(void* ptr, char* str, int from);
void* QByteArray_NewQByteArray();
void* QByteArray_NewQByteArray6(void* other);
void* QByteArray_NewQByteArray5(void* other);
void* QByteArray_NewQByteArray2(char* data, int size);
void* QByteArray_NewQByteArray3(int size, char* ch);
void* QByteArray_Append5(void* ptr, char* ch);
void* QByteArray_Append(void* ptr, void* ba);
void* QByteArray_Append2(void* ptr, char* str);
void* QByteArray_Append3(void* ptr, char* str);
void* QByteArray_Append4(void* ptr, char* str, int len);
int QByteArray_Capacity(void* ptr);
void QByteArray_Chop(void* ptr, int n);
int QByteArray_Contains3(void* ptr, char* ch);
int QByteArray_Contains(void* ptr, void* ba);
int QByteArray_Contains2(void* ptr, char* str);
int QByteArray_Count4(void* ptr);
int QByteArray_Count3(void* ptr, char* ch);
int QByteArray_Count(void* ptr, void* ba);
int QByteArray_Count2(void* ptr, char* str);
int QByteArray_EndsWith3(void* ptr, char* ch);
int QByteArray_EndsWith(void* ptr, void* ba);
int QByteArray_EndsWith2(void* ptr, char* str);
void* QByteArray_Fill(void* ptr, char* ch, int size);
void* QByteArray_QByteArray_FromBase64(void* base64);
void* QByteArray_QByteArray_FromBase642(void* base64, int options);
void* QByteArray_QByteArray_FromHex(void* hexEncoded);
void* QByteArray_QByteArray_FromPercentEncoding(void* input, char* percent);
void* QByteArray_QByteArray_FromRawData(char* data, int size);
int QByteArray_IndexOf4(void* ptr, char* ch, int from);
int QByteArray_IndexOf(void* ptr, void* ba, int from);
int QByteArray_IndexOf2(void* ptr, char* str, int from);
int QByteArray_IsEmpty(void* ptr);
int QByteArray_LastIndexOf4(void* ptr, char* ch, int from);
int QByteArray_LastIndexOf2(void* ptr, char* str, int from);
void* QByteArray_Left(void* ptr, int len);
void* QByteArray_LeftJustified(void* ptr, int width, char* fill, int truncate);
int QByteArray_Length(void* ptr);
void* QByteArray_Mid(void* ptr, int pos, int len);
void* QByteArray_QByteArray_Number(int n, int base);
void* QByteArray_Prepend4(void* ptr, char* ch);
void* QByteArray_Prepend(void* ptr, void* ba);
void* QByteArray_Prepend2(void* ptr, char* str);
void* QByteArray_Prepend3(void* ptr, char* str, int len);
void QByteArray_Push_back3(void* ptr, char* ch);
void QByteArray_Push_back(void* ptr, void* other);
void QByteArray_Push_back2(void* ptr, char* str);
void QByteArray_Push_front3(void* ptr, char* ch);
void QByteArray_Push_front(void* ptr, void* other);
void QByteArray_Push_front2(void* ptr, char* str);
void* QByteArray_Repeated(void* ptr, int times);
void QByteArray_Reserve(void* ptr, int size);
void QByteArray_Resize(void* ptr, int size);
void* QByteArray_Right(void* ptr, int len);
void* QByteArray_RightJustified(void* ptr, int width, char* fill, int truncate);
void* QByteArray_SetNum(void* ptr, int n, int base);
int QByteArray_Size(void* ptr);
void QByteArray_Squeeze(void* ptr);
int QByteArray_StartsWith3(void* ptr, char* ch);
int QByteArray_StartsWith(void* ptr, void* ba);
int QByteArray_StartsWith2(void* ptr, char* str);
void QByteArray_Swap(void* ptr, void* other);
void* QByteArray_ToBase64(void* ptr);
void* QByteArray_ToBase642(void* ptr, int options);
void* QByteArray_ToHex(void* ptr);
int QByteArray_ToInt(void* ptr, int ok, int base);
void* QByteArray_ToPercentEncoding(void* ptr, void* exclude, void* include, char* percent);
void QByteArray_Truncate(void* ptr, int pos);
void QByteArray_DestroyQByteArray(void* ptr);
void* QByteArray_Simplified(void* ptr);
void* QByteArray_ToLower(void* ptr);
void* QByteArray_ToUpper(void* ptr);
void* QByteArray_Trimmed(void* ptr);

#ifdef __cplusplus
}
#endif