#ifdef __cplusplus
extern "C" {
#endif

void* QMetaType_NewQMetaType(int typeId);
int QMetaType_Flags(void* ptr);
int QMetaType_QMetaType_IsRegistered(int ty);
int QMetaType_IsRegistered2(void* ptr);
int QMetaType_IsValid(void* ptr);
void* QMetaType_MetaObject(void* ptr);
void* QMetaType_QMetaType_MetaObjectForType(int ty);
int QMetaType_QMetaType_SizeOf(int ty);
int QMetaType_SizeOf2(void* ptr);
int QMetaType_QMetaType_Type2(void* typeName);
int QMetaType_QMetaType_Type(char* typeName);
int QMetaType_QMetaType_TypeFlags(int ty);
void QMetaType_DestroyQMetaType(void* ptr);

#ifdef __cplusplus
}
#endif