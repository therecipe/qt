#ifdef __cplusplus
extern "C" {
#endif

void* QQmlListReference_NewQQmlListReference();
void* QQmlListReference_NewQQmlListReference2(void* object, char* property, void* engine);
int QQmlListReference_Append(void* ptr, void* object);
void* QQmlListReference_At(void* ptr, int index);
int QQmlListReference_CanAppend(void* ptr);
int QQmlListReference_CanAt(void* ptr);
int QQmlListReference_CanClear(void* ptr);
int QQmlListReference_CanCount(void* ptr);
int QQmlListReference_Clear(void* ptr);
int QQmlListReference_Count(void* ptr);
int QQmlListReference_IsManipulable(void* ptr);
int QQmlListReference_IsReadable(void* ptr);
int QQmlListReference_IsValid(void* ptr);
void* QQmlListReference_ListElementType(void* ptr);
void* QQmlListReference_Object(void* ptr);

#ifdef __cplusplus
}
#endif