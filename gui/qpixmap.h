#ifdef __cplusplus
extern "C" {
#endif

int QPixmap_Depth(void* ptr);
int QPixmap_Height(void* ptr);
int QPixmap_IsNull(void* ptr);
int QPixmap_IsQBitmap(void* ptr);
int QPixmap_Width(void* ptr);
int QPixmap_ConvertFromImage(void* ptr, void* image, int flags);
int QPixmap_QPixmap_DefaultDepth();
void QPixmap_Detach(void* ptr);
double QPixmap_DevicePixelRatio(void* ptr);
void QPixmap_Fill(void* ptr, void* color);
int QPixmap_HasAlpha(void* ptr);
int QPixmap_HasAlphaChannel(void* ptr);
int QPixmap_Load(void* ptr, char* fileName, char* format, int flags);
int QPixmap_LoadFromData2(void* ptr, void* data, char* format, int flags);
int QPixmap_Save2(void* ptr, void* device, char* format, int quality);
int QPixmap_Save(void* ptr, char* fileName, char* format, int quality);
void QPixmap_Scroll2(void* ptr, int dx, int dy, void* rect, void* exposed);
void QPixmap_Scroll(void* ptr, int dx, int dy, int x, int y, int width, int height, void* exposed);
void QPixmap_SetDevicePixelRatio(void* ptr, double scaleFactor);
void QPixmap_SetMask(void* ptr, void* mask);
void QPixmap_Swap(void* ptr, void* other);
void QPixmap_DestroyQPixmap(void* ptr);

#ifdef __cplusplus
}
#endif