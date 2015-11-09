#ifdef __cplusplus
extern "C" {
#endif

void* QVariant_NewQVariant20(void* c);
void* QVariant_NewQVariant18(void* val);
void* QVariant_NewQVariant11(int val);
void* QVariant_NewQVariant16(void* val);
void* QVariant_NewQVariant15(void* val);
void* QVariant_NewQVariant21(void* val);
void* QVariant_NewQVariant23(void* val);
void* QVariant_NewQVariant39(void* val);
void* QVariant_NewQVariant45(void* val);
void* QVariant_NewQVariant46(void* val);
void* QVariant_NewQVariant44(void* val);
void* QVariant_NewQVariant43(void* val);
void* QVariant_NewQVariant31(void* val);
void* QVariant_NewQVariant32(void* val);
void* QVariant_NewQVariant35(void* l);
void* QVariant_NewQVariant41(void* val);
void* QVariant_NewQVariant42(void* val);
void* QVariant_NewQVariant29(void* val);
void* QVariant_NewQVariant30(void* val);
void* QVariant_NewQVariant33(void* val);
void* QVariant_NewQVariant34(void* val);
void* QVariant_NewQVariant36(void* regExp);
void* QVariant_NewQVariant37(void* re);
void* QVariant_NewQVariant27(void* val);
void* QVariant_NewQVariant28(void* val);
void* QVariant_NewQVariant17(char* val);
void* QVariant_NewQVariant19(char* val);
void* QVariant_NewQVariant22(void* val);
void* QVariant_NewQVariant38(void* val);
void* QVariant_NewQVariant40(void* val);
void* QVariant_NewQVariant5(void* p);
void* QVariant_NewQVariant14(char* val);
void* QVariant_NewQVariant7(int val);
void* QVariant_ToByteArray(void* ptr);
void* QVariant_ToDateTime(void* ptr);
void* QVariant_ToEasingCurve(void* ptr);
void* QVariant_ToRegExp(void* ptr);
void* QVariant_ToRegularExpression(void* ptr);
char* QVariant_ToStringList(void* ptr);
void QVariant_DestroyQVariant(void* ptr);
void* QVariant_NewQVariant();
void* QVariant_NewQVariant6(void* s);
void* QVariant_NewQVariant47(void* other);
void QVariant_Clear(void* ptr);
int QVariant_Convert(void* ptr, int targetTypeId);
int QVariant_IsNull(void* ptr);
int QVariant_IsValid(void* ptr);
void QVariant_Swap(void* ptr, void* other);
int QVariant_ToBool(void* ptr);
int QVariant_ToInt(void* ptr, int ok);
void* QVariant_ToJsonArray(void* ptr);
void* QVariant_ToJsonDocument(void* ptr);
void* QVariant_ToJsonObject(void* ptr);
void* QVariant_ToModelIndex(void* ptr);
double QVariant_ToReal(void* ptr, int ok);
char* QVariant_ToString(void* ptr);
int QVariant_UserType(void* ptr);

#ifdef __cplusplus
}
#endif