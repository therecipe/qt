#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QQmlNdefRecord_TypeNameFormat(QtObjectPtr ptr);
QtObjectPtr QQmlNdefRecord_NewQQmlNdefRecord(QtObjectPtr parent);
QtObjectPtr QQmlNdefRecord_NewQQmlNdefRecord2(QtObjectPtr record, QtObjectPtr parent);
void QQmlNdefRecord_ConnectRecordChanged(QtObjectPtr ptr);
void QQmlNdefRecord_DisconnectRecordChanged(QtObjectPtr ptr);
void QQmlNdefRecord_SetRecord(QtObjectPtr ptr, QtObjectPtr record);
void QQmlNdefRecord_SetType(QtObjectPtr ptr, char* newtype);
void QQmlNdefRecord_SetTypeNameFormat(QtObjectPtr ptr, int newTypeNameFormat);
char* QQmlNdefRecord_Type(QtObjectPtr ptr);
void QQmlNdefRecord_ConnectTypeChanged(QtObjectPtr ptr);
void QQmlNdefRecord_DisconnectTypeChanged(QtObjectPtr ptr);
void QQmlNdefRecord_ConnectTypeNameFormatChanged(QtObjectPtr ptr);
void QQmlNdefRecord_DisconnectTypeNameFormatChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif