#ifdef __cplusplus
extern "C" {
#endif

void* QImageWriter_NewQImageWriter();
void* QImageWriter_NewQImageWriter2(void* device, void* format);
void* QImageWriter_NewQImageWriter3(char* fileName, void* format);
int QImageWriter_CanWrite(void* ptr);
int QImageWriter_Compression(void* ptr);
void* QImageWriter_Device(void* ptr);
int QImageWriter_Error(void* ptr);
char* QImageWriter_ErrorString(void* ptr);
char* QImageWriter_FileName(void* ptr);
void* QImageWriter_Format(void* ptr);
int QImageWriter_OptimizedWrite(void* ptr);
int QImageWriter_ProgressiveScanWrite(void* ptr);
int QImageWriter_Quality(void* ptr);
void QImageWriter_SetCompression(void* ptr, int compression);
void QImageWriter_SetDevice(void* ptr, void* device);
void QImageWriter_SetFileName(void* ptr, char* fileName);
void QImageWriter_SetFormat(void* ptr, void* format);
void QImageWriter_SetOptimizedWrite(void* ptr, int optimize);
void QImageWriter_SetProgressiveScanWrite(void* ptr, int progressive);
void QImageWriter_SetQuality(void* ptr, int quality);
void QImageWriter_SetSubType(void* ptr, void* ty);
void QImageWriter_SetText(void* ptr, char* key, char* text);
void QImageWriter_SetTransformation(void* ptr, int transform);
void* QImageWriter_SubType(void* ptr);
int QImageWriter_SupportsOption(void* ptr, int option);
int QImageWriter_Transformation(void* ptr);
int QImageWriter_Write(void* ptr, void* image);
void QImageWriter_DestroyQImageWriter(void* ptr);

#ifdef __cplusplus
}
#endif