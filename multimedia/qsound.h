#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSound_SetLoops(QtObjectPtr ptr, int number);
QtObjectPtr QSound_NewQSound(char* filename, QtObjectPtr parent);
char* QSound_FileName(QtObjectPtr ptr);
int QSound_IsFinished(QtObjectPtr ptr);
int QSound_Loops(QtObjectPtr ptr);
int QSound_LoopsRemaining(QtObjectPtr ptr);
void QSound_Play2(QtObjectPtr ptr);
void QSound_QSound_Play(char* filename);
void QSound_Stop(QtObjectPtr ptr);
void QSound_DestroyQSound(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif