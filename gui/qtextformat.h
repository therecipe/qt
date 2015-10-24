#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextFormat_NewQTextFormat3(QtObjectPtr other);
void QTextFormat_SetObjectIndex(QtObjectPtr ptr, int index);
QtObjectPtr QTextFormat_NewQTextFormat();
QtObjectPtr QTextFormat_NewQTextFormat2(int ty);
int QTextFormat_BoolProperty(QtObjectPtr ptr, int propertyId);
void QTextFormat_ClearBackground(QtObjectPtr ptr);
void QTextFormat_ClearForeground(QtObjectPtr ptr);
void QTextFormat_ClearProperty(QtObjectPtr ptr, int propertyId);
int QTextFormat_HasProperty(QtObjectPtr ptr, int propertyId);
int QTextFormat_IntProperty(QtObjectPtr ptr, int propertyId);
int QTextFormat_IsBlockFormat(QtObjectPtr ptr);
int QTextFormat_IsCharFormat(QtObjectPtr ptr);
int QTextFormat_IsEmpty(QtObjectPtr ptr);
int QTextFormat_IsFrameFormat(QtObjectPtr ptr);
int QTextFormat_IsImageFormat(QtObjectPtr ptr);
int QTextFormat_IsListFormat(QtObjectPtr ptr);
int QTextFormat_IsTableCellFormat(QtObjectPtr ptr);
int QTextFormat_IsTableFormat(QtObjectPtr ptr);
int QTextFormat_IsValid(QtObjectPtr ptr);
int QTextFormat_LayoutDirection(QtObjectPtr ptr);
void QTextFormat_Merge(QtObjectPtr ptr, QtObjectPtr other);
int QTextFormat_ObjectIndex(QtObjectPtr ptr);
int QTextFormat_ObjectType(QtObjectPtr ptr);
char* QTextFormat_Property(QtObjectPtr ptr, int propertyId);
int QTextFormat_PropertyCount(QtObjectPtr ptr);
void QTextFormat_SetBackground(QtObjectPtr ptr, QtObjectPtr brush);
void QTextFormat_SetForeground(QtObjectPtr ptr, QtObjectPtr brush);
void QTextFormat_SetLayoutDirection(QtObjectPtr ptr, int direction);
void QTextFormat_SetObjectType(QtObjectPtr ptr, int ty);
void QTextFormat_SetProperty(QtObjectPtr ptr, int propertyId, char* value);
char* QTextFormat_StringProperty(QtObjectPtr ptr, int propertyId);
void QTextFormat_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QTextFormat_Type(QtObjectPtr ptr);
void QTextFormat_DestroyQTextFormat(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif