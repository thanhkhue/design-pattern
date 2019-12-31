package strategy

type RouteStrategy interface {
	BuildRoute(string, string) string
}

type Route struct {
	RouteStrategy RouteStrategy
}

func (r *Route) Build(src, dest string) string {
	return r.RouteStrategy.BuildRoute(src, dest)
}

type RoadStrategy struct{}

func (r *RoadStrategy) BuildRoute(src, dest string) string {
	return "road:" + src + "->" + dest
}

type RiverStrategy struct{}

func (r *RiverStrategy) BuildRoute(src, dest string) string {
	return "river:" + src + "->" + dest
}

type AirStrategy struct{}

func (r *AirStrategy) BuildRoute(src, dest string) string {
	return "air:" + src + "->" + dest
}

func init(){
	
}