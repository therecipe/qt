#ifdef __cplusplus
extern "C" {
#endif

void* QBackingStore_PaintDevice(void* ptr);
void* QBackingStore_NewQBackingStore(void* window);
void QBackingStore_BeginPaint(void* ptr, void* region);
void QBackingStore_EndPaint(void* ptr);
void QBackingStore_Flush(void* ptr, void* region, void* win, void* offset);
int QBackingStore_HasStaticContents(void* ptr);
void QBackingStore_Resize(void* ptr, void* size);
int QBackingStore_Scroll(void* ptr, void* area, int dx, int dy);
void QBackingStore_SetStaticContents(void* ptr, void* region);
void* QBackingStore_StaticContents(void* ptr);
void* QBackingStore_Window(void* ptr);
void QBackingStore_DestroyQBackingStore(void* ptr);

#ifdef __cplusplus
}
#endif