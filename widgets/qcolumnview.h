#ifdef __cplusplus
extern "C" {
#endif

int QColumnView_ResizeGripsVisible(void* ptr);
void QColumnView_SetResizeGripsVisible(void* ptr, int visible);
void* QColumnView_NewQColumnView(void* parent);
void* QColumnView_IndexAt(void* ptr, void* point);
void* QColumnView_PreviewWidget(void* ptr);
void QColumnView_ScrollTo(void* ptr, void* index, int hint);
void QColumnView_SelectAll(void* ptr);
void QColumnView_SetModel(void* ptr, void* model);
void QColumnView_SetPreviewWidget(void* ptr, void* widget);
void QColumnView_SetRootIndex(void* ptr, void* index);
void QColumnView_SetSelectionModel(void* ptr, void* newSelectionModel);
void QColumnView_ConnectUpdatePreviewWidget(void* ptr);
void QColumnView_DisconnectUpdatePreviewWidget(void* ptr);
void QColumnView_DestroyQColumnView(void* ptr);

#ifdef __cplusplus
}
#endif