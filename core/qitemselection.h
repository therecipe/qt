#ifdef __cplusplus
extern "C" {
#endif

void* QItemSelection_NewQItemSelection();
void* QItemSelection_NewQItemSelection2(void* topLeft, void* bottomRight);
int QItemSelection_Contains(void* ptr, void* index);
void QItemSelection_Merge(void* ptr, void* other, int command);
void QItemSelection_Select(void* ptr, void* topLeft, void* bottomRight);
void QItemSelection_QItemSelection_Split(void* ran, void* other, void* result);

#ifdef __cplusplus
}
#endif