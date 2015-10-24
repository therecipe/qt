#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QImageReader_NewQImageReader();
QtObjectPtr QImageReader_NewQImageReader2(QtObjectPtr device, QtObjectPtr format);
QtObjectPtr QImageReader_NewQImageReader3(char* fileName, QtObjectPtr format);
int QImageReader_AutoDetectImageFormat(QtObjectPtr ptr);
int QImageReader_AutoTransform(QtObjectPtr ptr);
int QImageReader_CanRead(QtObjectPtr ptr);
int QImageReader_CurrentImageNumber(QtObjectPtr ptr);
int QImageReader_DecideFormatFromContent(QtObjectPtr ptr);
QtObjectPtr QImageReader_Device(QtObjectPtr ptr);
int QImageReader_Error(QtObjectPtr ptr);
char* QImageReader_ErrorString(QtObjectPtr ptr);
char* QImageReader_FileName(QtObjectPtr ptr);
int QImageReader_ImageCount(QtObjectPtr ptr);
int QImageReader_ImageFormat(QtObjectPtr ptr);
int QImageReader_JumpToImage(QtObjectPtr ptr, int imageNumber);
int QImageReader_JumpToNextImage(QtObjectPtr ptr);
int QImageReader_LoopCount(QtObjectPtr ptr);
int QImageReader_NextImageDelay(QtObjectPtr ptr);
int QImageReader_Quality(QtObjectPtr ptr);
int QImageReader_Read2(QtObjectPtr ptr, QtObjectPtr image);
void QImageReader_SetAutoDetectImageFormat(QtObjectPtr ptr, int enabled);
void QImageReader_SetAutoTransform(QtObjectPtr ptr, int enabled);
void QImageReader_SetBackgroundColor(QtObjectPtr ptr, QtObjectPtr color);
void QImageReader_SetClipRect(QtObjectPtr ptr, QtObjectPtr rect);
void QImageReader_SetDecideFormatFromContent(QtObjectPtr ptr, int ignored);
void QImageReader_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QImageReader_SetFileName(QtObjectPtr ptr, char* fileName);
void QImageReader_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QImageReader_SetQuality(QtObjectPtr ptr, int quality);
void QImageReader_SetScaledClipRect(QtObjectPtr ptr, QtObjectPtr rect);
void QImageReader_SetScaledSize(QtObjectPtr ptr, QtObjectPtr size);
int QImageReader_SupportsAnimation(QtObjectPtr ptr);
int QImageReader_SupportsOption(QtObjectPtr ptr, int option);
char* QImageReader_Text(QtObjectPtr ptr, char* key);
char* QImageReader_TextKeys(QtObjectPtr ptr);
int QImageReader_Transformation(QtObjectPtr ptr);
void QImageReader_DestroyQImageReader(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif