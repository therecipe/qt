#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QJsonArray_Append(QtObjectPtr ptr, QtObjectPtr value);
int QJsonArray_Contains(QtObjectPtr ptr, QtObjectPtr value);
int QJsonArray_Count(QtObjectPtr ptr);
int QJsonArray_Empty(QtObjectPtr ptr);
int QJsonArray_IsEmpty(QtObjectPtr ptr);
void QJsonArray_Pop_back(QtObjectPtr ptr);
void QJsonArray_Pop_front(QtObjectPtr ptr);
void QJsonArray_Prepend(QtObjectPtr ptr, QtObjectPtr value);
void QJsonArray_Push_back(QtObjectPtr ptr, QtObjectPtr value);
void QJsonArray_Push_front(QtObjectPtr ptr, QtObjectPtr value);
void QJsonArray_RemoveAt(QtObjectPtr ptr, int i);
void QJsonArray_RemoveFirst(QtObjectPtr ptr);
void QJsonArray_RemoveLast(QtObjectPtr ptr);
int QJsonArray_Size(QtObjectPtr ptr);
void QJsonArray_DestroyQJsonArray(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif