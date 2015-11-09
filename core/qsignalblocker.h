#ifdef __cplusplus
extern "C" {
#endif

void QSignalBlocker_Reblock(void* ptr);
void QSignalBlocker_Unblock(void* ptr);
void QSignalBlocker_DestroyQSignalBlocker(void* ptr);

#ifdef __cplusplus
}
#endif