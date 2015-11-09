#ifdef __cplusplus
extern "C" {
#endif

void* QShortcutEvent_NewQShortcutEvent(void* key, int id, int ambiguous);
int QShortcutEvent_IsAmbiguous(void* ptr);
int QShortcutEvent_ShortcutId(void* ptr);
void QShortcutEvent_DestroyQShortcutEvent(void* ptr);

#ifdef __cplusplus
}
#endif