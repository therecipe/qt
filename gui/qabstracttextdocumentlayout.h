#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAbstractTextDocumentLayout_AnchorAt(QtObjectPtr ptr, QtObjectPtr position);
QtObjectPtr QAbstractTextDocumentLayout_Document(QtObjectPtr ptr);
QtObjectPtr QAbstractTextDocumentLayout_HandlerForObject(QtObjectPtr ptr, int objectType);
int QAbstractTextDocumentLayout_PageCount(QtObjectPtr ptr);
void QAbstractTextDocumentLayout_ConnectPageCountChanged(QtObjectPtr ptr);
void QAbstractTextDocumentLayout_DisconnectPageCountChanged(QtObjectPtr ptr);
QtObjectPtr QAbstractTextDocumentLayout_PaintDevice(QtObjectPtr ptr);
void QAbstractTextDocumentLayout_RegisterHandler(QtObjectPtr ptr, int objectType, QtObjectPtr component);
void QAbstractTextDocumentLayout_SetPaintDevice(QtObjectPtr ptr, QtObjectPtr device);
void QAbstractTextDocumentLayout_UnregisterHandler(QtObjectPtr ptr, int objectType, QtObjectPtr component);

#ifdef __cplusplus
}
#endif