#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QImage_ColorCount(QtObjectPtr ptr);
void QImage_Fill2(QtObjectPtr ptr, int color);
void QImage_Fill3(QtObjectPtr ptr, QtObjectPtr color);
int QImage_Height(QtObjectPtr ptr);
int QImage_IsNull(QtObjectPtr ptr);
void QImage_SetOffset(QtObjectPtr ptr, QtObjectPtr offset);
void QImage_SetText(QtObjectPtr ptr, char* key, char* text);
int QImage_Width(QtObjectPtr ptr);
int QImage_AllGray(QtObjectPtr ptr);
int QImage_BitPlaneCount(QtObjectPtr ptr);
int QImage_ByteCount(QtObjectPtr ptr);
int QImage_BytesPerLine(QtObjectPtr ptr);
int QImage_Depth(QtObjectPtr ptr);
int QImage_DotsPerMeterX(QtObjectPtr ptr);
int QImage_DotsPerMeterY(QtObjectPtr ptr);
int QImage_Format(QtObjectPtr ptr);
int QImage_HasAlphaChannel(QtObjectPtr ptr);
void QImage_InvertPixels(QtObjectPtr ptr, int mode);
int QImage_IsGrayscale(QtObjectPtr ptr);
int QImage_Load2(QtObjectPtr ptr, QtObjectPtr device, char* format);
int QImage_Load(QtObjectPtr ptr, char* fileName, char* format);
int QImage_LoadFromData2(QtObjectPtr ptr, QtObjectPtr data, char* format);
int QImage_PixelIndex(QtObjectPtr ptr, QtObjectPtr position);
int QImage_PixelIndex2(QtObjectPtr ptr, int x, int y);
int QImage_Save2(QtObjectPtr ptr, QtObjectPtr device, char* format, int quality);
int QImage_Save(QtObjectPtr ptr, char* fileName, char* format, int quality);
void QImage_SetColorCount(QtObjectPtr ptr, int colorCount);
void QImage_SetDotsPerMeterX(QtObjectPtr ptr, int x);
void QImage_SetDotsPerMeterY(QtObjectPtr ptr, int y);
void QImage_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QImage_Text(QtObjectPtr ptr, char* key);
char* QImage_TextKeys(QtObjectPtr ptr);
int QImage_QImage_ToImageFormat(QtObjectPtr format);
int QImage_Valid(QtObjectPtr ptr, QtObjectPtr pos);
int QImage_Valid2(QtObjectPtr ptr, int x, int y);
void QImage_DestroyQImage(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif