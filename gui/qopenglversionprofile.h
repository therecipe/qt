#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QOpenGLVersionProfile_NewQOpenGLVersionProfile();
QtObjectPtr QOpenGLVersionProfile_NewQOpenGLVersionProfile3(QtObjectPtr other);
QtObjectPtr QOpenGLVersionProfile_NewQOpenGLVersionProfile2(QtObjectPtr format);
int QOpenGLVersionProfile_HasProfiles(QtObjectPtr ptr);
int QOpenGLVersionProfile_IsLegacyVersion(QtObjectPtr ptr);
int QOpenGLVersionProfile_IsValid(QtObjectPtr ptr);
int QOpenGLVersionProfile_Profile(QtObjectPtr ptr);
void QOpenGLVersionProfile_SetProfile(QtObjectPtr ptr, int profile);
void QOpenGLVersionProfile_SetVersion(QtObjectPtr ptr, int majorVersion, int minorVersion);
void QOpenGLVersionProfile_DestroyQOpenGLVersionProfile(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif