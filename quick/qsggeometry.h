#ifdef __cplusplus
extern "C" {
#endif

void QSGGeometry_Allocate(void* ptr, int vertexCount, int indexCount);
int QSGGeometry_AttributeCount(void* ptr);
int QSGGeometry_IndexCount(void* ptr);
void* QSGGeometry_IndexData(void* ptr);
void* QSGGeometry_IndexData2(void* ptr);
int QSGGeometry_IndexDataPattern(void* ptr);
int QSGGeometry_IndexType(void* ptr);
void QSGGeometry_MarkIndexDataDirty(void* ptr);
void QSGGeometry_MarkVertexDataDirty(void* ptr);
void QSGGeometry_SetIndexDataPattern(void* ptr, int p);
void QSGGeometry_SetVertexDataPattern(void* ptr, int p);
int QSGGeometry_SizeOfIndex(void* ptr);
int QSGGeometry_SizeOfVertex(void* ptr);
void QSGGeometry_QSGGeometry_UpdateRectGeometry(void* g, void* rect);
void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(void* g, void* rect, void* textureRect);
int QSGGeometry_VertexCount(void* ptr);
void* QSGGeometry_VertexData(void* ptr);
void* QSGGeometry_VertexData2(void* ptr);
int QSGGeometry_VertexDataPattern(void* ptr);
void QSGGeometry_DestroyQSGGeometry(void* ptr);

#ifdef __cplusplus
}
#endif