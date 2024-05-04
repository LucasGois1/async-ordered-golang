package flavor

type Flavor int

const (
	Vanilla Flavor = iota + 2
	Chocolate
	Strawberry
	Mint
	Yogurt
	Grape
	Lemon
	Pistachio
)

func Name(f Flavor) string {
	switch f {
	case Vanilla:
		return "Vanilla"
	case Chocolate:
		return "Chocolate"
	case Strawberry:
		return "Strawberry"
	case Mint:
		return "Mint"
	case Yogurt:
		return "Yogurt"
	case Grape:
		return "Grape"
	case Lemon:
		return "Lemon"
	case Pistachio:
		return "Pistachio"
	default:
		return "Unknown"
	}
}
