#ifdef __cplusplus
extern "C" {
#endif

void* QTouchDevice_NewQTouchDevice();
int QTouchDevice_Capabilities(void* ptr);
int QTouchDevice_MaximumTouchPoints(void* ptr);
char* QTouchDevice_Name(void* ptr);
void QTouchDevice_SetCapabilities(void* ptr, int caps);
void QTouchDevice_SetMaximumTouchPoints(void* ptr, int max);
void QTouchDevice_SetName(void* ptr, char* name);
void QTouchDevice_SetType(void* ptr, int devType);
int QTouchDevice_Type(void* ptr);
void QTouchDevice_DestroyQTouchDevice(void* ptr);

#ifdef __cplusplus
}
#endif