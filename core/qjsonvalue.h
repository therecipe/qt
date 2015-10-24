#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QJsonValue_NewQJsonValue5(QtObjectPtr s);
QtObjectPtr QJsonValue_NewQJsonValue(int ty);
QtObjectPtr QJsonValue_NewQJsonValue2(int b);
QtObjectPtr QJsonValue_NewQJsonValue7(QtObjectPtr a);
QtObjectPtr QJsonValue_NewQJsonValue8(QtObjectPtr o);
QtObjectPtr QJsonValue_NewQJsonValue9(QtObjectPtr other);
QtObjectPtr QJsonValue_NewQJsonValue4(char* s);
QtObjectPtr QJsonValue_NewQJsonValue6(char* s);
QtObjectPtr QJsonValue_NewQJsonValue12(int n);
int QJsonValue_IsArray(QtObjectPtr ptr);
int QJsonValue_IsBool(QtObjectPtr ptr);
int QJsonValue_IsDouble(QtObjectPtr ptr);
int QJsonValue_IsNull(QtObjectPtr ptr);
int QJsonValue_IsObject(QtObjectPtr ptr);
int QJsonValue_IsString(QtObjectPtr ptr);
int QJsonValue_IsUndefined(QtObjectPtr ptr);
int QJsonValue_ToBool(QtObjectPtr ptr, int defaultValue);
int QJsonValue_ToInt(QtObjectPtr ptr, int defaultValue);
char* QJsonValue_ToString(QtObjectPtr ptr, char* defaultValue);
char* QJsonValue_ToVariant(QtObjectPtr ptr);
int QJsonValue_Type(QtObjectPtr ptr);
void QJsonValue_DestroyQJsonValue(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif