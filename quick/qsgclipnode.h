#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSGClipNode_NewQSGClipNode();
int QSGClipNode_IsRectangular(QtObjectPtr ptr);
void QSGClipNode_SetClipRect(QtObjectPtr ptr, QtObjectPtr rect);
void QSGClipNode_SetIsRectangular(QtObjectPtr ptr, int rectHint);
void QSGClipNode_DestroyQSGClipNode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif