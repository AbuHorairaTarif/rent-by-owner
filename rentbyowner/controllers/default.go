package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type RefineSearch struct {
	beego.Controller
	Location      string
	Checkin_date  string
	Checkout_date string
}

type PropertyDetail struct {
	beego.Controller
	Location string
}

type PropertyInfoDetail struct {
	beego.Controller
}

type StayInfo struct {
	CurrentPage int `json:"currentPage"`
	Data        []struct {
		Typename            string        `json:"__typename"`
		AcceptsWalletCredit bool          `json:"acceptsWalletCredit"`
		Badges              []interface{} `json:"badges"`
		BasicPropertyData   struct {
			Typename                   string      `json:"__typename"`
			AccommodationTypeID        int         `json:"accommodationTypeId"`
			AlternativeExternalReviews interface{} `json:"alternativeExternalReviews"`
			ExternalReviews            interface{} `json:"externalReviews"`
			ID                         int         `json:"id"`
			IsClosed                   bool        `json:"isClosed"`
			IsTestProperty             bool        `json:"isTestProperty"`
			Location                   struct {
				Typename    string `json:"__typename"`
				Address     string `json:"address"`
				City        string `json:"city"`
				CountryCode string `json:"countryCode"`
			} `json:"location"`
			PageName      string `json:"pageName"`
			PaymentConfig struct {
				Typename     string      `json:"__typename"`
				Installments interface{} `json:"installments"`
			} `json:"paymentConfig"`
			Photos struct {
				Typename string `json:"__typename"`
				Main     struct {
					Typename       string `json:"__typename"`
					HighResJpegURL struct {
						Typename    string `json:"__typename"`
						RelativeURL string `json:"relativeUrl"`
					} `json:"highResJpegUrl"`
					HighResURL struct {
						Typename    string `json:"__typename"`
						RelativeURL string `json:"relativeUrl"`
					} `json:"highResUrl"`
					LowResJpegURL struct {
						Typename    string `json:"__typename"`
						RelativeURL string `json:"relativeUrl"`
					} `json:"lowResJpegUrl"`
					LowResURL struct {
						Typename    string `json:"__typename"`
						RelativeURL string `json:"relativeUrl"`
					} `json:"lowResUrl"`
				} `json:"main"`
			} `json:"photos"`
			Reviews struct {
				Typename         string  `json:"__typename"`
				ReviewsCount     int     `json:"reviewsCount"`
				SecondaryScore   float64 `json:"secondaryScore"`
				SecondaryTextTag struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"secondaryTextTag"`
				ShowScore          bool `json:"showScore"`
				ShowSecondaryScore bool `json:"showSecondaryScore"`
				TotalScore         int  `json:"totalScore"`
				TotalScoreTextTag  struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"totalScoreTextTag"`
			} `json:"reviews"`
			StarRating struct {
				Typename string `json:"__typename"`
				Caption  struct {
					Typename    string      `json:"__typename"`
					Translation interface{} `json:"translation"`
				} `json:"caption"`
				ShowAdditionalInfoIcon bool   `json:"showAdditionalInfoIcon"`
				Symbol                 string `json:"symbol"`
				TocLink                struct {
					Typename    string      `json:"__typename"`
					Translation interface{} `json:"translation"`
				} `json:"tocLink"`
				Value int `json:"value"`
			} `json:"starRating"`
			Ufi int `json:"ufi"`
		} `json:"basicPropertyData"`
		Blocks []struct {
			Typename string `json:"__typename"`
			BlockID  struct {
				Typename      string `json:"__typename"`
				MealPlanID    int    `json:"mealPlanId"`
				Occupancy     int    `json:"occupancy"`
				PackageID     string `json:"packageId"`
				PolicyGroupID string `json:"policyGroupId"`
				RoomID        string `json:"roomId"`
			} `json:"blockId"`
			BlockMatchTags struct {
				Typename          string `json:"__typename"`
				ChildStaysForFree bool   `json:"childStaysForFree"`
			} `json:"blockMatchTags"`
			FinalPrice struct {
				Typename string  `json:"__typename"`
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"finalPrice"`
			FreeCancellationUntil time.Time   `json:"freeCancellationUntil"`
			HasCrib               bool        `json:"hasCrib"`
			OnlyXLeftMessage      interface{} `json:"onlyXLeftMessage"`
			OriginalPrice         struct {
				Typename string  `json:"__typename"`
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"originalPrice"`
			ThirdPartyInventoryContext struct {
				Typename   string `json:"__typename"`
				IsTpiBlock bool   `json:"isTpiBlock"`
			} `json:"thirdPartyInventoryContext"`
		} `json:"blocks"`
		BookerExperienceContentUIComponentProps []interface{} `json:"bookerExperienceContentUIComponentProps"`
		BundleRatesAvailable                    bool          `json:"bundleRatesAvailable"`
		CustomBadges                            struct {
			Typename                string      `json:"__typename"`
			ShowBhTravelCreditBadge bool        `json:"showBhTravelCreditBadge"`
			ShowIsWorkFriendly      bool        `json:"showIsWorkFriendly"`
			ShowOnlineCheckinBadge  interface{} `json:"showOnlineCheckinBadge"`
			ShowParkAndFly          bool        `json:"showParkAndFly"`
			ShowSkiToDoor           bool        `json:"showSkiToDoor"`
		} `json:"customBadges"`
		Description interface{} `json:"description"`
		DisplayName struct {
			Typename       string `json:"__typename"`
			Title          string `json:"text"`
			TranslationTag struct {
				Typename    string      `json:"__typename"`
				Translation interface{} `json:"translation"`
			} `json:"translationTag"`
		} `json:"displayName"`
		GeniusInfo            interface{} `json:"geniusInfo"`
		IDDetail              string      `json:"idDetail"`
		InferredLocationScore int         `json:"inferredLocationScore"`
		IsInCompanyBudget     interface{} `json:"isInCompanyBudget"`
		IsNewlyOpened         bool        `json:"isNewlyOpened"`
		LicenseDetails        interface{} `json:"licenseDetails"`
		Location              struct {
			Typename                           string        `json:"__typename"`
			BeachDistance                      string        `json:"beachDistance"`
			BeachWalkingTime                   interface{}   `json:"beachWalkingTime"`
			DisplayLocation                    string        `json:"displayLocation"`
			GeoDistanceMeters                  interface{}   `json:"geoDistanceMeters"`
			MainDistance                       string        `json:"mainDistance"`
			NearbyBeachNames                   []interface{} `json:"nearbyBeachNames"`
			PublicTransportDistanceDescription interface{}   `json:"publicTransportDistanceDescription"`
			SkiLiftDistance                    interface{}   `json:"skiLiftDistance"`
		} `json:"location"`
		MatchingUnitConfigurations struct {
			Typename            string `json:"__typename"`
			CommonConfiguration struct {
				Typename          string `json:"__typename"`
				BedConfigurations []struct {
					Typename string `json:"__typename"`
					Beds     []struct {
						Typename string `json:"__typename"`
						Count    int    `json:"count"`
						Type     int    `json:"type"`
					} `json:"beds"`
					NbAllBeds int `json:"nbAllBeds"`
				} `json:"bedConfigurations"`
				LocalizedArea struct {
					Typename      string `json:"__typename"`
					LocalizedArea string `json:"localizedArea"`
					Unit          string `json:"unit"`
				} `json:"localizedArea"`
				Name          interface{} `json:"name"`
				NbAllBeds     int         `json:"nbAllBeds"`
				NbBathrooms   int         `json:"nbBathrooms"`
				NbBedrooms    int         `json:"nbBedrooms"`
				NbKitchens    int         `json:"nbKitchens"`
				NbLivingrooms int         `json:"nbLivingrooms"`
				NbPools       int         `json:"nbPools"`
				NbUnits       int         `json:"nbUnits"`
				UnitID        int         `json:"unitId"`
				UnitTypeNames []struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"unitTypeNames"`
			} `json:"commonConfiguration"`
			UnitConfigurations []struct {
				Typename       string `json:"__typename"`
				ApartmentRooms []struct {
					Typename string `json:"__typename"`
					Config   struct {
						Typename  string `json:"__typename"`
						BedTypeID int    `json:"bedTypeId"`
						Count     int    `json:"count"`
						ID        int    `json:"id"`
						RoomType  string `json:"roomType"`
					} `json:"config"`
					Tag struct {
						Typename    string `json:"__typename"`
						Tag         string `json:"tag"`
						Translation string `json:"translation"`
					} `json:"tag"`
				} `json:"apartmentRooms"`
				BedConfigurations []struct {
					Typename string `json:"__typename"`
					Beds     []struct {
						Typename string `json:"__typename"`
						Count    int    `json:"count"`
						Type     int    `json:"type"`
					} `json:"beds"`
					NbAllBeds int `json:"nbAllBeds"`
				} `json:"bedConfigurations"`
				LocalizedArea struct {
					Typename      string `json:"__typename"`
					LocalizedArea string `json:"localizedArea"`
					Unit          string `json:"unit"`
				} `json:"localizedArea"`
				Name          string `json:"name"`
				NbAllBeds     int    `json:"nbAllBeds"`
				NbBathrooms   int    `json:"nbBathrooms"`
				NbBedrooms    int    `json:"nbBedrooms"`
				NbKitchens    int    `json:"nbKitchens"`
				NbLivingrooms int    `json:"nbLivingrooms"`
				NbPools       int    `json:"nbPools"`
				NbUnits       int    `json:"nbUnits"`
				UnitID        int    `json:"unitId"`
				UnitTypeNames []struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"unitTypeNames"`
			} `json:"unitConfigurations"`
		} `json:"matchingUnitConfigurations"`
		MealPlanIncluded interface{} `json:"mealPlanIncluded"`
		NbWishlists      int         `json:"nbWishlists"`
		Persuasion       struct {
			Typename            string      `json:"__typename"`
			Autoextended        bool        `json:"autoextended"`
			BookedXTimesMessage interface{} `json:"bookedXTimesMessage"`
			GeniusRateAvailable bool        `json:"geniusRateAvailable"`
			Highlighted         bool        `json:"highlighted"`
			NativeAdID          interface{} `json:"nativeAdId"`
			NativeAdsCpc        interface{} `json:"nativeAdsCpc"`
			NativeAdsTracking   interface{} `json:"nativeAdsTracking"`
			Preferred           bool        `json:"preferred"`
			PreferredPlus       bool        `json:"preferredPlus"`
			ShowNativeAdLabel   bool        `json:"showNativeAdLabel"`
			SponsoredAdsData    interface{} `json:"sponsoredAdsData"`
		} `json:"persuasion"`
		Policies struct {
			Typename                       string      `json:"__typename"`
			EnableJapaneseUsersSpecialCase interface{} `json:"enableJapaneseUsersSpecialCase"`
			ShowFreeCancellation           bool        `json:"showFreeCancellation"`
			ShowNoPrepayment               bool        `json:"showNoPrepayment"`
		} `json:"policies"`
		PriceDisplayInfoIrene struct {
			Typename    string        `json:"__typename"`
			Badges      []interface{} `json:"badges"`
			ChargesInfo struct {
				Typename    string      `json:"__typename"`
				Translation interface{} `json:"translation"`
			} `json:"chargesInfo"`
			Discounts []struct {
				Typename string `json:"__typename"`
				Amount   struct {
					Typename          string  `json:"__typename"`
					Amount            string  `json:"amount"`
					AmountRounded     string  `json:"amountRounded"`
					AmountUnformatted float64 `json:"amountUnformatted"`
					Currency          string  `json:"currency"`
				} `json:"amount"`
				Description struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"description"`
				ItemType string `json:"itemType"`
				Name     struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"name"`
				ProductID string `json:"productId"`
			} `json:"discounts"`
			DisplayPrice struct {
				Typename      string `json:"__typename"`
				AmountPerStay struct {
					Typename          string  `json:"__typename"`
					Amount            string  `json:"amount"`
					AmountRounded     string  `json:"amountRounded"`
					AmountUnformatted float64 `json:"amountUnformatted"`
					Currency          string  `json:"currency"`
				} `json:"amountPerStay"`
				Copy struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"copy"`
			} `json:"displayPrice"`
			ExcludedCharges struct {
				Typename                 string `json:"__typename"`
				ExcludeChargesAggregated struct {
					Typename      string `json:"__typename"`
					AmountPerStay struct {
						Typename          string  `json:"__typename"`
						Amount            string  `json:"amount"`
						AmountRounded     string  `json:"amountRounded"`
						AmountUnformatted float64 `json:"amountUnformatted"`
						Currency          string  `json:"currency"`
					} `json:"amountPerStay"`
					Copy struct {
						Typename    string      `json:"__typename"`
						Translation interface{} `json:"translation"`
					} `json:"copy"`
				} `json:"excludeChargesAggregated"`
				ExcludeChargesList []struct {
					Typename      string `json:"__typename"`
					AmountPerStay struct {
						Typename          string  `json:"__typename"`
						Amount            string  `json:"amount"`
						AmountRounded     string  `json:"amountRounded"`
						AmountUnformatted float64 `json:"amountUnformatted"`
						Currency          string  `json:"currency"`
					} `json:"amountPerStay"`
					ChargeInclusion string `json:"chargeInclusion"`
					ChargeMode      string `json:"chargeMode"`
					ChargeType      int    `json:"chargeType"`
				} `json:"excludeChargesList"`
			} `json:"excludedCharges"`
			PriceBeforeDiscount struct {
				Typename      string `json:"__typename"`
				AmountPerStay struct {
					Typename          string  `json:"__typename"`
					Amount            string  `json:"amount"`
					AmountRounded     string  `json:"amountRounded"`
					AmountUnformatted float64 `json:"amountUnformatted"`
					Currency          string  `json:"currency"`
				} `json:"amountPerStay"`
				Copy struct {
					Typename    string `json:"__typename"`
					Translation string `json:"translation"`
				} `json:"copy"`
			} `json:"priceBeforeDiscount"`
			Rewards struct {
				Typename          string `json:"__typename"`
				RewardsAggregated struct {
					Typename      string `json:"__typename"`
					AmountPerStay struct {
						Typename          string `json:"__typename"`
						Amount            string `json:"amount"`
						AmountRounded     string `json:"amountRounded"`
						AmountUnformatted int    `json:"amountUnformatted"`
						Currency          string `json:"currency"`
					} `json:"amountPerStay"`
					Copy struct {
						Typename    string `json:"__typename"`
						Translation string `json:"translation"`
					} `json:"copy"`
				} `json:"rewardsAggregated"`
				RewardsList []interface{} `json:"rewardsList"`
			} `json:"rewards"`
			TaxExceptions    []interface{} `json:"taxExceptions"`
			UseRoundedAmount bool          `json:"useRoundedAmount"`
		} `json:"priceDisplayInfoIrene"`
		PropertySustainability struct {
			Typename        string      `json:"__typename"`
			Certifications  interface{} `json:"certifications"`
			ChainProgrammes interface{} `json:"chainProgrammes"`
			Facilities      []struct {
				Ref string `json:"__ref"`
			} `json:"facilities"`
			IsSustainable bool   `json:"isSustainable"`
			LevelID       string `json:"levelId"`
			Tier          struct {
				Typename string `json:"__typename"`
				Type     string `json:"type"`
			} `json:"tier"`
		} `json:"propertySustainability"`
		RecommendedDate struct {
			Typename     string `json:"__typename"`
			Checkin      string `json:"checkin"`
			Checkout     string `json:"checkout"`
			LengthOfStay int    `json:"lengthOfStay"`
		} `json:"recommendedDate"`
		RecommendedDatesLabel  interface{}   `json:"recommendedDatesLabel"`
		RelocationMode         interface{}   `json:"relocationMode"`
		Ribbon                 interface{}   `json:"ribbon"`
		SeoThemes              []interface{} `json:"seoThemes"`
		ShowAdLabel            bool          `json:"showAdLabel"`
		ShowGeniusLoginMessage bool          `json:"showGeniusLoginMessage"`
		ShowPrivateHostMessage bool          `json:"showPrivateHostMessage"`
		SoldOutInfo            struct {
			Typename                 string        `json:"__typename"`
			AlternativeDatesMessages []interface{} `json:"alternativeDatesMessages"`
			IsSoldOut                bool          `json:"isSoldOut"`
			Messages                 []interface{} `json:"messages"`
		} `json:"soldOutInfo"`
		TrackOnView              []interface{} `json:"trackOnView"`
		VisibilityBoosterEnabled interface{}   `json:"visibilityBoosterEnabled"`
	} `json:"data"`
	Message          string `json:"message"`
	ResultsPerPage   int    `json:"resultsPerPage"`
	Status           bool   `json:"status"`
	TotalPages       int    `json:"totalPages"`
	TotalResultCount int    `json:"totalResultCount"`
}

type PropertyInfoDetails struct {
	Data struct {
		Ambiance []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
		} `json:"Ambiance"`
		BaseFacility []struct {
			Typename  string `json:"__typename"`
			GroupID   int    `json:"groupId"`
			ID        int    `json:"id"`
			Instances []struct {
				Typename   string `json:"__typename"`
				Attributes struct {
					Typename                  string        `json:"__typename"`
					ClosureScheduleDuringStay []interface{} `json:"closureScheduleDuringStay"`
					ExtendedAttributes        interface{}   `json:"extendedAttributes"`
					IsOffsite                 bool          `json:"isOffsite"`
					Name                      interface{}   `json:"name"`
					PaymentInfo               struct {
						Typename      string      `json:"__typename"`
						ChargeDetails interface{} `json:"chargeDetails"`
						ChargeMode    string      `json:"chargeMode"`
					} `json:"paymentInfo"`
					ScheduleEntry []interface{} `json:"scheduleEntry"`
				} `json:"attributes"`
				SubFacilities []interface{} `json:"subFacilities"`
				Title         string        `json:"title"`
			} `json:"instances"`
			Slug string `json:"slug"`
			Icon string `json:"icon,omitempty"`
		} `json:"BaseFacility"`
		BasicPropertyData []struct {
			Typename            string      `json:"__typename"`
			AccommodationTypeID int         `json:"accommodationTypeId"`
			ExternalReviews     interface{} `json:"externalReviews"`
			ID                  int         `json:"id"`
			Location            struct {
				Typename  string  `json:"__typename"`
				City      string  `json:"city"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
			Name     string `json:"name"`
			PageName string `json:"pageName"`
			Reviews  struct {
				Typename  string      `json:"__typename"`
				ShowScore interface{} `json:"showScore"`
			} `json:"reviews"`
			Ufi int `json:"ufi"`
		} `json:"BasicPropertyData"`
		BookingHomeFlags []struct {
			Typename string `json:"__typename"`
			IsBH8    bool   `json:"isBH8"`
			IsSUP    bool   `json:"isSUP"`
		} `json:"BookingHomeFlags"`
		ChildAndBedPolicy []struct {
			Typename         string `json:"__typename"`
			AllowChildren    bool   `json:"allowChildren"`
			AllowCribs       bool   `json:"allowCribs"`
			AllowExtraBeds   bool   `json:"allowExtraBeds"`
			ChildPolicyRules []struct {
				Ref string `json:"__ref"`
			} `json:"childPolicyRules"`
			IsFCR          bool `json:"isFCR"`
			MinChildrenAge int  `json:"minChildrenAge"`
		} `json:"ChildAndBedPolicy"`
		ChildPolicyRule []struct {
			Typename       string `json:"__typename"`
			AgeBandID      int    `json:"ageBandId"`
			AgeFrom        int    `json:"ageFrom"`
			AgeTo          int    `json:"ageTo"`
			LocalizedPrice string `json:"localizedPrice"`
			Price          string `json:"price"`
			PriceMode      string `json:"priceMode"`
			PriceType      string `json:"priceType"`
			RuleID         string `json:"ruleId"`
			RuleType       string `json:"ruleType"`
		} `json:"ChildPolicyRule"`
		Cuisine []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
		} `json:"Cuisine"`
		FacilityGroup []struct {
			Typename string      `json:"__typename"`
			ID       int         `json:"id"`
			Slug     interface{} `json:"slug"`
			Summary  interface{} `json:"summary"`
			Title    string      `json:"title"`
		} `json:"FacilityGroup"`
		FeaturedReview []struct {
			Typename           string `json:"__typename"`
			AverageScore       int    `json:"averageScore"`
			Completed          int    `json:"completed"`
			CustomerType       string `json:"customerType"`
			GuestCountryCode   string `json:"guestCountryCode"`
			GuestName          string `json:"guestName"`
			HasHideCountryFlag bool   `json:"hasHideCountryFlag"`
			ID                 int64  `json:"id"`
			IsAnonymous        bool   `json:"isAnonymous"`
			IsTranslatable     bool   `json:"isTranslatable"`
			Language           string `json:"language"`
			NegativeText       string `json:"negativeText"`
			PositiveText       string `json:"positiveText"`
			PurposeType        string `json:"purposeType"`
			ReviewURL          string `json:"reviewUrl"`
			Title              string `json:"title"`
			UserAvatarURL      string `json:"userAvatarUrl"`
			UserID             int    `json:"userId"`
		} `json:"FeaturedReview"`
		FreeWifiFacilityHighlight []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
			Level    string `json:"level"`
			Title    string `json:"title"`
		} `json:"FreeWifiFacilityHighlight"`
		GenericFacilityHighlight []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
			Level    string `json:"level"`
			Title    string `json:"title"`
		} `json:"GenericFacilityHighlight"`
		HealthSafety []struct {
			Typename                  string        `json:"__typename"`
			HealthAndSafetyCategories []interface{} `json:"healthAndSafetyCategories"`
		} `json:"HealthSafety"`
		HotelTranslation []struct {
			Typename    string      `json:"__typename"`
			Description string      `json:"description"`
			FinePrint   interface{} `json:"finePrint"`
			Name        string      `json:"name"`
		} `json:"HotelTranslation"`
		MealType []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
		} `json:"MealType"`
		OpenForMeal []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
		} `json:"OpenForMeal"`
		Property []struct {
			Typename string `json:"__typename"`

			FacilityGroups []struct {
				Ref string `json:"__ref"`
			} `json:"facilityGroups"`

			ID      int `json:"id"`
			Profile struct {
				Typename        string   `json:"__typename"`
				SpokenLanguages []string `json:"spokenLanguages"`
			} `json:"profile"`
			QuestionsAndAnswers struct {
				Typename                    string `json:"__typename"`
				PartnerAverageResponseTime  int    `json:"partnerAverageResponseTime"`
				PartnerOptedOutOfQAndA      bool   `json:"partnerOptedOutOfQAndA"`
				PropertyQuestionAnswerPairs []struct {
					Typename           string      `json:"__typename"`
					Answer             string      `json:"answer"`
					AnsweredTimestamp  string      `json:"answeredTimestamp"`
					AskedTimestamp     string      `json:"askedTimestamp"`
					Question           string      `json:"question"`
					QuestionID         int         `json:"questionId"`
					RoomID             int         `json:"roomId"`
					ShowTranslations   bool        `json:"showTranslations"`
					TranslatedAnswer   interface{} `json:"translatedAnswer"`
					TranslatedQuestion interface{} `json:"translatedQuestion"`
				} `json:"propertyQuestionAnswerPairs"`
				UserAlreadyAskedQuestion bool `json:"userAlreadyAskedQuestion"`
			} `json:"questionsAndAnswers"`
			Reviews struct {
				Typename  string `json:"__typename"`
				Questions []struct {
					Typename     string  `json:"__typename"`
					Name         string  `json:"name"`
					ReviewsCount int     `json:"reviewsCount"`
					Score        float64 `json:"score"`
				} `json:"questions"`
				ReviewsCount int `json:"reviewsCount"`
			} `json:"reviews"`
			Surroundings struct {
				Typename  string `json:"__typename"`
				IsOldTown bool   `json:"isOldTown"`
			} `json:"surroundings"`
		} `json:"Property"`
		PropertyReview []struct {
			Typename        string `json:"__typename"`
			FeaturedReviews []struct {
				Ref string `json:"__ref"`
			} `json:"featuredReviews"`
			Questions []struct {
				Typename     string `json:"__typename"`
				CustomerType string `json:"customerType"`
				Question     struct {
					Typename   string `json:"__typename"`
					Question   string `json:"question"`
					QuestionID int    `json:"questionId"`
				} `json:"question"`
				ScoreSegment struct {
					Typename     string  `json:"__typename"`
					ReviewsCount int     `json:"reviewsCount"`
					Score        float64 `json:"score"`
				} `json:"scoreSegment"`
			} `json:"questions"`
			TotalScore struct {
				Typename     string  `json:"__typename"`
				ReviewsCount int     `json:"reviewsCount"`
				Score        float64 `json:"score"`
			} `json:"totalScore"`
		} `json:"PropertyReview"`
		PropertySustainability []struct {
			Typename        string      `json:"__typename"`
			Certifications  interface{} `json:"certifications"`
			ChainProgrammes interface{} `json:"chainProgrammes"`
			Facilities      interface{} `json:"facilities"`
			IsSustainable   bool        `json:"isSustainable"`
			Tier            struct {
				Ref string `json:"__ref"`
			} `json:"tier"`
		} `json:"PropertySustainability"`
		PropertySustainabilityTier []struct {
			Typename string `json:"__typename"`
			Type     string `json:"type"`
		} `json:"PropertySustainabilityTier"`
		PropertyTripTypeLabelsInfo []struct {
			Typename                 string   `json:"__typename"`
			EnabledDestinationLabels []string `json:"enabledDestinationLabels"`
			EnabledPropertyLabels    []string `json:"enabledPropertyLabels"`
			IsSkiSeason              bool     `json:"isSkiSeason"`
		} `json:"PropertyTripTypeLabelsInfo"`
		ROOTQUERY struct {
			Typename string `json:"__typename"`
			Checkin  string `json:"checkin"`
			Checkout string `json:"checkout"`
		} `json:"ROOT_QUERY"`
		RTBreakfastScore []struct {
			Typename   string `json:"__typename"`
			NumReviews int    `json:"numReviews"`
			Score      int    `json:"score"`
			ScoreClass string `json:"scoreClass"`
		} `json:"RTBreakfastScore"`
		RTHotelFacilityHighlight []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
		} `json:"RTHotelFacilityHighlight"`
		RTRoomCard []struct {
			Typename          string        `json:"__typename"`
			ApartmentRooms    []interface{} `json:"apartmentRooms"`
			BathroomCount     int           `json:"bathroomCount"`
			BedConfigurations []struct {
				Typename string `json:"__typename"`
				Beds     []struct {
					Typename  string `json:"__typename"`
					BedTypeID int    `json:"bedTypeId"`
					Count     int    `json:"count"`
				} `json:"beds"`
				ConfigurationID int `json:"configurationId"`
			} `json:"bedConfigurations"`
			CategorizedFacilities []struct {
				Typename   string `json:"__typename"`
				Category   string `json:"category"`
				Facilities []int  `json:"facilities"`
			} `json:"categorizedFacilities"`
			Description      string `json:"description"`
			Facilities       []int  `json:"facilities"`
			HasRoomInventory bool   `json:"hasRoomInventory"`
			HasRoomPage      bool   `json:"hasRoomPage"`
			HasSubUnits      bool   `json:"hasSubUnits"`
			Highlights       []struct {
				Typename     string `json:"__typename,omitempty"`
				PrivacyLevel string `json:"privacyLevel,omitempty"`
				AreaValue    int    `json:"areaValue,omitempty"`
				Ref          string `json:"__ref,omitempty"`
			} `json:"highlights"`
			IsSmoking                                     bool        `json:"isSmoking"`
			Name                                          string      `json:"name"`
			NrStaysWithCheapestPriceForRoomLevelInventory interface{} `json:"nrStaysWithCheapestPriceForRoomLevelInventory"`
			Occupancy                                     struct {
				Typename    string `json:"__typename"`
				MaxChildren int    `json:"maxChildren"`
				MaxGuests   int    `json:"maxGuests"`
				MaxPersons  int    `json:"maxPersons"`
			} `json:"occupancy"`
			OnlyXLeftMessageDetails interface{} `json:"onlyXLeftMessageDetails"`
			ParkingInfo             struct {
				Typename   string `json:"__typename"`
				HasParking bool   `json:"hasParking"`
			} `json:"parkingInfo"`
			Photos []struct {
				Ref string `json:"__ref"`
			} `json:"photos"`
			RoomAvBlocks                    []interface{} `json:"roomAvBlocks"`
			RoomID                          int           `json:"roomId"`
			RoomIndex                       int           `json:"roomIndex"`
			RoomSize                        int           `json:"roomSize"`
			RoomTypeID                      int           `json:"roomTypeId"`
			ShowIsLargestAvailableRoomBadge bool          `json:"showIsLargestAvailableRoomBadge"`
		} `json:"RTRoomCard"`
		RTRoomFacilityHighlight []struct {
			Typename string `json:"__typename"`
			ID       int    `json:"id"`
		} `json:"RTRoomFacilityHighlight"`
		RTRoomPhoto []struct {
			Typename     string `json:"__typename"`
			ID           int    `json:"id"`
			PhotoURI     string `json:"photoUri"`
			Ranking      int    `json:"ranking"`
			ThumbnailURI string `json:"thumbnailUri"`
		} `json:"RTRoomPhoto"`
		ReviewTopicTranslation []struct {
			Typename string `json:"__typename"`
			ID       string `json:"id"`
			Name     string `json:"name"`
		} `json:"ReviewTopicTranslation"`
		RoomTableQueryResult []struct {
			Typename              string `json:"__typename"`
			AccommodationTypeID   int    `json:"accommodationTypeId"`
			AreaMeasurementUnit   string `json:"areaMeasurementUnit"`
			BedComfortReviewScore struct {
				Typename     string  `json:"__typename"`
				ReviewsCount int     `json:"reviewsCount"`
				Score        float64 `json:"score"`
			} `json:"bedComfortReviewScore"`
			BreakfastScore struct {
				Ref string `json:"__ref"`
			} `json:"breakfastScore"`
			BreakfastType                      int           `json:"breakfastType"`
			Genius                             interface{}   `json:"genius"`
			HasRoomsWithDifferentPrivacyLevels bool          `json:"hasRoomsWithDifferentPrivacyLevels"`
			HotelID                            int           `json:"hotelId"`
			IsBookingHome                      bool          `json:"isBookingHome"`
			IsSoldOut                          bool          `json:"isSoldOut"`
			MealPlans                          []interface{} `json:"mealPlans"`
			PropertyName                       string        `json:"propertyName"`
			RoomCards                          []struct {
				Ref string `json:"__ref"`
			} `json:"roomCards"`
			Stars int `json:"stars"`
		} `json:"RoomTableQueryResult"`
		TopicFilter []struct {
			Typename    string `json:"__typename"`
			ID          int    `json:"id"`
			IsSelected  bool   `json:"isSelected"`
			Name        string `json:"name"`
			Translation struct {
				Ref string `json:"__ref"`
			} `json:"translation"`
		} `json:"TopicFilter"`
		TripTypesHotelPage []struct {
			Typename string `json:"__typename"`
			Beaches  []struct {
				Typename                   string      `json:"__typename"`
				BeachID                    int         `json:"beachId"`
				LocalizedDistanceFormatted string      `json:"localizedDistanceFormatted"`
				Name                       string      `json:"name"`
				PhotoURL                   interface{} `json:"photoUrl"`
				ReviewScore                int         `json:"reviewScore"`
				ReviewScoreFormatted       string      `json:"reviewScoreFormatted"`
			} `json:"beaches"`
			ClosestTopBeach struct {
				Typename             string `json:"__typename"`
				BeachID              int    `json:"beachId"`
				Name                 string `json:"name"`
				ReviewScore          int    `json:"reviewScore"`
				ReviewScoreFormatted string `json:"reviewScoreFormatted"`
			} `json:"closestTopBeach"`
			NearbySkiLifts interface{} `json:"nearbySkiLifts"`
		} `json:"TripTypesHotelPage"`
	} `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *RefineSearch) GetRefineSearch() {

	rapidApiHost, _ := beego.AppConfig.String("rapidapi-host-name")
	rapidApiHostValue, _ := beego.AppConfig.String("rapidapi-host-value")

	rapidApiKey, _ := beego.AppConfig.String("rapidapi-key-name")
	rapidApiKeyValue, _ := beego.AppConfig.String("rapidapi-key-value")

	location := c.GetString("location")
	checkIn := c.GetString("t-start")
	checkOut := c.GetString("t-end")

	url := "https://booking-com13.p.rapidapi.com/stays/properties/list-v2" +
		"?location=" + location +
		"&checkin_date=" + checkIn +
		"&checkout_date=" + checkOut +
		"&language_code=en-us&currency_code=USD"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add(rapidApiHost, rapidApiHostValue)
	req.Header.Add(rapidApiKey, rapidApiKeyValue)

	dataChannel := make(chan StayInfo)

	go func() {
		// Make the HTTP request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.Data["Error"] = "Error making the request"

		}
		defer res.Body.Close()

		// Check the HTTP status code
		if res.StatusCode != http.StatusOK {
			c.Data["Error"] = "API request failed with status code " + fmt.Sprint(res.StatusCode)

		}

		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.Data["Error"] = "Error reading the response"

		}

		// Parse the JSON response into a struct
		var stayResponse StayInfo

		err = json.Unmarshal(body, &stayResponse)
		if err != nil {
			c.Data["Error"] = "Error parsing JSON response"
		}

		// Send the extracted data to the channel
		dataChannel <- stayResponse
	}()

	// Receiving data from channel
	extractedData := <-dataChannel

	c.Data["Stay"] = extractedData

	c.Data["refine"] = &RefineSearch{

		Location:      "Dhaka,Bangladesh",
		Checkin_date:  checkIn,
		Checkout_date: checkOut,
	}
	c.TplName = "refine.tpl"
}

func (c *PropertyInfoDetail) GetPropertyDetails() {

	checkIn := c.GetString("t-start")
	checkOut := c.GetString("t-end")

	id := c.Ctx.Input.Param(":countrycode") + "/" + c.Ctx.Input.Param(":idname")
	if id == "" {
		log.Fatal("Data not found")
	}

	// var amenityIcons = map[string]string{
	// 	"star": 			  "fa fa-star,
	// }
	// c.Data["Amenities"] = amenityIcons

	url := "https://booking-com13.p.rapidapi.com/stays/properties/details/" + id + "&checkin_date=" + checkIn +
		"&checkout_date=" + checkOut +
		"&language_code=en-us&currency_code=USD"

	propertyDetailsChan := make(chan PropertyInfoDetails)

	go func() {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making the response")
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading the response")
		}

		var propertyDetailInfo PropertyInfoDetails
		if err := json.Unmarshal(body, &propertyDetailInfo); err != nil {
			fmt.Println("Error parsing JSON response")
		}

		propertyDetailsChan <- propertyDetailInfo
	}()

	extractedData := <-propertyDetailsChan
	c.Data["PropertyDetails"] = extractedData
	c.TplName = "property_detail.tpl"
}

func (c *PropertyDetail) GetPropertyDetail() {
	c.Data["property"] = &PropertyDetail{
		Location: "Ctg, BD",
	}
	c.TplName = "property-detail.tpl"
}
