#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QImageWriter_NewQImageWriter();
QtObjectPtr QImageWriter_NewQImageWriter2(QtObjectPtr device, QtObjectPtr format);
QtObjectPtr QImageWriter_NewQImageWriter3(char* fileName, QtObjectPtr format);
int QImageWriter_CanWrite(QtObjectPtr ptr);
int QImageWriter_Compression(QtObjectPtr ptr);
QtObjectPtr QImageWriter_Device(QtObjectPtr ptr);
int QImageWriter_Error(QtObjectPtr ptr);
char* QImageWriter_ErrorString(QtObjectPtr ptr);
char* QImageWriter_FileName(QtObjectPtr ptr);
int QImageWriter_OptimizedWrite(QtObjectPtr ptr);
int QImageWriter_ProgressiveScanWrite(QtObjectPtr ptr);
int QImageWriter_Quality(QtObjectPtr ptr);
void QImageWriter_SetCompression(QtObjectPtr ptr, int compression);
void QImageWriter_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QImageWriter_SetFileName(QtObjectPtr ptr, char* fileName);
void QImageWriter_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
void QImageWriter_SetOptimizedWrite(QtObjectPtr ptr, int optimize);
void QImageWriter_SetProgressiveScanWrite(QtObjectPtr ptr, int progressive);
void QImageWriter_SetQuality(QtObjectPtr ptr, int quality);
void QImageWriter_SetSubType(QtObjectPtr ptr, QtObjectPtr ty);
void QImageWriter_SetText(QtObjectPtr ptr, char* key, char* text);
void QImageWriter_SetTransformation(QtObjectPtr ptr, int transform);
int QImageWriter_SupportsOption(QtObjectPtr ptr, int option);
int QImageWriter_Transformation(QtObjectPtr ptr);
int QImageWriter_Write(QtObjectPtr ptr, QtObjectPtr image);
void QImageWriter_DestroyQImageWriter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif