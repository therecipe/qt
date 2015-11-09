#ifdef __cplusplus
extern "C" {
#endif

int QPagedPaintDevice_NewPage(void* ptr);
int QPagedPaintDevice_PageSize(void* ptr);
int QPagedPaintDevice_SetPageLayout(void* ptr, void* newPageLayout);
int QPagedPaintDevice_SetPageMargins(void* ptr, void* margins);
int QPagedPaintDevice_SetPageMargins2(void* ptr, void* margins, int units);
int QPagedPaintDevice_SetPageOrientation(void* ptr, int orientation);
int QPagedPaintDevice_SetPageSize(void* ptr, void* pageSize);
void QPagedPaintDevice_SetPageSize2(void* ptr, int size);
void QPagedPaintDevice_SetPageSizeMM(void* ptr, void* size);
void QPagedPaintDevice_DestroyQPagedPaintDevice(void* ptr);

#ifdef __cplusplus
}
#endif