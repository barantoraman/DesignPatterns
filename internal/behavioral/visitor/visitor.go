package visitor

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForrectangle(*Rectangle)
}
