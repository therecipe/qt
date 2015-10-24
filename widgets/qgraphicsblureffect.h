#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsBlurEffect_BlurHints(QtObjectPtr ptr);
void QGraphicsBlurEffect_SetBlurHints(QtObjectPtr ptr, int hints);
QtObjectPtr QGraphicsBlurEffect_NewQGraphicsBlurEffect(QtObjectPtr parent);
void QGraphicsBlurEffect_ConnectBlurHintsChanged(QtObjectPtr ptr);
void QGraphicsBlurEffect_DisconnectBlurHintsChanged(QtObjectPtr ptr);
void QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif