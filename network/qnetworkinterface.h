#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNetworkInterface_NewQNetworkInterface();
QtObjectPtr QNetworkInterface_NewQNetworkInterface2(QtObjectPtr other);
int QNetworkInterface_Flags(QtObjectPtr ptr);
char* QNetworkInterface_HardwareAddress(QtObjectPtr ptr);
char* QNetworkInterface_HumanReadableName(QtObjectPtr ptr);
int QNetworkInterface_Index(QtObjectPtr ptr);
int QNetworkInterface_IsValid(QtObjectPtr ptr);
char* QNetworkInterface_Name(QtObjectPtr ptr);
void QNetworkInterface_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QNetworkInterface_DestroyQNetworkInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif