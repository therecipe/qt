#ifdef __cplusplus
extern "C" {
#endif

void QJsonArray_Append(void* ptr, void* value);
int QJsonArray_Contains(void* ptr, void* value);
int QJsonArray_Count(void* ptr);
int QJsonArray_Empty(void* ptr);
void* QJsonArray_QJsonArray_FromStringList(char* list);
int QJsonArray_IsEmpty(void* ptr);
void QJsonArray_Pop_back(void* ptr);
void QJsonArray_Pop_front(void* ptr);
void QJsonArray_Prepend(void* ptr, void* value);
void QJsonArray_Push_back(void* ptr, void* value);
void QJsonArray_Push_front(void* ptr, void* value);
void QJsonArray_RemoveAt(void* ptr, int i);
void QJsonArray_RemoveFirst(void* ptr);
void QJsonArray_RemoveLast(void* ptr);
int QJsonArray_Size(void* ptr);
void QJsonArray_DestroyQJsonArray(void* ptr);

#ifdef __cplusplus
}
#endif