package common


func GetLocalIp()(string,error){
	netFace,err := net.Interfaces()
	if err != nil{
		slog.Warn("Get interfaces error:", err.Error())
		return nil, err 
	}
	
}