#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomCharacterData_NewQDomCharacterData();
QtObjectPtr QDomCharacterData_NewQDomCharacterData2(QtObjectPtr x);
void QDomCharacterData_AppendData(QtObjectPtr ptr, char* arg);
char* QDomCharacterData_Data(QtObjectPtr ptr);
int QDomCharacterData_Length(QtObjectPtr ptr);
int QDomCharacterData_NodeType(QtObjectPtr ptr);
void QDomCharacterData_SetData(QtObjectPtr ptr, char* v);

#ifdef __cplusplus
}
#endif