#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLPixelTransferOptions_NewQOpenGLPixelTransferOptions();
int QOpenGLPixelTransferOptions_Alignment(QtObjectPtr ptr);
int QOpenGLPixelTransferOptions_ImageHeight(QtObjectPtr ptr);
int QOpenGLPixelTransferOptions_IsLeastSignificantBitFirst(QtObjectPtr ptr);
int QOpenGLPixelTransferOptions_IsSwapBytesEnabled(QtObjectPtr ptr);
void QOpenGLPixelTransferOptions_SetAlignment(QtObjectPtr ptr, int alignment);
void QOpenGLPixelTransferOptions_SetImageHeight(QtObjectPtr ptr, int imageHeight);
void QOpenGLPixelTransferOptions_SetSkipImages(QtObjectPtr ptr, int skipImages);
void QOpenGLPixelTransferOptions_SetSkipPixels(QtObjectPtr ptr, int skipPixels);
void QOpenGLPixelTransferOptions_SetSkipRows(QtObjectPtr ptr, int skipRows);
int QOpenGLPixelTransferOptions_SkipImages(QtObjectPtr ptr);
int QOpenGLPixelTransferOptions_SkipPixels(QtObjectPtr ptr);
int QOpenGLPixelTransferOptions_SkipRows(QtObjectPtr ptr);
int QOpenGLPixelTransferOptions_RowLength(QtObjectPtr ptr);
void QOpenGLPixelTransferOptions_SetLeastSignificantByteFirst(QtObjectPtr ptr, int lsbFirst);
void QOpenGLPixelTransferOptions_SetRowLength(QtObjectPtr ptr, int rowLength);
void QOpenGLPixelTransferOptions_SetSwapBytesEnabled(QtObjectPtr ptr, int swapBytes);
void QOpenGLPixelTransferOptions_DestroyQOpenGLPixelTransferOptions(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif