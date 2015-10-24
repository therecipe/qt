#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QByteArray_Clear(QtObjectPtr ptr);
int QByteArray_IndexOf3(QtObjectPtr ptr, char* str, int from);
int QByteArray_IsNull(QtObjectPtr ptr);
int QByteArray_LastIndexOf(QtObjectPtr ptr, QtObjectPtr ba, int from);
int QByteArray_LastIndexOf3(QtObjectPtr ptr, char* str, int from);
QtObjectPtr QByteArray_NewQByteArray();
QtObjectPtr QByteArray_NewQByteArray6(QtObjectPtr other);
QtObjectPtr QByteArray_NewQByteArray5(QtObjectPtr other);
QtObjectPtr QByteArray_NewQByteArray2(char* data, int size);
QtObjectPtr QByteArray_NewQByteArray3(int size, char* ch);
int QByteArray_Capacity(QtObjectPtr ptr);
void QByteArray_Chop(QtObjectPtr ptr, int n);
int QByteArray_Contains3(QtObjectPtr ptr, char* ch);
int QByteArray_Contains(QtObjectPtr ptr, QtObjectPtr ba);
int QByteArray_Contains2(QtObjectPtr ptr, char* str);
int QByteArray_Count4(QtObjectPtr ptr);
int QByteArray_Count3(QtObjectPtr ptr, char* ch);
int QByteArray_Count(QtObjectPtr ptr, QtObjectPtr ba);
int QByteArray_Count2(QtObjectPtr ptr, char* str);
int QByteArray_EndsWith3(QtObjectPtr ptr, char* ch);
int QByteArray_EndsWith(QtObjectPtr ptr, QtObjectPtr ba);
int QByteArray_EndsWith2(QtObjectPtr ptr, char* str);
int QByteArray_IndexOf4(QtObjectPtr ptr, char* ch, int from);
int QByteArray_IndexOf(QtObjectPtr ptr, QtObjectPtr ba, int from);
int QByteArray_IndexOf2(QtObjectPtr ptr, char* str, int from);
int QByteArray_IsEmpty(QtObjectPtr ptr);
int QByteArray_LastIndexOf4(QtObjectPtr ptr, char* ch, int from);
int QByteArray_LastIndexOf2(QtObjectPtr ptr, char* str, int from);
int QByteArray_Length(QtObjectPtr ptr);
void QByteArray_Push_back3(QtObjectPtr ptr, char* ch);
void QByteArray_Push_back(QtObjectPtr ptr, QtObjectPtr other);
void QByteArray_Push_back2(QtObjectPtr ptr, char* str);
void QByteArray_Push_front3(QtObjectPtr ptr, char* ch);
void QByteArray_Push_front(QtObjectPtr ptr, QtObjectPtr other);
void QByteArray_Push_front2(QtObjectPtr ptr, char* str);
void QByteArray_Reserve(QtObjectPtr ptr, int size);
void QByteArray_Resize(QtObjectPtr ptr, int size);
int QByteArray_Size(QtObjectPtr ptr);
void QByteArray_Squeeze(QtObjectPtr ptr);
int QByteArray_StartsWith3(QtObjectPtr ptr, char* ch);
int QByteArray_StartsWith(QtObjectPtr ptr, QtObjectPtr ba);
int QByteArray_StartsWith2(QtObjectPtr ptr, char* str);
void QByteArray_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QByteArray_ToInt(QtObjectPtr ptr, int ok, int base);
void QByteArray_Truncate(QtObjectPtr ptr, int pos);
void QByteArray_DestroyQByteArray(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif