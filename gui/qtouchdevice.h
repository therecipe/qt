#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTouchDevice_NewQTouchDevice();
int QTouchDevice_Capabilities(QtObjectPtr ptr);
int QTouchDevice_MaximumTouchPoints(QtObjectPtr ptr);
char* QTouchDevice_Name(QtObjectPtr ptr);
void QTouchDevice_SetCapabilities(QtObjectPtr ptr, int caps);
void QTouchDevice_SetMaximumTouchPoints(QtObjectPtr ptr, int max);
void QTouchDevice_SetName(QtObjectPtr ptr, char* name);
void QTouchDevice_SetType(QtObjectPtr ptr, int devType);
int QTouchDevice_Type(QtObjectPtr ptr);
void QTouchDevice_DestroyQTouchDevice(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif