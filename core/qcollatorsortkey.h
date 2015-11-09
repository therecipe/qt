#ifdef __cplusplus
extern "C" {
#endif

void* QCollatorSortKey_NewQCollatorSortKey(void* other);
void QCollatorSortKey_Swap(void* ptr, void* other);
void QCollatorSortKey_DestroyQCollatorSortKey(void* ptr);
int QCollatorSortKey_Compare(void* ptr, void* otherKey);

#ifdef __cplusplus
}
#endif