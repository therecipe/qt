#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSGGeometry_Allocate(QtObjectPtr ptr, int vertexCount, int indexCount);
int QSGGeometry_AttributeCount(QtObjectPtr ptr);
int QSGGeometry_IndexCount(QtObjectPtr ptr);
void QSGGeometry_IndexData(QtObjectPtr ptr);
void QSGGeometry_IndexData2(QtObjectPtr ptr);
int QSGGeometry_IndexDataPattern(QtObjectPtr ptr);
int QSGGeometry_IndexType(QtObjectPtr ptr);
void QSGGeometry_MarkIndexDataDirty(QtObjectPtr ptr);
void QSGGeometry_MarkVertexDataDirty(QtObjectPtr ptr);
void QSGGeometry_SetIndexDataPattern(QtObjectPtr ptr, int p);
void QSGGeometry_SetVertexDataPattern(QtObjectPtr ptr, int p);
int QSGGeometry_SizeOfIndex(QtObjectPtr ptr);
int QSGGeometry_SizeOfVertex(QtObjectPtr ptr);
void QSGGeometry_QSGGeometry_UpdateRectGeometry(QtObjectPtr g, QtObjectPtr rect);
void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(QtObjectPtr g, QtObjectPtr rect, QtObjectPtr textureRect);
int QSGGeometry_VertexCount(QtObjectPtr ptr);
void QSGGeometry_VertexData(QtObjectPtr ptr);
void QSGGeometry_VertexData2(QtObjectPtr ptr);
int QSGGeometry_VertexDataPattern(QtObjectPtr ptr);
void QSGGeometry_DestroyQSGGeometry(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif