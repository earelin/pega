package file_reader

type Control struct {
	ElectionType                                     int  `position:"0" length:"2"`
	Year                                             int  `position:"2" length:"4"`
	Month                                            int  `position:"6" length:"2"`
	Round                                            int  `position:"8" length:"1"`
	ControlFile                                      bool `position:"9" length:"1"`
	IdentificationFile                               bool `position:"10" length:"1"`
	CandidaturesFile                                 bool `position:"11" length:"1"`
	CandidatesListFile                               bool `position:"12" length:"1"`
	MunicipalitiesCommonDataFile                     bool `position:"13" length:"1"`
	MunicipalitiesCandidaturesDataFile               bool `position:"14" length:"1"`
	MunicipalitiesSuperiorScopeCommonDataFile        bool `position:"15" length:"1"`
	MunicipalitiesSuperiorScopeCandidaturesDataFile  bool `position:"16" length:"1"`
	TablesAndCeraCommonDataFile                      bool `position:"17" length:"1"`
	TablesAndCeraCandidaturesDataFile                bool `position:"18" length:"1"`
	MunicipalitiesSmallerThan250CommonDataFile       bool `position:"19" length:"1"`
	MunicipalitiesSmallerThan250CandidaturesDataFile bool `position:"20" length:"1"`
	JudicialDistrictCommonDataFile                   bool `position:"21" length:"1"`
	JudicialDistrictCandidaturesDataFile             bool `position:"22" length:"1"`
	ProvincialCouncilCommonDataFile                  bool `position:"23" length:"1"`
	ProvincialCouncilCandidaturesDataFile            bool `position:"24" length:"1"`
}

type Identification struct {
	Type                           int    `position:"0" length:"2"`
	Year                           int    `position:"2" length:"4"`
	Month                          int    `position:"6" length:"2"`
	Round                          int    `position:"8" length:"1"`
	ScopeType                      string `position:"9" length:"1"`
	TerritorialScope               int    `position:"10" length:"2"`
	CelebrationDay                 int    `position:"12" length:"2"`
	CelebrationMonth               int    `position:"14" length:"2"`
	CelebrationYear                int    `position:"16" length:"4"`
	PollStationOpeningTime         string `position:"20" length:"5"`
	PollStationClosingTime         string `position:"25" length:"5"`
	FirstParticipationAdvanceTime  string `position:"30" length:"5"`
	SecondParticipationAdvanceTime string `position:"35" length:"5"`
}

type Candidacy struct {
}