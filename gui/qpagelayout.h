#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPageLayout_NewQPageLayout();
QtObjectPtr QPageLayout_NewQPageLayout3(QtObjectPtr other);
QtObjectPtr QPageLayout_NewQPageLayout2(QtObjectPtr pageSize, int orientation, QtObjectPtr margins, int units, QtObjectPtr minMargins);
int QPageLayout_IsEquivalentTo(QtObjectPtr ptr, QtObjectPtr other);
int QPageLayout_IsValid(QtObjectPtr ptr);
int QPageLayout_Mode(QtObjectPtr ptr);
int QPageLayout_Orientation(QtObjectPtr ptr);
int QPageLayout_SetMargins(QtObjectPtr ptr, QtObjectPtr margins);
void QPageLayout_SetMinimumMargins(QtObjectPtr ptr, QtObjectPtr minMargins);
void QPageLayout_SetMode(QtObjectPtr ptr, int mode);
void QPageLayout_SetOrientation(QtObjectPtr ptr, int orientation);
void QPageLayout_SetPageSize(QtObjectPtr ptr, QtObjectPtr pageSize, QtObjectPtr minMargins);
void QPageLayout_SetUnits(QtObjectPtr ptr, int units);
void QPageLayout_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QPageLayout_Units(QtObjectPtr ptr);
void QPageLayout_DestroyQPageLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif