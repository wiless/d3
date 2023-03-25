package d3

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"

	"github.com/jszwec/csvutil"
)

var er = func(err error) {
	if err != nil {
		log.Println("D3:Error ", err)
	}
}

func Range(start, stop float64, step float64) []float64 {
	var size = int(math.Floor((stop - start) / step))
	res := make([]float64, size)
	for i := range res {
		res[i] = start + step*float64(i)
	}
	return res

}

func Map(array interface{}, fn interface{}) interface{} {
	// log.Printf("\n Map()..")
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)

	// hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function
	// resultValue := reflect.New(tOfArray)
	fnVal := reflect.ValueOf(fn)

	// resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	// fmt.Printf("\n INPUT = arg1 = %v and arg2 =  %v  ", tOfArray, tOffn)
	// fmt.Printf("\n Kind arg1 %s ", tOfArray.Kind())
	// fmt.Printf("\n Kind arg2  %s ", tOffn.Kind())

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray && tOffn.Kind() == reflect.Func) {
		log.Println("Mismatching Arguments..")
		return nil
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()
	var fnOutputType reflect.Type

	// fmt.Printf("\n ARRAY of type %s ", elemType)
	// fmt.Printf("\n Value / Handle of Function   %v ", fnVal)
	// fmt.Printf("\n Fn : Input Args:%d =", tOffn.NumIn())
	// for i := 0; i < tOffn.NumIn(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.In(i))
	// }
	// fmt.Printf("\n Fn : Output Args:%d =", tOffn.NumOut())
	// for i := 0; i < tOffn.NumOut(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.Out(i))
	// }

	if tOffn.NumIn() == 0 || tOffn.NumOut() == 0 {
		fmt.Println("MAP needs Fn with 1 input and 1 output arg \nYou might want to use d3.ForEach() instead.")
		return nil
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 && tOffn.In(0) == elemType {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int && tOffn.In(1) == elemType {
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return nil
	}
	fnOutputType = tOffn.Out(0)

	arrayvalues := reflect.ValueOf(array)
	toFnResults := reflect.SliceOf(fnOutputType)
	// fmt.Printf("\nMap(): Result Type %v", toFnResults)

	resultValue := reflect.MakeSlice(toFnResults, 0, arrayvalues.Cap())

	// }

	// avalue := reflect.ValueOf(array)
	//
	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		var resultObj []reflect.Value
		if fnType == 1 {
			resultObj = fnVal.Call([]reflect.Value{elemValue})
		} else {
			var indx = reflect.ValueOf(i)
			resultObj = fnVal.Call([]reflect.Value{indx, elemValue})
		}

		// 	for _, val := range response {
		// 		fmt.Println("Return of fn is of Type ", val.Type(), val.Kind())
		// 		// if found {
		resultValue = reflect.Append(resultValue, resultObj[0])
		// 		// 	cnt++

		// 		// }

		// 	}

	}

	// // log.Printf("Total Found %d out of %d", cnt, avalue.Len())
	// // fmt.Printf("\n Result %#v ", resultValue)
	// // fmt.Printf("Result is %#v", resultValue)
	return resultValue.Interface()
}

func ForEach(array interface{}, fn interface{}) {
	// log.Printf("\n ForEach()..")
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)

	// hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function
	// resultValue := reflect.New(tOfArray)
	fnVal := reflect.ValueOf(fn)

	// resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	// fmt.Printf("\n INPUT = arg1 = %v and arg2 =  %v  ", tOfArray, tOffn)
	// fmt.Printf("\n Kind arg1 %s ", tOfArray.Kind())
	// fmt.Printf("\n Kind arg2  %s ", tOffn.Kind())

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray && tOffn.Kind() == reflect.Func) {
		log.Println("Mismatching Arguments..")
		return
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()

	// fmt.Printf("\n ARRAY of type %s ", elemType)
	// fmt.Printf("\n Value / Handle of Function   %v ", fnVal)
	// fmt.Printf("\n Fn : Input Args:%d =", tOffn.NumIn())
	// for i := 0; i < tOffn.NumIn(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.In(i))
	// }
	// fmt.Printf("\n Fn : Output Args:%d =", tOffn.NumOut())
	// for i := 0; i < tOffn.NumOut(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.Out(i))
	// }

	if tOffn.NumIn() == 0 {
		fmt.Println("ForEach needs Fn with 1 or 2 input args")
		return
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 && tOffn.In(0) == elemType {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int && tOffn.In(1) == elemType {
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return
	}

	arrayvalues := reflect.ValueOf(array)

	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		if fnType == 1 {
			fnVal.Call([]reflect.Value{elemValue})
		} else {
			var indx = reflect.ValueOf(i)
			fnVal.Call([]reflect.Value{indx, elemValue})
		}
	}

}

// func Sort(array interface{}, fn interface{}) interface{} {

// }

func Filter(array interface{}, fn interface{}) interface{} {
	// log.Printf("\n Map()..")
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)

	// hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function
	// resultValue := reflect.New(tOfArray)
	fnVal := reflect.ValueOf(fn)

	// resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	// fmt.Printf("\n INPUT = arg1 = %v and arg2 =  %v  ", tOfArray, tOffn)
	// fmt.Printf("\n Kind arg1 %s ", tOfArray.Kind())
	// fmt.Printf("\n Kind arg2  %s ", tOffn.Kind())

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray && tOffn.Kind() == reflect.Func) {
		log.Println("Mismatching Arguments..")
		return nil
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()
	var fnOutputType reflect.Type

	// fmt.Printf("\n ARRAY of type %s ", elemType)
	// fmt.Printf("\n Value / Handle of Function   %v ", fnVal)
	// fmt.Printf("\n Fn : Input Args:%d =", tOffn.NumIn())
	// for i := 0; i < tOffn.NumIn(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.In(i))
	// }
	// fmt.Printf("\n Fn : Output Args:%d =", tOffn.NumOut())
	// for i := 0; i < tOffn.NumOut(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.Out(i))
	// }

	if tOffn.NumIn() == 0 || tOffn.NumOut() == 0 {
		fmt.Println("Filter needs Fn with 1 input and 1 output arg \nYou might want to use d3.ForEach() instead.")
		return nil
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 && tOffn.In(0) == elemType {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int && tOffn.In(1) == elemType {
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return nil
	}
	fnOutputType = tOffn.Out(0)

	// Check that FnOutput to be bool !!
	if fnOutputType.Kind() != reflect.Bool {
		fmt.Printf("\nError : Filter(): Function must return Bool!", fnOutputType)
		return nil
	}

	arrayvalues := reflect.ValueOf(array)
	resultValue := reflect.MakeSlice(tOfArray, 0, arrayvalues.Cap())
	// fmt.Printf("\nFilter(): Result Type %#v", resultValue)
	// }

	// avalue := reflect.ValueOf(array)
	//
	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		var resultObj []reflect.Value
		if fnType == 1 {
			resultObj = fnVal.Call([]reflect.Value{elemValue})
		} else {
			var indx = reflect.ValueOf(i)
			resultObj = fnVal.Call([]reflect.Value{indx, elemValue})
		}

		found := resultObj[0].Bool()
		if found {
			resultValue = reflect.Append(resultValue, elemValue)
		}

	}

	return resultValue.Interface()
}

func FilterIndex(array interface{}, fn interface{}) []int {
	log.Printf("\n FilterIndex..")
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)

	// hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function
	// resultValue := reflect.New(tOfArray)
	fnVal := reflect.ValueOf(fn)

	// resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	// fmt.Printf("\n INPUT = arg1 = %v and arg2 =  %v  ", tOfArray, tOffn)
	// fmt.Printf("\n Kind arg1 %s ", tOfArray.Kind())
	// fmt.Printf("\n Kind arg2  %s ", tOffn.Kind())

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray && tOffn.Kind() == reflect.Func) {
		log.Println("Mismatching Arguments..")
		return nil
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()
	var fnOutputType reflect.Type

	// fmt.Printf("\n ARRAY of type %s ", elemType)
	// fmt.Printf("\n Value / Handle of Function   %v ", fnVal)
	// fmt.Printf("\n Fn : Input Args:%d =", tOffn.NumIn())
	// for i := 0; i < tOffn.NumIn(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.In(i))
	// }
	// fmt.Printf("\n Fn : Output Args:%d =", tOffn.NumOut())
	// for i := 0; i < tOffn.NumOut(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.Out(i))
	// }

	if tOffn.NumIn() == 0 || tOffn.NumOut() == 0 {
		fmt.Println("FilterIndex needs Fn with 1 input and 1 output arg \nYou might want to use d3.ForEach() instead.")
		return nil
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 && tOffn.In(0) == elemType {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int && tOffn.In(1) == elemType {
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return nil
	}
	fnOutputType = tOffn.Out(0)

	// Check that FnOutput to be bool !!
	if fnOutputType.Kind() != reflect.Bool {
		fmt.Printf("\nError : Filter(): Function must return Bool!", fnOutputType)
		return nil
	}

	arrayvalues := reflect.ValueOf(array)
	var resultValue []int

	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		var resultObj []reflect.Value
		if fnType == 1 {
			resultObj = fnVal.Call([]reflect.Value{elemValue})
		} else {
			var indx = reflect.ValueOf(i)
			resultObj = fnVal.Call([]reflect.Value{indx, elemValue})
		}

		found := resultObj[0].Bool()
		if found {
			resultValue = append(resultValue, i)
		}

	}

	return resultValue
}

func FlatMap(array interface{}, fieldname string) interface{} {
	// log.Printf("\n FlatMap()..")
	tOfArray := reflect.TypeOf(array)

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray) {
		log.Println("Mismatching Arguments..")
		return nil
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()
	field, ok := elemType.FieldByName(fieldname)
	if !ok {
		fmt.Printf("Element not Type of struct | %v, field %v", elemType.Kind(), field)
		return nil
	}

	arrayvalues := reflect.ValueOf(array)

	resultValue := reflect.MakeSlice(reflect.SliceOf(field.Type), 0, arrayvalues.Cap())
	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		val := elemValue.FieldByName(fieldname)
		resultValue = reflect.Append(resultValue, val)
		// 	if fnType == 1 {
		// 		fnVal.Call([]reflect.Value{elemValue})
		// 	} else {
		// 		var indx = reflect.ValueOf(i)
		// 		fnVal.Call([]reflect.Value{indx, elemValue})
		// 	}
	}
	return resultValue.Interface()
}

func FindFirstIndex(array interface{}, fn interface{}) int {
	var index int = -1
	// log.Printf(" FindFirstIndex()..")
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)

	fnVal := reflect.ValueOf(fn)

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray && tOffn.Kind() == reflect.Func) {
		log.Println("Mismatching Arguments..")
		return -1
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()
	var fnOutputType reflect.Type

	if tOffn.NumIn() == 0 || tOffn.NumOut() == 0 {
		fmt.Println("FindFirstIndex needs Fn with 1 input and 1 output arg \nYou might want to use d3.ForEach() instead.")
		return -1
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 && tOffn.In(0) == elemType {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int && tOffn.In(1) == elemType {
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return -1
	}
	fnOutputType = tOffn.Out(0)

	// Check that FnOutput to be bool !!
	if fnOutputType.Kind() != reflect.Bool {
		fmt.Printf("\nError : Filter(): Function must return Bool!", fnOutputType)
		return -1
	}

	arrayvalues := reflect.ValueOf(array)

	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		var resultObj []reflect.Value
		if fnType == 1 {
			resultObj = fnVal.Call([]reflect.Value{elemValue})
		} else {
			var indx = reflect.ValueOf(i)
			resultObj = fnVal.Call([]reflect.Value{indx, elemValue})
		}
		found := resultObj[0].Bool()
		if found {
			index = i
			break
		}

	}

	return index
}

func FindFirst(array interface{}, fn interface{}) interface{} {
	// log.Printf("\n FindFirst()..")
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)

	fnVal := reflect.ValueOf(fn)

	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	if !(isArray && tOffn.Kind() == reflect.Func) {
		log.Println("Mismatching Arguments..")
		return nil
	}
	var elemType reflect.Type
	elemType = tOfArray.Elem()
	var fnOutputType reflect.Type

	if tOffn.NumIn() == 0 || tOffn.NumOut() == 0 {
		fmt.Println("FindFirst needs Fn with 1 input and 1 output arg \nYou might want to use d3.ForEach() instead.")
		return nil
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 && tOffn.In(0) == elemType {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int && tOffn.In(1) == elemType {
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return nil
	}
	fnOutputType = tOffn.Out(0)

	// Check that FnOutput to be bool !!
	if fnOutputType.Kind() != reflect.Bool {
		fmt.Printf("\nError : Filter(): Function must return Bool!", fnOutputType)
		return nil
	}

	arrayvalues := reflect.ValueOf(array)
	var resultValue reflect.Value
	resultValue = reflect.New(elemType)

	var found bool
	for i := 0; i < arrayvalues.Len(); i++ {

		elemValue := arrayvalues.Index(i)
		var resultObj []reflect.Value
		if fnType == 1 {
			resultObj = fnVal.Call([]reflect.Value{elemValue})
		} else {
			var indx = reflect.ValueOf(i)
			resultObj = fnVal.Call([]reflect.Value{indx, elemValue})
		}
		found = resultObj[0].Bool()
		if found {
			resultValue = elemValue
			break
		}

	}
	if !found {
		return resultValue.Elem().Interface()

	}
	return resultValue.Interface()
}

func ForEachParse(fname string, fn interface{}) {

	// log.Printf("\n ForEach()..")
	tOffn := reflect.TypeOf(fn)

	fnVal := reflect.ValueOf(fn)

	// resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	// fmt.Printf("\n INPUT = arg1 = %v and arg2 =  %v  ", tOfArray, tOffn)
	// fmt.Printf("\n Kind arg1 %s ", tOfArray.Kind())
	// fmt.Printf("\n Kind arg2  %s ", tOffn.Kind())

	// fmt.Printf("\n ARRAY of type %s ", elemType)
	// fmt.Printf("\n Value / Handle of Function   %v ", fnVal)
	// fmt.Printf("\n Fn : Input Args:%d =", tOffn.NumIn())
	// for i := 0; i < tOffn.NumIn(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.In(i))
	// }
	// fmt.Printf("\n Fn : Output Args:%d =", tOffn.NumOut())
	// for i := 0; i < tOffn.NumOut(); i++ {
	// 	fmt.Printf("\t %d=>%v,", i, tOffn.Out(i))
	// }

	if tOffn.NumIn() == 0 {
		fmt.Println("ForEach needs Fn with 1 or 2 input args")
		return
	}

	var fnType int = 1
	/// Function Argument must match Element type of the Array {
	if tOffn.NumIn() == 1 {
		fnType = 1
		// fmt.Printf("\nFunction TYPE 1 : (elem) : Element Type %v matches with Fn arg1 %v", elemType, tOffn.In(0), tOffn.In(1))
	} else if tOffn.NumIn() == 2 && tOffn.In(0).Kind() == reflect.Int {
		// Expect second argument of type "struct"
		fnType = 2
		// fmt.Printf("\nFunction TYPE 2 : (ind,elem)  : Element Type %v matches with Fn arg2 %v, arg 1=%v", elemType, tOffn.In(1), tOffn.In(0))
	} else {
		// fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
		return
	}

	var elemType reflect.Type
	elemType = tOffn.In(0)

	fd, err := os.Open(fname)
	er(err)
	csvReader := csv.NewReader(fd)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	header := dec.Header()
	_ = header
	var i int = 0
	// _ = fnVal
	elemValue := reflect.New(elemType)
	// fmt.Printf("\nType of Fn Arg is %v", elemType)
	// fmt.Printf("\nType of New Variable is %v | kind =%v", elemValue.Type(), elemType.Kind())

	for {

		// u := User{OtherData: make(map[string]string)}
		// element := reflect.New()

		u := elemValue.Interface()

		if err := dec.Decode(u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("\nRead File : %#v", u)
		obj := elemValue.Elem()
		// fmt.Printf("\n OBJ %#v", obj)
		if fnType == 1 {
			fnVal.Call([]reflect.Value{obj})
		} else {
			var indx = reflect.ValueOf(i)
			fnVal.Call([]reflect.Value{indx, obj})
			i++
		}

	}

}

// SubStruct creates array of objs with selected properties "fields" from the input array of objects
func SubStruct(v interface{}, fields ...string) interface{} {
	// fmt.Printf("\n Input : %#v", v)
	tOfv := reflect.TypeOf(v)
	var subfields []reflect.StructField
	var fnames []string
	for _, f := range fields {
		ftype, ok := tOfv.FieldByName(f)
		if ok {
			subfields = append(subfields, ftype)
			fnames = append(fnames, f)
		}
	}
	resultType := reflect.StructOf(subfields)
	elemVal := reflect.ValueOf(v)
	result := reflect.New(resultType)

	for _, f := range fnames {
		inpval := elemVal.FieldByName(f)
		// fmt.Printf("\n\nField  %v is %v ", f, inpval)
		newfield := result.Elem().FieldByName(f)
		// fmt.Printf("\nBefore Setting  %v is %#v ", f, newfield)
		if newfield.CanSet() {
			newfield.Set(inpval)
			// fmt.Printf("\nSetting  %v is %#v ", f, newfield)
		}

	}

	retobj := result.Elem()

	// fmt.Printf("\n Created : %#v", retobj)

	return retobj.Interface()
	// for i := 0; i < N; i++ {
	// 	tOfv.FieldByName()
	// }
}

func CSV(fname string, v interface{}) interface{} {

	fid, err := os.Open(fname)
	er(err)

	data, err := os.ReadFile(fname)
	er(err)

	err = csvutil.Unmarshal(data, v)
	er(err)

	defer fid.Close()

	fid.Close()
	return v
}

func CSVsave(fname string, v interface{}) {
	fmt.Println(reflect.TypeOf(v))
	fid, err := os.Create(fname)
	er(err)
	defer fid.Close()
	ba, _ := csvutil.Marshal(v)
	fid.Write(ba)
}

// func CSVsave(fname string, v interface{}) {

// 	fid, err := os.Create(fname)
// 	er(err)
// 	defer fid.Close()

// 	// str, _ := vlib.Struct2HeaderLine(v)
// 	// fid.WriteString("\n" + str)

// 	for i := 0; i < count; i++ {

// 	}, func(d interface{}) {
// 		fmt.Println("\nd=", d)
// 		str, _ := vlib.Struct2String(d)
// 		fid.WriteString("\n" + str)
// 	})

// }
