#ifdef __cplusplus
extern "C" {
#endif

void* QEventLoopLocker_NewQEventLoopLocker();
void* QEventLoopLocker_NewQEventLoopLocker2(void* loop);
void* QEventLoopLocker_NewQEventLoopLocker3(void* thread);
void QEventLoopLocker_DestroyQEventLoopLocker(void* ptr);

#ifdef __cplusplus
}
#endif