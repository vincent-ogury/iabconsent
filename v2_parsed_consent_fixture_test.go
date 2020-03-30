package iabconsent_test

import (
	"time"

	"github.com/LiveRamp/iabconsent"
)

var v2TestTime = time.Unix(1583436280, 9*nsPerDs).UTC()

// Consent fixtures have been generated using the tool at https://iabtcf.com/encode.
// To inspect a consent string used in the test u can enter it into: https://iabtcf.com/#/decode.

var v2ConsentFixtures = map[string]*iabconsent.V2ParsedConsent{
	// COvzTO5OvzTO5B7ABCENAPCYAKdAADkAAIqIFhwBAAGAAXAFGAsMAhYAgAMAAegBYAEKAAA.IFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUw.QE5QAwCvgHyATkA
	"COvzTO5OvzTO5B7ABCENAPCYAKdAADkAAIqIFhwBAAGAAXAFGAsMAhYAgAMAAegBYAEKAAA.IFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUw.QE5QAwCvgHyATkA": {
		Version:              2,
		Created:              v2TestTime,
		LastUpdated:          v2TestTime,
		CMPID:                123,
		CMPVersion:           1,
		ConsentScreen:        2,
		ConsentLanguage:      "EN",
		VendorListVersion:    15,
		TCFPolicyVersion:     2,
		IsServiceSpecific:    false,
		UseNonStandardStacks: true,
		SpecialFeaturesOptIn: map[int]bool{1: true},
		PurposesConsent: map[int]bool{
			1:  true,
			3:  true,
			6:  true,
			7:  true,
			8:  true,
			10: true,
		},
		PurposesLITransparency: map[int]bool{
			3: true,
			4: true,
			5: true,
			8: true,
		},
		PurposeOneTreatment:      true,
		PublisherCC:              "FR",
		MaxConsentVendorID:       707,
		IsConsentRangeEncoding:   true,
		ConsentedVendors:         nil,
		NumConsentEntries:        4,
		ConsentedVendorsRange:    []*iabconsent.RangeEntry{
			{StartVendorID:12, EndVendorID:12},
			{StartVendorID:23, EndVendorID:23},
			{StartVendorID:163, EndVendorID:163},
			{StartVendorID:707, EndVendorID:707},

		},
		MaxInterestsVendorID:     133,
		IsInterestsRangeEncoding: true,
		InterestsVendors:         nil,
		NumInterestsEntries:      4,
		InterestsVendorsRange:    []*iabconsent.RangeEntry{
			{StartVendorID:48, EndVendorID:48},
			{StartVendorID:61, EndVendorID:61},
			{StartVendorID:88, EndVendorID:88},
			{StartVendorID:133, EndVendorID:133},

		},
		NumPubRestrictions:       0,
		PubRestrictionEntries:    make([]*iabconsent.PubRestrictionEntry, 0),
		OOBDisclosedVendors:      &iabconsent.OOBVendorList{
			SegmentType:1,
			MaxVendorID:720,
			IsRangeEncoding:false,
			Vendors: map[int]bool{
				2:true,
				6:true,
				8:true,
				12:true,
				18:true,
				23:true,
				37:true,
				42:true,
				47:true,
				48:true,
				53:true,
				61:true,
				65:true,
				66:true,
				72:true,
				88:true,
				98:true,
				127:true,
				128:true,
				129:true,
				133:true,
				153:true,
				163:true,
				192:true,
				205:true,
				215:true,
				224:true,
				243:true,
				248:true,
				281:true,
				294:true,
				304:true,
				350:true,
				351:true,
				358:true,
				371:true,
				422:true,
				424:true,
				440:true,
				447:true,
				467:true,
				486:true,
				498:true,
				502:true,
				512:true,
				516:true,
				553:true,
				556:true,
				571:true,
				587:true,
				612:true,
				613:true,
				618:true,
				626:true,
				648:true,
				653:true,
				656:true,
				657:true,
				665:true,
				676:true,
				681:true,
				683:true,
				684:true,
				686:true,
				687:true,
				688:true,
				690:true,
				691:true,
				694:true,
				702:true,
				703:true,
				707:true,
				708:true,
				711:true,
				712:true,
				714:true,
				716:true,
				719:true,
				720:true,
			},
		},
		OOBAllowedVendors:        &iabconsent.OOBVendorList{
			SegmentType:     2,
			MaxVendorID:     626,
			IsRangeEncoding: true,
			NumEntries:      3,
			VendorEntries: []*iabconsent.RangeEntry{
				{StartVendorID:351, EndVendorID:351},
				{StartVendorID:498, EndVendorID:498},
				{StartVendorID:626, EndVendorID:626},
			},
		},
		PublisherTCEntry:         nil,
	},
	// COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA
	// COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFADBQAQA9hAAAcAA
	"COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFADBQAQA9hAAAcAA": {
		Version:                2,
		Created: v2TestTime,
		LastUpdated: v2TestTime,
		CMPID:                  81,
		CMPVersion:             0,
		ConsentScreen:          0,
		ConsentLanguage:        "EN",
		VendorListVersion:      15,
		TCFPolicyVersion:       2,
		IsServiceSpecific:      true,
		UseNonStandardStacks:   false,
		SpecialFeaturesOptIn:   map[int]bool{1: true},
		PurposesConsent:        map[int]bool{1: true, 3: true, 4: true, 7: true},
		PurposesLITransparency: map[int]bool{3: true, 4: true, 5: true},
		PurposeOneTreatment:    false,
		PublisherCC:            "AA",
		MaxConsentVendorID:     61,
		IsConsentRangeEncoding: true,
		NumConsentEntries:      3,
		ConsentedVendorsRange:  []*iabconsent.RangeEntry{
			{StartVendorID:2, EndVendorID:2},
			{StartVendorID:37, EndVendorID:37},
			{StartVendorID:61, EndVendorID:61},
		},
		MaxInterestsVendorID:8,
		IsInterestsRangeEncoding:false,
		InterestsVendors: map[int]bool{
			2:true,
			6:true,
			8:true,
		},
		NumPubRestrictions:       3,
		PubRestrictionEntries:    []*iabconsent.PubRestrictionEntry{
			{
				PurposeID: 1,
				RestrictionType: iabconsent.RequireConsent,
				NumEntries: 1,
				RestrictionsRange: []*iabconsent.RangeEntry{
					{
						StartVendorID: 123,
						EndVendorID: 123,
					},
				},
			},
			{
				PurposeID: 2,
				RestrictionType: iabconsent.PurposeFlatlyNotAllowed,
				RestrictionsRange: []*iabconsent.RangeEntry{},
			},
			{
				PurposeID: 3,
				RestrictionType: iabconsent.RequireLegitimateInterest,
				RestrictionsRange: []*iabconsent.RangeEntry{},
			},
		},
	},
	// COvzTO5OvzTO5BZAFMENAPCgAAAAAAAAAAwIFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUwLQIoghAAQhhARggACAIAAAAcQAAEAQAAAAgAQBAIAAEIAAAABAAgCAAAAAAAMCABAAAAAAAAKAAIEAABAAAgAiAIgAAAAASAAQABAAAAwgIAAAhMBACFuyAxmpgAA
	"COvzTO5OvzTO5BZAFMENAPCgAAAAAAAAAAwIFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUwLQIoghAAQhhARggACAIAAAAcQAAEAQAAAAgAQBAIAAEIAAAABAAgCAAAAAAAMCABAAAAAAAAKAAIEAABAAAgAiAIgAAAAASAAQABAAAAwgIAAAhMBACFuyAxmpgAA": {
		Version:2,
		Created: v2TestTime,
		LastUpdated: v2TestTime,
		CMPID:89,
		CMPVersion:5,
		ConsentScreen:12,
		ConsentLanguage: "EN",
		VendorListVersion:15,
		TCFPolicyVersion:2,
		IsServiceSpecific:true,
		UseNonStandardStacks:false,
		SpecialFeaturesOptIn:   map[int]bool{},
		PurposesConsent:map[int]bool{},
		PurposesLITransparency:map[int]bool{},
		PurposeOneTreatment:false,
		PublisherCC: "GB",
		MaxConsentVendorID:720,
		IsConsentRangeEncoding:false,
		ConsentedVendors:map[int]bool{
			2:true,
			6:true,
			8:true,
			12:true,
			18:true,
			23:true,
			37:true,
			42:true,
			47:true,
			48:true,
			53:true,
			61:true,
			65:true,
			66:true,
			72:true,
			88:true,
			98:true,
			127:true,
			128:true,
			129:true,
			133:true,
			153:true,
			163:true,
			192:true,
			205:true,
			215:true,
			224:true,
			243:true,
			248:true,
			281:true,
			294:true,
			304:true,
			350:true,
			351:true,
			358:true,
			371:true,
			422:true,
			424:true,
			440:true,
			447:true,
			467:true,
			486:true,
			498:true,
			502:true,
			512:true,
			516:true,
			553:true,
			556:true,
			571:true,
			587:true,
			612:true,
			613:true,
			618:true,
			626:true,
			648:true,
			653:true,
			656:true,
			657:true,
			665:true,
			676:true,
			681:true,
			683:true,
			684:true,
			686:true,
			687:true,
			688:true,
			690:true,
			691:true,
			694:true,
			702:true,
			703:true,
			707:true,
			708:true,
			711:true,
			712:true,
			714:true,
			716:true,
			719:true,
			720:true,
		},
		MaxInterestsVendorID:720,
		IsInterestsRangeEncoding:false,
		InterestsVendors:map[int]bool{
			2:true,
			6:true,
			8:true,
			12:true,
			18:true,
			23:true,
			37:true,
			42:true,
			47:true,
			48:true,
			53:true,
			61:true,
			65:true,
			66:true,
			72:true,
			88:true,
			98:true,
			127:true,
			128:true,
			129:true,
			133:true,
			153:true,
			163:true,
			192:true,
			205:true,
			215:true,
			224:true,
			243:true,
			248:true,
			281:true,
			294:true,
			304:true,
			350:true,
			351:true,
			358:true,
			371:true,
			422:true,
			424:true,
			440:true,
			447:true,
			467:true,
			486:true,
			498:true,
			502:true,
			512:true,
			516:true,
			553:true,
			556:true,
			571:true,
			587:true,
			612:true,
			613:true,
			618:true,
			626:true,
			648:true,
			653:true,
			656:true,
			657:true,
			665:true,
			676:true,
			681:true,
			683:true,
			684:true,
			686:true,
			687:true,
			688:true,
			690:true,
			691:true,
			694:true,
			702:true,
			703:true,
			707:true,
			708:true,
			711:true,
			712:true,
			714:true,
			716:true,
			719:true,
			720:true,
		},
		NumPubRestrictions:       0,
		PubRestrictionEntries:    make([]*iabconsent.PubRestrictionEntry, 0),
	},
}