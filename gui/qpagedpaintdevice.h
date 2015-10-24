#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPagedPaintDevice_NewPage(QtObjectPtr ptr);
int QPagedPaintDevice_PageSize(QtObjectPtr ptr);
int QPagedPaintDevice_SetPageLayout(QtObjectPtr ptr, QtObjectPtr newPageLayout);
int QPagedPaintDevice_SetPageMargins(QtObjectPtr ptr, QtObjectPtr margins);
int QPagedPaintDevice_SetPageMargins2(QtObjectPtr ptr, QtObjectPtr margins, int units);
int QPagedPaintDevice_SetPageOrientation(QtObjectPtr ptr, int orientation);
int QPagedPaintDevice_SetPageSize(QtObjectPtr ptr, QtObjectPtr pageSize);
void QPagedPaintDevice_SetPageSize2(QtObjectPtr ptr, int size);
void QPagedPaintDevice_SetPageSizeMM(QtObjectPtr ptr, QtObjectPtr size);
void QPagedPaintDevice_DestroyQPagedPaintDevice(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif