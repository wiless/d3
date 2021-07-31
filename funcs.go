package d3

import (
	"fmt"
	"log"
	"reflect"
)

type FilterFn func(interface{}) bool
type MapFn func(interface{}) interface{}
type ForEachFn func(interface{})

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
	fnOutputType = tOffn.Out(0)

	arrayvalues := reflect.ValueOf(array)
	toFnResults := reflect.SliceOf(fnOutputType)
	fmt.Printf("\nForEach(): Result Type %v", toFnResults)

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

func ForEach2(array interface{}, fn ForEachFn) {

	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)
	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function
	// resultValue := reflect.New(tOfArray)
	avalue := reflect.ValueOf(array)
	fnVal := reflect.ValueOf(fn)
	if !isArray {
		log.Println("Filter : Either arg1 is not an Array or arg of FilterFn mismatches")
		return
	}

	//resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	fmt.Printf("\n Processing Array of %v and Function %#v ", tOfArray, hasFunction)
	fmt.Printf("\n TypeOf array is  %s ", tOfArray)

	fmt.Printf("\n Definition of Function   %v ", fnVal)
	fmt.Printf("\n Input Args : Num : %d  ", tOffn.NumIn())
	for i := 0; i < tOffn.NumIn(); i++ {
		fmt.Printf("\n \t Args Types : %v  ", tOffn.In(i))
	}
	fmt.Printf("\n Output Args : Num : %d  ", tOffn.NumOut())
	for i := 0; i < tOffn.NumOut(); i++ {
		fmt.Printf("\n \t Response Types : %v  ", tOffn.Out(i))
	}

	for i := 0; i < avalue.Len(); i++ {

		elemValue := avalue.Index(i)
		fnVal.Call([]reflect.Value{elemValue})
		// for _, val := range response {
		// 	found = val.Bool()
		// 	if found {
		// 		resultValue = reflect.Append(resultValue, elemValue)
		// 		cnt++

		// 	}

		// }

	}

}

func Filter(array interface{}, fn FilterFn) interface{} {
	// if len(array) == 0 {
	// 	return false, -1
	// }
	log.Printf("\n Filter()..")
	var found bool
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)
	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function

	// resultValue := reflect.New(tOfArray)
	avalue := reflect.ValueOf(array)
	fnVal := reflect.ValueOf(fn)
	cnt := 0
	if !isArray {
		log.Println("Filter : Either arg1 is not an Array or arg of FilterFn mismatches")
		return nil
	}

	resultValue := reflect.MakeSlice(tOfArray, 0, avalue.Cap())

	fmt.Printf("\n Processing Array of %v and Function %#v ", tOfArray, hasFunction)
	fmt.Printf("\n TypeOf array is  %s ", tOfArray)

	for i := 0; i < avalue.Len(); i++ {

		elemValue := avalue.Index(i)
		response := fnVal.Call([]reflect.Value{elemValue})
		for _, val := range response {
			found = val.Bool()
			if found {
				resultValue = reflect.Append(resultValue, elemValue)
				cnt++

			}

		}

	}

	// log.Printf("Total Found %d out of %d", cnt, avalue.Len())
	// fmt.Printf("\n Result %#v ", resultValue)
	// fmt.Printf("Result is %#v", resultValue)
	return resultValue.Interface()
}

func FilterIndex(array interface{}, fn FilterFn) []int {
	// if len(array) == 0 {
	// 	return false, -1
	// }
	var found bool
	tOfArray := reflect.TypeOf(array)
	tOffn := reflect.TypeOf(fn)
	isArray := tOfArray.Kind() == reflect.Slice || tOfArray.Kind() == reflect.Array
	hasFunction := tOffn.Kind()
	// tOArg := // Find Argument type of function

	// resultValue := reflect.New(tOfArray)
	avalue := reflect.ValueOf(array)
	fnVal := reflect.ValueOf(fn)
	cnt := 0
	if !isArray {
		log.Println("Filter : Either arg1 is not an Array or arg of FilterFn mismatches")
		return nil
	}

	resultValue := make([]int, 0)

	fmt.Printf("\n Processing Array of %v and Function %#v ", tOfArray, hasFunction)
	fmt.Printf("\n TypeOf array is  %s ", tOfArray)

	for i := 0; i < avalue.Len(); i++ {

		elemValue := avalue.Index(i)
		response := fnVal.Call([]reflect.Value{elemValue})
		for _, val := range response {
			found = val.Bool()
			if found {
				resultValue = append(resultValue, i)
				cnt++
			}

		}

	}

	// log.Printf("Total Found %d out of %d", cnt, avalue.Len())
	// fmt.Printf("\n Result %#v ", resultValue)
	// fmt.Printf("Result is %#v", resultValue)
	return resultValue
}
