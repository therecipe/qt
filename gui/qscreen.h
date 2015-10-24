#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QScreen_Depth(QtObjectPtr ptr);
char* QScreen_Name(QtObjectPtr ptr);
int QScreen_NativeOrientation(QtObjectPtr ptr);
int QScreen_Orientation(QtObjectPtr ptr);
int QScreen_PrimaryOrientation(QtObjectPtr ptr);
int QScreen_AngleBetween(QtObjectPtr ptr, int a, int b);
int QScreen_IsLandscape(QtObjectPtr ptr, int o);
int QScreen_IsPortrait(QtObjectPtr ptr, int o);
void QScreen_ConnectOrientationChanged(QtObjectPtr ptr);
void QScreen_DisconnectOrientationChanged(QtObjectPtr ptr);
int QScreen_OrientationUpdateMask(QtObjectPtr ptr);
void QScreen_ConnectPrimaryOrientationChanged(QtObjectPtr ptr);
void QScreen_DisconnectPrimaryOrientationChanged(QtObjectPtr ptr);
void QScreen_SetOrientationUpdateMask(QtObjectPtr ptr, int mask);
void QScreen_DestroyQScreen(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif