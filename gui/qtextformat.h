#ifdef __cplusplus
extern "C" {
#endif

void* QTextFormat_NewQTextFormat3(void* other);
void QTextFormat_SetObjectIndex(void* ptr, int index);
void* QTextFormat_NewQTextFormat();
void* QTextFormat_NewQTextFormat2(int ty);
void* QTextFormat_Background(void* ptr);
int QTextFormat_BoolProperty(void* ptr, int propertyId);
void* QTextFormat_BrushProperty(void* ptr, int propertyId);
void QTextFormat_ClearBackground(void* ptr);
void QTextFormat_ClearForeground(void* ptr);
void QTextFormat_ClearProperty(void* ptr, int propertyId);
void* QTextFormat_ColorProperty(void* ptr, int propertyId);
double QTextFormat_DoubleProperty(void* ptr, int propertyId);
void* QTextFormat_Foreground(void* ptr);
int QTextFormat_HasProperty(void* ptr, int propertyId);
int QTextFormat_IntProperty(void* ptr, int propertyId);
int QTextFormat_IsBlockFormat(void* ptr);
int QTextFormat_IsCharFormat(void* ptr);
int QTextFormat_IsEmpty(void* ptr);
int QTextFormat_IsFrameFormat(void* ptr);
int QTextFormat_IsImageFormat(void* ptr);
int QTextFormat_IsListFormat(void* ptr);
int QTextFormat_IsTableCellFormat(void* ptr);
int QTextFormat_IsTableFormat(void* ptr);
int QTextFormat_IsValid(void* ptr);
int QTextFormat_LayoutDirection(void* ptr);
void QTextFormat_Merge(void* ptr, void* other);
int QTextFormat_ObjectIndex(void* ptr);
int QTextFormat_ObjectType(void* ptr);
void* QTextFormat_Property(void* ptr, int propertyId);
int QTextFormat_PropertyCount(void* ptr);
void QTextFormat_SetBackground(void* ptr, void* brush);
void QTextFormat_SetForeground(void* ptr, void* brush);
void QTextFormat_SetLayoutDirection(void* ptr, int direction);
void QTextFormat_SetObjectType(void* ptr, int ty);
void QTextFormat_SetProperty(void* ptr, int propertyId, void* value);
char* QTextFormat_StringProperty(void* ptr, int propertyId);
void QTextFormat_Swap(void* ptr, void* other);
int QTextFormat_Type(void* ptr);
void QTextFormat_DestroyQTextFormat(void* ptr);

#ifdef __cplusplus
}
#endif