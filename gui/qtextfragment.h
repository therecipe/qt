#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextFragment_NewQTextFragment();
QtObjectPtr QTextFragment_NewQTextFragment3(QtObjectPtr other);
int QTextFragment_CharFormatIndex(QtObjectPtr ptr);
int QTextFragment_Contains(QtObjectPtr ptr, int position);
int QTextFragment_IsValid(QtObjectPtr ptr);
int QTextFragment_Length(QtObjectPtr ptr);
int QTextFragment_Position(QtObjectPtr ptr);
char* QTextFragment_Text(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif