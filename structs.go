package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	Username          string             `bson:"username"`
	Paying            bool               `bson:"paying"`
	Provider          string             `bson:"provider"`
	Level             float32            `bson:"level"`
	BannedExercises   []string           `bson:"bannedExer"`
	BannedStretches   []string           `bson:"bannedStr"`
	BannedParts       []int              `bson:"bannedParts"`
	PlyoTolerance     int                `bson:"plyoToler"`
	ExerFavoriteRates map[string]float32 `bson:"exerfavs"`
	ExerModifications map[string]float32 `bson:"exermods"`
	TypeModifications map[string]float32 `bson:"typemods"`
	RoundEndurance    map[int]float32    `bson:"roundendur"`
	TimeEndurance     map[int]float32    `bson:"timeendur"`
	PushupSetting     string             `bson:"pushsetting"`
	LastMinutes       float32            `bson:"lastmins"`
	LastDifficulty    int                `bson:"lastdiff"`
	Assessed          bool               `bson:"assessed"`
	Badges            []string           `bson:"badges"`    //New
	CompletedCount    int                `bson:"completed"` //New
}

type Workout struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Name            string               `bson:"name"`
	UserID          string               `bson:"userid"`
	Username        string               `bson:"username"`
	Created         primitive.DateTime   `bson:"date"`     //PtlNew
	LastStarted     primitive.DateTime   `bson:"lastdate"` //New
	DateList        []primitive.DateTime `bson:"datelist"` //New
	Status          string               `bson:"status"`
	Minutes         float32              `bson:"minutes"`
	StretchTimes    StretchTimes         `bson:"stretchtimes"`
	PausedTime      float32              `bson:"paused"`
	LevelAtStart    float32              `bson:"level"`
	Difficulty      int                  `bson:"difficulty"`
	Dynamics        []string             `bson:"dynamics"`
	Statics         []string             `bson:"statics"`
	Exercises       [9]WorkoutRound      `bson:"exercises"`
	CardioRatings   [9]float32           `bson:"cardioratings"`
	CardioRating    float32              `bson:"cardiorating"`
	GeneralTypeVals [3]float32           `bson:"gentypevals"`
	IsIntro         bool                 `bson:"intro"`
	IsPinned        bool                 `bson:"pinned"`     //New
	AvgRating       float32              `bson:"rating"`     //New
	AvgFaves        float32              `bson:"faves"`      //New
	LastRating      int                  `bson:"lastrating"` //New
	LastFaves       int                  `bson:"lastfaves"`  //New
	RatedCount      int                  `bson:"ratedct"`    //New
	StartedCount    int                  `bson:"startedct"`  //New
}

type StretchTimes struct {
	DynamicPerSet []float32 `bson:"dynamicperset"`
	StaticPerSet  []float32 `bson:"staticperset"`
	DynamicSets   int       `bson:"dynamicsets"`
	StaticSets    int       `bson:"staticsets"`
	DynamicRest   float32   `bson:"dynamicrest"`
	FullRound     float32   `bson:"fullround"`
}

type ExerciseTimes struct {
	ExercisePerSet float32 `bson:"exerciseperset"`
	RestPerSet     float32 `bson:"restperset"`
	Sets           int     `bson:"sets"`
	RestPerRound   float32 `bson:"restperround"`
	FullRound      float32 `bson:"fullround"`
	ComboExers     int     `bson:"comboexers"`
}

type WorkoutRound struct {
	ExerciseIDs []string      `bson:"exerids"`
	Reps        []float32     `bson:"reps"`
	Pairs       []bool        `bson:"pairs"`
	Status      string        `bson:"status"`
	Times       ExerciseTimes `bson:"times"`
	AvgRating   float32       `bson:"rating"` //PtlNew
	AvgFaves    float32       `bson:"faves"`  //New
}

type StretchWorkout struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	UserID       string               `bson:"userid"`
	Created      primitive.DateTime   `bson:"date"`     //PtlNew
	LastStarted  primitive.DateTime   `bson:"lastdate"` //New
	DateList     []primitive.DateTime `bson:"datelist"` //New
	Status       string               `bson:"status"`
	StretchTimes StretchTimes         `bson:"stretchtimes"`
	LevelAtStart float32              `bson:"level"`
	PausedTime   float32              `bson:"paused"`
	Minutes      float32              `bson:"minutes"`
	Dynamics     []string             `bson:"dynamics"`
	Statics      []string             `bson:"statics"`
	IsPinned     bool                 `bson:"pinned"`    //New
	StartedCount int                  `bson:"startedct"` //New
}

type Exercise struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Parent       string             `bson:"parent"`
	MinLevel     float32            `bson:"minlevel"`
	MaxLevel     float32            `bson:"maxlevel"`
	MinReps      int                `bson:"minreps"`
	PlyoRating   int                `bson:"plyorating"`
	StartQuality float32            `bson:"startquality"`
	BodyParts    []int              `bson:"bodyparts"`
	RepVars      [3]float32         `bson:"repvars"`
	InSplits     bool               `bson:"insplits"`
	InPairs      bool               `bson:"inpairs"`
	UnderCombos  bool               `bson:"undercombos"`
	CardioRating float32            `bson:"cardiorating"`
	PushupType   string             `bson:"pushuptype"`
	GeneralType  []string           `bson:"generaltype"`
	SinglesGroup int                `bson:"singlesgroup"`
}

type Stretch struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	MinLevel     float32            `bson:"minlevel"`
	Status       string             `bson:"status"`
	BodyParts    []int              `bson:"bodyparts"`
	InPairs      bool               `bson:"inpairs"`
	DynamicPairs []string           `bson:"dynamicpairs"`
	Weight       float32            `bson:"weight"`
	ReqGroup     int                `bson:"reqgroup"`
}

type TypeMatrix struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Matrix [11][11]float32    `bson:"matrix"`
}

type DBToken struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID string             `bson:"user"`
	Token  string             `bson:"token"`
}

type StoredRating struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        string             `bson:"user"`
	WorkoutID     string             `bson:"workoutid"`
	Date          primitive.DateTime `bson:"date"`
	Minutes       float32            `bson:"minutes"`
	LevelAtStart  float32            `bson:"level"`
	Difficulty    int                `bson:"difficulty"`
	OverallRating int                `bson:"rating"`
	OverallFave   int                `bson:"fave"`
	RoundRatings  [9]RoundRating     `bson:"roundratings"`
}

type RoundRating struct {
	ActualRound WorkoutRound `bson:"round"`
	Rating      int          `bson:"rating"`
	Fave        int          `bson:"fave"`
	HasRating   bool         `bson:"hasrating"`
	HasFave     bool         `bson:"hasfave"`
}

type AnyWorkout interface {
	Display()
}

func (t Workout) Display() {
	fmt.Println("Workout: ")
	fmt.Println(t)
}

func (t StretchWorkout) Display() {
	fmt.Println("Stretch Workout: ")
	fmt.Println(t)
}
