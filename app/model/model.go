package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Password string             `json:"password" bson:"password"`

	FIO      string   `json:"fio" bson:"fio"`
	PhotoURL string   `json:"photoURL" bson:"photoURL"`
	SNILS    string   `json:"SNILS" bson:"SNILS"`
	About    string   `json:"about" bson:"about"`
	Phone    string   `json:"phone" bson:"phone"`
	Socials  string   `json:"socials" bson:"socials"`
	Email    string   `json:"email" bson:"email" binding:"required"`
	Skills   []string `json:"skills" bson:"skills"`

	WorkExperience []struct {
		CompanyName      string   `json:"companyName" bson:"companyName"`
		Position         string   `json:"position" bson:"position"`
		DateStart        string   `json:"dateStart" bson:"dateStart"`
		DateEnd          string   `json:"dateEnd" bson:"dateEnd"`
		Responsibilities []string `json:"responsibilities"  bson:"responsibilities"`
	} `json:"workExperience" bson:"workExperience"`

	Education []struct {
		Name      string `json:"name" bson:"name"`
		Specialty string `json:"specialty" bson:"specialty"`
		Degree    string `json:"degree" bson:"degree"`
		DateStart string `json:"dateStart" bson:"dateStart"`
		DateEnd   string `json:"dateEnd" bson:"dateEnd"`
	} `json:"education" bson:"education"`

	RegisteredEvents []primitive.ObjectID `bson:"registeredEvents" json:"registeredEvents"`
}

type Event struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	PhotoURL        string             `json:"photoURL" bson:"photoURL"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	Type            string             `json:"type" bson:"type"`
	Date            string             `json:"date" bson:"date"`
	Time            string             `json:"time" bson:"time"`
	FullDescription string             `json:"fullDescription" bson:"fullDescription"`
	Location        string             `json:"location" bson:"location"`
	Email           string             `json:"email" bson:"email"`
	Website         string             `json:"website" bson:"website"`
}

type Project struct {
	Area     string `json:"area" bson:"area"`
	Name     string `json:"name" bson:"name"`
	PhotoURL string `json:"photoURL" bson:"photoURL"`

	TeamCapitanID primitive.ObjectID   `json:"teamCapitan" bson:"teamCapitan"`
	TeamIDs       []primitive.ObjectID `json:"teamIDs" bson:"teamIDs"`

	Description     string `json:"description" bson:"description"`
	UniqueAdvantage string `json:"uniqueAdvantage" bson:"uniqueAdvantage"`
	ReadyStage      string `json:"readyStage" bson:"readyStage"`

	IntellectualProperty          string   `json:"intellectualProperty" bson:"intellectualProperty"`
	AdditionalMaterials           []string `json:"additionalMaterials" bson:"additionalMaterials"`
	Needs                         string   `json:"needs" bson:"needs"`
	MarketApplication             string   `json:"marketApplication" bson:"marketApplication"`
	MarketCapacityTargetCustomers string   `json:"marketCapacityTargetCustomers" bson:"marketCapacityTargetCustomers"`
	Competitors                   string   `json:"competitors" bson:"competitors"`

	DateStart                  string `json:"dateStart" bson:"dateStart"`
	LeadershipExperience       string `json:"leadershipExperience" bson:"leadershipExperience"`
	ResourcesAndInfrastructure string `json:"resourcesAndInfrastructure" bson:"resourcesAndInfrastructure"`
	CurrentProjectStatus       string `json:"currentProjectStatus" bson:"currentProjectStatus"`
	ImplementationModel        string `json:"implementationModel" bson:"implementationModel"`
	DevelopmentPlan            string `json:"developmentPlan" bson:"developmentPlan"`
}
