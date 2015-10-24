#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVideoFrame_NewQVideoFrame();
QtObjectPtr QVideoFrame_NewQVideoFrame2(QtObjectPtr buffer, QtObjectPtr size, int format);
QtObjectPtr QVideoFrame_NewQVideoFrame4(QtObjectPtr image);
QtObjectPtr QVideoFrame_NewQVideoFrame5(QtObjectPtr other);
QtObjectPtr QVideoFrame_NewQVideoFrame3(int bytes, QtObjectPtr size, int bytesPerLine, int format);
int QVideoFrame_BytesPerLine(QtObjectPtr ptr);
int QVideoFrame_BytesPerLine2(QtObjectPtr ptr, int plane);
int QVideoFrame_FieldType(QtObjectPtr ptr);
char* QVideoFrame_Handle(QtObjectPtr ptr);
int QVideoFrame_HandleType(QtObjectPtr ptr);
int QVideoFrame_Height(QtObjectPtr ptr);
int QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(int format);
int QVideoFrame_IsMapped(QtObjectPtr ptr);
int QVideoFrame_IsReadable(QtObjectPtr ptr);
int QVideoFrame_IsValid(QtObjectPtr ptr);
int QVideoFrame_IsWritable(QtObjectPtr ptr);
int QVideoFrame_Map(QtObjectPtr ptr, int mode);
int QVideoFrame_MapMode(QtObjectPtr ptr);
int QVideoFrame_MappedBytes(QtObjectPtr ptr);
char* QVideoFrame_MetaData(QtObjectPtr ptr, char* key);
int QVideoFrame_PixelFormat(QtObjectPtr ptr);
int QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(int format);
int QVideoFrame_PlaneCount(QtObjectPtr ptr);
void QVideoFrame_SetFieldType(QtObjectPtr ptr, int field);
void QVideoFrame_SetMetaData(QtObjectPtr ptr, char* key, char* value);
void QVideoFrame_Unmap(QtObjectPtr ptr);
int QVideoFrame_Width(QtObjectPtr ptr);
void QVideoFrame_DestroyQVideoFrame(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif