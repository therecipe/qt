#ifdef __cplusplus
extern "C" {
#endif

void* QImageReader_NewQImageReader();
void* QImageReader_NewQImageReader2(void* device, void* format);
void* QImageReader_NewQImageReader3(char* fileName, void* format);
int QImageReader_AutoDetectImageFormat(void* ptr);
int QImageReader_AutoTransform(void* ptr);
void* QImageReader_BackgroundColor(void* ptr);
int QImageReader_CanRead(void* ptr);
int QImageReader_CurrentImageNumber(void* ptr);
int QImageReader_DecideFormatFromContent(void* ptr);
void* QImageReader_Device(void* ptr);
int QImageReader_Error(void* ptr);
char* QImageReader_ErrorString(void* ptr);
char* QImageReader_FileName(void* ptr);
void* QImageReader_Format(void* ptr);
int QImageReader_ImageCount(void* ptr);
void* QImageReader_QImageReader_ImageFormat3(void* device);
void* QImageReader_QImageReader_ImageFormat2(char* fileName);
int QImageReader_ImageFormat(void* ptr);
int QImageReader_JumpToImage(void* ptr, int imageNumber);
int QImageReader_JumpToNextImage(void* ptr);
int QImageReader_LoopCount(void* ptr);
int QImageReader_NextImageDelay(void* ptr);
int QImageReader_Quality(void* ptr);
int QImageReader_Read2(void* ptr, void* image);
void QImageReader_SetAutoDetectImageFormat(void* ptr, int enabled);
void QImageReader_SetAutoTransform(void* ptr, int enabled);
void QImageReader_SetBackgroundColor(void* ptr, void* color);
void QImageReader_SetClipRect(void* ptr, void* rect);
void QImageReader_SetDecideFormatFromContent(void* ptr, int ignored);
void QImageReader_SetDevice(void* ptr, void* device);
void QImageReader_SetFileName(void* ptr, char* fileName);
void QImageReader_SetFormat(void* ptr, void* format);
void QImageReader_SetQuality(void* ptr, int quality);
void QImageReader_SetScaledClipRect(void* ptr, void* rect);
void QImageReader_SetScaledSize(void* ptr, void* size);
void* QImageReader_SubType(void* ptr);
int QImageReader_SupportsAnimation(void* ptr);
int QImageReader_SupportsOption(void* ptr, int option);
char* QImageReader_Text(void* ptr, char* key);
char* QImageReader_TextKeys(void* ptr);
int QImageReader_Transformation(void* ptr);
void QImageReader_DestroyQImageReader(void* ptr);

#ifdef __cplusplus
}
#endif