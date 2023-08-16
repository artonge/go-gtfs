package gtfs

// GTFS -
type GTFS struct {
	Path           string // The path to the containing directory
	Agency         Agency
	Agencies       []Agency
	Attributions   []Attribution
	Calendars      []Calendar
	CalendarDates  []CalendarDate
	FareAttributes []FareAttribute
	FareRules      []FareRule
	FeedInfos      []FeedInfo
	Frequencies    []Frequency
	Levels         []Level
	Routes         []Route
	Pathways       []Pathway
	Shapes         []Shape
	Stops          []Stop
	StopsTimes     []StopTime
	Trips          []Trip
	Transfers      []Transfer
}

// Route -
type Route struct {
	ID        string `csv:"route_id"`
	AgencyID  string `csv:"agency_id"`
	ShortName string `csv:"route_short_name"`
	LongName  string `csv:"route_long_name"`
	Type      int    `csv:"route_type"`
	Desc      string `csv:"route_desc"`
	URL       string `csv:"route_url"`
	Color     string `csv:"route_color"`
	TextColor string `csv:"route_text_color"`
}

// Trip -
type Trip struct {
	ID          string `csv:"trip_id"`
	Name        string `csv:"trip_short_name"`
	RouteID     string `csv:"route_id"`
	ServiceID   string `csv:"service_id"`
	ShapeID     string `csv:"shape_id"`
	DirectionID string `csv:"direction_id"`
	Headsign    string `csv:"trip_headsign"`
}

// Stop -
type Stop struct {
	ID          string  `csv:"stop_id"`
	Code        string  `csv:"stop_code"`
	Name        string  `csv:"stop_name"`
	Description string  `csv:"stop_desc"`
	Latitude    float64 `csv:"stop_lat"`
	Longitude   float64 `csv:"stop_lon"`
	Type        string  `csv:"location_type"`
	Parent      string  `csv:"parent_station"`
	ZoneId      string  `csv:"zone_id"`
}

// StopTime -
type StopTime struct {
	StopID       string  `csv:"stop_id"`
	StopSeq      uint32  `csv:"stop_sequence"`
	StopHeadSign string  `csv:"stop_headsign"`
	TripID       string  `csv:"trip_id"`
	Shape        float64 `csv:"shape_dist_traveled"`
	Departure    string  `csv:"departure_time"`
	Arrival      string  `csv:"arrival_time"`
}

// Calendar -
type Calendar struct {
	ServiceID string `csv:"service_id"`
	Monday    int    `csv:"monday"`
	Tuesday   int    `csv:"tuesday"`
	Wednesday int    `csv:"wednesday"`
	Thursday  int    `csv:"thursday"`
	Friday    int    `csv:"friday"`
	Saturday  int    `csv:"saturday"`
	Sunday    int    `csv:"sunday"`
	Start     string `csv:"start_date"`
	End       string `csv:"end_date"`
}

// CalendarDate -
type CalendarDate struct {
	ServiceID     string `csv:"service_id"`
	Date          string `csv:"date"`
	ExceptionType int    `csv:"exception_type"`
}

// Transfer -
type Transfer struct {
	FromStopID string `csv:"from_stop_id"`
	ToStopID   string `csv:"to_stop_id"`
	Type       int    `csv:"transfer_type"`
	MinTime    int    `csv:"min_transfer_time"`
}

// Agency -
type Agency struct {
	ID       string `csv:"agency_id"`
	Name     string `csv:"agency_name"`
	URL      string `csv:"agency_url"`
	Timezone string `csv:"agency_timezone"`
	Langue   string `csv:"agency_lang"`
	Phone    string `csv:"agency_phone"`
}

// Attribution -
type Attribution struct {
	ID               string `csv:"attribution_id"`
	AgencyID         string `csv:"agency_id"`
	RouteID          string `csv:"route_id"`
	TripID           string `csv:"trip_id"`
	OrganizationName string `csv:"organization_name"`
	IsProducer       int    `csv:"is_producer"`
	IsOperator       int    `csv:"is_operator"`
	IsAuthority      int    `csv:"is_authority"`
	Url              string `csv:"attribution_url"`
	Email            string `csv:"attribution_email"`
	Phone            string `csv:"attribution_phone"`
}

// Frequency -
type Frequency struct {
	TripId         string `csv:"trip_id"`
	StartTime      string `csv:"start_time"`
	EndTime        string `csv:"end_time"`
	HeadwaySeconds uint32 `csv:"headway_secs"`
	ExactTimes     string `csv:"exact_times"`
}

// FeedInfo -
type FeedInfo struct {
	PublisherName   string `csv:"feed_publisher_name"`
	PublisherUrl    string `csv:"feed_publisher_url"`
	Language        string `csv:"feed_lang"`
	DefaultLanguage string `csv:"default_lang"`
	StartDate       string `csv:"feed_start_date"`
	EndDate         string `csv:"feed_end_date"`
	Version         string `csv:"feed_version"`
	ContactEmail    string `csv:"feed_contact_email"`
	ContactUrl      string `csv:"feed_contact_url"`
}

// Level -
type Level struct {
	ID    string  `csv:"level_id"`
	Index float64 `csv:"level_index"`
	Name  string  `csv:"level_name"`
}

// FareAttribute -
type FareAttribute struct {
	ID               string  `csv:"fare_id"`
	Price            float64 `csv:"price"`
	CurrencyType     string  `csv:"currency_type"`
	PaymentMethod    int     `csv:"payment_method"`
	Transfers        int     `csv:"transfers"`
	AgencyID         string  `csv:"agency_id"`
	TransferDuration string  `csv:"transfer_duration"`
}

// FareRule -
type FareRule struct {
	ID            string `csv:"fare_id"`
	RouteID       string `csv:"route_id"`
	OriginID      string `csv:"origin_id"`
	DestinationID string `csv:"destination_id"`
	ContainsID    string `csv:"contains_id"`
}

// Pathway -
type Pathway struct {
	ID                   string  `csv:"pathway_id"`
	FromStopID           string  `csv:"from_stop_id"`
	ToStopID             string  `csv:"to_stop_id"`
	PathwayMode          int     `csv:"pathway_mode"`
	IsBidirectional      int     `csv:"is_bidirectional"`
	Length               float64 `csv:"length"`
	TraversalTime        uint32  `csv:"traversal_time"`
	StairCount           int     `csv:"stair_count"`
	MaxSlope             float64 `csv:"max_slope"`
	MinWidth             float64 `csv:"min_width"`
	SignpostedAs         string  `csv:"signposted_as"`
	ReversedSignpostedAs string  `csv:"reversed_signposted_as"`
}

// Shape -
type Shape struct {
	ID               string  `csv:"shape_id"`
	PointLatitude    float64 `csv:"shape_pt_lat"`
	PointLongitude   float64 `csv:"shape_pt_lon"`
	PointSequence    int64   `csv:"shape_pt_sequence"`
	DistanceTraveled float64 `csv:"shape_dist_traveled"`
}
