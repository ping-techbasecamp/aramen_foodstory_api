// Code generated by ent, DO NOT EDIT.

package branches

import (
	"aramen-api/cmd/api/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ShortDescription applies equality check predicate on the "short_description" field. It's identical to ShortDescriptionEQ.
func ShortDescription(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShortDescription), v))
	})
}

// FullDescription applies equality check predicate on the "full_description" field. It's identical to FullDescriptionEQ.
func FullDescription(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullDescription), v))
	})
}

// Telephone applies equality check predicate on the "telephone" field. It's identical to TelephoneEQ.
func Telephone(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// GooleMapURL applies equality check predicate on the "goole_map_url" field. It's identical to GooleMapURLEQ.
func GooleMapURL(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGooleMapURL), v))
	})
}

// DineIn applies equality check predicate on the "dine_in" field. It's identical to DineInEQ.
func DineIn(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDineIn), v))
	})
}

// Delivery applies equality check predicate on the "delivery" field. It's identical to DeliveryEQ.
func Delivery(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDelivery), v))
	})
}

// TakeAway applies equality check predicate on the "take_away" field. It's identical to TakeAwayEQ.
func TakeAway(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTakeAway), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ShortDescriptionEQ applies the EQ predicate on the "short_description" field.
func ShortDescriptionEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionNEQ applies the NEQ predicate on the "short_description" field.
func ShortDescriptionNEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionIn applies the In predicate on the "short_description" field.
func ShortDescriptionIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldShortDescription), v...))
	})
}

// ShortDescriptionNotIn applies the NotIn predicate on the "short_description" field.
func ShortDescriptionNotIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldShortDescription), v...))
	})
}

// ShortDescriptionGT applies the GT predicate on the "short_description" field.
func ShortDescriptionGT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionGTE applies the GTE predicate on the "short_description" field.
func ShortDescriptionGTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionLT applies the LT predicate on the "short_description" field.
func ShortDescriptionLT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionLTE applies the LTE predicate on the "short_description" field.
func ShortDescriptionLTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionContains applies the Contains predicate on the "short_description" field.
func ShortDescriptionContains(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionHasPrefix applies the HasPrefix predicate on the "short_description" field.
func ShortDescriptionHasPrefix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionHasSuffix applies the HasSuffix predicate on the "short_description" field.
func ShortDescriptionHasSuffix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionEqualFold applies the EqualFold predicate on the "short_description" field.
func ShortDescriptionEqualFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldShortDescription), v))
	})
}

// ShortDescriptionContainsFold applies the ContainsFold predicate on the "short_description" field.
func ShortDescriptionContainsFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldShortDescription), v))
	})
}

// FullDescriptionEQ applies the EQ predicate on the "full_description" field.
func FullDescriptionEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionNEQ applies the NEQ predicate on the "full_description" field.
func FullDescriptionNEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionIn applies the In predicate on the "full_description" field.
func FullDescriptionIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFullDescription), v...))
	})
}

// FullDescriptionNotIn applies the NotIn predicate on the "full_description" field.
func FullDescriptionNotIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFullDescription), v...))
	})
}

// FullDescriptionGT applies the GT predicate on the "full_description" field.
func FullDescriptionGT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionGTE applies the GTE predicate on the "full_description" field.
func FullDescriptionGTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionLT applies the LT predicate on the "full_description" field.
func FullDescriptionLT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionLTE applies the LTE predicate on the "full_description" field.
func FullDescriptionLTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionContains applies the Contains predicate on the "full_description" field.
func FullDescriptionContains(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionHasPrefix applies the HasPrefix predicate on the "full_description" field.
func FullDescriptionHasPrefix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionHasSuffix applies the HasSuffix predicate on the "full_description" field.
func FullDescriptionHasSuffix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionEqualFold applies the EqualFold predicate on the "full_description" field.
func FullDescriptionEqualFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFullDescription), v))
	})
}

// FullDescriptionContainsFold applies the ContainsFold predicate on the "full_description" field.
func FullDescriptionContainsFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFullDescription), v))
	})
}

// TelephoneEQ applies the EQ predicate on the "telephone" field.
func TelephoneEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// TelephoneNEQ applies the NEQ predicate on the "telephone" field.
func TelephoneNEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelephone), v))
	})
}

// TelephoneIn applies the In predicate on the "telephone" field.
func TelephoneIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTelephone), v...))
	})
}

