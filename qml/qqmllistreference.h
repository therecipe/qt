#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlListReference_NewQQmlListReference();
QtObjectPtr QQmlListReference_NewQQmlListReference2(QtObjectPtr object, char* property, QtObjectPtr engine);
int QQmlListReference_Append(QtObjectPtr ptr, QtObjectPtr object);
QtObjectPtr QQmlListReference_At(QtObjectPtr ptr, int index);
int QQmlListReference_CanAppend(QtObjectPtr ptr);
int QQmlListReference_CanAt(QtObjectPtr ptr);
int QQmlListReference_CanClear(QtObjectPtr ptr);
int QQmlListReference_CanCount(QtObjectPtr ptr);
int QQmlListReference_Clear(QtObjectPtr ptr);
int QQmlListReference_Count(QtObjectPtr ptr);
int QQmlListReference_IsManipulable(QtObjectPtr ptr);
int QQmlListReference_IsReadable(QtObjectPtr ptr);
int QQmlListReference_IsValid(QtObjectPtr ptr);
QtObjectPtr QQmlListReference_ListElementType(QtObjectPtr ptr);
QtObjectPtr QQmlListReference_Object(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif