#ifdef __cplusplus
extern "C" {
#endif

void QSocketNotifier_ConnectActivated(void* ptr);
void QSocketNotifier_DisconnectActivated(void* ptr);
int QSocketNotifier_IsEnabled(void* ptr);
void QSocketNotifier_SetEnabled(void* ptr, int enable);
int QSocketNotifier_Type(void* ptr);
void QSocketNotifier_DestroyQSocketNotifier(void* ptr);

#ifdef __cplusplus
}
#endif