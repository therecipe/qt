#ifdef __cplusplus
extern "C" {
#endif

void* QScrollEvent_NewQScrollEvent(void* contentPos, void* overshootDistance, int scrollState);
int QScrollEvent_ScrollState(void* ptr);
void QScrollEvent_DestroyQScrollEvent(void* ptr);

#ifdef __cplusplus
}
#endif