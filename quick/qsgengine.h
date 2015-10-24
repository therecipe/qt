#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSGEngine_NewQSGEngine(QtObjectPtr parent);
QtObjectPtr QSGEngine_CreateRenderer(QtObjectPtr ptr);
QtObjectPtr QSGEngine_CreateTextureFromImage(QtObjectPtr ptr, QtObjectPtr image, int options);
void QSGEngine_Initialize(QtObjectPtr ptr, QtObjectPtr context);
void QSGEngine_Invalidate(QtObjectPtr ptr);
void QSGEngine_DestroyQSGEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif