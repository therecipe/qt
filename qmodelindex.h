#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
int QModelIndex_Column(QtObjectPtr ptr);
char* QModelIndex_Data_Int(QtObjectPtr ptr, int role);
int QModelIndex_Flags(QtObjectPtr ptr);
int QModelIndex_IsValid(QtObjectPtr ptr);
int QModelIndex_Row(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
