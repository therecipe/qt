#ifdef __cplusplus
extern "C" {
#endif

int QScreen_Depth(void* ptr);
double QScreen_DevicePixelRatio(void* ptr);
double QScreen_LogicalDotsPerInch(void* ptr);
double QScreen_LogicalDotsPerInchX(void* ptr);
double QScreen_LogicalDotsPerInchY(void* ptr);
char* QScreen_Name(void* ptr);
int QScreen_NativeOrientation(void* ptr);
int QScreen_Orientation(void* ptr);
double QScreen_PhysicalDotsPerInch(void* ptr);
double QScreen_PhysicalDotsPerInchX(void* ptr);
double QScreen_PhysicalDotsPerInchY(void* ptr);
int QScreen_PrimaryOrientation(void* ptr);
double QScreen_RefreshRate(void* ptr);
int QScreen_AngleBetween(void* ptr, int a, int b);
int QScreen_IsLandscape(void* ptr, int o);
int QScreen_IsPortrait(void* ptr, int o);
void QScreen_ConnectOrientationChanged(void* ptr);
void QScreen_DisconnectOrientationChanged(void* ptr);
int QScreen_OrientationUpdateMask(void* ptr);
void QScreen_ConnectPrimaryOrientationChanged(void* ptr);
void QScreen_DisconnectPrimaryOrientationChanged(void* ptr);
void QScreen_SetOrientationUpdateMask(void* ptr, int mask);
void QScreen_DestroyQScreen(void* ptr);

#ifdef __cplusplus
}
#endif