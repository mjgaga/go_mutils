package util

type CommonChanStruct struct {
	Data  interface{}
	Error error
}

type CommonChan chan *CommonChanStruct

func (this CommonChan) Next() (interface{}, error) {
	ccs := <-this
	return ccs.Data, ccs.Error
}

func (this CommonChan) Push(data interface{}, err error) {
	this <- NewCommonChanStruct(data, err)
}

func (this CommonChan) Close() {
	close(this)
}

func MakeCommonChan(len int) CommonChan {
	return make(CommonChan, len)
}
func MakeACommonChan() CommonChan {
	return MakeCommonChan(1)
}
func NewCommonChanStruct(data interface{}, err error) *CommonChanStruct {
	return &CommonChanStruct{Data: data, Error: err}
}
