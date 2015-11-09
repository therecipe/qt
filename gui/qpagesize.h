#ifdef __cplusplus
extern "C" {
#endif

void* QPageSize_NewQPageSize();
void* QPageSize_NewQPageSize2(int pageSize);
void* QPageSize_NewQPageSize5(void* other);
void* QPageSize_NewQPageSize3(void* pointSize, char* name, int matchPolicy);
void* QPageSize_NewQPageSize4(void* size, int units, char* name, int matchPolicy);
int QPageSize_QPageSize_DefinitionUnits2(int pageSizeId);
int QPageSize_DefinitionUnits(void* ptr);
int QPageSize_QPageSize_Id2(void* pointSize, int matchPolicy);
int QPageSize_QPageSize_Id3(void* size, int units, int matchPolicy);
int QPageSize_QPageSize_Id4(int windowsId);
int QPageSize_Id(void* ptr);
int QPageSize_IsEquivalentTo(void* ptr, void* other);
int QPageSize_IsValid(void* ptr);
char* QPageSize_QPageSize_Key2(int pageSizeId);
char* QPageSize_Key(void* ptr);
char* QPageSize_QPageSize_Name2(int pageSizeId);
char* QPageSize_Name(void* ptr);
void QPageSize_Swap(void* ptr, void* other);
int QPageSize_QPageSize_WindowsId2(int pageSizeId);
int QPageSize_WindowsId(void* ptr);
void QPageSize_DestroyQPageSize(void* ptr);

#ifdef __cplusplus
}
#endif