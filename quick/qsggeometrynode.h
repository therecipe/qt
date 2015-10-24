#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSGGeometryNode_NewQSGGeometryNode();
QtObjectPtr QSGGeometryNode_Material(QtObjectPtr ptr);
QtObjectPtr QSGGeometryNode_OpaqueMaterial(QtObjectPtr ptr);
void QSGGeometryNode_SetMaterial(QtObjectPtr ptr, QtObjectPtr material);
void QSGGeometryNode_SetOpaqueMaterial(QtObjectPtr ptr, QtObjectPtr material);
void QSGGeometryNode_DestroyQSGGeometryNode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif