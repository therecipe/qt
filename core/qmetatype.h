#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMetaType_NewQMetaType(int typeId);
int QMetaType_Flags(QtObjectPtr ptr);
int QMetaType_QMetaType_IsRegistered(int ty);
int QMetaType_IsRegistered2(QtObjectPtr ptr);
int QMetaType_IsValid(QtObjectPtr ptr);
QtObjectPtr QMetaType_MetaObject(QtObjectPtr ptr);
QtObjectPtr QMetaType_QMetaType_MetaObjectForType(int ty);
int QMetaType_QMetaType_SizeOf(int ty);
int QMetaType_SizeOf2(QtObjectPtr ptr);
int QMetaType_QMetaType_Type2(QtObjectPtr typeName);
int QMetaType_QMetaType_Type(char* typeName);
int QMetaType_QMetaType_TypeFlags(int ty);
void QMetaType_DestroyQMetaType(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif