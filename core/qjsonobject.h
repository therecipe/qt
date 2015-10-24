#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QJsonObject_Contains(QtObjectPtr ptr, char* key);
int QJsonObject_Count(QtObjectPtr ptr);
int QJsonObject_Empty(QtObjectPtr ptr);
int QJsonObject_IsEmpty(QtObjectPtr ptr);
char* QJsonObject_Keys(QtObjectPtr ptr);
int QJsonObject_Length(QtObjectPtr ptr);
int QJsonObject_Size(QtObjectPtr ptr);
void QJsonObject_DestroyQJsonObject(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif