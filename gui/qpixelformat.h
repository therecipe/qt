#ifdef __cplusplus
extern "C" {
#endif

void* QPixelFormat_NewQPixelFormat();
int QPixelFormat_AlphaPosition(void* ptr);
int QPixelFormat_AlphaUsage(void* ptr);
int QPixelFormat_ByteOrder(void* ptr);
int QPixelFormat_ColorModel(void* ptr);
int QPixelFormat_Premultiplied(void* ptr);
int QPixelFormat_TypeInterpretation(void* ptr);
int QPixelFormat_YuvLayout(void* ptr);

#ifdef __cplusplus
}
#endif