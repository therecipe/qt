#ifdef __cplusplus
extern "C" {
#endif

void* QObjectCleanupHandler_NewQObjectCleanupHandler();
void* QObjectCleanupHandler_Add(void* ptr, void* object);
void QObjectCleanupHandler_Clear(void* ptr);
int QObjectCleanupHandler_IsEmpty(void* ptr);
void QObjectCleanupHandler_DestroyQObjectCleanupHandler(void* ptr);

#ifdef __cplusplus
}
#endif