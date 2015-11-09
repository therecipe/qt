#ifdef __cplusplus
extern "C" {
#endif

void* QInputMethodQueryEvent_NewQInputMethodQueryEvent(int queries);
int QInputMethodQueryEvent_Queries(void* ptr);
void QInputMethodQueryEvent_SetValue(void* ptr, int query, void* value);
void* QInputMethodQueryEvent_Value(void* ptr, int query);

#ifdef __cplusplus
}
#endif