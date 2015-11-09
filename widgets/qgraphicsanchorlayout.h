#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(void* parent);
void* QGraphicsAnchorLayout_AddAnchor(void* ptr, void* firstItem, int firstEdge, void* secondItem, int secondEdge);
void QGraphicsAnchorLayout_AddAnchors(void* ptr, void* firstItem, void* secondItem, int orientations);
void QGraphicsAnchorLayout_AddCornerAnchors(void* ptr, void* firstItem, int firstCorner, void* secondItem, int secondCorner);
void* QGraphicsAnchorLayout_Anchor(void* ptr, void* firstItem, int firstEdge, void* secondItem, int secondEdge);
int QGraphicsAnchorLayout_Count(void* ptr);
double QGraphicsAnchorLayout_HorizontalSpacing(void* ptr);
void QGraphicsAnchorLayout_Invalidate(void* ptr);
void* QGraphicsAnchorLayout_ItemAt(void* ptr, int index);
void QGraphicsAnchorLayout_RemoveAt(void* ptr, int index);
void QGraphicsAnchorLayout_SetGeometry(void* ptr, void* geom);
void QGraphicsAnchorLayout_SetHorizontalSpacing(void* ptr, double spacing);
void QGraphicsAnchorLayout_SetSpacing(void* ptr, double spacing);
void QGraphicsAnchorLayout_SetVerticalSpacing(void* ptr, double spacing);
double QGraphicsAnchorLayout_VerticalSpacing(void* ptr);
void QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(void* ptr);

#ifdef __cplusplus
}
#endif