package d3

import (
	"fmt"
	"log"
	"reflect"
)

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
	fmt.Printf("\nMap(): Result Type %v", toFnResults)

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
	log.Printf("\n ForEach()..")
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
		// fmt.Printf("\nError : Array Element %v DOES NOT MATCH with Fn arg %v", elemType, tOffn.In(0))
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
	log.Printf("\n FlatMap()..")
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
	log.Printf(" FindFirstIndex()..")
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
	log.Printf("\n FindFirst()..")
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
		return nil
	}
	return resultValue.Interface()
}
