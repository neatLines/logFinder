package models

var (
	Hosts map[string]*Set
	empty Empty
)

func init() {
	Hosts = make(map[string]*Set)
	// Hosts["hjkhsbnmn123"] = &Set{map[string]Empty{}}
	// Hosts["mjjkxsxsaa23"] = &Set{map[string]Empty{}}
	// Hosts["hjkhsbnmn123"].S["aa"] = empty
	// Hosts["mjjkxsxsaa23"].S["bb"] = empty
	// Hosts["mjjkxsxsaa23"].S["cc"] = empty

}

//
func AddOne(host Host) (success bool) {
	if Hosts[host.Appname] != nil {
		Hosts[host.Appname].S[host.Hostname] = empty
	} else {
		Hosts[host.Appname] = &Set{map[string]Empty{}}
		Hosts[host.Appname].S[host.Hostname] = empty
	}
	return true
}

//
func GetAll() map[string]*Set {
	return Hosts
}
