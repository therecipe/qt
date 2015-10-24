#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNdefRecord_NewQNdefRecord();
QtObjectPtr QNdefRecord_NewQNdefRecord2(QtObjectPtr other);
int QNdefRecord_IsEmpty(QtObjectPtr ptr);
void QNdefRecord_SetId(QtObjectPtr ptr, QtObjectPtr id);
void QNdefRecord_SetPayload(QtObjectPtr ptr, QtObjectPtr payload);
void QNdefRecord_SetType(QtObjectPtr ptr, QtObjectPtr ty);
void QNdefRecord_SetTypeNameFormat(QtObjectPtr ptr, int typeNameFormat);
int QNdefRecord_TypeNameFormat(QtObjectPtr ptr);
void QNdefRecord_DestroyQNdefRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif