#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QStandardPaths_QStandardPaths_SetTestModeEnabled(int testMode);
char* QStandardPaths_QStandardPaths_FindExecutable(char* executableName, char* paths);
char* QStandardPaths_QStandardPaths_Locate(int ty, char* fileName, int options);
char* QStandardPaths_QStandardPaths_LocateAll(int ty, char* fileName, int options);
char* QStandardPaths_QStandardPaths_DisplayName(int ty);
char* QStandardPaths_QStandardPaths_StandardLocations(int ty);
char* QStandardPaths_QStandardPaths_WritableLocation(int ty);

#ifdef __cplusplus
}
#endif