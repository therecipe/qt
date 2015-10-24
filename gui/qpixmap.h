#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPixmap_Depth(QtObjectPtr ptr);
int QPixmap_Height(QtObjectPtr ptr);
int QPixmap_IsNull(QtObjectPtr ptr);
int QPixmap_IsQBitmap(QtObjectPtr ptr);
int QPixmap_Width(QtObjectPtr ptr);
int QPixmap_ConvertFromImage(QtObjectPtr ptr, QtObjectPtr image, int flags);
int QPixmap_QPixmap_DefaultDepth();
void QPixmap_Detach(QtObjectPtr ptr);
void QPixmap_Fill(QtObjectPtr ptr, QtObjectPtr color);
int QPixmap_HasAlpha(QtObjectPtr ptr);
int QPixmap_HasAlphaChannel(QtObjectPtr ptr);
int QPixmap_Load(QtObjectPtr ptr, char* fileName, char* format, int flags);
int QPixmap_LoadFromData2(QtObjectPtr ptr, QtObjectPtr data, char* format, int flags);
int QPixmap_Save2(QtObjectPtr ptr, QtObjectPtr device, char* format, int quality);
int QPixmap_Save(QtObjectPtr ptr, char* fileName, char* format, int quality);
void QPixmap_Scroll2(QtObjectPtr ptr, int dx, int dy, QtObjectPtr rect, QtObjectPtr exposed);
void QPixmap_Scroll(QtObjectPtr ptr, int dx, int dy, int x, int y, int width, int height, QtObjectPtr exposed);
void QPixmap_SetMask(QtObjectPtr ptr, QtObjectPtr mask);
void QPixmap_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPixmap_DestroyQPixmap(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif