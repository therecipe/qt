#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBackingStore_PaintDevice(QtObjectPtr ptr);
QtObjectPtr QBackingStore_NewQBackingStore(QtObjectPtr window);
void QBackingStore_BeginPaint(QtObjectPtr ptr, QtObjectPtr region);
void QBackingStore_EndPaint(QtObjectPtr ptr);
void QBackingStore_Flush(QtObjectPtr ptr, QtObjectPtr region, QtObjectPtr win, QtObjectPtr offset);
int QBackingStore_HasStaticContents(QtObjectPtr ptr);
void QBackingStore_Resize(QtObjectPtr ptr, QtObjectPtr size);
int QBackingStore_Scroll(QtObjectPtr ptr, QtObjectPtr area, int dx, int dy);
void QBackingStore_SetStaticContents(QtObjectPtr ptr, QtObjectPtr region);
QtObjectPtr QBackingStore_Window(QtObjectPtr ptr);
void QBackingStore_DestroyQBackingStore(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif