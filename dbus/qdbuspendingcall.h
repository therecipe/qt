#ifdef __cplusplus
extern "C" {
#endif

void* QDBusPendingCall_NewQDBusPendingCall(void* other);
void QDBusPendingCall_Swap(void* ptr, void* other);
void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr);

#ifdef __cplusplus
}
#endif