#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QKeyEvent_Matches(QtObjectPtr ptr, int key);
int QKeyEvent_Count(QtObjectPtr ptr);
int QKeyEvent_IsAutoRepeat(QtObjectPtr ptr);
int QKeyEvent_Key(QtObjectPtr ptr);
int QKeyEvent_Modifiers(QtObjectPtr ptr);
char* QKeyEvent_Text(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif