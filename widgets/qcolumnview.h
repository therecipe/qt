#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QColumnView_ResizeGripsVisible(QtObjectPtr ptr);
void QColumnView_SetResizeGripsVisible(QtObjectPtr ptr, int visible);
QtObjectPtr QColumnView_NewQColumnView(QtObjectPtr parent);
QtObjectPtr QColumnView_IndexAt(QtObjectPtr ptr, QtObjectPtr point);
QtObjectPtr QColumnView_PreviewWidget(QtObjectPtr ptr);
void QColumnView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint);
void QColumnView_SelectAll(QtObjectPtr ptr);
void QColumnView_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QColumnView_SetPreviewWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QColumnView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index);
void QColumnView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr newSelectionModel);
void QColumnView_ConnectUpdatePreviewWidget(QtObjectPtr ptr);
void QColumnView_DisconnectUpdatePreviewWidget(QtObjectPtr ptr);
void QColumnView_DestroyQColumnView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif