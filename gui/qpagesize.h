#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPageSize_NewQPageSize();
QtObjectPtr QPageSize_NewQPageSize2(int pageSize);
QtObjectPtr QPageSize_NewQPageSize5(QtObjectPtr other);
QtObjectPtr QPageSize_NewQPageSize3(QtObjectPtr pointSize, char* name, int matchPolicy);
QtObjectPtr QPageSize_NewQPageSize4(QtObjectPtr size, int units, char* name, int matchPolicy);
int QPageSize_QPageSize_DefinitionUnits2(int pageSizeId);
int QPageSize_DefinitionUnits(QtObjectPtr ptr);
int QPageSize_QPageSize_Id2(QtObjectPtr pointSize, int matchPolicy);
int QPageSize_QPageSize_Id3(QtObjectPtr size, int units, int matchPolicy);
int QPageSize_QPageSize_Id4(int windowsId);
int QPageSize_Id(QtObjectPtr ptr);
int QPageSize_IsEquivalentTo(QtObjectPtr ptr, QtObjectPtr other);
int QPageSize_IsValid(QtObjectPtr ptr);
char* QPageSize_QPageSize_Key2(int pageSizeId);
char* QPageSize_Key(QtObjectPtr ptr);
char* QPageSize_QPageSize_Name2(int pageSizeId);
char* QPageSize_Name(QtObjectPtr ptr);
void QPageSize_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QPageSize_QPageSize_WindowsId2(int pageSizeId);
int QPageSize_WindowsId(QtObjectPtr ptr);
void QPageSize_DestroyQPageSize(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif