package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	//"bytes"
	"reflect"
	"unsafe"
)

func myReverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func myToCharArray(x string) []string {
	arr := make([]string, len([]rune(x)), len([]rune(x)))
	for i := 0; i < len([]rune(x)); i++ {
		arr[i] = string(x[i])
	}
	return arr
}

func myCopyArrDblByte(arr1 [][]byte, pos1 int, arr2 [][]byte, pos2 int, count int) [][]byte {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([][]byte, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([][]byte, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([][]byte, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrByte(arr1 []byte, pos1 int, arr2 []byte, pos2 int, count int) []byte {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]byte, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]byte, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]byte, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrInt64(arr1 []int64, pos1 int, arr2 []int64, pos2 int, count int) []int64 {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]int64, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]int64, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]int64, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrBool(arr1 []bool, pos1 int, arr2 []bool, pos2 int, count int) []bool {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]bool, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]bool, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]bool, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrString(arr1 []string, pos1 int, arr2 []string, pos2 int, count int) []string {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]string, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]string, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]string, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrInt(arr1 []int, pos1 int, arr2 []int, pos2 int, count int) []int {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]int, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]int, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]int, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrEventTypeSchema(arr1 []EventTypeSchema, pos1 int, arr2 []EventTypeSchema, pos2 int, count int) []EventTypeSchema {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]EventTypeSchema, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]EventTypeSchema, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]EventTypeSchema, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

func myCopyArrEventType(arr1 []EventType, pos1 int, arr2 []EventType, pos2 int, count int) []EventType {
	var err error
	if count > len(arr1) {
		err = errors.New("count>len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]EventType, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos1 >= len(arr1) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]EventType, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	if pos2 >= len(arr2) {
		err = errors.New("pos1>=len(arr1)")
	}
	if err != nil {
		fmt.Println("errors")
		errorArr := make([]EventType, len(arr2), len(arr2))
		copy(arr2, errorArr)
		return arr2
	}

	k := 0
	j := pos1
	for i := pos2; i < len(arr2); i++ {
		arr2[i] = arr1[j]
		j++
		k++
		if k == count {
			break
		}
	}
	return arr2
}

//////////////////////////////////////////Transmogrifier.clas////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////GrammarCache.class//////////////////////////////////////////////////////////////////////////////////
var m_fragmentGrammar Grammar //fix problem

type GrammarCache struct {
	m_schema EXISchema
	//m_documentGrammar DocumentGrammar
	//m_fragmentGrammar Grammar
	elementFragmentGrammar          ElementFragmentGrammar
	exiGrammars                     []EXIGrammar
	exiGrammarUses                  []EXIGrammarUse
	grammarOptions                  int
	m_builtinElementGrammarTemplate BuiltinElementGrammar
}

func (g *GrammarCache) getEXISchema() EXISchema { //public EXISchema getEXISchema() {
	return g.m_schema
}
func (g *GrammarCache) GrammarCache3(var1 EXISchema, var2 int) { //public GrammarCache(EXISchema var1, short var2) {
	g.m_schema = var1                                      //this.m_schema = var1;
	g.grammarOptions = var2                                //this.grammarOptions = var2;
	if !(reflect.DeepEqual(g.m_schema, nilVAREXISchema)) { //if (this.m_schema != null) {
		var var3 int = var1.getTotalGrammarCount()                                                         //int var3 = var1.getTotalGrammarCount();
		g.exiGrammars = make([]EXIGrammar, var3, var3)                                                     //this.exiGrammars = new EXIGrammar[var3];
		var var4 []int = g.m_schema.getElems()                                                             //int[] var4 = this.m_schema.getElems();
		g.exiGrammarUses = make([]EXIGrammarUse, var1.getElemCountOfSchema(), var1.getElemCountOfSchema()) //this.exiGrammarUses = new EXIGrammarUse[var1.getElemCountOfSchema()];
		var var6 int = 0                                                                                   //int var6 = 0;

		var var5 int                             //int var5;
		var var8 int                             //int var8;
		for var5 = 0; var6 < len(var4); var5++ { //for(var5 = 0; var6 < var4.length; ++var5) {
			var var7 int = var1.getTypeOfElem(var6) //int var7 = var1.getTypeOfElem(var6);
			if var1.isSimpleType(var7) {            //var8 = var1.isSimpleType(var7) ? var7 : var1.getContentDatatypeOfComplexType(var7);
				var8 = var7
			} else {
				var8 = var1.getContentDatatypeOfComplexType(var7)
			}
			IVAREXIGrammarUse := VAREXIGrammarUse
			IVAREXIGrammarUse.EXIGrammarUse1(var8, *g) //to do!!!!(*g or g)!!!!!!!!!!!!!!!!!!
			g.exiGrammarUses[var5] = IVAREXIGrammarUse //this.exiGrammarUses[var5] = new EXIGrammarUse(var8, this);
			var6 = var6 + 4                            //var6 += 4;
		}

		var var14 []int = g.m_schema.getGrammars() //int[] var14 = this.m_schema.getGrammars();

		var var9 int                                                                   //int var9;
		for var8 = 0; var8 < len(var14); var8 = var8 + getSizeOfGrammar(var8, var14) { //for(var8 = 0; var8 < var14.length; var8 += EXISchema.getSizeOfGrammar(var8, var14)) {
			var9 = var1.getSerialOfGrammar(var8) //var9 = var1.getSerialOfGrammar(var8);

			//if g.exiGrammars[var9] == nil {//assert this.exiGrammars[var9] == null;

			g.exiGrammars[var9] = superEXIGrammar(g) //this.exiGrammars[var9] = new EXIGrammar(this);//super()!!!
		}

		for var8 = 0; var8 < len(var14); var8 += getSizeOfGrammar(var8, var14) { //for(var8 = 0; var8 < var14.length; var8 += EXISchema.getSizeOfGrammar(var8, var14)) {
			var9 = var1.getSerialOfGrammar(var8)          //var9 = var1.getSerialOfGrammar(var8);
			g.exiGrammars[var9].substantiate(var8, false) //this.exiGrammars[var9].substantiate(var8, false);
		}

		var var15 []EXIGrammar
		if var2 == 1 { //EXIGrammar[] var15 = var2 == 1 ? new EXIGrammar[3 * var3] : null
			var15 = make([]EXIGrammar, 3*var3, 3*var3)
		}
		var6 = 0 //var6 = 0;

		for var5 = 0; var6 < len(var4); var5++ { //for(var5 = 0; var6 < var4.length; ++var5) {
			if var2 == 1 { //if (var2 == 1) {
				var9 = 0                          //var9 = 0;
				if var1.isNillableElement(var6) { //if (var1.isNillableElement(var6)) {
					var9 |= 1 //var9 |= 1;
				}

				if var1.isTypableType(var1.getTypeOfElem(var6)) { //if (var1.isTypableType(var1.getTypeOfElem(var6))) {
					var9 |= 2 //var9 |= 2;
				}

				var var11 int = var1.getSerialOfGrammar(var1.getGrammarOfType(var1.getTypeOfElem(var6))) //int var11 = var1.getSerialOfGrammar(var1.getGrammarOfType(var1.getTypeOfElem(var6)));
				var var10 EXIGrammar                                                                     //EXIGrammar var10;
				if var9 == 0 {                                                                           //if (var9 == 0) {
					var10 = g.exiGrammars[var11] //var10 = this.exiGrammars[var11];
				} else { //} else {
					var var12 int = var3 * (var9 - 1)               ///int var12 = var3 * (var9 - 1);
					var var13 int = var12 + var11                   //int var13 = var12 + var11;
					if reflect.DeepEqual(var10, nilVAREXIGrammar) { //if ((var10 = var15[var13]) == null) {///to do??????????????????????????????????????????????????????????????????????
						var10 = superEXIGrammar(g)
						var10 = var15[var13]
						//var10 = new EXIGrammar(this);
						var10.substantiate(var6, true) //var10.substantiate(var6, true);
						var15[var13] = var10           //var15[var13] = var10;
					}
				}

				if reflect.DeepEqual(var10, nilVAREXIGrammar) { //assert var10 != null;
					panic("GrammarCache.class, func GrammarCache3, var10 == nil")
				}
				g.exiGrammarUses[var5].exiGrammar = var10 //this.exiGrammarUses[var5].exiGrammar = var10;
			} else { //} else {
				if (var2 & 1) != 0 { //assert (var2 & 1) == 0;
					panic("GrammarCache.class, func GrammarCache3, (var2 & 1) != 0 ")
				}
				var9 = var1.getTypeOfElem(var6)                                                                         //var9 = var1.getTypeOfElem(var6);
				g.exiGrammarUses[var5].exiGrammar = g.exiGrammars[var1.getSerialOfGrammar(var1.getGrammarOfType(var9))] //this.exiGrammarUses[var5].exiGrammar = this.exiGrammars[var1.getSerialOfGrammar(var1.getGrammarOfType(var9))];
			}

			var6 = var6 + 4 //var6 += 4;
		}

		g.elementFragmentGrammar = ElementFragmentGrammar(g) //this.elementFragmentGrammar = new ElementFragmentGrammar(this);
		m_fragmentGrammar = FragmentGrammar(g)               //this.m_fragmentGrammar = new FragmentGrammar(this);
	} else {
		g.elementFragmentGrammar = nilVARElementFragmentGrammar                     //this.elementFragmentGrammar = null;
		m_fragmentGrammar = varBuiltinFragmentGrammar.funcBuiltinFragmentGrammar(g) //this.m_fragmentGrammar = new BuiltinFragmentGrammar(this);
		g.exiGrammars = nil                                                         //this.exiGrammars = null;
		g.exiGrammarUses = nil                                                      //this.exiGrammarUses = null;
	}

	g.m_documentGrammar = DocumentGrammar(g) //this.m_documentGrammar = new DocumentGrammar(this);
	g.m_builtinElementGrammarTemplate        //this.m_builtinElementGrammarTemplate = new BuiltinElementGrammar("", this);
}

//BuiltinFragmentGrammar.class//
type BuiltinFragmentGrammar struct {
	m_eventTypeLists []EventTypeList  ///private final EventTypeList[] m_eventTypeLists;
	m_eventCodes     []EventCodeTuple //private final EventCodeTuple[] m_eventCodes;
	m_eventTypesInit []EventType      //private static final EventType[] m_eventTypesInit = new EventType[10];
}

var varBuiltinFragmentGrammar BuiltinFragmentGrammar = BuiltinFragmentGrammar{
	m_eventTypesInit: make([]EventType, 10, 10),
}

func (b *BuiltinFragmentGrammar) funcBuiltinFragmentGrammar(var1 GrammarCache) Grammar { //BuiltinFragmentGrammar(GrammarCache var1) {
	//super((byte)1, var1);
	var var2 int = var1.grammarOptions                                                                  //short var2 = var1.grammarOptions;
	b.m_eventTypeLists = make([]EventTypeList, 2, 2)                                                    //this.m_eventTypeLists = new EventTypeList[2];
	b.m_eventCodes = make([]EventCodeTuple, 2, 2)                                                       //this.m_eventCodes = new EventCodeTuple[2];
	var var3 ArrayEventTypeList                                                                         //ArrayEventTypeList var3;
	b.m_eventTypeLists[0] = (*(*EventTypeList)(unsafe.Pointer(&var3)))                                  //var3     //this.m_eventTypeLists[0] = var3 = new ArrayEventTypeList();
	var var4 ArrayEventCodeTuple                                                                        //ArrayEventCodeTuple var4;
	b.m_eventCodes[0] = (*(*EventCodeTuple)(unsafe.Pointer(&var4)))                                     //var4     //this.m_eventCodes[0] = var4 = new ArrayEventCodeTuple();
	var var5 []EventType = []EventType{createStartDocument((*(*EventTypeList)(unsafe.Pointer(&var3))))} //EventType[] var5 = new EventType[]{EventTypeFactory.createStartDocument(var3)};
	var3.setItems(var5)                                                                                 //var3.setItems(var5);
	var4.setItems((*(*([]EventCode))(unsafe.Pointer(&([]EventType{var5[0]})))))                         //var4.setItems(new EventType[]{var5[0]});
	b.populateContentGrammar(var2)                                                                      //this.populateContentGrammar(var2);
}

///end BuiltinFragmentGrammar.class//
//ArrayEventCodeTuple.class///!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
type ArrayEventCodeTuple struct {
	myArrParentEventCode [][]EventCode //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}

func (a *ArrayEventCodeTuple) setItems(var1 []EventCode) { //final void setItems(EventCode[] var1) {!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//assert var1 != null && var1.length > 0;
	var myvar1 EventCodeTuple = VAREventCodeTuple
	myvar1.eventCodes = var1                   //this.eventCodes = var1;
	myvar1.itemsCount = len(myvar1.eventCodes) //this.itemsCount = this.eventCodes.length;
	var var3 int = 0                           //int var3 = 0;

	for var2 := len(myvar1.eventCodes) - 1; var2 > 0; var3++ { //for(int var2 = this.eventCodes.length - 1; var2 > 0; ++var3) {
		var2 >>= 1 //var2 >>= 1;
	}

	myvar1.width = var3 //this.width = var3;

	var myArrParentEventCode []EventCode
	for var4 := 0; var4 < len(var1); var4++ { //for(int var4 = 0; var4 < var1.length; ++var4) {
		var myvarParentEventCode EventCode
		var var5 EventCode = var1[var4] //EventCode var5 = var1[var4];
		if var5 != nilVAREventCode {    //if (var5 != null) {
			myvarParentEventCode = var5.setParentalContext(var4, var5) //var5.setParentalContext(var4, this);
			if var5.itemType != -1 {                                   //if (var5.itemType != -1) {
				(*(*EventType)(unsafe.Pointer(&var5))).computeItemPath() //((EventType)var5).computeItemPath();//(*(*EventCode)(unsafe.Pointer(&var23)))
			}
		}
		myArrParentEventCode = append(myArrParentEventCode, myvarParentEventCode)
	}

	myvar1.headItem = var1[0]                                                     //this.headItem = var1[0];
	a.myArrParentEventCode = append(a.myArrParentEventCode, myArrParentEventCode) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}

///end ArrayEventCodeTuple.class///
//FragmentGrammar.class//
type FragmentGrammar struct {
	m_fragmentElems  []int                //private final int[] m_fragmentElems;
	m_eventTypes     [][]EventType        //private final EventType[][] m_eventTypes;
	m_eventCodes     []EventCodeTuple     //private final EventCodeTuple[] m_eventCodes;
	m_eventTypeLists []ArrayEventTypeList //private final ArrayEventTypeList[] m_eventTypeLists;
}

func (f *FragmentGrammar) thisFragmentGrammar(var1 GrammarCache) FragmentGrammar { //FragmentGrammar(GrammarCache var1) {
	//super((byte)3, var1);
	var var2 int = var1.grammarOptions //short var2 = var1.grammarOptions;
	var var3 int
	var myvarVARGrammar Grammar = VARGrammar

	if !(reflect.DeepEqual(myvarVARGrammar.schema, nilVAREXISchema)) { //int var3 = this.schema != null ? this.schema.getFragmentElemCount() : 0;
		var3 = myvarVARGrammar.schema.getFragmentElemCount()
	} else {
		var3 = 0
	}
	f.m_eventTypes = make([][]EventType, 2, 2)            //this.m_eventTypes = new EventType[2][];
	f.m_eventCodes = make([]EventCodeTuple, 2, 2)         //this.m_eventCodes = new EventCodeTuple[2];
	f.m_eventTypeLists = make([]ArrayEventTypeList, 2, 2) //this.m_eventTypeLists = new ArrayEventTypeList[2];

	var var4 int //int var4;
	//for(var4 = 0; var4 < 2; ++var4) {
	//this.m_eventTypeLists[var4] = new ArrayEventTypeList();
	//}

	var var11 []EventType = []EventType{createStartDocument((*(*EventTypeList)(unsafe.Pointer(&(f.m_eventTypeLists[0])))))} //EventType[] var11 = new EventType[]{EventTypeFactory.createStartDocument(this.m_eventTypeLists[0])};
	var var12 ArrayEventCodeTuple                                                                                           //ArrayEventCodeTuple var12 = new ArrayEventCodeTuple();
	var12.setItems(*(*([]EventCode))(unsafe.Pointer(&([]EventType{var11[0]}))))                                             //var12.setItems(new EventType[]{var11[0]});
	f.m_eventCodes[0] = (*(*EventCodeTuple)(unsafe.Pointer(&var12)))                                                        //this.m_eventCodes[0] = var12;
	f.m_eventTypes[0] = var11                                                                                               //this.m_eventTypes[0] = var11;
	var var5 bool = false                                                                                                   //boolean var5 = false;
	var var13 ArrayEventCodeTuple                                                                                           //ArrayEventCodeTuple var13 = null;
	var var7 int = var3 + 2                                                                                                 //int var7 = var3 + 2;
	var var8 int = 0                                                                                                        //int var8 = 0;
	var var9 bool                                                                                                           //boolean var9;
	if hasCM(var2) {                                                                                                        //if (var9 = GrammarOptions.hasCM(var2)) {
		var9 = hasCM(var2)
		var8++ //++var8;
	}

	var var10 bool   //boolean var10;
	if hasPI(var2) { //if (var10 = GrammarOptions.hasPI(var2)) {
		var10 = hasPI(var2)
		var8++ //++var8;
	}

	if var8 != 0 { //if (var8 != 0) {
		var5 = true //var5 = true;
		var7++      //++var7;
	}

	if !(reflect.DeepEqual(myvarVARGrammar.schema, nilVAREXISchema)) { //this.m_fragmentElems = this.schema != null ? this.schema.getFragmentINodes() : null;
		f.m_fragmentElems = myvarVARGrammar.schema.getFragmentINodes()
	} else {
		f.m_fragmentElems = nil
	}
	var var6 int
	if var5 { //int var6 = var5 ? var8 + var7 - 1 : var7;
		var6 = var8 + var7 - 1
	} else {
		var6 = var7
	}
	var11 = make([]EventType, var6, var6)                 //var11 = new EventType[var6];
	var var14 []EventCode = make([]EventCode, var7, var7) //EventCode[] var14 = new EventCode[var7];
	var var15 []EventCode                                 //EventCode[] var15 = null;

	var var16 int                          //int var16;
	for var16 = 0; var16 < var3; var16++ { //for(var16 = 0; var16 < var3; ++var16) {
		var var17 bool = true                                         //   boolean var17 = true;
		var var18 int                                                 //   int var18;
		if f.m_fragmentElems[var16]&-2147483648 /*0x80000000*/ != 0 { //   if (((var18 = this.m_fragmentElems[var16]) & Integer.MIN_VALUE) != 0) {
			var18 = f.m_fragmentElems[var16]
			var17 = false    //	  var17 = false;
			var18 = (^var18) //	  var18 = ~var18;
		}

		var var19 EXIGrammarUse //= nil //EXIGrammarUse var19 = null;
		if var17 {              //if (var17) {
			var19 = var1.exiGrammarUses[myvarVARGrammar.schema.getSerialOfElem(var18)] //var19 = var1.exiGrammarUses[this.schema.getSerialOfElem(var18)];
		}

		var var20 int = myvarVARGrammar.schema.getUriOfElem(var18)                                                                                                                                                               //int var20 = this.schema.getUriOfElem(var18);
		var var21 int = myvarVARGrammar.schema.getLocalNameOfElem(var18)                                                                                                                                                         //int var21 = this.schema.getLocalNameOfElem(var18);
		var var22 EventTypeElement = createStartElement(var20, var21, myvarVARGrammar.schema.uris[var20], myvarVARGrammar.schema.localNames[var20][var21], (*(*EventTypeList)(unsafe.Pointer(&(f.m_eventTypeLists[1])))), var19) //EventTypeElement var22 = EventTypeFactory.createStartElement(var20, var21, this.schema.uris[var20], this.schema.localNames[var20][var21], this.m_eventTypeLists[1], var19);
		var14[var16] = (*(*EventCode)(unsafe.Pointer(&(var11[var16]))))                                                                                                                                                          //var11[var16]                                                                                                                                                      //var14[var16] = var11[var16] = var22;
		var11[var16] = (*(*EventType)(unsafe.Pointer(&var22)))                                                                                                                                                                   //var22
	}

	var14[var16] = (*(*EventCode)(unsafe.Pointer(&(var11[var16]))))                                                               //var11[var16]
	var11[var16] = thisEventType1(byte(1), (*(*EventTypeList)(unsafe.Pointer(&(f.m_eventTypeLists[1])))), 5 /*, (IGrammar)null*/) //var14[var16] = var11[var16] = new EventType((byte)1, this.m_eventTypeLists[1], (byte)5, (IGrammar)null);
	var16++                                                                                                                       //++var16;
	var14[var16] = (*(*EventCode)(unsafe.Pointer(&(var11[var16]))))                                                               //var11[var16]
	var11[var16] = createEndDocument((*(*EventTypeList)(unsafe.Pointer(&(f.m_eventTypeLists[1])))))                               //var14[var16] = var11[var16] = EventTypeFactory.createEndDocument(this.m_eventTypeLists[1]);
	var16++                                                                                                                       //++var16;
	if var5 {                                                                                                                     //if (var5) {
		var var15 []EventCode = make([]EventCode, var8, var8) //var15 = new EventCode[var8];
		var myvar1 ArrayEventCodeTuple
		var13 = myvar1                                         //var13 = new ArrayEventCodeTuple();
		var14[var16] = (*(*EventCode)(unsafe.Pointer(&var13))) //var13 //var14[var16] = var13;
		var var23 int = 0                                      //int var23 = 0;
		var var24 EventType                                    //EventType var24;
		if var9 {                                              //if (var9) {
			var24 = thisEventType1(byte(2), (*(*EventTypeList)(unsafe.Pointer(&(f.m_eventTypeLists[1])))), 1 /*, (IGrammar)null*/) //var24 = new EventType((byte)2, this.m_eventTypeLists[1], (byte)1, (IGrammar)null);
			var11[var16] = var24                                                                                                   //var11[var16++] = var24;
			var16++
			var15[var23] = (*(*EventCode)(unsafe.Pointer(&var24))) //var24 //var15[var23++] = var24;
			var23++
		}

		if var10 { //if (var10) {
			var24 = thisEventType1(byte(2), (*(*EventTypeList)(unsafe.Pointer(&(f.m_eventTypeLists[1])))), 0 /*, (IGrammar)null*/) //var24 = new EventType((byte)2, this.m_eventTypeLists[1], (byte)0, (IGrammar)null);
			var11[var16] = var24                                                                                                   //var11[var16++] = var24;
			var16++
			var15[var23] = (*(*EventCode)(unsafe.Pointer(&var24))) //var15[var23++] = var24;
			var23++
		}
	}

	var myvar2 ArrayEventCodeTuple
	var12 = myvar2      //var12 = new ArrayEventCodeTuple();
	if len(var11) > 0 { //if (var11.length > 0) {
		var12.setItems(var14) //var12.setItems(var14);
		if var5 {             //if (var5) {
			var13.setItems(var15) //var13.setItems(var15);
		}
	}

	f.m_eventCodes[1] = (*(*EventCodeTuple)(unsafe.Pointer(&var12))) //var12 //this.m_eventCodes[1] = var12;
	f.m_eventTypes[1] = var11                                        //this.m_eventTypes[1] = var11;

	for var4 = 0; var4 < 2; var4++ { //for(var4 = 0; var4 < 2; ++var4) {
		f.m_eventTypeLists[var4].setItems(f.m_eventTypes[var4]) //this.m_eventTypeLists[var4].setItems(this.m_eventTypes[var4]);
	}

	return *f ///to do
}

//end FragmentGrammar.class//
//ElementFragmentGrammar.class///
type ElementFragmentGrammar struct {
	m_fragmentINodes []int                ///private final int[] m_fragmentINodes;
	m_eventTypes     [][]EventType        ///private final EventType[][] m_eventTypes;
	m_eventCodes     []EventCodeTuple     ///private final EventCodeTuple[] m_eventCodes;
	m_eventTypeLists []ArrayEventTypeList ///private final ArrayEventTypeList[] m_eventTypeLists;
}

var nilVARElementFragmentGrammar ElementFragmentGrammar = ElementFragmentGrammar{}

func (e *ElementFragmentGrammar) ElementFragmentGrammar(var1 GrammarCache) ElementFragmentGrammar { //ElementFragmentGrammar(GrammarCache var1) {
	//super((byte)4, var1);
	var myVARGrammar Grammar = VARGrammar
	e.m_fragmentINodes = myVARGrammar.schema.getFragmentINodes() //this.m_fragmentINodes = this.schema.getFragmentINodes();
	var var2 int = myVARGrammar.schema.getFragmentElemCount()    //int var2 = this.schema.getFragmentElemCount();
	var var3 int = len(e.m_fragmentINodes) - var2                //int var3 = this.m_fragmentINodes.length - var2;
	e.m_eventTypes = make([][]EventType, 4, 4)                   //this.m_eventTypes = new EventType[4][];
	e.m_eventCodes = make([]EventCodeTuple, 4, 4)                //this.m_eventCodes = new EventCodeTuple[4];
	e.m_eventTypeLists = make([]ArrayEventTypeList, 4, 4)        //this.m_eventTypeLists = new ArrayEventTypeList[4];

	var var4 int //int var4;
	//for(var4 = 0; var4 < 4; ++var4) {
	//this.m_eventTypeLists[var4] = new ArrayEventTypeList();
	//}

	var var5 bool = isPermitDeviation(var1.grammarOptions) //boolean var5 = GrammarOptions.isPermitDeviation(var1.grammarOptions);
	var var8 []string                                      //ArrayList var8 = null;
	var var6 []string = make([]string, 0, 1)               //ArrayList var6 = new ArrayList();
	if var5 {                                              //if (var5) {
		var8 = make([]string, 0, 1) //var8 = new ArrayList();
	}

	var var9 EventTypeSchema            //EventTypeSchema var9;
	var var10 int                       //int var10;
	var var11 bool                      //boolean var11;
	var var12 int                       //int var12;
	for var4 = 0; var4 < var3; var4++ { //for(var4 = 0; var4 < var3; ++var4) {
		var10 = var2 + var4                                                  //var10 = var2 + var4;
		if ((e.m_fragmentINodes[var10]) & -2147483648 /*0x80000000*/) != 0 { //if (var11 = ((var12 = this.m_fragmentINodes[var10]) & Integer.MIN_VALUE) != 0) {
			var11 = (e.m_fragmentINodes[var10] & -2147483648 /*0x80000000*/)
			var12 = e.m_fragmentINodes[var10]
			var12 = (^var12) //var12 = ~var12;
		}

		var9 = createAttributeEventType(var12, var11, (*(*EventTypeList)(unsafe.Pointer(&e.m_eventTypeLists[0]))), myVARGrammar) //var9 = this.createAttributeEventType(var12, var11, this.m_eventTypeLists[0]);!!!!!!!!!!!!!!!!!!!add new var in function!!!!!!!!!!!!!!!myVARGrammar!!!!!!!!!!!!!!!!!!!!!!!!!!!!1
		var6 = append(var6, var9)                                                                                                //var6.add(var9);
		if var5 {                                                                                                                //if (var5) {
			var8 = append(var8, createEventTypeSchemaAttributeInvalid((*(*EventType)(unsafe.Pointer(&var9))), (*(*EventTypeList)(unsafe.Pointer(&e.m_eventTypeLists[0]))))) //var8.add(this.createEventTypeSchemaAttributeInvalid(var9, this.m_eventTypeLists[0]));
		}
	}

	var6 = append(var6, thisEventType1(byte(1), e.m_eventTypeLists[0], 17 /*, (IGrammar)null)*/)) //var6.add(new EventType((byte)1, this.m_eventTypeLists[0], (byte)17, (IGrammar)null));

	var var13 int                       //int var 13;
	var var14 EventTypeElement          //EventTypeElement var14;
	var var15 EXIGrammarUse             //EXIGrammarUse var15;
	var var16 int                       //int var16;
	for var4 = 0; var4 < var2; var4++ { //for(var4 = 0; var4 < var2; ++var4) {
		if ((e.m_fragmentINodes[var4]) & -2147483648 /*0x80000000*/) != 0 { //if (((var16 = this.m_fragmentINodes[var4]) & Integer.MIN_VALUE) != 0) {

			var16 = e.m_fragmentINodes[var4]
			var16 = (^var16) //var16 = ~var16;
			//var15 = null;
		} else { //} else {
			var15 = var1.exiGrammarUses[myVARGrammar.schema.getSerialOfElem(var16)] //var15 = var1.exiGrammarUses[this.schema.getSerialOfElem(var16)];
		}

		var12 = myVARGrammar.schema.getUriOfElem(var16)                                                                                                                                               //var12 = this.schema.getUriOfElem(var16);
		var13 = myVARGrammar.schema.getLocalNameOfElem(var16)                                                                                                                                         //var13 = this.schema.getLocalNameOfElem(var16);
		var14 = createStartElement(var12, var13, myVARGrammar.schema.uris[var12], myVARGrammar.schema.localNames[var12][var13], (*(*EventTypeList)(unsafe.Pointer(&(e.m_eventTypeLists[0])))), var15) //var14 = EventTypeFactory.createStartElement(var12, var13, this.schema.uris[var12], this.schema.localNames[var12][var13], this.m_eventTypeLists[0], var15);
		var6 = append(var6, var14)                                                                                                                                                                    //var6.add(var14);
	}

	var6 = append(var6, thisEventType1(byte(1), e.m_eventTypeLists[0], 5 /*, (IGrammar)null)*/))                   //var6.add(new EventType((byte)1, this.m_eventTypeLists[0], (byte)5, (IGrammar)null));
	var6 = append(var6, EventTypeFactory.creatEndElement(byte(1), e.m_eventTypeLists[0]))                          //var6.add(EventTypeFactory.creatEndElement((byte)1, this.m_eventTypeLists[0]));
	var6 = append(var6, thisEventType1(byte(1), e.m_eventTypeLists[0], 3 /*, (IGrammar)null)*/))                   //var6.add(new EventType((byte)1, this.m_eventTypeLists[0], (byte)3, (IGrammar)null));
	var var7 EventCodeTupleSink                                                                                    //EventCodeTupleSink var7;
	e.createEventCodeTuple1(var6, var1.grammarOptions, var7, var8, e.m_eventTypeLists[0], true, -1, false, -1, -1) //this.createEventCodeTuple(var6, var1.grammarOptions, var7 = new EventCodeTupleSink(), var8, this.m_eventTypeLists[0], true, -1, false, -1, -1);
	e.m_eventCodes[0] = var7.eventCodeTuple                                                                        //this.m_eventCodes[0] = var7.eventCodeTuple;
	e.m_eventTypes[0] = (*(*[]EventType)(unsafe.Pointer(&var7.eventTypes)))                                        //var7.eventTypes       //this.m_eventTypes[0] = var7.eventTypes;
	e.m_eventTypeLists[0].setItems(*(*[]EventType)(unsafe.Pointer(&var7.eventTypes)))                              //(var7.eventTypes)     //this.m_eventTypeLists[0].setItems(var7.eventTypes);
	var6 = make([]string, 0, 1)                                                                                    //var6 = new ArrayList();

	for var4 = 0; var4 < var2; var4++ { //for(var4 = 0; var4 < var2; ++var4) {
		if ((e.m_fragmentINodes[var4]) & -2147483648 /*0x80000000*/) != 0 { //if (((var16 = this.m_fragmentINodes[var4]) & Integer.MIN_VALUE) != 0) {
			var16 = e.m_fragmentINodes[var4]
			var16 = (^var16) //var16 = ~var16;
			//var15 = null;
		} else {
			var15 = var1.exiGrammarUses[myVARGrammar.schema.getSerialOfElem(var16)] //var15 = var1.exiGrammarUses[this.schema.getSerialOfElem(var16)];
		}

		var12 = myVARGrammar.schema.getUriOfElem(var16)                                                                                                                                               //var12 = this.schema.getUriOfElem(var16);
		var13 = myVARGrammar.schema.getLocalNameOfElem(var16)                                                                                                                                         //var13 = this.schema.getLocalNameOfElem(var16);
		var14 = createStartElement(var12, var13, myVARGrammar.schema.uris[var12], myVARGrammar.schema.localNames[var12][var13], (*(*EventTypeList)(unsafe.Pointer(&(e.m_eventTypeLists[1])))), var15) //var14 = EventTypeFactory.createStartElement(var12, var13, this.schema.uris[var12], this.schema.localNames[var12][var13], this.m_eventTypeLists[1], var15);
		var6 = append(var6, var14)                                                                                                                                                                    //var6.add(var14);
	}

	var6 = append(var6, thisEventType1(byte(1), e.m_eventTypeLists[1], 5 /*, (IGrammar)null*/)) //var6.add(new EventType((byte)1, this.m_eventTypeLists[1], (byte)5, (IGrammar)null));
	var6 = append(var6, EventTypeFactory.creatEndElement(byte(1), e.m_eventTypeLists[1]))       //var6.add(EventTypeFactory.creatEndElement((byte)1, this.m_eventTypeLists[1]));
	var6 = append(var6, thisEventType1(byte(1), e.m_eventTypeLists[1], 3 /*, (IGrammar)null*/)) //var6.add(new EventType((byte)1, this.m_eventTypeLists[1], (byte)3, (IGrammar)null));
	var myvar EventCodeTupleSink
	var7 = myvar
	e.createEventCodeTuple(var6, var1.grammarOptions, var7, e.m_eventTypeLists[1])    //this.createEventCodeTuple(var6, var1.grammarOptions, var7 = new EventCodeTupleSink(), this.m_eventTypeLists[1]);
	e.m_eventCodes[1] = var7.eventCodeTuple                                           //this.m_eventCodes[1] = var7.eventCodeTuple;
	e.m_eventTypes[1] = (*(*[]EventType)(unsafe.Pointer(&var7.eventTypes)))           //var7.eventTypes    //this.m_eventTypes[1] = var7.eventTypes;
	e.m_eventTypeLists[1].setItems(*(*[]EventType)(unsafe.Pointer(&var7.eventTypes))) //(var7.eventTypes)                                //this.m_eventTypeLists[1].setItems(var7.eventTypes);
	var6 = make([]string, 0, 1)                                                       //var6 = new ArrayList();
	if var5 {                                                                         //if (var5) {
		var8.clear() //var8.clear();
	}

	for var4 = 0; var4 < var3; var4++ { //for(var4 = 0; var4 < var3; ++var4) {
		var10 = var2 + var4                                                // var10 = var2 + var4;
		if (e.m_fragmentINodes[var10] & -2147483648 /*0x80000000*/) != 0 { //if (var11 = ((var12 = this.m_fragmentINodes[var10]) & Integer.MIN_VALUE) != 0) {
			var12 = e.m_fragmentINodes[var10]
			var11 = ((e.m_fragmentINodes[var10]) & -2147483648 /*0x80000000*/)
			var12 = (^var12) //var12 = ~var12;
		}

		var9 = createAttributeEventType(var12, var11, (*(*EventTypeList)(unsafe.Pointer(&(e.m_eventTypeLists[2])))), myVARGrammar) //var9 = this.createAttributeEventType(var12, var11, this.m_eventTypeLists[2]);!!!!!!!!!!!!!!!!!!!!!!!!!!!!!add new war in func!!!!!!!!!!!!myVARGrammar!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		var6 = append(var6, var9)                                                                                                  //var6.add(var9);
		if var5 {                                                                                                                  //if (var5) {
			var8 = append(var8, e.createEventTypeSchemaAttributeInvalid(var9, e.m_eventTypeLists[2])) //var8.add(this.createEventTypeSchemaAttributeInvalid(var9, this.m_eventTypeLists[2]));
		}
	}

	var6 = append(var6, thisEventType1(byte(1), e.m_eventTypeLists[2], 17 /*, (IGrammar)null*/)) //var6.add(new EventType((byte)1, this.m_eventTypeLists[2], (byte)17, (IGrammar)null));
	var6 = append(var6, creatEndElement(byte(1), e.m_eventTypeLists[2]))                         //var6.add(EventTypeFactory.creatEndElement((byte)1, this.m_eventTypeLists[2]));
	var7 = myvar
	e.createEventCodeTuple(var6, var1.grammarOptions, var7, var8, e.m_eventTypeLists[2], ((var1.grammarOptions & 1) == 0), -1, false, -1, -1) //this.createEventCodeTuple(var6, var1.grammarOptions, var7 = new EventCodeTupleSink(), var8, this.m_eventTypeLists[2], (var1.grammarOptions & 1) == 0, -1, false, -1, -1);
	e.m_eventCodes[2] = var7.eventCodeTuple                                                                                                   //this.m_eventCodes[2] = var7.eventCodeTuple;
	e.m_eventTypes[2] = (*(*[]EventType)(unsafe.Pointer(&var7.eventTypes)))                                                                   //var7.eventTypes      //this.m_eventTypes[2] = var7.eventTypes;
	e.m_eventTypeLists[2].setItems(*(*[]EventType)(unsafe.Pointer(&var7.eventTypes)))                                                         //(var7.eventTypes)    //this.m_eventTypeLists[2].setItems(var7.eventTypes);
	var6 = make([]string, 0, 1)                                                                                                               //var6 = new ArrayList();
	var6 = append(var6)                                                                                                                       //var6.add(EventTypeFactory.creatEndElement((byte)1, this.m_eventTypeLists[3]));
	var7 = myvar
	e.createEventCodeTuple(var6, var1.grammarOptions, var7, e.m_eventTypeLists[3])    //this.createEventCodeTuple(var6, var1.grammarOptions, var7 = new EventCodeTupleSink(), this.m_eventTypeLists[3]);
	e.m_eventCodes[3] = var7.eventCodeTuple                                           //this.m_eventCodes[3] = var7.eventCodeTuple;
	e.m_eventTypes[3] = (*(*[]EventType)(unsafe.Pointer(&var7.eventTypes)))           //var7.eventTypes   //this.m_eventTypes[3] = var7.eventTypes;
	e.m_eventTypeLists[3].setItems(*(*[]EventType)(unsafe.Pointer(&var7.eventTypes))) //(var7.eventTypes)    //this.m_eventTypeLists[3].setItems(var7.eventTypes);
	return *e
}

func createAttributeEventType(var1 int, var2 bool, var3 EventTypeList, myVARGrammar Grammar) EventTypeSchema { //private EventTypeSchema createAttributeEventType(int var1, boolean var2, EventTypeList var3) { !!!!!!!!!!!!!!!!add new var in func !!!!!!!!!!!!myVARGrammar!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	var var4 int = myVARGrammar.schema.getUriOfAttr(var1)       //int var4 = this.schema.getUriOfAttr(var1);
	var var5 int = myVARGrammar.schema.getLocalNameOfAttr(var1) //int var5 = this.schema.getLocalNameOfAttr(var1);
	if var2 {
		-1
	}
	x := (EventTypeSchema1(myVARGrammar.schema.getTypeOfAttr(var1), myVARGrammar.schema.uris[var4], myVARGrammar.schema.localNames[var4][var5], var4, var5, byte(1), var3, byte(16), nilVAREXIGrammar))
	return (*(*EventTypeSchema)(unsafe.Pointer(&x)))
	//return new EventTypeSchema(var2 ? -1 : this.schema.getTypeOfAttr(var1), this.schema.uris[var4], this.schema.localNames[var4][var5], var4, var5, (byte)1, var3, (byte)16, (EXIGrammar)null);
}

//end ElementFragmentGrammar.class///
////////////////////EXISchema.class/////////////////////////
/////////Characters.class////////
type Characters struct { //public final class Characters {
	isVolatile bool     //	public boolean isVolatile;
	characters []string //	public char[] characters;
	startIndex int      //	public int startIndex;
	length     int      //	public final int length;
	ucsCount   int      //	public final int ucsCount;
	m_hashCode int      //	private final int m_hashCode;
	//	public static final Characters CHARACTERS_EMPTY = new Characters("".toCharArray(), 0, 0, false);
}

var CHARACTERS_EMPTY Characters = Characters{
	isVolatile: false,
	characters: []string{""},
	startIndex: 0,
	length:     0,
	ucsCount:   0,
	m_hashCode: 0,
}

////to do  проверить на правельность заполнение!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (c *Characters) Characters1(var1 []string, var2 int, var3 int, var4 bool) Characters { //	public Characters(char[] var1, int var2, int var3, boolean var4) {
	c.isVolatile = var4                    //	   this.isVolatile = var4;
	c.characters = var1                    //	   this.characters = var1;
	c.startIndex = var2                    //	   this.startIndex = var2;
	c.length = var3                        //       this.length = var3;
	var var5 int = 0                       //	   int var5 = 0;
	var var6 int = c.startIndex + c.length //	   int var6 = this.startIndex + this.length;
	var var7 int = 0                       //	   int var7 = 0;

	for var8 := c.startIndex; var8 < var6; var7++ { // for(int var8 = this.startIndex; var8 < var6; ++var7) {

		var var9 string = c.characters[var8] // char var9 = this.characters[var8++];
		var8++
		var99, err := strconv.Atoi(var9)
		if err != nil {
			panic(err)
		} else {
			var5 = (var5 * 31) + var99 // var5 = var5 * 31 + var9;
			x, err := strconv.Atoi(var9)
			if err != nil {
				fmt.Println(err)
			}
			if (x&'\ufc00' /*[64512]*/) == 55296 && var8 < var6 { //  if ((var9 & '\ufc00') == 55296 && var8 < var6) {
				var9 = c.characters[var8]
				y, err := strconv.Atoi(var9) //			 var9 = this.characters[var8];
				if err != nil {
					fmt.Println(err)
				}
				if (y & '\ufc00' /*[64512]*/) == 56320 { //			 if ((var9 & '\ufc00') == 56320) {
					var8++ //				++var8;
				}
			}
		}
	}

	c.ucsCount = var7   // this.ucsCount = var7;
	c.m_hashCode = var5 // this.m_hashCode = var5;
	var result Characters = Characters{
		isVolatile: var4,
		characters: var1,
		startIndex: var2,
		length:     var3,
		ucsCount:   var7,
		m_hashCode: var5,
	}
	return result
}

/////////end Characters.class////////

type EXISchema struct {
	COOKIE                      []byte
	NIL_NODE                    int
	NIL_VALUE                   int
	EMPTY_STRING                int
	NIL_GRAM                    int
	EVENT_AT_WILDCARD           int
	EVENT_SE_WILDCARD           int
	EVENT_CH_UNTYPED            int
	EVENT_CH_TYPED              int
	MIN_EVENT_ID                int
	EVENT_TYPE_AT               byte
	EVENT_TYPE_SE               byte
	EVENT_TYPE_AT_WILDCARD_NS   byte
	EVENT_TYPE_SE_WILDCARD_NS   byte
	TRUE_VALUE                  int
	FALSE_VALUE                 int
	UNBOUNDED_OCCURS            int
	CONSTRAINT_NONE             int
	CONSTRAINT_DEFAULT          int
	CONSTRAINT_FIXED            int
	WHITESPACE_PRESERVE         int
	WHITESPACE_REPLACE          int
	WHITESPACE_COLLAPSE         int
	VARIANT_STRING              byte
	VARIANT_FLOAT               byte
	VARIANT_DECIMAL             byte
	VARIANT_INTEGER             byte
	VARIANT_INT                 byte
	VARIANT_LONG                byte
	VARIANT_DATETIME            byte
	VARIANT_DURATION            byte
	VARIANT_BASE64              byte
	VARIANT_BOOLEAN             byte
	VARIANT_HEXBIN              byte
	INTEGER_CODEC_DEFAULT       int
	INTEGER_CODEC_NONNEGATIVE   int
	ANCESTRY_IDS                []int //[]byte
	ELEMENT_NAMES               []string
	DEFAULT_TYPABLES            []bool
	UR_SIMPLE_TYPE              byte
	ATOMIC_SIMPLE_TYPE          byte
	LIST_SIMPLE_TYPE            byte
	UNION_SIMPLE_TYPE           byte
	m_elems                     []int
	m_attrs                     []int
	m_types                     []int
	uris                        []string
	localNames                  [][]string
	m_localNames                [][]int
	m_names                     []string
	m_strings                   []string
	m_ints                      []int
	m_mantissas                 []int
	m_exponents                 []int
	m_signs                     []bool
	m_integralDigits            []string
	m_reverseFractionalDigits   []string
	m_integers                  []int //big.int
	m_longs                     []int
	m_datetimes                 []string //XSDateTime[]
	m_durations                 []int64  //[]Duration
	m_binaries                  [][]byte
	m_variantTypes              []byte
	m_variants                  []int
	m_computedDatetimes         []string     //XSDateTime[]
	m_variantCharacters         []Characters //string //[]Characters
	m_n_stypes                  int
	ancestryIds                 []int //[]byte
	m_stypes_end                int
	m_grammars                  []int
	m_productions               []int
	m_eventTypes                []byte
	m_eventData                 []int
	m_grammarCount              int
	m_fragmentINodes            []int
	m_n_fragmentElems           int
	m_globalElementsDirectory   map[string][]int
	m_globalAttributesDirectory map[string][]int
	m_globalTypesDirectory      map[string][]int
	m_buitinTypes               []int
	m_globalElems               []int
	m_globalAttrs               []int
	//datatypeFactory             DatatypeFactory///to do
}

var VAREXISchema = EXISchema{
	COOKIE:                    []byte{36, 51, 43, 45},
	NIL_NODE:                  -1,
	NIL_VALUE:                 -1,
	EMPTY_STRING:              0,
	NIL_GRAM:                  -1,
	EVENT_AT_WILDCARD:         -1,
	EVENT_SE_WILDCARD:         -2,
	EVENT_CH_UNTYPED:          -3,
	EVENT_CH_TYPED:            -4,
	MIN_EVENT_ID:              -4,
	EVENT_TYPE_AT:             0,
	EVENT_TYPE_SE:             1,
	EVENT_TYPE_AT_WILDCARD_NS: 2,
	EVENT_TYPE_SE_WILDCARD_NS: 3,
	TRUE_VALUE:                1,
	FALSE_VALUE:               0,
	UNBOUNDED_OCCURS:          -1,
	CONSTRAINT_NONE:           0,
	CONSTRAINT_DEFAULT:        1,
	CONSTRAINT_FIXED:          2,
	WHITESPACE_PRESERVE:       0,
	WHITESPACE_REPLACE:        1,
	WHITESPACE_COLLAPSE:       2,
	VARIANT_STRING:            0,
	VARIANT_FLOAT:             1,
	VARIANT_DECIMAL:           2,
	VARIANT_INTEGER:           3,
	VARIANT_INT:               4,
	VARIANT_LONG:              5,
	VARIANT_DATETIME:          6,
	VARIANT_DURATION:          7,
	VARIANT_BASE64:            8,
	VARIANT_BOOLEAN:           9,
	VARIANT_HEXBIN:            10,
	INTEGER_CODEC_DEFAULT:     255,
	INTEGER_CODEC_NONNEGATIVE: 254,
	ANCESTRY_IDS:              []int{-1, -1, 2, 3, 4, 5, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, -1, -1, 21}, //[]byte
	ELEMENT_NAMES:             []string{"", "", "StringType", "BooleanType", "DecimalType", "FloatType", "FloatType", "DurationType", "DateTimeType", "TimeType", "DateType", "GYearMonthType", "GYearType", "GMonthDayType", "GDayType", "GMonthType", "HexBinaryType", "Base64BinaryType", "AnyURIType", "QNameType", "QNameType", "IntegerType"},
	DEFAULT_TYPABLES:          []bool{true, true, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, false, true, false, true, true, false, true, false, true, true, true, true, false, true, true, false, false, false, false, false, false, false},
	UR_SIMPLE_TYPE:            0,
	ATOMIC_SIMPLE_TYPE:        1,
	LIST_SIMPLE_TYPE:          2,
	UNION_SIMPLE_TYPE:         3,
}
var nilVAREXISchema = EXISchema{
	COOKIE:                      make([]byte, 0, 0),
	NIL_NODE:                    0,
	NIL_VALUE:                   0,
	EMPTY_STRING:                0,
	NIL_GRAM:                    0,
	EVENT_AT_WILDCARD:           0,
	EVENT_SE_WILDCARD:           0,
	EVENT_CH_UNTYPED:            0,
	EVENT_CH_TYPED:              0,
	MIN_EVENT_ID:                0,
	EVENT_TYPE_AT:               0,
	EVENT_TYPE_SE:               0,
	EVENT_TYPE_AT_WILDCARD_NS:   0,
	EVENT_TYPE_SE_WILDCARD_NS:   0,
	TRUE_VALUE:                  0,
	FALSE_VALUE:                 0,
	UNBOUNDED_OCCURS:            0,
	CONSTRAINT_NONE:             0,
	CONSTRAINT_DEFAULT:          0,
	CONSTRAINT_FIXED:            0,
	WHITESPACE_PRESERVE:         0,
	WHITESPACE_REPLACE:          0,
	WHITESPACE_COLLAPSE:         0,
	VARIANT_STRING:              0,
	VARIANT_FLOAT:               0,
	VARIANT_DECIMAL:             0,
	VARIANT_INTEGER:             0,
	VARIANT_INT:                 0,
	VARIANT_LONG:                0,
	VARIANT_DATETIME:            0,
	VARIANT_DURATION:            0,
	VARIANT_BASE64:              0,
	VARIANT_BOOLEAN:             0,
	VARIANT_HEXBIN:              0,
	INTEGER_CODEC_DEFAULT:       0,
	INTEGER_CODEC_NONNEGATIVE:   0,
	ANCESTRY_IDS:                make([]int, 0, 0), //[]byte
	ELEMENT_NAMES:               make([]string, 0, 0),
	DEFAULT_TYPABLES:            make([]bool, 0, 0),
	UR_SIMPLE_TYPE:              0,
	ATOMIC_SIMPLE_TYPE:          0,
	LIST_SIMPLE_TYPE:            0,
	UNION_SIMPLE_TYPE:           0,
	m_elems:                     make([]int, 0, 0),
	m_attrs:                     make([]int, 0, 0),
	m_types:                     make([]int, 0, 0),
	uris:                        make([]string, 0, 0),
	localNames:                  make([][]string, 0, 0),
	m_localNames:                make([][]int, 0, 0),
	m_names:                     make([]string, 0, 0),
	m_strings:                   make([]string, 0, 0),
	m_ints:                      make([]int, 0, 0),
	m_mantissas:                 make([]int, 0, 0),
	m_exponents:                 make([]int, 0, 0),
	m_signs:                     make([]bool, 0, 0),
	m_integralDigits:            make([]string, 0, 0),
	m_reverseFractionalDigits:   make([]string, 0, 0),
	m_integers:                  make([]int, 0, 0), //big.int
	m_longs:                     make([]int, 0, 0),
	m_datetimes:                 make([]string, 0, 0), //XSDateTime[]
	m_durations:                 make([]int64, 0, 0),  //[]Duration
	m_binaries:                  make([][]byte, 0, 0),
	m_variantTypes:              make([]byte, 0, 0),
	m_variants:                  make([]int, 0, 0),
	m_computedDatetimes:         make([]string, 0, 0),     //XSDateTime[]
	m_variantCharacters:         make([]Characters, 0, 0), //string //[]Characters
	m_n_stypes:                  0,
	ancestryIds:                 make([]int, 0, 0), //[]byte
	m_stypes_end:                0,
	m_grammars:                  make([]int, 0, 0),
	m_productions:               make([]int, 0, 0),
	m_eventTypes:                make([]byte, 0, 0),
	m_eventData:                 make([]int, 0, 0),
	m_grammarCount:              0,
	m_fragmentINodes:            make([]int, 0, 0),
	m_n_fragmentElems:           0,
	m_globalElementsDirectory:   make(map[string][]int),
	m_globalAttributesDirectory: make(map[string][]int),
	m_globalTypesDirectory:      make(map[string][]int),
	m_buitinTypes:               make([]int, 0, 0),
	m_globalElems:               make([]int, 0, 0),
	m_globalAttrs:               make([]int, 0, 0),
	//datatypeFactory    :         DatatypeFactory///to do
}

func (e *EXISchema) getFragmentINodes() []int { //public int[] getFragmentINodes() {
	return e.m_fragmentINodes //return this.m_fragmentINodes;
}

func (e *EXISchema) getFragmentElemCount() int { //public int getFragmentElemCount() {
	return e.m_n_fragmentElems //return this.m_n_fragmentElems;
}

func (e *EXISchema) isTypableType(var1 int) bool { //public boolean isTypableType(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func isTypableType, 0 > var1")
	}
	return e.m_types[var1+3] != 0 //return this.m_types[var1 + 3] != 0;
}

func (e *EXISchema) getTypeEmptyGrammarOfGrammar(var1 int) int { //public int getTypeEmptyGrammarOfGrammar(int var1) {
	return _getTypeEmptyGrammarOfGrammar(var1, e.m_grammars) //return _getTypeEmptyGrammarOfGrammar(var1, this.m_grammars);
}

func _getTypeEmptyGrammarOfGrammar(var0 int, var1 []int) int { //public static int _getTypeEmptyGrammarOfGrammar(int var0, int[] var1) {
	if _hasEmptyGrammar(var0, var1) {
		return var1[var0+3]
	} else {
		return -1
	}
	//return _hasEmptyGrammar(var0, var1) ? var1[var0 + 3] : -1;
}

func (e *EXISchema) hasEndElement(var1 int) bool { //public boolean hasEndElement(int var1) {
	return (e.m_grammars[var1+1] & 65536) != 0 //return (this.m_grammars[var1 + 1] & 65536) != 0;
}

func (e *EXISchema) getEventType(var1 int) byte { //public byte getEventType(int var1) {
	return e.m_eventTypes[var1]
}

func (e *EXISchema) getGrammarOfProduction(var1 int) int { //public int getGrammarOfProduction(int var1) {
	return e.m_productions[var1+1] //return this.m_productions[var1 + 1];
}

func (e *EXISchema) getProductionCountOfGrammar(var1 int) int { //public int getProductionCountOfGrammar(int var1) {
	return _getProductionCountOfGrammar(var1, e.m_grammars) //return _getProductionCountOfGrammar(var1, this.m_grammars);
}

func (e *EXISchema) getContentGrammarOfGrammar(var1 int) int { //public int getContentGrammarOfGrammar(int var1) {
	return _getContentGrammarOfGrammar(var1, e.m_grammars)
}

func _getContentGrammarOfGrammar(var0 int, var1 []int) int { //public static int _getContentGrammarOfGrammar(int var0, int[] var1) {

	if _hasContentGrammar(var0, var1) { //return _hasContentGrammar(var0, var1) ? var1[var0 + 2] : -1;
		return var1[var0+2]
	} else {
		return -1
	}
}

func (e *EXISchema) getGrammarOfType(var1 int) int { //public int getGrammarOfType(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getGrammarOfType, 0 > var1")
	}
	return e.m_types[var1+4] //return this.m_types[var1 + 4];
}

func (e *EXISchema) getSerialOfGrammar(var1 int) int { //public int getSerialOfGrammar(int var1) {
	return e.m_grammars[var1+0]
}

func getSizeOfGrammar(var0 int, var1 []int) int { //public static int getSizeOfGrammar(int var0, int[] var1) {
	var var2 int = 2                    //int var2 = 2;
	if _hasContentGrammar(var0, var1) { //if (_hasContentGrammar(var0, var1)) {
		var2++                            //++var2;
		if _hasEmptyGrammar(var0, var1) { //if (_hasEmptyGrammar(var0, var1)) {
			var2++ //++var2;
		}
	}

	return var2 + _getProductionCountOfGrammar(var0, var1)
}

func _getProductionCountOfGrammar(var0 int, var1 []int) int { //public static int _getProductionCountOfGrammar(int var0, int[] var1) {
	return var1[var0+1] & '\uffff' //return var1[var0 + 1] & '\uffff'
}

func _hasContentGrammar(var0 int, var1 []int) bool { //public static boolean _hasContentGrammar(int var0, int[] var1) {
	return (var1[var0+1] & 131072) != 0
}

func _hasEmptyGrammar(var0 int, var1 []int) bool { //public static boolean _hasEmptyGrammar(int var0, int[] var1) {
	var var2 bool = (var1[var0+1] & 262144) != 0 //boolean var2 = (var1[var0 + 1] & 262144) != 0;

	if var2 || !_hasContentGrammar(var0, var1) { //assert !var2 || _hasContentGrammar(var0, var1);
		panic("EXISchema.class, func  _hasEmptyGrammar, var2 || !_hasContentGrammar(var0, var1)")
	}
	return var2
}

func (e *EXISchema) getGrammars() []int { //public int[] getGrammars() {
	return e.m_grammars
}

func (e *EXISchema) getTotalGrammarCount() int { //public int getTotalGrammarCount() {
	return e.m_grammarCount
}

func (e *EXISchema) EXISchema(var1 []int, var2 int, var3 []int, var4 int, var5 []int, var6 int, var7 []string, var8 int, var9 []string, var10 int, var11 [][]int, var12 []string, var13 int, var14 []int, var15 int, var16 []int, var17 []int, var18 int, var19 []bool, var20 []string, var21 []string, var22 int, var23 []int /*big.int*/, var24 int, var25 []int, var26 int, var27 []string /*[]XSDateTime*/, var28 int, var29 []int64 /*[]Duration*/, var30 int, var31 [][]byte, var32 int, var33 []byte, var34 []int, var35 int, var36 []int, var37 int, var38 int, var39 []int, var40 int, var41 []byte, var42 []int, var43 int, var44 int) {
	e.m_elems = make([]int, var2, var2)
	copy(myCopyArrInt(var1, 0, e.m_elems, 0, var2), e.m_elems) //copy(e.m_elems, var1) //System.arraycopy(var1, 0, this.m_elems, 0, var2);

	e.m_attrs = make([]int, var4, var4)
	copy(myCopyArrInt(var3, 0, e.m_attrs, 0, var4), e.m_attrs) //copy(e.m_attrs, var3) //System.arraycopy(var3, 0, this.m_attrs, 0, var4);

	e.m_types = make([]int, var6, var6)
	copy(myCopyArrInt(var5, 0, e.m_types, 0, var6), e.m_types) //copy(e.m_types, var5) //System.arraycopy(var5, 0, this.m_types, 0, var6);

	e.uris = make([]string, var8, var8)
	copy(myCopyArrString(var7, 0, e.uris, 0, var8), e.uris) //copy(e.uris, var7)    //System.arraycopy(var7, 0, this.uris, 0, var8);

	e.m_localNames = make([][]int, len(var11), len(var11)) //this.m_localNames = new int[var11.length][];
	copy(e.m_localNames, var11)                            /*
										for(int var45 = 0; var45 < var11.length; ++var45) {
		        							this.m_localNames[var45] = new int[var11[var45].length];
		         							System.arraycopy(var11[var45], 0, this.m_localNames[var45], 0, var11[var45].length);
		      							}
	*/

	e.m_names = make([]string, var10, var10)
	copy(myCopyArrString(var9, 0, e.m_names, 0, var10), e.m_names) //copy(e.m_names, var9)    //System.arraycopy(var9, 0, this.m_names, 0, var10);

	e.m_strings = make([]string, var13, var13)                          // String[var13];
	copy(myCopyArrString(var12, 0, e.m_strings, 0, var13), e.m_strings) //copy(e.m_strings, var12) //System.arraycopy(var12, 0, this.m_strings, 0, var13);

	e.m_ints = make([]int, var15, var15)                       //int[var15];
	copy(myCopyArrInt(var14, 0, e.m_ints, 0, var15), e.m_ints) //copy(e.m_ints, var14)    //System.arraycopy(var14, 0, this.m_ints, 0, var15);

	e.m_mantissas = make([]int, var18, var18)                            //long[var18];
	copy(myCopyArrInt(var16, 0, e.m_mantissas, 0, var18), e.m_mantissas) //copy(e.m_mantissas, var16) //System.arraycopy(var16, 0, this.m_mantissas, 0, var18);

	e.m_exponents = make([]int, var18, var18)                            //int[var18];
	copy(myCopyArrInt(var17, 0, e.m_exponents, 0, var18), e.m_exponents) //copy(e.m_exponents, var17) //System.arraycopy(var17, 0, this.m_exponents, 0, var18);

	e.m_signs = make([]bool, var22, var22)                        //boolean[var22];
	copy(myCopyArrBool(var19, 0, e.m_signs, 0, var22), e.m_signs) //copy(e.m_signs, var19)      //System.arraycopy(var19, 0, this.m_signs, 0, var22);

	e.m_integralDigits = make([]string, var22, var22)                                 //String[var22];
	copy(myCopyArrString(var20, 0, e.m_integralDigits, 0, var22), e.m_integralDigits) //copy(e.m_integralDigits, var20) //System.arraycopy(var20, 0, this.m_integralDigits, 0, var22);

	e.m_reverseFractionalDigits = make([]string, var22, var22)                                          //String[var22];
	copy(myCopyArrString(var21, 0, e.m_reverseFractionalDigits, 0, var22), e.m_reverseFractionalDigits) //copy(e.m_reverseFractionalDigits, var21) //System.arraycopy(var21, 0, this.m_reverseFractionalDigits, 0, var22);

	e.m_integers = make([]int /*big.int*/, var24, var24)               //BigInteger[var24];
	copy(myCopyArrInt(var23, 0, e.m_integers, 0, var24), e.m_integers) //copy(e.m_integers, var23)  //System.arraycopy(var23, 0, this.m_integers, 0, var24);

	e.m_longs = make([]int, var26, var26)                        //long[var26];
	copy(myCopyArrInt(var25, 0, e.m_longs, 0, var26), e.m_longs) //copy(e.m_longs, var25)     //System.arraycopy(var25, 0, this.m_longs, 0, var26);

	e.m_datetimes = make([]string /*XSDateTime*/, var28, var28)             //new XSDateTime[var28];
	copy(myCopyArrString(var27, 0, e.m_datetimes, 0, var28), e.m_datetimes) //copy(e.m_datetimes, var27) //System.arraycopy(var27, 0, this.m_datetimes, 0, var28);

	e.m_durations = make([]int64 /*Duration*/, var30, var30)               //Duration[var30];
	copy(myCopyArrInt64(var29, 0, e.m_durations, 0, var30), e.m_durations) //copy(e.m_durations, var29) //System.arraycopy(var29, 0, this.m_durations, 0, var30);

	e.m_binaries = make([][]byte, var32, var32)                            //byte[var32][];
	copy(myCopyArrDblByte(var31, 0, e.m_binaries, 0, var32), e.m_binaries) //copy(e.m_binaries, var31)   //System.arraycopy(var31, 0, this.m_binaries, 0, var32);

	e.m_variants = make([]int, var35, var35)                           //int[var35];
	copy(myCopyArrInt(var34, 0, e.m_variants, 0, var35), e.m_variants) //copy(e.m_variants, var34)   //System.arraycopy(var34, 0, this.m_variants, 0, var35);

	e.m_variantTypes = make([]byte, var35, var35)                               //byte[var35];
	copy(myCopyArrByte(var33, 0, e.m_variantTypes, 0, var35), e.m_variantTypes) //copy(e.m_variantTypes, var33) //System.arraycopy(var33, 0, this.m_variantTypes, 0, var35);

	e.m_grammars = make([]int, var37, var37)                           //int[var37];
	copy(myCopyArrInt(var36, 0, e.m_grammars, 0, var37), e.m_grammars) //copy(e.m_grammars, var36)    //System.arraycopy(var36, 0, this.m_grammars, 0, var37);

	e.m_productions = make([]int, var40, var40)                              //int[var40];
	copy(myCopyArrInt(var39, 0, e.m_productions, 0, var40), e.m_productions) //copy(e.m_productions, var39) //System.arraycopy(var39, 0, this.m_productions, 0, var40);

	e.m_eventTypes = make([]byte, var43)                                    //byte[var43];
	copy(myCopyArrByte(var41, 0, e.m_eventTypes, 0, var43), e.m_eventTypes) //copy(e.m_eventTypes, var41)  //System.arraycopy(var41, 0, this.m_eventTypes, 0, var43);

	e.m_eventData = make([]int, var43, var43)                            //int[var43];
	copy(myCopyArrInt(var42, 0, e.m_eventData, 0, var43), e.m_eventData) //copy(e.m_eventData, var42)   //System.arraycopy(var42, 0, this.m_eventData, 0, var43);

	e.m_n_stypes = var44
	e.m_grammarCount = var38
	e.setUp()
}

func (e *EXISchema) setUp() {
	e.computeAncestryIds()
	e.populateLocalNames()
	e.computeGlobalDirectory()
	e.computeVariantCharacters()
	e.computeDateTimes()
	e.buildFragmentsArray()
}

func (e *EXISchema) buildFragmentsArray() { //private void buildFragmentsArray() {
	var var6 []int             //ArrayList var6 = new ArrayList();
	var var3 int = -1          //int var3 = -1;
	var var2 int = -1          //int var2 = -1;
	var var4 int = -2147483648 /*0x80000000*/ //Integer.MIN_VALUE //int var4 = Integer.MIN_VALUE;
	var var5 bool = true       //boolean var5 = true;

	var var1 int                                           //int var1;
	var var8 int                                           //int var8;
	var var9 int                                           //int var9;
	for var1 = 0; var1 < len(e.m_elems); var1 = var1 + 4 { //for(var1 = 0; var1 < this.m_elems.length; var1 += 4) {
		var8 = e.m_elems[var1+1]          //var8 = this.m_elems[var1 + 1];
		var9 = e.m_elems[var1+0]          //var9 = this.m_elems[var1 + 0];
		if var2 == var8 && var3 == var9 { //if (var2 == var8 && var3 == var9) {
			if var4 == var1 && var8 != var2 && var9 != var3 { //assert var4 != var1 && var8 == var2 && var9 == var3;
				panic("EXISchema.class, func  buildFragmentsArray, var4 == var1 && var8 != var2 && var9 != var3")
			}
			if e.getTypeOfElem(var1) != e.getTypeOfElem(var4) || e.isNillableElement(var1) != e.isNillableElement(var4) { // if (this.getTypeOfElem(var1) != this.getTypeOfElem(var4) || this.isNillableElement(var1) != this.isNillableElement(var4)) {
				var5 = false //var5 = false;
			}
		} else {
			if var1 != 0 { //if (var1 != 0) {
				if var5 { //var6.add(var5 ? var4 : 0 - var4 - 1);
					var6 = append(var6, var4)
				} else {
					var6 = append(var6, (0 - var4 - 1))
				}
			}

			var4 = var1 //var4 = var1;
			var5 = true //var5 = true;
			var2 = var8 //var2 = var8;
			var3 = var9 //var3 = var9;
		}
	}

	if var1 != 0 { //if (var1 != 0) {
		if var5 {
			var6 = append(var6, var4) //var6.add(var5 ? var4 : 0 - var4 - 1);
		} else {
			var6 = append(var6, (0 - var4 - 1))
		}
	}

	var var7 []int     //ArrayList var7 = new ArrayList();
	var3 = -1          //var3 = -1;
	var2 = -1          //var2 = -1;
	var4 = -2147483648 /*0x80000000*/ //Integer.MIN_VALUE //var4 = Integer.MIN_VALUE;
	var5 = true        //var5 = true;

	for var1 = 0; var1 < len(e.m_attrs); var1 = var1 + 3 { //for(var1 = 0; var1 < this.m_attrs.length; var1 += 3) {
		var9 = e.m_attrs[var1+1]           //var9 = this.m_attrs[var1 + 1];
		var var10 int = e.m_attrs[var1+0]  //int var10 = this.m_attrs[var1 + 0];
		if var2 == var9 && var3 == var10 { //if (var2 == var9 && var3 == var10) {
			if var4 == var1 && var9 != var2 && var10 != var3 { //assert var4 != var1 && var9 == var2 && var10 == var3;
				panic("EXISchema.class, func  buildFragmentsArray, var4 == var1 && var9 != var2 && var10 != var3")
			}
			if e.getTypeOfAttr(var1) != e.getTypeOfAttr(var4) { //if (this.getTypeOfAttr(var1) != this.getTypeOfAttr(var4)) {
				var5 = false //var5 = false;
			}
		} else {
			if var1 != 0 { //if (var1 != 0) {
				if var5 { //var7.add(var5 ? var4 : 0 - var4 - 1);
					var7 = append(var7, var4)
				} else {
					var7 = append(var7, (0 - var4 - 1))
				}
			}

			var4 = var1  //var4 = var1;
			var5 = true  //var5 = true;
			var2 = var9  //var2 = var9;
			var3 = var10 //var3 = var10;
		}
	}

	if var1 != 0 { //if (var1 != 0) {
		if var5 { //var7.add(var5 ? var4 : 0 - var4 - 1);
			var7 = append(var7, var4)
		} else {
			var7 = append(var7, (0 - var4 - 1))
		}
	}

	var8 = 0                                                                                       //var8 = 0;
	e.m_n_fragmentElems = len(var6)                                                                //this.m_n_fragmentElems = var6.size();
	e.m_fragmentINodes = make([]int, e.m_n_fragmentElems+len(var7), e.m_n_fragmentElems+len(var7)) //this.m_fragmentINodes = new int[this.m_n_fragmentElems + var7.size()];

	//var var11 Iterator //Iterator var11;
	for _, elem := range var6 { //for(var11 = var6.iterator(); var11.hasNext(); this.m_fragmentINodes[var8++] = (Integer)var11.next()) {
		e.m_fragmentINodes[var8] = elem
		var8++
	}

	if var8 != e.m_n_fragmentElems { //assert var8 == this.m_n_fragmentElems;
		panic("EXISchema.class, func  buildFragmentsArray, var8 != e.m_n_fragmentElems")
	}

	for _, elem := range var7 { //for(var11 = var7.iterator(); var11.hasNext(); this.m_fragmentINodes[var8++] = (Integer)var11.next()) {
		e.m_fragmentINodes[var8] = elem
		var8++
	}

}

func (e *EXISchema) getContentDatatypeOfComplexType(var1 int) int { //public int getContentDatatypeOfComplexType(int var1) {
	if 0 > var1 && e.m_types[var1+5] < 0 { //assert 0 <= var1 && this.m_types[var1 + 5] >= 0;
		panic("EXISchema.class, func getContentDatatypeOfComplexType, 0 > var1 && e.m_types[var1 + 5] < 0 ")
	}
	var var2 int                //int var2;
	if e.m_types[var1+5] != 0 { //return (var2 = this.m_types[var1 + 5]) != 0 ? var2 : -1;
		var2 = var2
	} else {
		var2 = (-1)
	}
	return var2
}

func (e *EXISchema) getElemCountOfSchema() int { //public int getElemCountOfSchema() {
	return len(e.m_elems) / 4
}

func (e *EXISchema) computeDateTimes() { //private void computeDateTimes() {
	var var1 int = len(e.m_datetimes)                               //int var1 = this.m_datetimes.length;
	var var2 []string = /*XSDateTime[]*/ make([]string, var1, var1) //XSDateTime[] var2 = new XSDateTime[var1];

	for var3 := 0; var3 < var1; var3++ { //for(int var3 = 0; var3 < var1; ++var3) {
		var var4 string //XSDateTime var4;
		var4 = e.m_datetimes[var3]
		var2[var3] = var4 //var2[var3] = var4 = new XSDateTime(this.m_datetimes[var3]);
		//var4.normalize();
	}

	e.m_computedDatetimes = var2 //this.m_computedDatetimes = var2;
}

func (e *EXISchema) computeAncestryIds() {
	e.ancestryIds = make([]int, e.m_n_stypes+1, e.m_n_stypes+1)
	e.ancestryIds[0] = -1
	var var1 int = 6
	for var2 := 0; var2 < e.m_n_stypes; var1 += _getTypeSize(var1, e.m_types, e.ancestryIds) { //for(int var2 = 0; var2 < this.m_n_stypes; var1 += _getTypeSize(var1, this.m_types, this.ancestryIds)) {
		var var3 int = e.getSerialOfType(var1) //	int var3 = this.getSerialOfType(var1);

		if 0 >= var3 && var3 > e.m_n_stypes && e.isSimpleType(var1) { //	assert 0 < var3 && var3 <= this.m_n_stypes && this.isSimpleType(var1);
			panic("EXISchema.class, func computeAncestryIds, 0 >= var3 && var3 > e.m_n_stypes && e.isSimpleType(var1)")
		}
		var2++
		var var4 int                             //byte
		if e.getVarietyOfSimpleType(var1) == 1 { //	if (this.getVarietyOfSimpleType(var1) == 1) {
			var var6 int //	   int var6;

			var var5 int                   //
			var6 = e.getSerialOfType(var5) //

			for var5 = var1; var6 >= 22; var5 = e.getBaseTypeOfSimpleType(var5) {
			} //for(int var5 = var1; (var6 = this.getSerialOfType(var5)) >= 22; var5 = this.getBaseTypeOfSimpleType(var5)) {

			if var6 < 2 { //	   assert var6 >= 2;
				panic("EXISchema.class, func computeAncestryIds, var6 < 2")
			}
			var4 = VAREXISchema.ANCESTRY_IDS[var6]
		} else {
			var4 = -1
		}

		e.ancestryIds[var3] = var4
	}

}

//todo мaccив проверить на коректность заполнения !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (e *EXISchema) populateLocalNames() { //   private void populateLocalNames() {
	e.localNames = make([][]string, len(e.m_localNames), len(e.m_localNames)) //	this.localNames = new String[this.m_localNames.length][];
	for var1 := 0; var1 < len(e.m_localNames); var1++ {                       //	for(int var1 = 0; var1 < this.m_localNames.length; ++var1) {
		var var2 []string = make([]string, len(e.m_localNames[var1]), len(e.m_localNames[var1])) //String[] var2 = new String[this.m_localNames[var1].length];
		/*
					for(int var3 = 0; var3 < var2.length; ++var3) {
			            var2[var3] = this.m_names[this.m_localNames[var1][var3]];
			         }
		*/
		for var3 := 0; var3 < len(var2); var3++ {
			var2[var3] = e.m_names[e.m_localNames[var1][var3]]
		}
		//copy(var2, (e.m_names[e.m_localNames[var1]]))
		//copy(var2, e.m_names)
		e.localNames[var1] = var2 //var2[var3] = this.m_names[this.m_localNames[var1][var3]];
	}
}

//todo-проверить правельность заполнение массива/мапы!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (e *EXISchema) computeGlobalDirectory() { //private void computeGlobalDirectory() {
	e.m_globalElementsDirectory = make(map[string][]int)   //this.m_globalElementsDirectory = new HashMap();
	e.m_globalAttributesDirectory = make(map[string][]int) //this.m_globalAttributesDirectory = new HashMap();
	e.m_globalTypesDirectory = make(map[string][]int)      ///this.m_globalTypesDirectory = new HashMap();
	var var1 []int                                         //ArrayList var1 = new ArrayList();

	var var2 int                                     //int var2;
	for var2 = 0; var2 < len(e.m_elems); var2 += 4 { //for(var2 = 0; var2 < this.m_elems.length; var2 += 4) {
		if e.isGlobalElem(var2) { // if (this.isGlobalElem(var2)) {
			var var3 string = e.getNameOfElem(var2)                       //String var3 = this.getNameOfElem(var2);
			var var4 []int                                                //int[] var4;
			if value, inMap := e.m_globalElementsDirectory[var3]; inMap { //if ((var4 = (int[])this.m_globalElementsDirectory.get(var3)) != null) {
				//myValue := make([]int, 1, 1)
				//myValue[0] = value //какова длинна массива!!!???!!!????!!!?!?!?!?!?!?!??!?!?!?!?
				var4 = value
				var5 := make([]int, len(var4)+1, len(var4)+1) //int[] var5 = new int[var4.length + 1];
				copy(var5, var4)                              //System.arraycopy(var4, 0, var5, 0, var4.length);
				var5[len(var4)] = var2                        //var5[var4.length] = var2;
				var4 = var5                                   //var4 = var5;
			} else {
				var4 = []int{var2} //var4 = new int[]{var2};
			}

			e.m_globalElementsDirectory[var3] = var4 //this.m_globalElementsDirectory.put(var3, var4);
			var1 = append(var1, var2)                //var1.add(var2);
		}
	}

	e.m_globalElems = make([]int, len(var1), len(var1)) //this.m_globalElems = new int[var1.size()];

	if len(e.m_globalElems) != e.getGlobalElemCountOfSchema() { //assert this.m_globalElems.length == this.getGlobalElemCountOfSchema();//зачем это!!!?????????????????????????????
		panic("EXISchema.class, func computeGlobalDirectory, len(e.m_globalElems) != e.getGlobalElemCountOfSchema()")
	}
	for var2 = 0; var2 < len(e.m_globalElems); var2++ { //for(var2 = 0; var2 < this.m_globalElems.length; ++var2) {
		e.m_globalElems[var2] = int(var1[var2]) //this.m_globalElems[var2] = (Integer)var1.get(var2);
	}

	var var9 []int //ArrayList var9 = new ArrayList();

	var var7 []int                                      //int[] var7;
	var var10 int                                       //int var10;
	var var11 int                                       //int var11;
	for var10 = 0; var10 < len(e.m_attrs); var10 += 3 { //for(var10 = 0; var10 < this.m_attrs.length; var10 += 3) {
		if e.isGlobalAttr(var10) { //if (this.isGlobalAttr(var10)) {
			var11 = e.m_attrs[var10+0]                                //var11 = this.m_attrs[var10 + 0];
			var var12 int = e.m_attrs[10+1]                           //int var12 = this.m_attrs[var10 + 1];
			var var6 string = e.m_names[e.m_localNames[var12][var11]] //String var6 = this.m_names[this.m_localNames[var12][var11]];

			if value, inMap := e.m_globalAttributesDirectory[var6]; inMap { //if ((var7 = (int[])this.m_globalAttributesDirectory.get(var6)) != null) {
				var7 = value
				var var8 []int = make([]int, len(var7)+1, len(var7)+1) //int[] var8 = new int[var7.length + 1];
				copy(var8, var7)                                       //System.arraycopy(var7, 0, var8, 0, var7.length);
				var8[len(var7)] = var10                                //var8[var7.length] = var10;
				var7 = var8                                            //var7 = var8;
			} else { //} else {
				var7 = []int{var10} //var7 = new int[]{var10};
			}

			e.m_globalAttributesDirectory[var6] = var7 //this.m_globalAttributesDirectory.put(var6, var7);
			var9 = append(var9, var10)                 //var9.add(var10);
		}
	}

	e.m_globalAttrs = make([]int, len(var9), len(var9)) //this.m_globalAttrs = new int[var9.size()];

	for var10 = 0; var10 < len(e.m_globalAttrs); var10++ { //for(var10 = 0; var10 < this.m_globalAttrs.length; ++var10) {
		e.m_globalAttrs[var10] = int(var9[var10]) //this.m_globalAttrs[var10] = (Integer)var9.get(var10);
	}

	e.m_buitinTypes = make([]int, 46, 46) //this.m_buitinTypes = new int[46];
	var10 = 0                             //var10 = 0;

	for var11 = 0; var10 < len(e.m_types); var11++ { //for(var11 = 0; var10 < this.m_types.length; ++var11) {
		var var13 string = e.getNameOfType(var10) //String var13 = this.getNameOfType(var10);
		if var11 < 46 {                           //if (var11 < 46) {
			if e.getUriOfType(var10) != 3 && len(([]rune(var13))) == 0 { //assert this.getUriOfType(var10) == 3 && var13.length() != 0;//add []rune
				panic("EXISchema.class, func computeGlobalDirectory, e.getUriOfType(var10) != 3 && len(([]rune(var13))) == 0 ")
			}
			e.m_buitinTypes[var11] = var10 //this.m_buitinTypes[var11] = var10;
		}

		if !("" == var13) { //if (!"".equals(var13)) {
			var var14 []int                                             //int[] var14;
			if value, inMap := e.m_globalTypesDirectory[var13]; inMap { //if ((var14 = (int[])this.m_globalTypesDirectory.get(var13)) != null) {
				var14 = value
				var7 = make([]int, len(var14)+1, len(var14)+1) //var7 = new int[var14.length + 1];
				copy(var7, var14)                              //System.arraycopy(var14, 0, var7, 0, var14.length);
				var7[len(var14)] = var10                       //var7[var14.length] = var10;
				var14 = var7                                   //var14 = var7;
			} else {
				var14 = []int{var10} //var14 = new int[]{var10};
			}

			e.m_globalAttributesDirectory[var13] = var14 //this.m_globalTypesDirectory.put(var13, var14);
		}

		var10 += _getTypeSize(var10, e.m_types, e.ancestryIds) //var10 += _getTypeSize(var10, this.m_types, this.ancestryIds);
	}

}

func (e *EXISchema) getUriOfEventType(var1 int) int { //public int getUriOfEventType(int var1) {
	if e.m_eventTypes[var1] != 2 || e.m_eventTypes[var1] != 3 { //assert this.m_eventTypes[var1] == 2 || this.m_eventTypes[var1] == 3;
		panic("EXISchema.class, func getUriOfEventType, e.m_eventTypes[var1] != 2 || e.m_eventTypes[var1] != 3")
	}
	return e.m_eventData[var1] //return this.m_eventData[var1];
}

func (e *EXISchema) getSerialOfElem(var1 int) int { //public int getSerialOfElem(int var1) {
	if 0 > var1 && var1%4 != 0 { //assert 0 <= var1 && var1 % 4 == 0;
		panic("EXISchema.class, func getSerialOfElem, 0 > var1")
	}
	return var1 / 4 //return var1 / 4;
}

func (e *EXISchema) getLocalNameOfElem(var1 int) int { //public int getLocalNameOfElem(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getLocalNameOfElem, 0 > var1")
	}
	return e.m_elems[var1+0]
}

func (e *EXISchema) getUriOfElem(var1 int) int { //public int getUriOfElem(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getUriOfElem, 0 > var1")
	}
	return e.m_elems[var1+1]
}

func (e *EXISchema) getLocalNameOfAttr(var1 int) int { //public int getLocalNameOfAttr(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getLocalNameOfAttr, 0 > var1")
	}
	return e.m_attrs[var1+0]
}

func (e *EXISchema) getUriOfAttr(var1 int) int { //public int getUriOfAttr(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getUriOfAttr, 0 > var1")
	}
	return e.m_attrs[var1+1]
}

func (e *EXISchema) getNodeOfEventType(var1 int) int { //public int getNodeOfEventType(int var1) {
	if e.m_eventTypes[var1] != 0 || e.m_eventTypes[var1] != 1 { //assert this.m_eventTypes[var1] == 0 || this.m_eventTypes[var1] == 1;
		panic("EXISchema.class, func getNodeOfEventType, e.m_eventTypes[var1] != 0 || e.m_eventTypes[var1] != 1")
	}

	return e.m_eventData[var1]
}

func (e *EXISchema) getEventOfProduction(var1 int) int { //public int getEventOfProduction(int var1) {
	return e.m_productions[var1+0]
}

func (e *EXISchema) getProductionOfGrammar(var1 int, var2 int) int { //public int getProductionOfGrammar(int var1, int var2) {
	var var3 int = e.getProductionCountOfGrammar(var1) //int var3 = this.getProductionCountOfGrammar(var1);
	/*if (var2 >= 0 && var2 < var3) { //if (var2 >= 0 && var2 < var3) {
		return _getProductionOfGrammar(var1, var2, e.m_grammars) //return _getProductionOfGrammar(var1, var2, this.m_grammars);
	} else {
	   //throw new EXISchemaRuntimeException(1, new String[]{String.valueOf(var2), String.valueOf(0), String.valueOf(var3 - 1)});
	}*/
	if var2 < 0 && var2 >= var3 {
		panic("EXISchema class, func getProductionOfGrammar, wrong dev")
	}
	return _getProductionOfGrammar(var1, var2, e.m_grammars)
}

func _getProductionOfGrammar(var0 int, var1 int, var2 []int) int { //public static int _getProductionOfGrammar(int var0, int var1, int[] var2) {
	var var3 int = var0 + 2 //int var3 = var0 + 2;
	if _hasContentGrammar(var0, var2) {
		var3++ //++var3;
		if _hasEmptyGrammar(var0, var2) {
			var3++ //++var3;
		}
	}

	return var2[var3+var1]
}

func (e *EXISchema) computeVariantCharacters() { //private void computeVariantCharacters() {
	var var1 int = len(e.m_variants)                                  //int var1 = this.m_variants.length;
	var var2 []Characters = /*string*/ make([]Characters, var1, var1) //Characters[] var2 = new Characters[var1];

	for var3 := 0; var3 < var1; var3++ {
		var var4 string = e.computeVariantCharacters1(var3)                                           //String var4 = this.computeVariantCharacters(var3);
		var2[var3] = CHARACTERS_EMPTY.Characters1(myToCharArray(var4), 0, len(([]rune(var4))), false) //var2[var3] = new Characters(var4.toCharArray(), 0, var4.length(), false); //add []rune
	}

	e.m_variantCharacters = var2 //this.m_variantCharacters = var2;
}

func (e *EXISchema) computeVariantCharacters1(var1 int) string { //private String computeVariantCharacters(int var1) {
	var var2 string = ""            //String var2 = "";
	var var3 []byte                 //byte[] var3;
	switch e.m_variantTypes[var1] { //switch(this.m_variantTypes[var1]) {
	case 0: //case 0:
		var2 = e.getStringValueOfVariant(var1) //var2 = this.getStringValueOfVariant(var1);
		//break;
	case 1: //case 1:
		var var4 int = e.m_variants[var1]  //int var4 = this.m_variants[var1];
		var var5 int = e.m_mantissas[var4] //long var5 = this.m_mantissas[var4];
		var var7 int                       //int var7;
		if e.m_exponents[var4] != 0 {      //if ((var7 = this.m_exponents[var4]) != 0) {
			var7 = e.m_exponents[var4]
			if var7 == -16384 { //if (var7 == -16384) {
				if var5 == int(1) { //var2 = var5 == 1L ? "INF" : (var5 == -1L ? "-INF" : "NaN");
					var2 = "INF"
				} else if var5 == int(-1) {
					var2 = "-INF"
				} else {
					var2 = "NaN"
				}
			} else {
				var2 = strconv.Itoa(var5) + "E" + strconv.Itoa(var7) //var2 = Long.toString(var5) + "E" + Integer.toString(var7);//
			}
		} else {
			var2 = strconv.Itoa(var5) //var2 = Long.toString(var5);
		}
		//break;
	case 2: //case 2:
		var var8 bool = e.getSignOfDecimalVariant(var1)                                  //boolean var8 = this.getSignOfDecimalVariant(var1);
		var var9 string = e.getIntegralDigitsOfDecimalVariant(var1)                      //String var9 = this.getIntegralDigitsOfDecimalVariant(var1);
		var var10 string = myReverse(e.getReverseFractionalDigitsOfDecimalVariant(var1)) //String var10 = (new StringBuilder(this.getReverseFractionalDigitsOfDecimalVariant(var1))).reverse().toString();
		var var11 bool                                                                   //boolean var11;
		if "0" == (var10) && "0" == (var9) {                                             //if ((var11 = "0".equals(var10)) && "0".equals(var9)) {
			var11 = ("0" == (var10))
			var2 = "0" //var2 = "0";
		} else {
			var11 = ("0" == (var10)) //my add!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
			if var8 {                //if (var8) {
				var2 = var2 + `-` //var2 = var2 + '-';
			}

			var2 = var2 + var9 //var2 = var2 + var9;
			if !var11 {        //if (!var11) {
				var2 = var2 + `.`   //var2 = var2 + '.';
				var2 = var2 + var10 //var2 = var2 + var10;
			}
		}
		//break;
	case 3: //case 3:
		var var12 int = /*BigInteger*/ e.getIntegerValueOfVariant(var1) //BigInteger var12 = this.getIntegerValueOfVariant(var1);
		var2 = strconv.Itoa(var12)                                      //var2 = var12.toString();
		//break;
	case 4: //case 4:
		var var13 int = e.getIntValueOfVariant(var1) //int var13 = this.getIntValueOfVariant(var1);
		var2 = strconv.Itoa(var13)                   //var2 = Integer.toString(var13);
		//break;
	case 5: //case 5:
		var var14 int = e.getLongValueOfVariant(var1) //long var14 = this.getLongValueOfVariant(var1);
		var2 = strconv.Itoa(var14)                    //var2 = Long.toString(var14);
		//break;
	case 6: //case 6:
		var var16 string = /*XSDateTime*/ e.getDateTimeValueOfVariant(var1) //XSDateTime var16 = this.getDateTimeValueOfVariant(var1);
		var2 = var16                                                        //var2 = var16.toString();
		//break;
	case 7: //case 7:
		var2 = strconv.Itoa(int(e.getDurationValueOfVariant(var1))) //var2 = this.getDurationValueOfVariant(var1).toString();
		//break;
	case 8: //case 8:
		var3 = e.getBinaryValueOfVariant(var1) //var3 = this.getBinaryValueOfVariant(var1);
		var var17 int = len(var3) / 3 << 2     //int var17 = var3.length / 3 << 2;
		if len(var3)%3 != 0 {                  //if (var3.length % 3 != 0) {
			var17 += 4 //var17 += 4;
		}

		var17 += var17 / 76                                        //var17 += var17 / 76;
		var var18 []string = make([]string, var17, var17)          //char[] var18 = new char[var17];
		var var19 int = Base64encode(var3, 0, len(var3), var18, 0) //int var19 = Base64.encode(var3, 0, var3.length, var18, 0);

		var2 = ""                      //var2 = new String(var18, 0, var19);
		for idx, elem := range var18 { //var2 = new String(var18, 0, var19);
			var2 += elem
			if idx >= var19 {
				break
			}
		}
		//break;
	case 9: //case 9:
		panic("EXISchema.class, func computeVariantCharacters, case 9") //assert false;

		return "nil" //return null
	case 10: //case 10:
		var3 = e.getBinaryValueOfVariant(var1) //var3 = this.getBinaryValueOfVariant(var1);
		var var20 strings.Builder              //StringBuffer var20 = new StringBuffer();
		HexBinencode(var3, len(var3), &var20)  //HexBin.encode(var3, var3.length, var20);
		var2 = var20.String()                  //var2 = var20.toString();
		//break;
	default: //default:
		panic("EXISchema.class, func computeVariantCharacters, default") //assert false;
	}

	return var2 //return var2;
}

func (e *EXISchema) getElems() []int { //public int[] getElems() {
	return e.m_elems
}

func (e *EXISchema) getTypeOfAttr(var1 int) int { //public int getTypeOfAttr(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getTypeOfAttr, 0 > var1")
	}
	result := e._getTypeOfAttr(var1, e.m_attrs)
	return result //return _getTypeOfAttr(var1, this.m_attrs);
}

func (e *EXISchema) _getTypeOfAttr(var0 int, var1 []int) int { //public static int _getTypeOfAttr(int var0, int[] var1) {
	var var2 int = var1[var0+2] //int var2 = var1[var0 + 2];
	var result int
	if (var2 & -2147483648 /*0x80000000*/ /*Integer.MIN_VALUE*/) != 0 {
		result = (^var2)
	} else {
		result = var2
	}
	return result //return (var2 & Integer.MIN_VALUE) != 0 ? ~var2 : var2;
}

func (e *EXISchema) isNillableElement(var1 int) bool { //public boolean isNillableElement(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func isNillableElement, 0 > var1")
	}
	return e.m_elems[var1+3] != 0 //return this.m_elems[var1 + 3] != 0;
}

func (e *EXISchema) getTypeOfElem(var1 int) int { //public int getTypeOfElem(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getTypeOfElem, 0 > var1")
	}
	var var2 int = e.m_elems[var1+2] //int var2 = this.m_elems[var1 + 2];
	if (var2 & -2147483648 /*0x80000000*/ /*Integer.MIN_VALUE*/) != 0 {
		return (^var2)
	}
	return var2 //return (var2 & Integer.MIN_VALUE) != 0 ? ~var2 : var2;
}

func (e *EXISchema) getBinaryValueOfVariant(var1 int) []byte { //public byte[] getBinaryValueOfVariant(int var1) {
	var var2 byte = e.m_variantTypes[var1] //byte var2 = this.m_variantTypes[var1];

	if 0 > var1 && (var2 != 8 || var2 != 10) { //assert 0 <= var1 && (var2 == 8 || var2 == 10);
		panic("EXISchema.class, func  getBinaryValueOfVariant, 0 > var1 && (var2 != 8 || var2 != 10)")
	}
	return e.m_binaries[e.m_variants[var1]]
}

func (e *EXISchema) getDurationValueOfVariant(var1 int) int64 /*Duration*/ { //public Duration getDurationValueOfVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 7 { //assert 0 <= var1 && this.m_variantTypes[var1] == 7;
		panic("EXISchema.class, func getDurationValueOfVariant, 0 > var1 && e.m_variantTypes[var1] != 7")
	}
	return e.m_durations[e.m_variants[var1]]
}

func (e *EXISchema) getDateTimeValueOfVariant(var1 int) string /*XSDateTime*/ { //public XSDateTime getDateTimeValueOfVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 6 { //assert 0 <= var1 && this.m_variantTypes[var1] == 6;
		panic("EXISchema.class, func getDateTimeValueOfVariant, 0 > var1 && e.m_variantTypes[var1] != 6 ")
	}
	return e.m_datetimes[e.m_variants[var1]]
}

func (e *EXISchema) getLongValueOfVariant(var1 int) int { //public long getLongValueOfVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 5 { //assert 0 <= var1 && this.m_variantTypes[var1] == 5;
		panic("EXISchema.class, func getLongValueOfVariant, 0 > var1 && e.m_variantTypes[var1] != 5")
	}
	return e.m_longs[e.m_variants[var1]]
}

func (e *EXISchema) getIntValueOfVariant(var1 int) int { //public int getIntValueOfVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 4 { //assert 0 <= var1 && this.m_variantTypes[var1] == 4;
		panic("EXISchema.class, func getIntValueOfVariant, 0 > var1 && e.m_variantTypes[var1] != 4")
	}
	return e.m_ints[e.m_variants[var1]]
}

func (e *EXISchema) getIntegerValueOfVariant(var1 int) int /*BigInteger*/ { //public BigInteger getIntegerValueOfVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 3 { // assert 0 <= var1 && this.m_variantTypes[var1] == 3;
		panic("EXISchema.class, func getIntegerValueOfVariant, 0 > var1 && e.m_variantTypes[var1] != 3")
	}
	return e.m_integers[e.m_variants[var1]]
}

func (e *EXISchema) getReverseFractionalDigitsOfDecimalVariant(var1 int) string { //public String getReverseFractionalDigitsOfDecimalVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 2 { //assert 0 <= var1 && this.m_variantTypes[var1] == 2;
		panic("EXISchema.class, func getReverseFractionalDigitsOfDecimalVariant, 0 > var1 && e.m_variantTypes[var1] != 2 ")
	}
	return e.m_reverseFractionalDigits[e.m_variants[var1]]
}

func (e *EXISchema) getIntegralDigitsOfDecimalVariant(var1 int) string { //public String getIntegralDigitsOfDecimalVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 2 { //assert 0 <= var1 && this.m_variantTypes[var1] == 2;
		panic("EXISchema.class, func getIntegralDigitsOfDecimalVariant, 0 > var1 && e.m_variantTypes[var1] != 2 ")
	}
	return e.m_integralDigits[e.m_variants[var1]]
}

func (e *EXISchema) getSignOfDecimalVariant(var1 int) bool { //public boolean getSignOfDecimalVariant(int var1) {
	if 0 > var1 && e.m_variantTypes[var1] != 2 { //assert 0 <= var1 && this.m_variantTypes[var1] == 2;
		panic("EXISchema.class, func getSignOfDecimalVariant, 0 > var1 && e.m_variantTypes[var1] != 2 ")
	}
	return e.m_signs[e.m_variants[var1]]
}

func (e *EXISchema) getStringValueOfVariant(var1 int) string { //public String getStringValueOfVariant(int var1) {
	if (0 > var1) && (e.m_variantTypes[var1] != 0) { //assert 0 <= var1 && this.m_variantTypes[var1] == 0
		panic("EXISchema.class, func getStringValueOfVariant, 0 > var1 && this.m_variantTypes[var1] != 0")
	}
	return e.m_strings[e.m_variants[var1]]
}

func (e *EXISchema) getUriOfType(var1 int) int { //public int getUriOfType(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getUriOfType, 0 > var1")
	}
	return e.m_types[var1+1]
}

func (e *EXISchema) getNameOfType(var1 int) string { //public String getNameOfType(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getNameOfType, 0 > var1")
	}
	var var2 = e.m_types[var1+0] //int var2 = this.m_types[var1 + 0];
	if var2 != -1 {
		var var3 = e.m_types[var1+1]                 //int var3 = this.m_types[var1 + 1];
		return e.m_names[e.m_localNames[var3][var2]] //return this.m_names[this.m_localNames[var3][var2]];
	} else {
		return ""
	}
}

func (e *EXISchema) isGlobalAttr(var1 int) bool { //public boolean isGlobalAttr(int var1) {
	var var2 int = e.m_attrs[var1+2]
	return (var2 & -2147483648 /*0x80000000*/ /*Integer.MIN_VALUE*/) != 0
}

func (e *EXISchema) getGlobalElemCountOfSchema() int { //public int getGlobalElemCountOfSchema() {
	return len(e.m_globalElems)
}

func (e *EXISchema) isGlobalElem(var1 int) bool { //public boolean isGlobalElem(int var1) {
	var var2 int = e.m_elems[var1+2]
	return (var2 & -2147483648 /*0x80000000*/ /*Integer.MIN_VALUE*/) != 0
}

func (e *EXISchema) getNameOfElem(var1 int) string { //public String getNameOfElem(int var1) {
	if 0 > var1 { //assert 0 <= var1;
		panic("EXISchema.class, func getNameOfElem, 0 > var1")
	}
	var var2 int = e.m_elems[var1+0] //int var2 = this.m_elems[var1 + 0];
	if var2 != -1 {
		var var3 int = e.m_elems[var1+1]
		return e.m_names[e.m_localNames[var3][var2]]
	} else {
		return ""
	}
}

func (e *EXISchema) getBaseTypeOfSimpleType(var1 int) int { //public int getBaseTypeOfSimpleType(int var1) {
	if 0 > var1 && e.m_types[var1+5] >= 0 { //	assert 0 <= var1 && this.m_types[var1 + 5] < 0;
		panic("EXISchema.class, func getBaseTypeOfSimpleType, 0 > var1 && e.m_types[var1 + 5] >= 0")
	}
	return e.m_types[var1+6]
}

func (e *EXISchema) getVarietyOfSimpleType(var1 int) byte { //public byte getVarietyOfSimpleType(int var1) {
	if 0 > var1 && e.m_types[var1+5] >= 0 { //  assert 0 <= var1 && this.m_types[var1 + 5] < 0;
		panic("EXISchema.class, func getVarietyOfSimpleType, 0 > var1 && e.m_types[var1 + 5] >= 0")
	}
	return _getVarietyOfSimpleType(var1, e.m_types)
}

func (e *EXISchema) isSimpleType(var1 int) bool { //   public boolean isSimpleType(int var1) {
	return (e.m_types[var1+5] & -2147483648 /*0x80000000*/ /*Integer.MIN_VALUE*/) != 0
}

func (e *EXISchema) getSerialOfType(var1 int) int { //public int getSerialOfType(int var1) {
	if 0 > var1 { //	assert 0 <= var1;
		panic("EXISchema.class, func getSerialOfType, 0 > var1")
	}
	return e.m_types[var1+2]
}

func _getTypeSize(var0 int, var1 []int, var2 []int /*[]byte*/) int { //public static int _getTypeSize(int var0, int[] var1, byte[] var2) {
	if var0 == -1 {
		panic("EXISchema.class, func _getTypeSize, var0==-1")
	} //      assert var0 != -1;

	if (var1[var0+5] & -2147483648 /*0x80000000*/ /*Integer.MIN_VALUE*/) != 0 { //      return (var1[var0 + 5] & Integer.MIN_VALUE) != 0 ? _getSizeOfSimpleType(var0, var1, var2) : 6;
		return _getSizeOfSimpleType(var0, var1, var2)
	}
	return 6
}

func _getSizeOfSimpleType(var0 int, var1 []int, var2 []int /*[]byte*/) int { //private static int _getSizeOfSimpleType(int var0, int[] var1, byte[] var2) {
	var var3 int = 8                                    //      int var3 = 8;
	var var4 byte = _getVarietyOfSimpleType(var0, var1) //      byte var4 = _getVarietyOfSimpleType(var0, var1);
	if var4 == 1 {                                      //      if (var4 == 1) {
		if _isEnumeratedAtomicSimpleType(var0, var1) { //         if (_isEnumeratedAtomicSimpleType(var0, var1)) {
			var3 += (1 + _getEnumerationFacetCountOfAtomicSimpleType(var0, var1, var2)) //            var3 += 1 + _getEnumerationFacetCountOfAtomicSimpleType(var0, var1, var2);
		}

		if var2[var1[var0+2]] == 2 { //         if (var2[var1[var0 + 2]] == 2) {
			var3 += _getRestrictedCharacterCountOfStringSimpleType(var0, var1, var2) //            var3 += _getRestrictedCharacterCountOfStringSimpleType(var0, var1, var2);
		}
	}

	return var3
}
func _getVarietyOfSimpleType(var0 int, var1 []int) byte { //public static byte _getVarietyOfSimpleType(int var0, int[] var1) {
	return byte(var1[var0+5] & 3)
}

func _isEnumeratedAtomicSimpleType(var0 int, var1 []int) bool { //private static boolean _isEnumeratedAtomicSimpleType(int var0, int[] var1) {
	return (var1[var0+5] & 4) != 0
}

func _getEnumerationFacetCountOfAtomicSimpleType(var0 int, var1 []int, var2 []int /*[]byte*/) int { //private static int _getEnumerationFacetCountOfAtomicSimpleType(int var0, int[] var1, byte[] var2) {
	if _isEnumeratedAtomicSimpleType(var0, var1) {
		var var3 int
		if var2[var1[var0+2]] == 2 { //     int var3 = var2[var1[var0 + 2]] == 2 ? _getRestrictedCharacterCountOfStringSimpleType(var0, var1, var2) : 0;
			var3 = _getRestrictedCharacterCountOfStringSimpleType(var0, var1, var2)
		} else {
			var3 = 0
		}
		return var1[var0+8+var3]
	}

	return 0
}

func _getRestrictedCharacterCountOfStringSimpleType(var0 int, var1 []int, var2 []int /*[]byte)*/) int { //   private static int _getRestrictedCharacterCountOfStringSimpleType(int var0, int[] var1, byte[] var2) {
	var var3 int = (var1[var0+5] & 8160) >> 5 //int var3 = (var1[var0 + 5] & 8160) >> 5;

	if 0 > var3 && var3 >= 256 { //      assert 0 <= var3 && var3 < 256;
		panic("EXISchema.class, func _getRestrictedCharacterCountOfStringSimpleType, (0 > var3 && var3 >= 256)")
	}

	return var3
}

//////////////////Base64.class//////////////

type Base64 struct { //public final class Base64 {
	BASE64_ASCIIS []string //private static final char[] BASE64_ASCIIS;
	BASE64CHARS   []byte   //private static final byte[] BASE64CHARS = new byte[8192];
	m_octets      []byte   //private static final byte[] m_octets;
}

var myVarBase64 Base64 = Base64{
	BASE64CHARS: make([]byte, 8192, 8192),
}

//private Base64() {
//}

func Base64encode(var0 []byte, var1 int, var2 int, var3 []string, var4 int) int { //public static int encode(byte[] var0, int var1, int var2, char[] var3, int var4) {
	var var5 int = var1 + var2 //int var5 = var1 + var2;
	var var6 = var4            //int var6 = var4;
	var var7 = var1            //int var7 = var1;

	for var8 := 0; var7 < var5; var8++ { //for(int var8 = 0; var7 < var5; ++var8) {
		var var9 int                                    //int var9;
		for var9 = 0; var9 < 3 && var7 < var5; var9++ { //for(var9 = 0; var9 < 3 && var7 < var5; ++var9) {
			var7++ //++var7;
		}

		if var9 != 1 || var9 != 2 || var9 != 3 { //assert var9 == 1 || var9 == 2 || var9 == 3;
			panic("Base64.class, Base64encode, var9 != 1 || var9 != 2 || var9 != 3")
		}
		var var13 byte = 64 //byte var13 = 64;
		var var14 byte = 64 //byte var14 = 64;
		var var11 byte      //byte var11;
		var11 = byte(var0[var7] >> 2)
		if (byte(var0[var7] >> 2)) < 0 { //if ((var11 = (byte)(var0[var7] >> 2)) < 0) {
			var11 = byte(var11 ^ 192) //var11 = (byte)(var11 ^ 192);
		}

		var var12 byte //byte var12;
		var12 = byte(var0[var7+1] >> 4)
		if var9 > 1 { //if (var9 > 1) {
			if (byte(var0[var7+1] >> 4)) < 0 { //if ((var12 = (byte)(var0[var7 + 1] >> 4)) < 0) {
				var12 = byte(var12 ^ 240) //var12 = (byte)(var12 ^ 240);
			}

			var12 = (byte)((var0[var7]&3)<<4 | var12) //var12 = (byte)((var0[var7] & 3) << 4 | var12);
			if var9 > 2 {                             //if (var9 > 2) {
				var13 = byte(var0[var7+2] >> 6)
				if (byte(var0[var7+2] >> 6)) < 0 { //if ((var13 = (byte)(var0[var7 + 2] >> 6)) < 0) {
					var13 = byte(var13 ^ 252) //var13 = (byte)(var13 ^ 252);
				}

				var13 = byte((var0[var7+1]&15)<<2 | var13) //var13 = (byte)((var0[var7 + 1] & 15) << 2 | var13);
				var14 = byte(var0[var7+2] & 63)            //var14 = (byte)(var0[var7 + 2] & 63);
			} else { //} else {
				var13 = byte((var0[var7+1] & 15) << 2) //var13 = (byte)((var0[var7 + 1] & 15) << 2);
			}
		} else { //} else {
			var12 = byte((var0[var7] & 3) << 4) //var12 = (byte)((var0[var7] & 3) << 4);
		}

		var3[var6] = myVarBase64.BASE64_ASCIIS[var11]   //var3[var6++] = BASE64_ASCIIS[var11];
		var3[var6+1] = myVarBase64.BASE64_ASCIIS[var12] //var3[var6++] = BASE64_ASCIIS[var12];
		var3[var6+2] = myVarBase64.BASE64_ASCIIS[var13] //var3[var6++] = BASE64_ASCIIS[var13];
		var3[var6+3] = myVarBase64.BASE64_ASCIIS[var14] //var3[var6++] = BASE64_ASCIIS[var14];
		var6 = var6 + 4

		if var8%19 == 18 { //if (var8 % 19 == 18) {
			var3[var6] = "\n" //var3[var6++] = '\n';
			var6 = var6 + 1
		}
	}

	return var6 - var4 //return var6 - var4;
}

//to do
/*
	static {
	   var Base64var0 string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=" //String var0 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
	   var Base64var1 int = 0 //int var1 = 0;

	   var Base64var2 int //int var2;
	   	for Base64var2 = len(var0); Base64var1 < Base64var2; Base64var1++ { //for(var2 = var0.length(); var1 < var2; ++var1) {
			var var3 string = string(Base64var0[Base64var1]) //char var3 = var0.charAt(var1);
			var var10000 []byte = BASE64CHARS //byte[] var10000 = BASE64CHARS;
			var10000[var3 / 8] = byte(var10000[var3 / 8] | 1 << 7 - var3 % 8) //var10000[var3 / 8] = (byte)(var10000[var3 / 8] | 1 << 7 - var3 % 8);
		}

	   BASE64_ASCIIS = make([]string, len(Base64var0), len(Base64var0)) //BASE64_ASCIIS = new char[var0.length()];
	   Base64var1 = 0 //var1 = 0;

	   for Base64var2 = len(Base64var0); Base64var1 < Base64var2; Base64var1++ { //for(var2 = var0.length(); var1 < var2; ++var1) {
		  BASE64_ASCIIS[Base64var1] = string(Base64var0[Base64var1]) //BASE64_ASCIIS[var1] = var0.charAt(var1);
	   }

	   m_octets = make([]byte,256,256) //m_octets = new byte[256];

	   var Base64var4 int //int var4;
	   for Base64var4 = 90; Base64var4 >= 65; Base64var4 = Base64var4-1  { //for(var4 = 90; var4 >= 65; --var4) {
		  m_octets[Base64var4] = byte(Base64var4 - 65) //m_octets[var4] = (byte)(var4 - 65);
	   }

	   for Base64var4 = 122; Base64var4 >= 97; Base64var4 = Base64var4 - 1 { //for(var4 = 122; var4 >= 97; --var4) {
		  m_octets[Base64var4] = byte(Base64var4 - 97 + 26) //m_octets[var4] = (byte)(var4 - 97 + 26);
	   }

	   for Base64var4 = 57; Base64var4 >= 48; Base64var4 = Base64var4 - 1 { //for(var4 = 57; var4 >= 48; --var4) {
		  m_octets[vaBase644] = byte(Base64var4 - 48 + 52) //m_octets[var4] = (byte)(var4 - 48 + 52);
	   }

	   m_octets[43] = 62 //m_octets[43] = 62;
	   m_octets[47] = 63 //m_octets[47] = 63;
	}
*/

////////////////end Base64.class/////////
////////////////HexBin.class//////////
var HEXBIN_ASCIIS string = "0123456789ABCDEF"

func HexBinencode(var0 []byte, var1 int, var2 *strings.Builder) { //public static void encode(byte[] var0, int var1, StringBuffer var2) {
	if var0 != nil && (*var2).String() != "" { //if (var0 != null && var2 != null) {
		for var3 := 0; var3 < var1; var3++ { //for(int var3 = 0; var3 < var1; ++var3) {
			var var4 int //int var4;

			var4 = int(var0[var3] >> 4)
			if (var0[var3] >> 4) < 0 { //if ((var4 = var0[var3] >> 4) < 0) {
				var4 &= 15 //var4 &= 15;
			}

			var var5 int = int(var0[var3] & 15)              //int var5 = var0[var3] & 15;
			(*var2).WriteString(string(HEXBIN_ASCIIS[var4])) //var2.append("0123456789ABCDEF".charAt(var4));
			(*var2).WriteString(string(HEXBIN_ASCIIS[var5])) //var2.append("0123456789ABCDEF".charAt(var5));
		}
	}

}

////////////////end HexBin.class/////

/////////////////end EXISchema.class////////////////////////

//////////EventType.class/////////////////////////////////

type EventType struct {
	ITEM_PI                      byte
	ITEM_CM                      byte
	ITEM_ER                      byte
	ITEM_CH                      byte
	ITEM_ED                      byte
	ITEM_SE_WC                   byte
	ITEM_SC                      byte
	ITEM_NS                      byte
	ITEM_AT_WC_ANY_UNTYPED       byte
	ITEM_EE                      byte
	ITEM_DTD                     byte
	ITEM_SE                      byte
	ITEM_AT                      byte
	ITEM_SD                      byte
	ITEM_SCHEMA_WC_ANY           byte
	ITEM_SCHEMA_WC_NS            byte
	ITEM_SCHEMA_AT               byte
	ITEM_SCHEMA_AT_WC_ANY        byte
	ITEM_SCHEMA_AT_WC_NS         byte
	ITEM_SCHEMA_CH               byte
	ITEM_SCHEMA_CH_MIXED         byte
	ITEM_SCHEMA_NIL              byte
	ITEM_SCHEMA_TYPE             byte
	ITEM_SCHEMA_AT_INVALID_VALUE byte
	depth                        byte
	uri                          string
	name                         string
	//subsequentGrammar            IGrammar
	m_path      []EventCode
	m_uriId     int
	m_nameId    int
	m_eventKind int //byte
	m_index     int
	m_ownerList EventTypeList
}

var VAREventType EventType = EventType{
	ITEM_PI:                      0,
	ITEM_CM:                      1,
	ITEM_ER:                      2,
	ITEM_CH:                      3,
	ITEM_ED:                      4,
	ITEM_SE_WC:                   5,
	ITEM_SC:                      6,
	ITEM_NS:                      7,
	ITEM_AT_WC_ANY_UNTYPED:       8,
	ITEM_EE:                      9,
	ITEM_DTD:                     10,
	ITEM_SE:                      11,
	ITEM_AT:                      12,
	ITEM_SD:                      13,
	ITEM_SCHEMA_WC_ANY:           14,
	ITEM_SCHEMA_WC_NS:            15,
	ITEM_SCHEMA_AT:               16,
	ITEM_SCHEMA_AT_WC_ANY:        17,
	ITEM_SCHEMA_AT_WC_NS:         18,
	ITEM_SCHEMA_CH:               19,
	ITEM_SCHEMA_CH_MIXED:         20,
	ITEM_SCHEMA_NIL:              21,
	ITEM_SCHEMA_TYPE:             22,
	ITEM_SCHEMA_AT_INVALID_VALUE: 23,
}
var nilVAREventType EventType

func (e *EventType) getNameId() int { //public final int getNameId() {
	return e.m_nameId
}

func (e *EventType) getURIId() int { //public final int getURIId() {
	return e.m_uriId
}

func (e *EventType) myEventTypeToString() string {
	var var1 string = fmt.Sprintf("%d", e.ITEM_PI)
	var var2 string = fmt.Sprintf("%d", e.ITEM_CM)
	var var3 string = fmt.Sprintf("%d", e.ITEM_ER)
	var var4 string = fmt.Sprintf("%d", e.ITEM_CH)
	var var5 string = fmt.Sprintf("%d", e.ITEM_ED)
	var var6 string = fmt.Sprintf("%d", e.ITEM_SE_WC)
	var var7 string = fmt.Sprintf("%d", e.ITEM_SC)
	var var8 string = fmt.Sprintf("%d", e.ITEM_NS)
	var var9 string = fmt.Sprintf("%d", e.ITEM_AT_WC_ANY_UNTYPED)
	var var10 string = fmt.Sprintf("%d", e.ITEM_EE)
	var var11 string = fmt.Sprintf("%d", e.ITEM_DTD)
	var var12 string = fmt.Sprintf("%d", e.ITEM_SE)
	var var13 string = fmt.Sprintf("%d", e.ITEM_AT)
	var var14 string = fmt.Sprintf("%d", e.ITEM_SD)
	var var15 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_WC_ANY)
	var var16 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_WC_NS)
	var var17 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_AT)
	var var18 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_AT_WC_ANY)
	var var19 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_AT_WC_NS)
	var var20 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_CH)
	var var21 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_CH_MIXED)
	var var22 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_NIL)
	var var23 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_TYPE)
	var var24 string = fmt.Sprintf("%d", e.ITEM_SCHEMA_AT_INVALID_VALUE)
	var var25 string = fmt.Sprintf("%d", e.depth)
	var var26 string = e.uri
	var var27 string = e.name

	var var28 string
	for i := 0; i < len(e.m_path); i++ {
		var28 += e.m_path[i].myEventCodeToString()
	}

	//var arr28 string = e.m_path.myEventCodeToString()
	var var29 string = fmt.Sprintf("%d", e.m_uriId)
	var var30 string = fmt.Sprintf("%d", e.m_nameId)
	var var31 string = fmt.Sprintf("%d", e.m_eventKind)
	var var32 string = fmt.Sprintf("%d", e.m_index)
	var var33 string = e.m_ownerList.myEventTypeListToString()

	var str string
	str = str + var1 + var2 + var3 + var4 + var5 + var6 + var7 + var8 + var9 + var10 + var11 + var12 + var13 + var14 + var15 + var16 + var17 + var18 + var19 + var20 + var21 + var22 + var23 + var24 + var25 + var26 + var27 + var28 + var29 + var30 + var31 + var32 + var33
	return str
}

func thisEventType1(var1 byte, var2 EventTypeList, var3 int /*, var4 IGrammar*/) EventType { //public EventType(byte var1, EventTypeList var2, byte var3, IGrammar var4) {
	return EventType3("", "", -1, -1, var1, var2, var3, -1 /*, var4*/) //this((String)null, (String)null, -1, -1, var1, var2, var3, (byte)-1, var4);
}

func thisEventType2(var1 string, var2 string, var3 int, var4 int, var5 byte, var6 EventTypeList, var7 int /*, var8 IGrammar*/) EventType { //public EventType(String var1, String var2, int var3, int var4, byte var5, EventTypeList var6, byte var7, IGrammar var8) {
	return EventType3(var1, var2, var3, var4, var5, var6, var7, -1 /*, var8*/) //this(var1, var2, var3, var4, var5, var6, var7, (byte)-1, var8);
}

func EventType3(var1 string, var2 string, var3 int, var4 int, var5 byte, var6 EventTypeList, var7 byte, var8 int /*var9 IGrammar*/) EventType { //public EventType(String var1, String var2, int var3, int var4, byte var5, EventTypeList var6, byte var7, byte var8, IGrammar var9) {
	y := VAREventCode //	super(var7);
	y.itemType = var7

	var x EventType
	x.depth = var5                           //	this.depth = var5;
	x.m_path = make([]EventCode, var5, var5) //	this.m_path = new EventCode[var5];
	for i := 0; i < len(x.m_path); i++ {
		x.m_path[i] = y
	}

	x.uri = var1         //	this.uri = var1;
	x.name = var2        //	this.name = var2;
	x.m_uriId = var3     //	this.m_uriId = var3;
	x.m_nameId = var4    //	this.m_nameId = var4;
	x.m_index = -1       //	this.m_index = -1;
	x.m_ownerList = var6 //	this.m_ownerList = var6;
	//x.subsequentGrammar = var9 //	this.subsequentGrammar = var9;
	x.m_eventKind = var8 //	this.m_eventKind = var8;
	return x
}

func (e *EventType) computeItemPath() { //public final void computeItemPath() {
	var var1 int = e.getDepth() //int var1 = this.getDepth();
	var var2 int = 0            //int var2 = 0;

	//for(Object var3 = this; var2 < var1; ++var2) {
	//this.m_path[var1 - 1 - var2] = (EventCode)var3;
	//var3 = ((EventCode)var3).parent;
	//}

}

///////EventCode.class/////////
type EventCode struct {
	ITEM_TUPLE             int //byte
	EVENT_CODE_DEPTH_ONE   byte
	EVENT_CODE_DEPTH_TWO   byte
	EVENT_CODE_DEPTH_THREE byte
	//parent EventCode
	position int
	itemType byte
}

var VAREventCode EventCode = EventCode{
	ITEM_TUPLE:             -1,
	EVENT_CODE_DEPTH_ONE:   1,
	EVENT_CODE_DEPTH_TWO:   2,
	EVENT_CODE_DEPTH_THREE: 3,
	//parent = null;
	position: -1,
}

var varParentEventCode EventCode

var nilVAREventCode EventCode = EventCode{}

func (e *EventCode) myEventCodeToString() string {
	var var1 string = fmt.Sprintf("%d", e.ITEM_TUPLE)
	var var2 string = fmt.Sprintf("%d", e.EVENT_CODE_DEPTH_ONE)
	var var3 string = fmt.Sprintf("%d", e.EVENT_CODE_DEPTH_TWO)
	var var4 string = fmt.Sprintf("%d", e.EVENT_CODE_DEPTH_THREE)
	//var var5 string = fmt.Sprintf("%d", e.parent)
	var var6 string = fmt.Sprintf("%d", e.position)
	var var7 string = fmt.Sprintf("%d", e.itemType)
	var str string
	str = str + var1 + var2 + var3 + var4 /*+ var5*/ + var6 + var7
	return str
}

func (e *EventCode) setParentalContext(var1 int, var2 EventCode) EventCode { //public final void setParentalContext(int var1, EventCode var2) {
	e.position = var1                        //this.position = var1;
	var varParentEventCode1 EventCode = var2 //this.parent = var2;
	return varParentEventCode1
}

/////end EventCode.class///////
///////EventTypeList.class/////
type EventTypeList struct {
	isReverse bool
	//public static final EventTypeList EMPTY = new 1(false);
}

func (e *EventTypeList) myEventTypeListToString() string {
	var str string = strconv.FormatBool(e.isReverse)
	return str
}

//////eng EventTypeList.class//

////EventTypeElement.class///
type EventTypeElement struct {
	ensuingGrammar Grammar
}

var VAREventTypeElement EventTypeElement

func (e *EventTypeElement) EventTypeElement1(var1 int, var2 string, var3 int, var4 string, var5 EventTypeList, var6 Grammar /*, var7 IGrammar*/) EventType {
	e.ensuingGrammar = var6
	return EventType3(var2, var4, var1, var3, byte(1), var5, byte(11), 2 /*, var7*/) //super(var2, var4, var1, var3, (byte)1, var5, (byte)11, (byte)2, var7);
	//this.ensuingGrammar = var6;
}

///end EventTypeElement.class///

///////end EventType.class/////////////////

//////////////EXIGrammar.class/////////////
type EXIGrammar struct { //final class EXIGrammar extends SchemaInformedGrammar implements IGrammar {
	m_eventTypes    []EventType
	m_eventCode     EventCodeTuple
	m_eventTypeList ArrayEventTypeList
}

var nilVAREXIGrammar EXIGrammar = EXIGrammar{}

var VAREXIGrammar EXIGrammar = EXIGrammar{
	m_eventCode:     VAREventCodeTuple,
	m_eventTypeList: VARArrayEventTypeList,
}

func superEXIGrammar(var1 GrammarCache) {
	superSchemaInformedGrammar(byte(5), var1)
}

func (e *EXIGrammar) substantiate(var1 int, var2 bool) { //void substantiate(int var1, boolean var2) {
	var var3 int //int var3;
	if var2 {    //if (var2) {
		if var1 == -1 { //assert var1 != -1;
			panic("EXIGrammar class, func substantiate, var1 == -1")
		}
		var var4 int = VARGrammar.schema.getTypeOfElem(var1) //int var4 = this.schema.getTypeOfElem(var1);
		var3 = VARGrammar.schema.getGrammarOfType(var4)      //var3 = this.schema.getGrammarOfType(var4);
	} else {
		var3 = var1 //var3 = var1;
	}

	var x ArrayEventTypeList
	e.m_eventTypeList = x //this.m_eventTypeList = new ArrayEventTypeList();

	var var12 []string                                                 //to doпроверить на правильность        //ArrayList var12 = new ArrayList();
	var var5 int = VARGrammar.schema.getProductionCountOfGrammar(var3) //int var5 = this.schema.getProductionCountOfGrammar(var3);
	var var6 int = VARGrammar.schema.getContentGrammarOfGrammar(var3)  //int var6 = this.schema.getContentGrammarOfGrammar(var3);

	var var7 []string
	firstvar7elem := 0
	if isPermitDeviation(VARGrammar.m_grammarCache.grammarOptions) && (var6 != -1) {
		var myvar7 []string = make([]string, 0, 1)
		var7 = myvar7
	}
	//ArrayList var7 = GrammarOptions.isPermitDeviation(this.m_grammarCache.grammarOptions) && var6 != -1 ? new ArrayList() : null;

	var var8 int                        //int var8;
	var var9 int                        //int var9;
	for var8 = 0; var8 < var5; var8++ { //for(var8 = 0; var8 < var5; ++var8) {
		var9 = VARGrammar.schema.getProductionOfGrammar(var3, var8)      //var9 = this.schema.getProductionOfGrammar(var3, var8);
		var var10 EventType = e.createEventType(var9, e.m_eventTypeList) //EventType var10 = this.createEventType(var9, this.m_eventTypeList);
		var12 = append(var12, var10.myEventTypeToString())               //var12.add(var10);

		if var7 != nil && var10.itemType == 16 { //if (var7 != null && var10.itemType == 16) {
			if firstvar7elem == 0 {
				x := createEventTypeSchemaAttributeInvalid(EventTypeSchema(var10), e.m_eventTypeList)
				var7 = append(var7, x.myEventTypeToString())
				firstvar7elem++
			}
			y := createEventTypeSchemaAttributeInvalid(EventTypeSchema(var10), e.m_eventTypeList)
			var7 = append(var7, y.myEventTypeToString()) //var7.add(this.createEventTypeSchemaAttributeInvalid((EventTypeSchema)var10, this.m_eventTypeList));
		}
	}

	if VARGrammar.schema.hasEndElement(var3) { //if (this.schema.hasEndElement(var3)) {
		var var13 EventType = /*EventTypeFactory.*/ creatEndElement(byte(1), e.m_eventTypeList) //EventType var13 = EventTypeFactory.creatEndElement((byte)1, this.m_eventTypeList);
		var9 = len(var12)                                                                       //var9 = var12.size();
		var var15 int = 0                                                                       //int var15 = 0;
		if var9 != 0 {                                                                          //if (var9 != 0) {
			var var11 EventType = var12[var9-1] //EventType var11 = (EventType)var12.get(var9 - 1);
			if var11.itemType == 20 {           //var15 = var11.itemType == 20 ? var9 - 1 : var9;
				var15 = var9 - 1
			} else {
				var15 = var9
			}
		}

		var12 = append(var12, var15, var13) //var12.add(var15, var13);
	}

	var8 = VARGrammar.schema.getTypeEmptyGrammarOfGrammar(var3)                                                                                                     //var8 = this.schema.getTypeEmptyGrammarOfGrammar(var3);
	var var14 EventCodeTupleSink                                                                                                                                    //EventCodeTupleSink var14;
	e.createEventCodeTuple(var12, e.m_grammarCache.grammarOptions, var14, var7, e.m_eventTypeList, VARGrammar.schema.hasEmptyGrammar(var3), var1, var2, var8, var6) //this.createEventCodeTuple(var12, this.m_grammarCache.grammarOptions, var14 = new EventCodeTupleSink(), var7, this.m_eventTypeList, this.schema.hasEmptyGrammar(var3), var1, var2, var8, var6);
	e.m_eventCode = var14.eventCodeTuple                                                                                                                            //this.m_eventCode = var14.eventCodeTuple;
	e.m_eventTypes = var14.eventTypes                                                                                                                               //this.m_eventTypes = var14.eventTypes;
	e.m_eventTypeList.setItems(var14.eventTypes)                                                                                                                    //this.m_eventTypeList.setItems(var14.eventTypes);
}

func (e *EXIGrammar) createEventType(var1 int, var2 EventTypeList) EventType { //private EventType createEventType(int var1, EventTypeList var2) {
	var var3 int = VARGrammar.schema.getGrammarOfProduction(var1) //int var3 = this.schema.getGrammarOfProduction(var1);
	var var4 EXIGrammar = VARGrammar.retrieveEXIGrammar(var3)     //EXIGrammar var4 = this.retrieveEXIGrammar(var3);
	var var5 int                                                  //int var5;
	var5 = VARGrammar.schema.getEventOfProduction(var1)
	switch var5 { //switch(var5 = this.schema.getEventOfProduction(var1)) {
	case -4: //case -4:

		return thisEventType1(byte(1), var2, 19 /*, var4*/) //return new EventType((byte)1, var2, (byte)19, var4);
	case -3: //case -3:
		return thisEventType1(byte(1), var2, 20 /*, var4*/) //return new EventType((byte)1, var2, (byte)20, var4);
	case -2: //case -2:
		return thisEventType1(byte(1), var2, 14 /*, var4*/) //return new EventType((byte)1, var2, (byte)14, var4);
	case -1: //case -1:
		return thisEventType1(byte(1), var2, 17 /*, var4*/) //return new EventType((byte)1, var2, (byte)17, var4);
	default: //default:
		var var6 int                                  //int var6;
		var var7 int                                  //int var7;
		switch VARGrammar.schema.getEventType(var5) { //switch(this.schema.getEventType(var5)) {
		case 0: //case 0:
			var var8 int = VARGrammar.schema.getNodeOfEventType(var5)                                                                                                                   //int var8 = this.schema.getNodeOfEventType(var5);
			var6 = VARGrammar.schema.getUriOfAttr(var8)                                                                                                                                 //var6 = this.schema.getUriOfAttr(var8);
			var7 = VARGrammar.schema.getLocalNameOfAttr(var8)                                                                                                                           //var7 = this.schema.getLocalNameOfAttr(var8);
			return EventTypeSchema1(VARGrammar.schema.getTypeOfAttr(var8), VARGrammar.schema.uris[var6], VARGrammar.schema.localNames[var6][var7], var6, var7, byte(1), var2, 16, var4) //return new EventTypeSchema(this.schema.getTypeOfAttr(var8), this.schema.uris[var6], this.schema.localNames[var6][var7], var6, var7, (byte)1, var2, (byte)16, var4);
		case 1: //case 1:
			var var9 int = VARGrammar.schema.getNodeOfEventType(var5)                                                                                                //int var9 = this.schema.getNodeOfEventType(var5);
			var6 = VARGrammar.schema.getUriOfElem(var9)                                                                                                              //var6 = this.schema.getUriOfElem(var9);
			var7 = VARGrammar.schema.getLocalNameOfElem(var9)                                                                                                        //var7 = this.schema.getLocalNameOfElem(var9);
			var var10 EXIGrammarUse = VARGrammar.m_grammarCache.exiGrammarUses[VARGrammar.schema.getSerialOfElem(var9)]                                              //EXIGrammarUse var10 = this.m_grammarCache.exiGrammarUses[this.schema.getSerialOfElem(var9)];
			return VAREventTypeElement.EventTypeElement1(var6, VARGrammar.schema.uris[var6], var7, VARGrammar.schema.localNames[var6][var7], var2, var10 /*, var4*/) //return new EventTypeElement(var6, this.schema.uris[var6], var7, this.schema.localNames[var6][var7], var2, var10, var4);
		case 2: //case 2:
			var6 = VARGrammar.schema.getUriOfEventType(var5)                                            //var6 = this.schema.getUriOfEventType(var5);
			return EventType3(VARGrammar.schema.uris[var6], "", var6, -1, byte(1), var2, 18 /*, var4*/) //return new EventType(this.schema.uris[var6], (String)null, var6, -1, (byte)1, var2, (byte)18, var4);
		case 3: //case 3:
			var6 = VARGrammar.schema.getUriOfEventType(var5)                                            //var6 = this.schema.getUriOfEventType(var5);
			return EventType3(VARGrammar.schema.uris[var6], "", var6, -1, byte(1), var2, 15 /*, var4*/) //return new EventType(this.schema.uris[var6], (String)null, var6, -1, (byte)1, var2, (byte)15, var4);
		default: //default:
			//assert false;
			panic("EXIGrammar.class, func createEventType, in default")
			//return null;
		}
	}
}

//EventCodeTupleSink.class///
type EventCodeTupleSink struct {
	eventCodeTuple EventCodeTuple
	eventTypes     EventType
}

//end EventCodeTupleSink.class///

///ArrayEventTypeList.class////
type ArrayEventTypeList struct { //final class ArrayEventTypeList extends EventTypeList {
	m_eventTypes                   []EventType       //private EventType[] m_eventTypes;
	SCHEMA_ATTRIBUTES_NONE         []EventTypeSchema //private static final EventTypeSchema[] SCHEMA_ATTRIBUTES_NONE = new EventTypeSchema[0];
	SCHEMA_ATTRIBUTES_INVALID_NONE []EventTypeSchema //private static final EventType[] SCHEMA_ATTRIBUTES_INVALID_NONE = new EventTypeSchema[0];
	SCHEMA_ATTRIBUTES_WILDCARD_NS  []EventType       //private static final EventType[] SCHEMA_ATTRIBUTES_WILDCARD_NS = new EventType[0];
	m_schemaAttributes             []EventTypeSchema //private EventTypeSchema[] m_schemaAttributes;
	m_schemaAttributesInvalid      []EventType       //private EventType[] m_schemaAttributesInvalid;
	m_n_schemaAttributes           int               //private int m_n_schemaAttributes;
	m_n_schemaAttributesInvalid    int               //private int m_n_schemaAttributesInvalid;
	m_schemaAttributeWildcard      EventType         //private EventType m_schemaAttributeWildcard;
	m_attributeWildcardAnyUntyped  EventType         ///private EventType m_attributeWildcardAnyUntyped;
	m_schemaAttributeWildcardNS    []EventType       //private EventType[] m_schemaAttributeWildcardNS;
	m_n_schemaAttributesWildcardNS int               //private int m_n_schemaAttributesWildcardNS;
	m_schemaCharacters             EventType         //private EventType m_schemaCharacters;
	m_characters                   EventType         //private EventType m_characters;
	m_namespaceDeclaration         EventType         //private EventType m_namespaceDeclaration;
}

var VARArrayEventTypeList ArrayEventTypeList = ArrayEventTypeList{
	SCHEMA_ATTRIBUTES_NONE:         make([]EventTypeSchema, 0, 0), //private static final EventTypeSchema[] SCHEMA_ATTRIBUTES_NONE = new EventTypeSchema[0];
	SCHEMA_ATTRIBUTES_INVALID_NONE: make([]EventTypeSchema, 0, 0), //private static final EventType[] SCHEMA_ATTRIBUTES_INVALID_NONE = new EventTypeSchema[0];
	SCHEMA_ATTRIBUTES_WILDCARD_NS:  make([]EventType, 0, 0),       //private static final EventType[] SCHEMA_ATTRIBUTES_WILDCARD_NS = new EventType[0];
}

func (a *ArrayEventTypeList) setItems(var1 []EventType) { //final void setItems(EventType[] var1) {
	//assert this.m_eventTypes == null;

	a.m_eventTypes = var1 //this.m_eventTypes = var1;
	var var2 int = 0      //int var2 = 0;

	for var3 := len(a.m_eventTypes); var2 < var3; var2++ { //for(int var3 = this.m_eventTypes.length; var2 < var3; ++var2) {
		var var4 EventType = var1[var2] //EventType var4 = var1[var2];
		var4.setIndex(var2)             //var4.setIndex(var2);
		var var5 int = var4.itemType    //short var5 = (short)var4.itemType;
		var var6 int                    //int var6;
		var var7 []EventType            //EventType[] var7;
		switch var5 {                   //switch(var5) {
		case 3: //case 3:
			if a.m_characters == nil { //if (this.m_characters == null) {
				a.m_characters = var4 //this.m_characters = var4;
			}
		case 4: //case 4:
		case 5: //case 5:
		case 6: //case 6:
		case 9: //case 9:
		case 10: //case 10:
		case 11: //case 11:
		case 12: //case 12:
		case 13: //case 13:
		case 14: //case 14:
		case 15: //case 15:
		case 21: //case 21:
		case 22: //case 22:
		default: //default:
			break //break;
		case 7: //case 7:
			a.m_namespaceDeclaration = var4 //this.m_namespaceDeclaration = var4;
			break                           //break;
		case 8: //case 8:
			if a.m_attributeWildcardAnyUntyped == nil { //if (this.m_attributeWildcardAnyUntyped == null) {
				a.m_attributeWildcardAnyUntyped = var4 //this.m_attributeWildcardAnyUntyped = var4;
			}
			break //break;
		case 16: //case 16:
			if len(a.m_schemaAttributes) == a.m_n_schemaAttributes { //if (this.m_schemaAttributes.length == this.m_n_schemaAttributes) {
				if a.m_n_schemaAttributes == 0 { //var6 = this.m_n_schemaAttributes == 0 ? 4 : 2 * this.m_n_schemaAttributes;
					var6 = 4
				} else {
					var6 = (2 * a.m_n_schemaAttributes)
				}
				var var8 []EventTypeSchema = make([]EventTypeSchema, var6, var6) //EventTypeSchema[] var8 = new EventTypeSchema[var6];
				if a.m_n_schemaAttributes != 0 {                                 //if (this.m_n_schemaAttributes != 0) {
					copy(myCopyArrEventTypeSchema(a.m_schemaAttributes, 0, var8, 0, a.m_n_schemaAttributes), var8) //System.arraycopy(this.m_schemaAttributes, 0, var8, 0, this.m_n_schemaAttributes);
				}

				a.m_schemaAttributes = var8 //this.m_schemaAttributes = var8;
			}

			a.m_schemaAttributes[a.m_n_schemaAttributes] = EventTypeSchema(var4) //this.m_schemaAttributes[this.m_n_schemaAttributes++] = (EventTypeSchema)var4;
			a.m_n_schemaAttributes++
			break //break;
		case 17: //case 17:
			if a.m_schemaAttributeWildcard == nil { //if (this.m_schemaAttributeWildcard == null) {
				a.m_schemaAttributeWildcard = var4 //this.m_schemaAttributeWildcard = var4;
			} else { //} else {
				//assert this.m_schemaAttributeWildcard.itemType == 17;
			}
			break //break;
		case 18: //case 18:
			if len(a.m_schemaAttributeWildcardNS) == a.m_n_schemaAttributesWildcardNS { //if (this.m_schemaAttributeWildcardNS.length == this.m_n_schemaAttributesWildcardNS) {
				if a.m_n_schemaAttributesWildcardNS == 0 { //var6 = this.m_n_schemaAttributesWildcardNS == 0 ? 4 : 2 * this.m_n_schemaAttributesWildcardNS;
					var6 = 4
				} else {
					var6 = 2 * a.m_n_schemaAttributesWildcardNS
				}
				var7 = make([]EventType, var6, var6)       //var7 = new EventType[var6];
				if a.m_n_schemaAttributesWildcardNS != 0 { //if (this.m_n_schemaAttributesWildcardNS != 0) {
					copy(myCopyArrEventType(a.m_schemaAttributeWildcardNS, 0, var7, 0, a.m_n_schemaAttributesWildcardNS), var7) //System.arraycopy(this.m_schemaAttributeWildcardNS, 0, var7, 0, this.m_n_schemaAttributesWildcardNS);
				}

				a.m_schemaAttributeWildcardNS = var7 //this.m_schemaAttributeWildcardNS = var7;
			}

			a.m_schemaAttributeWildcardNS[a.m_n_schemaAttributesWildcardNS] = var4 //this.m_schemaAttributeWildcardNS[this.m_n_schemaAttributesWildcardNS++] = var4;
			a.m_n_schemaAttributesWildcardNS++
			break //break;
		case 19: //case 19:
			a.m_schemaCharacters = var4 //this.m_schemaCharacters = var4;
			break                       //break;
		case 20: //case 20:
			//assert this.m_characters == null;

			a.m_characters = var4 //this.m_characters = var4;
			break                 //break;
		case 23: //case 23:
			if len(a.m_schemaAttributesInvalid) == a.m_n_schemaAttributesInvalid { //if (this.m_schemaAttributesInvalid.length == this.m_n_schemaAttributesInvalid) {
				if a.m_n_schemaAttributesInvalid == 0 { //var6 = this.m_n_schemaAttributesInvalid == 0 ? 4 : 2 * this.m_n_schemaAttributesInvalid;
					var6 = 4
				} else {
					var6 = 2 * a.m_n_schemaAttributesInvalid
				}
				var var7 []EventType = make([]EventType, var6, var6) //var7 = new EventType[var6];
				if a.m_n_schemaAttributesInvalid != 0 {              //if (this.m_n_schemaAttributesInvalid != 0) {
					copy(myCopyArrEventType(a.m_schemaAttributesInvalid, 0, var7, 0, a.m_n_schemaAttributesInvalid), var7) //System.arraycopy(this.m_schemaAttributesInvalid, 0, var7, 0, this.m_n_schemaAttributesInvalid);
				}

				a.m_schemaAttributesInvalid = var7 //this.m_schemaAttributesInvalid = var7;
			}

			a.m_schemaAttributesInvalid[a.m_n_schemaAttributesInvalid] = var4 //this.m_schemaAttributesInvalid[this.m_n_schemaAttributesInvalid++] = var4;
			a.m_n_schemaAttributesInvalid++
		}
	}

}

///EventTypeSchema.class///
type EventTypeSchema struct { //public final class EventTypeSchema extends EventType {
	nd int
}

var VAREventTypeSchema EventTypeSchema

func EventTypeSchema1(var1 int, var2 string, var3 string, var4 int, var5 int, var6 byte, var7 EventTypeList, var8 byte, var9 EXIGrammar) EventType { //EventTypeSchema(int var1, String var2, String var3, int var4, int var5, byte var6, EventTypeList var7, byte var8, EXIGrammar var9) {
	//super(var2, var3, var4, var5, var6, var7, var8, (byte)-1, var9);
	VAREventTypeSchema.nd = var1                                               //this.nd = var1;
	return EventType3(var2, var3, var4, var5, var6, var7, var8, -1 /*, var9*/) ////!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}

///end EventTypeSchema.class////

///end ArrayEventTypeList.class/////

///EventCodeTuple.class///
type EventCodeTuple struct { //public abstract class EventCodeTuple extends EventCode {
	reversed   bool
	eventCodes []EventCode
	width      int
	itemsCount int
	headItem   EventCode
}

var VAREventCodeTuple EventCodeTuple = EventCodeTuple{
	width:      0,
	itemsCount: 0,
}

func (e *EventCodeTuple) createEventCodeTuple(var1 []string, var2 int, var3 EventCodeTupleSink, var4 EventTypeList) { //protected final void createEventCodeTuple(ArrayList<EventType> var1, short var2, org.openexi.proc.grammars.SchemaInformedGrammar.EventCodeTupleSink var3, EventTypeList var4) {
	e.thisCreateEventCodeTuple(var1, var2, var3, nil, var4, false, -1, false, -1, -1) //this.createEventCodeTuple(var1, var2, var3, (ArrayList)null, var4, false, -1, false, -1, -1);
}

func (e *EventCodeTuple) thisCreateEventCodeTuple(var1 []string, var2 int, var3 EventCodeTupleSink, var4 []EventType, var5 EventTypeList, var6 bool, var7 int, var8 bool, var9 int, var10 int) { //protected final void createEventCodeTuple(ArrayList<EventType> var1, short var2, org.openexi.proc.grammars.SchemaInformedGrammar.EventCodeTupleSink var3, ArrayList<EventType> var4, EventTypeList var5, boolean var6, int var7, boolean var8, int var9, int var10) {
	var var11 int //int var11;
	var var12 int //int var12;
	var var13 int //int var13;
	if var8 {     //if (var8) {
		var12 = var7 //var12 = var7;

		if var7 == (-1) { //assert var7 != -1;
			panic("EventCodeTuple.class, func thisCreateEventCodeTuple, var7 == (-1)")
		}
		var13 = VARGrammar.schema.getTypeOfElem(var7)     //var13 = this.schema.getTypeOfElem(var7);
		var11 = VARGrammar.schema.getGrammarOfType(var13) //var11 = this.schema.getGrammarOfType(var13);

		if var11 == (-1) { //assert var11 != -1;
			panic("EventCodeTuple.class, func thisCreateEventCodeTuple, var11 == (-1)")
		}
	} else { //} else {
		var12 = -1   //var12 = -1;
		var11 = var7 //var11 = var7;
	}

	if var10 == -1 { //if (var10 == -1) {
		var10 = var11 //var10 = var11;
	}

	if var4 != nil { //var13 = var4 != null ? var4.size() : 0;
		var13 = len(var4)
	} else {
		var13 = 0
	}

	var var14 bool = isPermitDeviation(var2) //boolean var14 = GrammarOptions.isPermitDeviation(var2);

	if (var4 != nil) || (!var14) { //assert var4 == null || var14;
		panic("EventCodeTuple.class, func thisCreateEventCodeTuple, (var4 != null) || (!var14)")
	}

	var var15 bool = hasDTD(var2) //boolean var15 = GrammarOptions.hasDTD(var2);
	var var16 bool = hasCM(var2)  //boolean var16 = GrammarOptions.hasCM(var2);
	var var17 bool = hasPI(var2)  //boolean var17 = GrammarOptions.hasPI(var2);
	var var18 bool                //boolean var18;
	var var19 bool                //boolean var19;
	var var20 bool                //boolean var20;
	var var21 bool                //boolean var21;
	var var22 bool                //boolean var22;
	if var6 {                     //if (var6) {
		if VARGrammar.grammarType == 4 { //if (this.grammarType == 4) {
			var20 = true //var20 = true;
			var21 = true //var21 = true;
		} else { //} else {
			if VARGrammar.grammarType == 5 { //assert this.grammarType == 5;
				panic("EventCodeTuple.class, func thisCreateEventCodeTuple, e.grammarType == 5 ")
			}
			if !var8 { //if (!var8) {
				var22 = !isXsiNilTypeRestricted(var2) //var22 = !GrammarOptions.isXsiNilTypeRestricted(var2);
				var21 = var22                         //var21 = var22;
				var20 = var22                         //var20 = var22;
			} else { //} else {
				var22 = isXsiNilTypeRestricted(var2)                                                      //var22 = GrammarOptions.isXsiNilTypeRestricted(var2);
				var21 = !var22 || VARGrammar.schema.isTypableType(VARGrammar.schema.getTypeOfElem(var12)) //var21 = !var22 || this.schema.isTypableType(this.schema.getTypeOfElem(var12));
				var20 = !var22 || VARGrammar.schema.isNillableElement(var12)                              //var20 = !var22 || this.schema.isNillableElement(var12);
			}
		}

		var18 = hasNS(var2) //var18 = GrammarOptions.hasNS(var2);
		var19 = hasSC(var2) //var19 = GrammarOptions.hasSC(var2);

	} else { //} else {
		var19 = false //var19 = false;
		var18 = false //var18 = false;
		var20 = false //var20 = false;
		var21 = false //var21 = false;
	}

	var22 = false                         //var22 = false;
	var var23 EventType = nilVAREventType //EventType var23 = null;
	var var24 EventType = nilVAREventType //EventType var24 = null;
	var var25 EventType = nilVAREventType //EventType var25 = null;
	var var26 EventType = nilVAREventType //EventType var26 = null;
	var var27 EventType = nilVAREventType //EventType var27 = null;
	var var28 EventType = nilVAREventType //EventType var28 = null;
	var var29 EventType = nilVAREventType //EventType var29 = null;
	var var30 EventType = nilVAREventType //EventType var30 = null;
	var var31 EventType = nilVAREventType //EventType var31 = null;
	var var32 EventType = nilVAREventType //EventType var32 = null;
	var var33 EventType = nilVAREventType //EventType var33 = null;
	var var34 EventType = nilVAREventType //EventType var34 = null;
	var var35 bool = false                //boolean var35 = false;
	var var36 bool = false                //boolean var36 = false;
	var var37 bool = false                //boolean var37 = false;
	var var38 int = 0                     //int var38 = 0;
	var var39 int = 0                     //int var39 = 0;
	var var41 int = len(var1)             //int var41 = var1.size();
	var var40 int                         //int var40;
	if var14 {                            //if (var14) {
		for var40 = 0; var40 < var41; var40++ { //for(var40 = 0; var40 < var41; ++var40) {
			var var42 EventType = (*(*EventType)(unsafe.Pointer(&var1[var40]))) //EventType(var1.get(var40)) //EventType var42 = (EventType)var1.get(var40);
			if (*(*EventCode)(unsafe.Pointer(&var42))).itemType == 9 {          //if var42.itemType == 9 { //if (var42.itemType == 9) { ///EventCode
				break // break;
			}
		}

		if var40 == var41 { //if (var40 == var41) {
			var22 = true                           //var22 = true;
			var23 = creatEndElement(byte(2), var5) //var23 = EventTypeFactory.creatEndElement((byte)2, var5);
			var38++                                // ++var38;
		}

		var30 = thisEventType1(byte(2), var5, 5 /*, VARGrammar.retrieveEXIGrammar(var10)*/) //var30 = new EventType((byte)2, var5, (byte)5, this.retrieveEXIGrammar(var10));
		var31 = thisEventType1(byte(2), var5, 3 /*, VARGrammar.retrieveEXIGrammar(var10)*/) //var31 = new EventType((byte)2, var5, (byte)3, this.retrieveEXIGrammar(var10));
		var38 = var38 + 2                                                                   //var38 += 2;
		if var4 != nil {                                                                    //if (var4 != null) {
			var28 = thisEventType1(byte(2), var5, 17 /*, (IGrammar)null*/) //var28 = new EventType((byte)2, var5, (byte)17, (IGrammar)null);
			var29 = thisEventType1(byte(3), var5, 8 /*, (IGrammar)null*/)  //var29 = new EventType((byte)3, var5, (byte)8, (IGrammar)null);
			var38 = var38 + 2                                              //var38 += 2;
			var36 = true                                                   // var36 = true;
		}

		var35 = true // var35 = true;
	}

	if var21 { //if (var21) {
		var24 = createEventTypeXsiType(var5) // var24 = this.createEventTypeXsiType(var5);//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		var38++                              // ++var38;
		var35 = true                         // var35 = true;
	}

	if var20 { //if (var20) {
		var25 = createEventTypeXsiNil(var5 /*, e.retrieveEXIGrammar(var9)*/) //var25 = this.createEventTypeXsiNil(var5, this.retrieveEXIGrammar(var9));////e.retrieveEXIGrammar(var9)//!!!!!!!!!!!!!!!!!!!!!!!!!!!
		var38++                                                              //++var38;
		var35 = true                                                         // var35 = true;
	}

	if var18 { //if (var18) {
		var26 = thisEventType1(byte(2), var5, 7 /*, (IGrammar)null*/) //var26 = new EventType((byte)2, var5, (byte)7, (IGrammar)null);
		var38++                                                       //++var38;
		var35 = true                                                  // var35 = true;
	}

	if var19 { //if (var19) {
		var27 = thisEventType1(byte(2), var5, 6 /*, (IGrammar)null*/) // var27 = new EventType((byte)2, var5, (byte)6, (IGrammar)null);
		var38++                                                       // ++var38;
		var35 = true                                                  // var35 = true;
	}

	if var15 { //if (var15) {
		var32 = thisEventType1(byte(2), var5, 2 /*, this.retrieveEXIGrammar(var10)*/) // var32 = new EventType((byte)2, var5, (byte)2, this.retrieveEXIGrammar(var10));
		var38++                                                                       // ++var38;
		var35 = true                                                                  // var35 = true;
	}

	if var16 { //if (var16) {
		var33 = thisEventType1(byte(3), var5, 1 /*, this.retrieveEXIGrammar(var10)*/) //var33 = new EventType((byte)3, var5, (byte)1, this.retrieveEXIGrammar(var10));
		var39++                                                                       //++var39;
		var37 = true                                                                  // var37 = true;
	}

	if var17 { //if (var17) {
		var34 = thisEventType1(byte(3), var5, 0 /*, this.retrieveEXIGrammar(var10)*/) // var34 = new EventType((byte)3, var5, (byte)0, this.retrieveEXIGrammar(var10));
		var39++                                                                       // ++var39;
		var37 = true                                                                  /// var37 = true;
	}

	if var37 { //if (var37) {
		var38++      //++var38;
		var35 = true // var35 = true;
	}

	var var44 []EventCode = nil   //EventCode[] var44 = null;
	var var45 []EventCode = nil   //EventCode[] var45 = null;
	var var46 []EventCode = nil   //EventCode[] var46 = null;
	var var47 ArrayEventCodeTuple //ArrayEventCodeTuple var47 = null;
	var var48 ArrayEventCodeTuple //ArrayEventCodeTuple var48 = null;
	var var49 ArrayEventCodeTuple //ArrayEventCodeTuple var49 = null;

	var var50 int
	if var39 != 0 { //int var50 = var39 != 0 ? var41 + (var38 - 1) + var39 + var13 : var41 + var38 + var13;
		var50 = var41 + (var38 - 1) + var39 + var13
	} else {
		var50 = var41 + var38 + var13
	}

	var var55 ArrayEventCodeTuple                     //ArrayEventCodeTuple var55 = new ArrayEventCodeTuple();
	var var51 []EventType = make([]EventType, 50, 50) //EventType[] var51 = new EventType[var50];
	var var43 []EventCode
	if var35 { //EventCode[] var43 = new EventCode[var35 ? var41 + 1 : var41];
		var43 = make([]EventCode, var41+1, var41+1)
	} else {
		var43 = make([]EventCode, var41, var41)
	}

	var var52 int = 0 //int var52 = 0;

	if !(reflect.DeepEqual(var24, nilVAREventType)) { //nilVAREventType { //if (var24 != null) {
		var51[var52] = var24 // var51[var52++] = var24;
		var52++
	}

	if !(reflect.DeepEqual(var25, nilVAREventType)) { //var25 != nil { //if (var25 != null) {
		var51[var52] = var25 // var51[var52++] = var25;
		var52++
	}

	for var40 = 0; var40 < var41; var40++ { //for(var40 = 0; var40 < var41; ++var40) {
		var var53 EventType = (*(*EventType)(unsafe.Pointer(&(var1[var40])))) //EventType(var1[var40]) // EventType var53 = (EventType)var1.get(var40);
		var51[var52+var40] = var53                                            // var51[var52 + var40] = var53;
		var43[var40] = *(*EventCode)(unsafe.Pointer(&var53))                  // var43[var40] = var53;
	}

	var52 = var52 + var41 //var52 += var41;
	if var35 {            //if (var35) {
		var44 = make([]EventCode, var38, var38) //var44 = new EventCode[var38];
		var myvar ArrayEventCodeTuple
		var47 = myvar                                        //var47 = new ArrayEventCodeTuple();
		var43[var41] = *(*EventCode)(unsafe.Pointer(&var47)) //var43[var41] = var47;
		if var37 {                                           //if (var37) {
			var46 = make([]EventCode, var39, var39)                // var46 = new EventCode[var39];
			var49 = myvar                                          // var49 = new ArrayEventCodeTuple();
			var44[var38-1] = *(*EventCode)(unsafe.Pointer(&var49)) //  var44[var38 - 1] = var49;
		}
	}

	var var56 int = 0 //int var56 = 0;
	var var54 int = 0 //int var54 = 0;
	if var22 {        //if (var22) {
		var51[var52] = var23 // var51[var52++] = var23;
		var52++
		var44[var56] = (*(*EventCode)(unsafe.Pointer(&var23))) // var44[var56++] = var23;///полезная///gjktpyfz
		var56++
	}

	if var21 { //if (var21) {
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var24)) // var44[var56++] = var24;
		var56++
	}

	if var20 { //if (var20) {
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var25)) // var44[var56++] = var25;
		var56++
	}

	if var36 { //if (var36) {
		var51[var52] = var28 // var51[var52++] = var28;
		var52++
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var28)) // var44[var56++] = var28;
		var56++
		var45 = make([]EventCode, var13+1, var13+1) // var45 = new EventCode[var13 + 1];
		var myvar ArrayEventCodeTuple
		var48 = myvar                                        // var48 = new ArrayEventCodeTuple();
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var48)) // var44[var56++] = var48;
		var56++

		for var40 = 0; var40 < var13; var40++ { //for(var40 = 0; var40 < var13; ++var40) {
			var51[var52] = EventType(var4[var40]) // var51[var52++] = (EventType)var4.get(var40);
			var52++

			var45[var40] = *(*EventCode)(unsafe.Pointer(&(var4[var40]))) // var45[var40] = (EventCode)var4.get(var40);
		}

		var51[var52] = var29 // var51[var52++] = var29;
		var52++

		var45[var13] = *(*EventCode)(unsafe.Pointer(&var29)) // var45[var13] = var29;
	}

	if var18 { //if (var18) {
		var51[var52] = var26 // var51[var52++] = var26;
		var52++
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var26)) // var44[var56++] = var26;
		var56++
	}

	if var19 { //if (var19) {
		var51[var52] = var27 // var51[var52++] = var27;
		var52++
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var27)) // var44[var56++] = var27;*(*EventCode)(unsafe.Pointer(&var23))
		var56++
	}

	if var14 { //if (var14) {
		var51[var52] = var30 //var51[var52++] = var30;
		var52++
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var30)) //var44[var56++] = var30;
		var56++
		var51[var52] = var31 //var51[var52++] = var31;
		var52++
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var31)) //var44[var56++] = var31;
		var56++
	}

	if var15 { //if (var15) {
		var51[var52] = var32 // var51[var52++] = var32;
		var52++
		var44[var56] = *(*EventCode)(unsafe.Pointer(&var32)) // var44[var56++] = var32;
		var56++
	}

	if var16 { //if (var16) {
		var51[var52] = var33 // var51[var52++] = var33;
		var52++
		var46[var54] = *(*EventCode)(unsafe.Pointer(&var33)) // var46[var54++] = var33;
		var54++
	}

	if var17 { //if (var17) {
		var51[var52] = var34 //var51[var52++] = var34;
		var52++
		var46[var54] = *(*EventCode)(unsafe.Pointer(&var34)) //var46[var54++] = var34;
		var54++
	}

	if var37 { //assert var50 == var52 && (var37 ? var38 - 1 : var38) == var56 && var39 == var54;
		if (var50 != var52) && ((var38 - 1) != var56) && (var39 != var54) {
			panic("EventCodeTuple.class, func thisCreateEventCodeTuple, (var50 != var52) && ((var38 - 1 )!= var56) && (var39 != var54) ")
		}
	} else {
		if (var50 != var52) && (var38 != var56) && (var39 != var54) {
			panic("EventCodeTuple.class, func thisCreateEventCodeTuple, (var50 != var52) && (var38 != var56) && (var39 != var54)")
		}
	}

	var55.setItems(var43) //var55.setItems(var43);
	if var35 {            //if (var35) {
		var47.setItems(var44) //var47.setItems(var44);
		if var36 {            //if (var36) {
			var48.setItems(var45) //var48.setItems(var45);
		}

		if var37 { //if (var37) {
			var49.setItems(var46) //var49.setItems(var46);
		}
	}

	var3.eventCodeTuple = *(*EventCodeTuple)(unsafe.Pointer(&var55)) //var3.eventCodeTuple = var55;
	var3.eventTypes = *(*EventType)(unsafe.Pointer(&var51))          //var3.eventTypes = var51;
}

///end EventCodeTuple.class///
//
///SchemaInformedGrammar.class///
type SchemaInformedGrammar struct { //public abstract class SchemaInformedGrammar extends Grammar {
	ELEMENT_FRAGMENT_STATE_BASE          byte //protected static final byte ELEMENT_FRAGMENT_STATE_BASE = 4;
	ELEMENT_FRAGMENT_STATE_TAG           byte //public static final byte ELEMENT_FRAGMENT_STATE_TAG = 4;
	ELEMENT_FRAGMENT_STATE_CONTENT       byte //public static final byte ELEMENT_FRAGMENT_STATE_CONTENT = 5;
	ELEMENT_FRAGMENT_EMPTY_STATE_TAG     byte //public static final byte ELEMENT_FRAGMENT_EMPTY_STATE_TAG = 6;
	ELEMENT_FRAGMENT_EMPTY_STATE_CONTENT byte //public static final byte ELEMENT_FRAGMENT_EMPTY_STATE_CONTENT = 7;
}

var VARSchemaInformedGrammar SchemaInformedGrammar = SchemaInformedGrammar{
	ELEMENT_FRAGMENT_STATE_BASE:          4,
	ELEMENT_FRAGMENT_STATE_TAG:           4,
	ELEMENT_FRAGMENT_STATE_CONTENT:       5,
	ELEMENT_FRAGMENT_EMPTY_STATE_TAG:     6,
	ELEMENT_FRAGMENT_EMPTY_STATE_CONTENT: 7,
}

func superSchemaInformedGrammar(var1 byte, var2 GrammarCache) {
	superGrammar(var1, var2)
}

func createEventTypeSchemaAttributeInvalid(var1 EventType, var2 EventTypeList) EventType { //protected final EventType createEventTypeSchemaAttributeInvalid(EventType var1, EventTypeList var2) {
	var newEventType EventType
	newEventType = thisEventType2(var1.uri, var1.name, var1.getURIId(), var1.getNameId(), byte(3), var2, 23 /*, var1.subsequentGrammar*/)
	return newEventType //return new EventType(var1.uri, var1.name, var1.getURIId(), var1.getNameId(), (byte)3, var2, (byte)23, var1.subsequentGrammar);
}

func createEventTypeXsiNil(var1 EventTypeList /*, var2 IGrammar*/) EventType { //private EventType createEventTypeXsiNil(EventTypeList var1, IGrammar var2) {
	return thisEventType2("http://www.w3.org/2001/XMLSchema-instance", "nil", 2, 0, byte(2), var1, 21 /*, var2*/) //return new EventType("http://www.w3.org/2001/XMLSchema-instance", "nil", 2, 0, (byte)2, var1, (byte)21, var2);
}

func createEventTypeXsiType(var1 EventTypeList) EventType { //private EventType createEventTypeXsiType(EventTypeList var1) {
	return thisEventType2("http://www.w3.org/2001/XMLSchema-instance", "type", 2, 1, byte(2), var1, 22 /*, (IGrammar)null*/) //return new EventType("http://www.w3.org/2001/XMLSchema-instance", "type", 2, 1, (byte)2, var1, (byte)22, (IGrammar)null);
}

///end SchemaInformedGrammar.clas///
//
///Grammar.class//
type Grammar struct { //public abstract class Grammar {
	BUILTIN_GRAMMAR_ELEMENT             byte         //public static final byte BUILTIN_GRAMMAR_ELEMENT = 0;
	BUILTIN_GRAMMAR_FRAGMENT            byte         //public static final byte BUILTIN_GRAMMAR_FRAGMENT = 1;
	SCHEMA_GRAMMAR_DOCUMENT             byte         //public static final byte SCHEMA_GRAMMAR_DOCUMENT = 2;
	SCHEMA_GRAMMAR_FRAGMENT             byte         //public static final byte SCHEMA_GRAMMAR_FRAGMENT = 3;
	SCHEMA_GRAMMAR_ELEMENT_FRAGMENT     byte         //public static final byte SCHEMA_GRAMMAR_ELEMENT_FRAGMENT = 4;
	SCHEMA_GRAMMAR_ELEMENT_AND_TYPE     byte         //public static final byte SCHEMA_GRAMMAR_ELEMENT_AND_TYPE = 5;
	SCHEMA_GRAMMAR_ELEMENT_AND_TYPE_USE byte         //public static final byte SCHEMA_GRAMMAR_ELEMENT_AND_TYPE_USE = 6;
	DOCUMENT_STATE_BASE                 byte         //private static final byte DOCUMENT_STATE_BASE = 0;
	DOCUMENT_STATE_CREATED              byte         //protected static final byte DOCUMENT_STATE_CREATED = 0;
	DOCUMENT_STATE_DEPLETE              byte         //protected static final byte DOCUMENT_STATE_DEPLETE = 1;
	DOCUMENT_STATE_COMPLETED            byte         //public static final byte DOCUMENT_STATE_COMPLETED = 2;
	DOCUMENT_STATE_END                  byte         //public static final byte DOCUMENT_STATE_END = 3;
	grammarType                         byte         //public final byte grammarType;
	m_grammarCache                      GrammarCache //protected final GrammarCache  m_grammarCache;
	schema                              EXISchema    //protected final EXISchema schema;
}

var VARGrammar Grammar = Grammar{
	BUILTIN_GRAMMAR_ELEMENT:             0,
	BUILTIN_GRAMMAR_FRAGMENT:            1,
	SCHEMA_GRAMMAR_DOCUMENT:             2,
	SCHEMA_GRAMMAR_FRAGMENT:             3,
	SCHEMA_GRAMMAR_ELEMENT_FRAGMENT:     4,
	SCHEMA_GRAMMAR_ELEMENT_AND_TYPE:     5,
	SCHEMA_GRAMMAR_ELEMENT_AND_TYPE_USE: 6,
	DOCUMENT_STATE_BASE:                 0,
	DOCUMENT_STATE_CREATED:              0,
	DOCUMENT_STATE_DEPLETE:              1,
	DOCUMENT_STATE_COMPLETED:            2,
	DOCUMENT_STATE_END:                  3,
}

func superGrammar(var1 byte, var2 GrammarCache) {
	VARGrammar.grammarType = var1
	VARGrammar.m_grammarCache = var2
	VARGrammar.schema = VARGrammar.m_grammarCache.getEXISchema()
}

func (g *Grammar) retrieveEXIGrammar(var1 int) EXIGrammar { //protected final EXIGrammar retrieveEXIGrammar(int var1) {///!!!!todo nil
	if var1 != -1 { //if (var1 != -1) {
		var var2 int = g.schema.getSerialOfGrammar(var1) //int var2 = this.schema.getSerialOfGrammar(var1);
		var var3 EXIGrammar                              //EXIGrammar var3;
		var3 = g.m_grammarCache.exiGrammars[var2]
		if !(reflect.DeepEqual(g.m_grammarCache.exiGrammars[var2], nilVAREXIGrammar)) { //if ((var3 = this.m_grammarCache.exiGrammars[var2]) != null) {
			return var3 //return var3;
		} else { //} else {
			superEXIGrammar(g.m_grammarCache) //var3 = new EXIGrammar(this.m_grammarCache);
			var3 = nilVAREXIGrammar           //var3.substantiate(var1, false);
			return var3                       //return var3;
		}
	} else { //} else {
		return nilVAREXIGrammar //return null;
	}
}

///end Grammar.class///

///////////end EXIGrammar.class///////////

///EXIGrammarUse.class///
type EXIGrammarUse struct { //final class EXIGrammarUse extends Grammar implements IGrammar {
	exiGrammar      EXIGrammar //EXIGrammar exiGrammar;
	contentDatatype int        //final int contentDatatype
}

var VAREXIGrammarUse EXIGrammarUse = EXIGrammarUse{}

func (e *EXIGrammarUse) EXIGrammarUse1(var1 int, var2 GrammarCache) {
	//super((byte)6, var2);
	e.contentDatatype = var1
}

///end EXIGrammarUse.class///

////////////////////////////////////end GrammarCache.class////////////////////////////////////////////////////////////////////////////////

////SchemaId.class////////////////////
type SchemaId struct {
	m_value string
}

var VARSchemaId SchemaId = SchemaId{""}

func (s *SchemaId) SchemaId1(var1 string) {
	s.m_value = var1
}

func (s *SchemaId) getValue() string {
	return s.m_value
}

////end SchemaId.class/////////////////

type Transmogrifier struct {
	//m_xmlReader           *os.File //возможно string
	//m_saxHandler          string   // обработка SAX событий(событий XML)
	//m_outputOptions       HeaderOptionsOutputType
	m_exiOptions          EXIOptions
	SCHEMAID_NO_SCHEMA    SchemaId
	SCHEMAID_EMPTY_SCHEMA SchemaId
}

var VARTransmogrifier Transmogrifier = Transmogrifier{
	SCHEMAID_NO_SCHEMA:    VARSchemaId,
	SCHEMAID_EMPTY_SCHEMA: VARSchemaId,
}

func (t *Transmogrifier) setGrammarCache1(var1 GrammarCache) { //public final void setGrammarCache(GrammarCache var1) throws EXIOptionsException {
	t.setGrammarCache2(var1, SchemaId{""}) //this.setGrammarCache(var1, (SchemaId)null);
}

func (t *Transmogrifier) setGrammarCache2(var1 GrammarCache, var2 SchemaId) { //public final void setGrammarCache(GrammarCache var1, SchemaId var2) throws EXIOptionsException {
	var var3 EXISchema = var1.getEXISchema() //EXISchema var3 = var1.getEXISchema();
	if var2 == VARSchemaId {                 //if (var2 == null) {
		if reflect.DeepEqual(var3, nilVAREXISchema) { //if (var3 == null) {//todo if == else
			var2 = VARTransmogrifier.SCHEMAID_NO_SCHEMA //var2 = SCHEMAID_NO_SCHEMA;
		} else if reflect.DeepEqual(var3, nilVAREXISchema) { //} else if (var3 == EmptySchema.getEXISchema()) {
			var2 = VARTransmogrifier.SCHEMAID_EMPTY_SCHEMA //var2 = SCHEMAID_EMPTY_SCHEMA;
		}
	}

	t.m_exiOptions.setSchemaId(var2)                      //this.m_exiOptions.setSchemaId(var2);
	t.m_exiOptions.setGrammarOptions(var1.grammarOptions) //this.m_exiOptions.setGrammarOptions(var1.grammarOptions);
	//org.openexi.sax.Transmogrifier.SAXEventHandler.access$200(this.m_saxHandler, var1);
}

///////////EXIOptions.class////////////
type EXIOptions struct {
	ADD_LESSCOMMON                        int
	ADD_UNCOMMON                          int
	ADD_ALIGNMENT                         int
	ADD_PRESERVE                          int
	ADD_COMMON                            int
	ADD_VALUE_MAX_LENGTH                  int
	ADD_VALUE_PARTITION_CAPACITY          int
	ADD_FRAGMENT                          int
	ADD_DTRM                              int
	BLOCKSIZE_DEFAULT                     int
	VALUE_MAX_LENGTH_UNBOUNDED            int
	VALUE_PARTITION_CAPACITY_UNBOUNDED    int
	m_alignmentType                       AlignmentType
	m_isFragment                          bool
	m_isStrict                            bool
	m_preserveComments                    bool
	m_preservePIs                         bool
	m_preserveDTD                         bool
	m_preserveNS                          bool
	m_infuseSC                            bool
	m_schemaId                            SchemaId
	m_blockSize                           int
	m_valueMaxLength                      int
	m_valuePartitionCapacity              int
	m_preserveLexicalValues               bool
	m_n_datatypeRepresentationMapBindings int
	m_datatypeRepresentationMap           []QName
}

func (e *EXIOptions) EXIOptions() {
	//e.init()//init() используется для выполнения инициализации сервлета. Создает и загружает объекты, которые используются сервлетом при обработке его запроса.
	e.m_datatypeRepresentationMap = make([]QName, 16, 16)
	/*//думаю не нужно
	for var1 := 0; var1 < len(q.m_datatypeRepresentationMap); var1++ {
	   q.m_datatypeRepresentationMap[var1] = new QName();
	}
	*/
}

var VAREXIOptions EXIOptions = EXIOptions{
	ADD_LESSCOMMON:                     1,
	ADD_UNCOMMON:                       2,
	ADD_ALIGNMENT:                      4,
	ADD_PRESERVE:                       8,
	ADD_COMMON:                         16,
	ADD_VALUE_MAX_LENGTH:               32,
	ADD_VALUE_PARTITION_CAPACITY:       64,
	ADD_FRAGMENT:                       128,
	ADD_DTRM:                           256,
	BLOCKSIZE_DEFAULT:                  1000000,
	VALUE_MAX_LENGTH_UNBOUNDED:         -1,
	VALUE_PARTITION_CAPACITY_UNBOUNDED: -1,
}

func (e *EXIOptions) setSchemaId(var1 SchemaId) { //func (e *EXIOptions) setSchemaId(var1 SchemaId) *EXIOptions {
	e.m_schemaId = var1
	//return e
}

func (e *EXIOptions) setGrammarOptions(var1 int) { //throws EXIOptionsException
	e.m_isStrict = (var1 == 1) //if (this.m_isStrict = var1 == 1) {
	if e.m_isStrict {
		e.m_preserveComments = false
		e.m_preservePIs = false
		e.m_preserveDTD = false
		e.m_preserveNS = false
		e.m_infuseSC = false
	} else {
		e.m_preserveComments = VARGrammarOptions.hasCM(var1) //было без VAR...
		e.m_preservePIs = VARGrammarOptions.hasPI(var1)      //было без VAR...
		e.m_preserveDTD = VARGrammarOptions.hasDTD(var1)     //было без VAR...
		e.m_preserveNS = VARGrammarOptions.hasNS(var1)       //было без VAR...
		e.setInfuseSC(VARGrammarOptions.hasSC(var1))         //было без VAR...
	}
}

func (e *EXIOptions) setInfuseSC(var1 bool /*VARAlignmentType AlignmentType*/) (*EXIOptions, error) { //добавил VARAlignmentType AlignmentType//throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_infuseSC != var1 {
		if var1 {
			if e.m_alignmentType.compress == VARAlignmentType.compress { //опять таже штука с AlignmentType
				err = errors.New("selfContained option and compression option cannot be used together")
				e.m_infuseSC = var2
				return e, err
			}
			if e.m_alignmentType.preCompress == VARAlignmentType.preCompress { //опять таже штука с AlignmentType
				err = errors.New("selfContained option and pre-compression option cannot be used together")
				e.m_infuseSC = var2
				return e, err
			}
			if e.m_isStrict {
				err = errors.New("selfContained option and strict option cannot be used together")
				e.m_infuseSC = var2
				return e, err
			}
		}
		e.m_infuseSC = var1
		return e, err
	}
	err = errors.New("metod setInfuseSC: e.m_infuseSC == var1")
	e.m_infuseSC = var2
	return e, err
}

///QName.class////////////////////////
type QName struct {
	qName         string
	namespaceName string
	localName     string
	prefix        string
	uriId         int
	localNameId   int
}

func (q *QName) QName1() {}
func (q *QName) QName2(var1, var2 string) {
	q.qName = var1
	var var3 int = strings.Index(q.qName, ":")
	if var3 != -1 {
		var var4 int = len([]rune(q.qName)) //add []rune
		var var5 int = var4 - 1
		//label40:
		for var5 > 0 {
			if string(var1[var5]) == string('\t') {
			} else if string(var1[var5]) == string('\n') {
			} else if string(var1[var5]) == string('\r') {
			} else if string(var1[var5]) == string(' ') {
				var5 = (var5 - 1)
				break
			} else {
				break
			}
		}
		q.localName = q.qName[var3 : var5+1]
		var5 = 0

		//label31:
		for var5 < var4 {
			if string(var1[var5]) == string('\t') {
			} else if string(var1[var5]) == string('\n') {
			} else if string(var1[var5]) == string('\r') {
			} else if string(var1[var5]) == string(' ') {
				var5 = (var5 + 1)
				break
			} else {
				break
			}
		}
		q.prefix = q.qName[var5-1 : var3]
		q.namespaceName = var2

	} else {
		q.localName = q.qName
		if var2 != "" {
			q.namespaceName = var2
		} else {
			q.namespaceName = ""
		}
		q.prefix = ""
	}
}

func (q *QName) setValue1(var1, var2, var3 string) *QName {
	var var4 string = ""
	q.qName = var4
	q.namespaceName = var1
	q.localName = var2
	q.prefix = var3
	return q
}

func (q *QName) setValue2(var1, var2, var3, var4 string) *QName { //(*string, *string, *string, *string) {
	q.qName = var4
	q.namespaceName = var1
	q.localName = var2
	q.prefix = var3
	return q //&q.qName, &q.namespaceName, &q.localName, &q.prefix
}

/*нужно подумать нужно ли оно в принципе
public boolean equals(Object var1) {
	if (var1 != null && var1 instanceof QName) {
	   QName var2 = (QName)var1;
	   if (this.namespaceName != null) {
		  return this.namespaceName.equals(var2.namespaceName) && this.localName.equals(var2.localName);
	   }
	}

	return false;
}
*/
func (q *QName) toString() string {
	if q.namespaceName != "" {
		return ("{ " + q.namespaceName + " }" + q.localName)
	}
	return ("")
}
func (q *QName) isSame(var0 []QName, var1 int, var2 []QName, var3 int) bool {
	var var4 int = 2 * var3
	if var0 == nil {
		panic("met isSame QNane struct - var0 is nil")
	}
	if var1 == 0 {
		if var4 == 0 {
			return true
		}
	} else {
		var lenVar0 int = len(var0) //my var
		var lenVar2 int = len(var2) //my var
		var var5 int = 2 * var1

		if var5 == 0 {
			panic("met isSame QNane struct - var5 is 0")
		}
		if var4 != 0 && var4 == var5 {
			var var6 int
			for var6 = 0; var6 < var4; var6++ {
				if lenVar0 == var6 {
					if lenVar2 < var6 {
						break
					}
				} else if (lenVar2 == var6) || (!((var2[var6+1]) == (var0[var6+1]))) {
					break
				}
				/*
					var var7 QName = var0[var6]
					var var8 QName = var2[var6]
				*/ //убрал, ибо не вижу пока смысла от переменных
			}

			if var6 == var4 {
				return true
			}
		}
	}
	return false
}

//////////end QName.class/////////////

///////////AlignmentType.class////////
type AlignmentType struct {
	bitPacked   string //наверное
	byteAligned string //наверное
	preCompress string //наверное
	compress    string //наверное
}

var VARAlignmentType AlignmentType = AlignmentType{"bitPacked", "byteAligned", "preCompress", "compress"}

////////////end AlignmentType.class//////

///////end EXIOptions.class////////////
//////GrammarOptions.class/////////////////////////
type GrammarOptions struct {
	RESTRICT_XSI_NIL_TYPE_MASK int
	ADD_UNDECLARED_EA_MASK     int
	ADD_NS                     int
	ADD_SC                     int
	ADD_DTD                    int
	ADD_CM                     int
	ADD_PI                     int
	OPTIONS_UNUSED             int
	DEFAULT_OPTIONS            int
	STRICT_OPTIONS             int
}

var VARGrammarOptions GrammarOptions = GrammarOptions{
	RESTRICT_XSI_NIL_TYPE_MASK: 1,
	ADD_UNDECLARED_EA_MASK:     2,
	ADD_NS:                     4,
	ADD_SC:                     8,
	ADD_DTD:                    16,
	ADD_CM:                     32,
	ADD_PI:                     64,
	OPTIONS_UNUSED:             0,
	DEFAULT_OPTIONS:            2,
	STRICT_OPTIONS:             1,
}

func (g *GrammarOptions) GrammarOptions() {}

func (g *GrammarOptions) MyMetodIsPermitDeviation(var0 int) bool { //public static boolean isPermitDeviation(short var0) {
	return (var0 & 2) != 0 //return (var0 & 2) != 0;
}

func (g *GrammarOptions) restrictXsiNilType(var0 int, var1 bool) int {
	if var1 {
		return (var0 | 1)
	}
	return (var0 & -2)
}

func (g *GrammarOptions) isXsiNilTypeRestricted(var0 int) bool {
	return ((var0 & 1) != 0)
}
func isXsiNilTypeRestricted(var0 int) bool {
	return ((var0 & 1) != 0)
}

func isPermitDeviation(var0 int) bool {
	return ((var0 & 2) != 0)
}

func (g *GrammarOptions) hasNS(var0 int) bool {
	return ((var0 & 4) != 0)
}
func hasNS(var0 int) bool {
	return ((var0 & 4) != 0)
}

func (g *GrammarOptions) hasSC(var0 int) bool {
	return ((var0 & 8) != 0)
}
func hasSC(var0 int) bool {
	return ((var0 & 8) != 0)
}

func (g *GrammarOptions) hasDTD(var0 int) bool {
	return ((var0 & 16) != 0)
}

func hasDTD(var0 int) bool {
	return ((var0 & 16) != 0)
}

func (g *GrammarOptions) hasCM(var0 int) bool {
	return ((var0 & 32) != 0)
}
func hasCM(var0 int) bool {
	return ((var0 & 32) != 0)
}

func (g *GrammarOptions) hasPI(var0 int) bool {
	return ((var0 & 64) != 0)
}
func hasPI(var0 int) bool {
	return ((var0 & 64) != 0)
}

func (g *GrammarOptions) addNS(var0 int) int {
	return (var0 | 4)
}

func (g *GrammarOptions) addSC(var0 int) int {
	return (var0 | 8)
}

func (g *GrammarOptions) addDTD(var0 int) int {
	return (var0 | 16)
}

func (g *GrammarOptions) addCM(var0 int) int {
	return (var0 | 32)
}

func (g *GrammarOptions) addPI(var0 int) int {
	return (var0 | 64)
}

//////end GrammarOptions.class/////////////////////
///EventTypeFactory.class///

func creatEndElement(var0 byte, var1 EventTypeList) EventType { //static EventType creatEndElement(byte var0, EventTypeList var1) {
	return EventType3("", "", -1, -1, var0, var1, byte(9), 7 /*, (IGrammar)null*/) //return new EventType((String)null, (String)null, -1, -1, var0, var1, (byte)9, (byte)7, (IGrammar)null);
}

func createEndDocument(var0 EventTypeList) EventType { //static EventType createEndDocument(EventTypeList var0) {
	return EventType3("", "", -1, -1, byte(1), var0, byte(4), 1 /*, (IGrammar)null*/) //	return new EventType((String)null, (String)null, -1, -1, (byte)1, var0, (byte)4, (byte)1, (IGrammar)null);
}

func createStartDocument(var0 EventTypeList) EventType { // static EventType createStartDocument(EventTypeList var0) {
	return EventType3("", "", -1, -1, byte(1), var0, byte(13), 0 /*, (IGrammar)null*/) //	return new EventType((String)null, (String)null, -1, -1, (byte)1, var0, (byte)13, (byte)0, (IGrammar)null);
}

func createStartElement(var0 int, var1 int, var2 string, var3 string, var4 EventTypeList, var5 EXIGrammarUse) EventTypeElement { //static EventTypeElement createStartElement(int var0, int var1, String var2, String var3, EventTypeList var4, EXIGrammarUse var5) {
	var myvar EventType = VAREventTypeElement.EventTypeElement1(var0, var2, var1, var3, var4, var5 /*, (EXIGrammar)null*/) //return new EventTypeElement(var0, var2, var1, var3, var4, var5, (EXIGrammar)null);
	return *(*EventTypeElement)(unsafe.Pointer(&myvar))
}

///end EventTypeFactory.class///

///////////////////////////////////////end Transmogrifier.class///////////////////////////////////////////////////////////////////////////

func main() {
	infile, err := os.Open("build.txt")
	//var xml1 string
	if err != nil {
		err = errors.New("reading file error")
		fmt.Println(err)
		return
	}
	outfile, err := os.Create("exiBuild.txt")
	if err != nil {
		err = errors.New("out file error")
		fmt.Println(err)
		return
	}
	defer infile.Close()
	defer outfile.Close()

	var workgrammarCache GrammarCache

}
