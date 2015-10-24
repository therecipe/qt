#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QBoxLayout_Direction(QtObjectPtr ptr);
QtObjectPtr QBoxLayout_NewQBoxLayout(int dir, QtObjectPtr parent);
void QBoxLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item);
void QBoxLayout_AddLayout(QtObjectPtr ptr, QtObjectPtr layout, int stretch);
void QBoxLayout_AddSpacerItem(QtObjectPtr ptr, QtObjectPtr spacerItem);
void QBoxLayout_AddSpacing(QtObjectPtr ptr, int size);
void QBoxLayout_AddStretch(QtObjectPtr ptr, int stretch);
void QBoxLayout_AddStrut(QtObjectPtr ptr, int size);
void QBoxLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int stretch, int alignment);
int QBoxLayout_Count(QtObjectPtr ptr);
int QBoxLayout_ExpandingDirections(QtObjectPtr ptr);
int QBoxLayout_HasHeightForWidth(QtObjectPtr ptr);
int QBoxLayout_HeightForWidth(QtObjectPtr ptr, int w);
void QBoxLayout_InsertItem(QtObjectPtr ptr, int index, QtObjectPtr item);
void QBoxLayout_InsertLayout(QtObjectPtr ptr, int index, QtObjectPtr layout, int stretch);
void QBoxLayout_InsertSpacerItem(QtObjectPtr ptr, int index, QtObjectPtr spacerItem);
void QBoxLayout_InsertSpacing(QtObjectPtr ptr, int index, int size);
void QBoxLayout_InsertStretch(QtObjectPtr ptr, int index, int stretch);
void QBoxLayout_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch, int alignment);
void QBoxLayout_Invalidate(QtObjectPtr ptr);
QtObjectPtr QBoxLayout_ItemAt(QtObjectPtr ptr, int index);
int QBoxLayout_MinimumHeightForWidth(QtObjectPtr ptr, int w);
void QBoxLayout_SetDirection(QtObjectPtr ptr, int direction);
void QBoxLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr r);
void QBoxLayout_SetSpacing(QtObjectPtr ptr, int spacing);
void QBoxLayout_SetStretch(QtObjectPtr ptr, int index, int stretch);
int QBoxLayout_SetStretchFactor2(QtObjectPtr ptr, QtObjectPtr layout, int stretch);
int QBoxLayout_SetStretchFactor(QtObjectPtr ptr, QtObjectPtr widget, int stretch);
int QBoxLayout_Spacing(QtObjectPtr ptr);
int QBoxLayout_Stretch(QtObjectPtr ptr, int index);
QtObjectPtr QBoxLayout_TakeAt(QtObjectPtr ptr, int index);
void QBoxLayout_DestroyQBoxLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif