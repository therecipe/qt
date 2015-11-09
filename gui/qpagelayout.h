#ifdef __cplusplus
extern "C" {
#endif

void* QPageLayout_NewQPageLayout();
void* QPageLayout_NewQPageLayout3(void* other);
void* QPageLayout_NewQPageLayout2(void* pageSize, int orientation, void* margins, int units, void* minMargins);
int QPageLayout_IsEquivalentTo(void* ptr, void* other);
int QPageLayout_IsValid(void* ptr);
int QPageLayout_Mode(void* ptr);
int QPageLayout_Orientation(void* ptr);
int QPageLayout_SetBottomMargin(void* ptr, double bottomMargin);
int QPageLayout_SetLeftMargin(void* ptr, double leftMargin);
int QPageLayout_SetMargins(void* ptr, void* margins);
void QPageLayout_SetMinimumMargins(void* ptr, void* minMargins);
void QPageLayout_SetMode(void* ptr, int mode);
void QPageLayout_SetOrientation(void* ptr, int orientation);
void QPageLayout_SetPageSize(void* ptr, void* pageSize, void* minMargins);
int QPageLayout_SetRightMargin(void* ptr, double rightMargin);
int QPageLayout_SetTopMargin(void* ptr, double topMargin);
void QPageLayout_SetUnits(void* ptr, int units);
void QPageLayout_Swap(void* ptr, void* other);
int QPageLayout_Units(void* ptr);
void QPageLayout_DestroyQPageLayout(void* ptr);

#ifdef __cplusplus
}
#endif