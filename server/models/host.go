package models

var (
	Hosts map[string]*Set
	empty Empty
)

func init() {
	Hosts = make(map[string]*Set)
	Hosts["hjkhsbnmn123"] = &Set{map[string]Empty{}}
	Hosts["mjjkxsxsaa23"] = &Set{map[string]Empty{}}
	Hosts["hjkhsbnmn123"].S["aa"] = empty

}

//
func AddOne(appname string, hostname string) (success bool) {
	if Hosts[appname] != nil {
		Hosts[appname].S[hostname] = empty
	} else {
		Hosts[appname] = &Set{map[string]Empty{}}
		Hosts[appname].S[hostname] = empty
	}
	return true
}

//
func GetAll() map[string]*Set {
	return Hosts
}
