package async

import "reflect"

// const GOROUTE_COUTE int = 8

type OnProcessItem func(initItem interface{}) interface{}

// type ParallelTask struct {
// 	intChan           chan interface{}
// 	resultChan        chan interface{}
// 	goRouteFinishChan chan bool
// 	batchFinish       chan bool
// 	result            []interface{}
// }

func ConvertToInterfaceList(i interface{}) (initData []interface{}) {
	v := reflect.ValueOf(i)
	// There is no need to check, we want to panic if it's not slice or array
	intf := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		intf[i] = v.Index(i).Interface()
	}
	return intf
}

func StartParallelList(initData []interface{}, onProcessItem OnProcessItem, gorouteCount ...int) *[]interface{} {

	var GOROUTE_COUTE int = 8
	if len(gorouteCount) > 0 && gorouteCount[0] > 0 {
		GOROUTE_COUTE = gorouteCount[0]
	}

	intChan := make(chan interface{})
	resultChan := make(chan interface{}, GOROUTE_COUTE)
	result := make([]interface{}, 0)
	goRouteFinishChan := make(chan bool, GOROUTE_COUTE)
	batchFinish := make(chan bool, 1)

	go pushInitDataToChan(intChan, initData)

	for i := 0; i < GOROUTE_COUTE; i++ {
		go doCalc(intChan, resultChan, goRouteFinishChan, onProcessItem)
	}

	go checkGoRouteProcess(goRouteFinishChan, resultChan)

	go fatchResultFromChan(resultChan, &result, batchFinish)

	<-batchFinish

	return &result
}

func pushInitDataToChan(intChan chan interface{}, initData []interface{}) {
	for _, initItem := range initData {
		intChan <- initItem
	}
	close(intChan)
}

func doCalc(intChan chan interface{}, resultChan chan interface{}, finishChan chan bool, onProcessItem OnProcessItem) {
	for initItem := range intChan {
		itemResult := onProcessItem(initItem)
		resultChan <- itemResult
	}
	finishChan <- true
}

func checkGoRouteProcess(goRouteFinishChan chan bool, resultChan chan interface{}) {
	GOROUTE_COUTE := cap(goRouteFinishChan)
	for i := 0; i < GOROUTE_COUTE; i++ {
		<-goRouteFinishChan
	}
	close(resultChan)
	close(goRouteFinishChan)

	//logrus.Println("all goroutes finished")
}

func fatchResultFromChan(resultChan chan interface{}, returnResult *[]interface{}, flag chan bool) {
	for result := range resultChan {
		*returnResult = append(*returnResult, result)
	}
	flag <- true
}
