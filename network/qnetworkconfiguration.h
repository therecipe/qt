#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkConfiguration_NewQNetworkConfiguration();
QtObjectPtr QNetworkConfiguration_NewQNetworkConfiguration2(QtObjectPtr other);
int QNetworkConfiguration_BearerType(QtObjectPtr ptr);
int QNetworkConfiguration_BearerTypeFamily(QtObjectPtr ptr);
char* QNetworkConfiguration_BearerTypeName(QtObjectPtr ptr);
char* QNetworkConfiguration_Identifier(QtObjectPtr ptr);
int QNetworkConfiguration_IsRoamingAvailable(QtObjectPtr ptr);
int QNetworkConfiguration_IsValid(QtObjectPtr ptr);
char* QNetworkConfiguration_Name(QtObjectPtr ptr);
int QNetworkConfiguration_Purpose(QtObjectPtr ptr);
void QNetworkConfiguration_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QNetworkConfiguration_Type(QtObjectPtr ptr);
void QNetworkConfiguration_DestroyQNetworkConfiguration(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif