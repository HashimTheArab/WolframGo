package wolframgo

type baseResult struct {
	Result ComplexResult `json:"queryresult"`
}

type ComplexResult struct {
	Success       bool         `json:"success"`
	Error         bool         `json:"error"`
	NumPods       int          `json:"numpods"`
	DataTypes     string       `json:"datatypes"`
	TimedOut      string       `json:"timedout"`
	TimedOutPods  string       `json:"timedoutpods"`
	Timing        float64      `json:"timing"`
	ParseTiming   float64      `json:"parsetiming"`
	ParseTimedOut bool         `json:"parsetimedout"`
	Recalculate   string       `json:"recalculate"`
	ID            string       `json:"id"`
	Host          string       `json:"host"`
	Server        string       `json:"server"`
	Related       string       `json:"related"`
	Version       string       `json:"version"`
	InputString   string       `json:"inputstring"`
	Pods          []Pod        `json:"pods"`
	UserInfoused  UserInfoused `json:"userinfoused"`
	Sources       Sources      `json:"sources"`
}

type Pod struct {
	Title           string          `json:"title"`
	Scanner         string          `json:"scanner"`
	ID              string          `json:"id"`
	Position        int             `json:"position"`
	Error           bool            `json:"error"`
	NumSubPods      int             `json:"numsubpods"`
	Primary         bool            `json:"primary"`
	SubPods         []SubPods       `json:"subpods"`
	ExpressionTypes ExpressionTypes `json:"expressiontypes"`
}

type Img struct {
	Src             string `json:"src"`
	Alt             string `json:"alt"`
	Title           string `json:"title"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Type            string `json:"type"`
	Themes          string `json:"themes"`
	ColorInvertable bool   `json:"colorinvertable"`
}

type SubPods struct {
	Title     string `json:"title"`
	Img       Img    `json:"img"`
	Plaintext string `json:"plaintext"`
}

type ExpressionTypes struct {
	Name string `json:"name"`
}

type UserInfoused struct {
	Name string `json:"name"`
}

type Sources struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}
