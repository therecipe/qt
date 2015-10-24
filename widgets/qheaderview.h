#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QHeaderView_CascadingSectionResizes(QtObjectPtr ptr);
int QHeaderView_DefaultAlignment(QtObjectPtr ptr);
int QHeaderView_DefaultSectionSize(QtObjectPtr ptr);
int QHeaderView_HighlightSections(QtObjectPtr ptr);
int QHeaderView_IsSortIndicatorShown(QtObjectPtr ptr);
int QHeaderView_MaximumSectionSize(QtObjectPtr ptr);
int QHeaderView_MinimumSectionSize(QtObjectPtr ptr);
void QHeaderView_ResetDefaultSectionSize(QtObjectPtr ptr);
void QHeaderView_ResizeSection(QtObjectPtr ptr, int logicalIndex, int size);
void QHeaderView_SetCascadingSectionResizes(QtObjectPtr ptr, int enable);
void QHeaderView_SetDefaultAlignment(QtObjectPtr ptr, int alignment);
void QHeaderView_SetDefaultSectionSize(QtObjectPtr ptr, int size);
void QHeaderView_SetHighlightSections(QtObjectPtr ptr, int highlight);
void QHeaderView_SetMaximumSectionSize(QtObjectPtr ptr, int size);
void QHeaderView_SetMinimumSectionSize(QtObjectPtr ptr, int size);
void QHeaderView_SetOffset(QtObjectPtr ptr, int offset);
void QHeaderView_SetSortIndicatorShown(QtObjectPtr ptr, int show);
void QHeaderView_SetStretchLastSection(QtObjectPtr ptr, int stretch);
int QHeaderView_StretchLastSection(QtObjectPtr ptr);
int QHeaderView_Count(QtObjectPtr ptr);
void QHeaderView_ConnectGeometriesChanged(QtObjectPtr ptr);
void QHeaderView_DisconnectGeometriesChanged(QtObjectPtr ptr);
void QHeaderView_HeaderDataChanged(QtObjectPtr ptr, int orientation, int logicalFirst, int logicalLast);
int QHeaderView_HiddenSectionCount(QtObjectPtr ptr);
void QHeaderView_HideSection(QtObjectPtr ptr, int logicalIndex);
int QHeaderView_IsSectionHidden(QtObjectPtr ptr, int logicalIndex);
int QHeaderView_Length(QtObjectPtr ptr);
int QHeaderView_LogicalIndex(QtObjectPtr ptr, int visualIndex);
int QHeaderView_LogicalIndexAt3(QtObjectPtr ptr, QtObjectPtr pos);
int QHeaderView_LogicalIndexAt(QtObjectPtr ptr, int position);
int QHeaderView_LogicalIndexAt2(QtObjectPtr ptr, int x, int y);
void QHeaderView_MoveSection(QtObjectPtr ptr, int from, int to);
int QHeaderView_Offset(QtObjectPtr ptr);
int QHeaderView_Orientation(QtObjectPtr ptr);
void QHeaderView_Reset(QtObjectPtr ptr);
int QHeaderView_ResizeContentsPrecision(QtObjectPtr ptr);
void QHeaderView_ResizeSections(QtObjectPtr ptr, int mode);
int QHeaderView_RestoreState(QtObjectPtr ptr, QtObjectPtr state);
void QHeaderView_ConnectSectionClicked(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionClicked(QtObjectPtr ptr);
void QHeaderView_ConnectSectionCountChanged(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionCountChanged(QtObjectPtr ptr);
void QHeaderView_ConnectSectionDoubleClicked(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionDoubleClicked(QtObjectPtr ptr);
void QHeaderView_ConnectSectionEntered(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionEntered(QtObjectPtr ptr);
void QHeaderView_ConnectSectionHandleDoubleClicked(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionHandleDoubleClicked(QtObjectPtr ptr);
void QHeaderView_ConnectSectionMoved(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionMoved(QtObjectPtr ptr);
int QHeaderView_SectionPosition(QtObjectPtr ptr, int logicalIndex);
void QHeaderView_ConnectSectionPressed(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionPressed(QtObjectPtr ptr);
int QHeaderView_SectionResizeMode(QtObjectPtr ptr, int logicalIndex);
void QHeaderView_ConnectSectionResized(QtObjectPtr ptr);
void QHeaderView_DisconnectSectionResized(QtObjectPtr ptr);
int QHeaderView_SectionSize(QtObjectPtr ptr, int logicalIndex);
int QHeaderView_SectionSizeHint(QtObjectPtr ptr, int logicalIndex);
int QHeaderView_SectionViewportPosition(QtObjectPtr ptr, int logicalIndex);
int QHeaderView_SectionsClickable(QtObjectPtr ptr);
int QHeaderView_SectionsHidden(QtObjectPtr ptr);
int QHeaderView_SectionsMovable(QtObjectPtr ptr);
int QHeaderView_SectionsMoved(QtObjectPtr ptr);
void QHeaderView_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QHeaderView_SetOffsetToLastSection(QtObjectPtr ptr);
void QHeaderView_SetOffsetToSectionPosition(QtObjectPtr ptr, int visualSectionNumber);
void QHeaderView_SetResizeContentsPrecision(QtObjectPtr ptr, int precision);
void QHeaderView_SetSectionHidden(QtObjectPtr ptr, int logicalIndex, int hide);
void QHeaderView_SetSectionResizeMode(QtObjectPtr ptr, int mode);
void QHeaderView_SetSectionResizeMode2(QtObjectPtr ptr, int logicalIndex, int mode);
void QHeaderView_SetSectionsClickable(QtObjectPtr ptr, int clickable);
void QHeaderView_SetSectionsMovable(QtObjectPtr ptr, int movable);
void QHeaderView_SetSortIndicator(QtObjectPtr ptr, int logicalIndex, int order);
void QHeaderView_SetVisible(QtObjectPtr ptr, int v);
void QHeaderView_ShowSection(QtObjectPtr ptr, int logicalIndex);
void QHeaderView_ConnectSortIndicatorChanged(QtObjectPtr ptr);
void QHeaderView_DisconnectSortIndicatorChanged(QtObjectPtr ptr);
int QHeaderView_SortIndicatorOrder(QtObjectPtr ptr);
int QHeaderView_SortIndicatorSection(QtObjectPtr ptr);
int QHeaderView_StretchSectionCount(QtObjectPtr ptr);
void QHeaderView_SwapSections(QtObjectPtr ptr, int first, int second);
int QHeaderView_VisualIndex(QtObjectPtr ptr, int logicalIndex);
int QHeaderView_VisualIndexAt(QtObjectPtr ptr, int position);
void QHeaderView_DestroyQHeaderView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif