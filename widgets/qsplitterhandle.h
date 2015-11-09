#ifdef __cplusplus
extern "C" {
#endif

void* QSplitterHandle_NewQSplitterHandle(int orientation, void* parent);
int QSplitterHandle_OpaqueResize(void* ptr);
int QSplitterHandle_Orientation(void* ptr);
void QSplitterHandle_SetOrientation(void* ptr, int orientation);
void* QSplitterHandle_Splitter(void* ptr);
void QSplitterHandle_DestroyQSplitterHandle(void* ptr);

#ifdef __cplusplus
}
#endif