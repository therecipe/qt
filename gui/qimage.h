#ifdef __cplusplus
extern "C" {
#endif

int QImage_ColorCount(void* ptr);
void QImage_Fill2(void* ptr, int color);
void QImage_Fill3(void* ptr, void* color);
int QImage_Height(void* ptr);
int QImage_IsNull(void* ptr);
void QImage_SetOffset(void* ptr, void* offset);
void QImage_SetText(void* ptr, char* key, char* text);
int QImage_Width(void* ptr);
int QImage_AllGray(void* ptr);
int QImage_BitPlaneCount(void* ptr);
int QImage_ByteCount(void* ptr);
int QImage_BytesPerLine(void* ptr);
int QImage_Depth(void* ptr);
double QImage_DevicePixelRatio(void* ptr);
int QImage_DotsPerMeterX(void* ptr);
int QImage_DotsPerMeterY(void* ptr);
int QImage_Format(void* ptr);
int QImage_HasAlphaChannel(void* ptr);
void QImage_InvertPixels(void* ptr, int mode);
int QImage_IsGrayscale(void* ptr);
int QImage_Load2(void* ptr, void* device, char* format);
int QImage_Load(void* ptr, char* fileName, char* format);
int QImage_LoadFromData2(void* ptr, void* data, char* format);
int QImage_PixelIndex(void* ptr, void* position);
int QImage_PixelIndex2(void* ptr, int x, int y);
int QImage_Save2(void* ptr, void* device, char* format, int quality);
int QImage_Save(void* ptr, char* fileName, char* format, int quality);
void QImage_SetColorCount(void* ptr, int colorCount);
void QImage_SetDevicePixelRatio(void* ptr, double scaleFactor);
void QImage_SetDotsPerMeterX(void* ptr, int x);
void QImage_SetDotsPerMeterY(void* ptr, int y);
void QImage_Swap(void* ptr, void* other);
char* QImage_Text(void* ptr, char* key);
char* QImage_TextKeys(void* ptr);
int QImage_QImage_ToImageFormat(void* format);
int QImage_Valid(void* ptr, void* pos);
int QImage_Valid2(void* ptr, int x, int y);
void QImage_DestroyQImage(void* ptr);

#ifdef __cplusplus
}
#endif