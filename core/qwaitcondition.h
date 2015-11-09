#ifdef __cplusplus
extern "C" {
#endif

void* QWaitCondition_NewQWaitCondition();
void QWaitCondition_WakeAll(void* ptr);
void QWaitCondition_WakeOne(void* ptr);
void QWaitCondition_DestroyQWaitCondition(void* ptr);

#ifdef __cplusplus
}
#endif