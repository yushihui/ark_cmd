package ark

var ArkFunds = map[string]string{
	"ARKK": "ARK INNOVATION ~ 20+ B",
	"ARKQ": "ARK AUTONOMOUS TECHNOLOGY & ROBOTICS ~ 2.7 B",
	"ARKW": "ARK NEXT GENERATION INTERNET ~ 6.5 B",
	"ARKG": "ARK GENOMIC REVOLUTION MULTISECTOR ~ 11+ B",
	"ARKF": "ARK FINTECH INNOVATION ~ 2.6 B",
	"PRNT": "THE 3D RINTING ~ 0.26 B",
	"IZRL": "ARK ISRAEL INNOVATIVE TECHNOLOGY ~ 0.15 B",
}

// Security record
type Security struct {
	Fund       string
	Name       string
	TickerCode string
	Shares     float64
	Delta      float64
	Price      float64
	IsNew      bool
	Weight     float64
}

// ByValue sorter
type ByValue []*Security

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Delta*a[i].Price > a[j].Delta*a[j].Price }