// TelephoneNotIn applies the NotIn predicate on the "telephone" field.
func TelephoneNotIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTelephone), v...))
	})
}

// TelephoneGT applies the GT predicate on the "telephone" field.
func TelephoneGT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelephone), v))
	})
}

// TelephoneGTE applies the GTE predicate on the "telephone" field.
func TelephoneGTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelephone), v))
	})
}

// TelephoneLT applies the LT predicate on the "telephone" field.
func TelephoneLT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelephone), v))
	})
}

// TelephoneLTE applies the LTE predicate on the "telephone" field.
func TelephoneLTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelephone), v))
	})
}

// TelephoneContains applies the Contains predicate on the "telephone" field.
func TelephoneContains(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTelephone), v))
	})
}

// TelephoneHasPrefix applies the HasPrefix predicate on the "telephone" field.
func TelephoneHasPrefix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTelephone), v))
	})
}

// TelephoneHasSuffix applies the HasSuffix predicate on the "telephone" field.
func TelephoneHasSuffix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTelephone), v))
	})
}

// TelephoneEqualFold applies the EqualFold predicate on the "telephone" field.
func TelephoneEqualFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTelephone), v))
	})
}

// TelephoneContainsFold applies the ContainsFold predicate on the "telephone" field.
func TelephoneContainsFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTelephone), v))
	})
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatitude), v))
	})
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...int) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLatitude), v...))
	})
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...int) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLatitude), v...))
	})
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatitude), v))
	})
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatitude), v))
	})
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatitude), v))
	})
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatitude), v))
	})
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongitude), v))
	})
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...int) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLongitude), v...))
	})
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...int) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLongitude), v...))
	})
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongitude), v))
	})
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongitude), v))
	})
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongitude), v))
	})
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v int) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongitude), v))
	})
}

// GooleMapURLEQ applies the EQ predicate on the "goole_map_url" field.
func GooleMapURLEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLNEQ applies the NEQ predicate on the "goole_map_url" field.
func GooleMapURLNEQ(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLIn applies the In predicate on the "goole_map_url" field.
func GooleMapURLIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGooleMapURL), v...))
	})
}

// GooleMapURLNotIn applies the NotIn predicate on the "goole_map_url" field.
func GooleMapURLNotIn(vs ...string) predicate.Branches {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGooleMapURL), v...))
	})
}

// GooleMapURLGT applies the GT predicate on the "goole_map_url" field.
func GooleMapURLGT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLGTE applies the GTE predicate on the "goole_map_url" field.
func GooleMapURLGTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLLT applies the LT predicate on the "goole_map_url" field.
func GooleMapURLLT(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLLTE applies the LTE predicate on the "goole_map_url" field.
func GooleMapURLLTE(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLContains applies the Contains predicate on the "goole_map_url" field.
func GooleMapURLContains(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLHasPrefix applies the HasPrefix predicate on the "goole_map_url" field.
func GooleMapURLHasPrefix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLHasSuffix applies the HasSuffix predicate on the "goole_map_url" field.
func GooleMapURLHasSuffix(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLEqualFold applies the EqualFold predicate on the "goole_map_url" field.
func GooleMapURLEqualFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGooleMapURL), v))
	})
}

// GooleMapURLContainsFold applies the ContainsFold predicate on the "goole_map_url" field.
func GooleMapURLContainsFold(v string) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGooleMapURL), v))
	})
}

// DineInEQ applies the EQ predicate on the "dine_in" field.
func DineInEQ(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDineIn), v))
	})
}

// DineInNEQ applies the NEQ predicate on the "dine_in" field.
func DineInNEQ(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDineIn), v))
	})
}

// DeliveryEQ applies the EQ predicate on the "delivery" field.
func DeliveryEQ(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDelivery), v))
	})
}

// DeliveryNEQ applies the NEQ predicate on the "delivery" field.
func DeliveryNEQ(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDelivery), v))
	})
}

// TakeAwayEQ applies the EQ predicate on the "take_away" field.
func TakeAwayEQ(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTakeAway), v))
	})
}

// TakeAwayNEQ applies the NEQ predicate on the "take_away" field.
func TakeAwayNEQ(v bool) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTakeAway), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Branches) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Branches) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Branches) predicate.Branches {
	return predicate.Branches(func(s *sql.Selector) {
		p(s.Not())
	})
}
