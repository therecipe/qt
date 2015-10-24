#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(QtObjectPtr parent);
QtObjectPtr QGraphicsAnchorLayout_AddAnchor(QtObjectPtr ptr, QtObjectPtr firstItem, int firstEdge, QtObjectPtr secondItem, int secondEdge);
void QGraphicsAnchorLayout_AddAnchors(QtObjectPtr ptr, QtObjectPtr firstItem, QtObjectPtr secondItem, int orientations);
void QGraphicsAnchorLayout_AddCornerAnchors(QtObjectPtr ptr, QtObjectPtr firstItem, int firstCorner, QtObjectPtr secondItem, int secondCorner);
QtObjectPtr QGraphicsAnchorLayout_Anchor(QtObjectPtr ptr, QtObjectPtr firstItem, int firstEdge, QtObjectPtr secondItem, int secondEdge);
int QGraphicsAnchorLayout_Count(QtObjectPtr ptr);
void QGraphicsAnchorLayout_Invalidate(QtObjectPtr ptr);
QtObjectPtr QGraphicsAnchorLayout_ItemAt(QtObjectPtr ptr, int index);
void QGraphicsAnchorLayout_RemoveAt(QtObjectPtr ptr, int index);
void QGraphicsAnchorLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr geom);
void QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif