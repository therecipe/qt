#ifdef __cplusplus
extern "C" {
#endif

void QGraphicsAnchor_SetSizePolicy(void* ptr, int policy);
void QGraphicsAnchor_SetSpacing(void* ptr, double spacing);
int QGraphicsAnchor_SizePolicy(void* ptr);
double QGraphicsAnchor_Spacing(void* ptr);
void QGraphicsAnchor_UnsetSpacing(void* ptr);
void QGraphicsAnchor_DestroyQGraphicsAnchor(void* ptr);

#ifdef __cplusplus
}
#endif