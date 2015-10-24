#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPixelFormat_NewQPixelFormat();
int QPixelFormat_AlphaPosition(QtObjectPtr ptr);
int QPixelFormat_AlphaUsage(QtObjectPtr ptr);
int QPixelFormat_ByteOrder(QtObjectPtr ptr);
int QPixelFormat_ColorModel(QtObjectPtr ptr);
int QPixelFormat_Premultiplied(QtObjectPtr ptr);
int QPixelFormat_TypeInterpretation(QtObjectPtr ptr);
int QPixelFormat_YuvLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif