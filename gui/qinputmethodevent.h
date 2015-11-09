#ifdef __cplusplus
extern "C" {
#endif

void* QInputMethodEvent_NewQInputMethodEvent();
void* QInputMethodEvent_NewQInputMethodEvent3(void* other);
char* QInputMethodEvent_CommitString(void* ptr);
char* QInputMethodEvent_PreeditString(void* ptr);
int QInputMethodEvent_ReplacementLength(void* ptr);
int QInputMethodEvent_ReplacementStart(void* ptr);
void QInputMethodEvent_SetCommitString(void* ptr, char* commitString, int replaceFrom, int replaceLength);

#ifdef __cplusplus
}
#endif