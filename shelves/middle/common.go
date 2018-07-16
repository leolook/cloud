package middle

import (
	pb "cloud/shelves/protocol"
	"reflect"
)

func GetCommonReq(req *reflect.Value) {
	common := pb.CommonReq{
		UserId: 2018,
	}

	if req.Elem().CanSet() {
		req.Elem().Set(reflect.ValueOf(common))
	}
}
