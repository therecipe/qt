#ifdef __cplusplus
extern "C" {
#endif

void* QNdefFilter_NewQNdefFilter();
void* QNdefFilter_NewQNdefFilter2(void* other);
void QNdefFilter_Clear(void* ptr);
int QNdefFilter_OrderMatch(void* ptr);
int QNdefFilter_RecordCount(void* ptr);
void QNdefFilter_SetOrderMatch(void* ptr, int on);
void QNdefFilter_DestroyQNdefFilter(void* ptr);

#ifdef __cplusplus
}
#endif