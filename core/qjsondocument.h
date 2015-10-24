#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QJsonDocument_NewQJsonDocument();
QtObjectPtr QJsonDocument_NewQJsonDocument3(QtObjectPtr array);
QtObjectPtr QJsonDocument_NewQJsonDocument4(QtObjectPtr other);
QtObjectPtr QJsonDocument_NewQJsonDocument2(QtObjectPtr object);
int QJsonDocument_IsArray(QtObjectPtr ptr);
int QJsonDocument_IsEmpty(QtObjectPtr ptr);
int QJsonDocument_IsNull(QtObjectPtr ptr);
int QJsonDocument_IsObject(QtObjectPtr ptr);
void QJsonDocument_SetArray(QtObjectPtr ptr, QtObjectPtr array);
void QJsonDocument_SetObject(QtObjectPtr ptr, QtObjectPtr object);
char* QJsonDocument_ToVariant(QtObjectPtr ptr);
void QJsonDocument_DestroyQJsonDocument(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif