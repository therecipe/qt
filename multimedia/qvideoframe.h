#ifdef __cplusplus
extern "C" {
#endif

void* QVideoFrame_NewQVideoFrame();
void* QVideoFrame_NewQVideoFrame2(void* buffer, void* size, int format);
void* QVideoFrame_NewQVideoFrame4(void* image);
void* QVideoFrame_NewQVideoFrame5(void* other);
void* QVideoFrame_NewQVideoFrame3(int bytes, void* size, int bytesPerLine, int format);
int QVideoFrame_BytesPerLine(void* ptr);
int QVideoFrame_BytesPerLine2(void* ptr, int plane);
int QVideoFrame_FieldType(void* ptr);
void* QVideoFrame_Handle(void* ptr);
int QVideoFrame_HandleType(void* ptr);
int QVideoFrame_Height(void* ptr);
int QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(int format);
int QVideoFrame_IsMapped(void* ptr);
int QVideoFrame_IsReadable(void* ptr);
int QVideoFrame_IsValid(void* ptr);
int QVideoFrame_IsWritable(void* ptr);
int QVideoFrame_Map(void* ptr, int mode);
int QVideoFrame_MapMode(void* ptr);
int QVideoFrame_MappedBytes(void* ptr);
void* QVideoFrame_MetaData(void* ptr, char* key);
int QVideoFrame_PixelFormat(void* ptr);
int QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(int format);
int QVideoFrame_PlaneCount(void* ptr);
void QVideoFrame_SetFieldType(void* ptr, int field);
void QVideoFrame_SetMetaData(void* ptr, char* key, void* value);
void QVideoFrame_Unmap(void* ptr);
int QVideoFrame_Width(void* ptr);
void QVideoFrame_DestroyQVideoFrame(void* ptr);

#ifdef __cplusplus
}
#endif