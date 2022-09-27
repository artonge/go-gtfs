package gtfs

// Const used in GTFS
const (
	RouteTypeTram      = 0
	RouteTypeSubway    = 1
	RouteTypeRail      = 2
	RouteTypeBus       = 3
	RouteTypeFerry     = 4
	RouteTypeCableCar  = 5
	RouteTypeGondola   = 6
	RouteTypeFunicular = 7

	ExceptionTypeAdded   = 1
	ExceptionTypeRemoved = 2

	PaymentMethodOnBoard        = 0
	PaymentMethodBeforeBoarding = 1

	TransfersNotPermitted = 0
	TransfersAtMostOnce   = 1
	TransfersAtMostTwice  = 2

	AttributionHasNoRole = 0
	AttributionHasRole   = 1

	PathwayWalkway        = 1
	PathwayStairs         = 2
	PathwayMovingSidewalk = 3
	PathwayEscalator      = 4
	PathwayElevator       = 5
	PathwayFareGate       = 6
	PathwayExitGate       = 7

	PathwayUnidirectional = 0
	PathwayBidirectional  = 1
)
