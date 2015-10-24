#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSGTransformNode_NewQSGTransformNode();
void QSGTransformNode_SetMatrix(QtObjectPtr ptr, QtObjectPtr matrix);
void QSGTransformNode_DestroyQSGTransformNode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif