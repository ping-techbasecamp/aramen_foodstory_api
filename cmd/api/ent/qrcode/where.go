// Code generated by ent, DO NOT EDIT.

package qrcode

import (
	"aramen-api/cmd/api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// QrcodeID applies equality check predicate on the "qrcode_id" field. It's identical to QrcodeIDEQ.
func QrcodeID(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQrcodeID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderStatus applies equality check predicate on the "order_status" field. It's identical to OrderStatusEQ.
func OrderStatus(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderStatus), v))
	})
}

// BranchName applies equality check predicate on the "branch_name" field. It's identical to BranchNameEQ.
func BranchName(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBranchName), v))
	})
}

// OrderChannel applies equality check predicate on the "order_channel" field. It's identical to OrderChannelEQ.
func OrderChannel(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderChannel), v))
	})
}

// TotalPaid applies equality check predicate on the "total_paid" field. It's identical to TotalPaidEQ.
func TotalPaid(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalPaid), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// QrcodeIDEQ applies the EQ predicate on the "qrcode_id" field.
func QrcodeIDEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDNEQ applies the NEQ predicate on the "qrcode_id" field.
func QrcodeIDNEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDIn applies the In predicate on the "qrcode_id" field.
func QrcodeIDIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldQrcodeID), v...))
	})
}

// QrcodeIDNotIn applies the NotIn predicate on the "qrcode_id" field.
func QrcodeIDNotIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldQrcodeID), v...))
	})
}

// QrcodeIDGT applies the GT predicate on the "qrcode_id" field.
func QrcodeIDGT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDGTE applies the GTE predicate on the "qrcode_id" field.
func QrcodeIDGTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDLT applies the LT predicate on the "qrcode_id" field.
func QrcodeIDLT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDLTE applies the LTE predicate on the "qrcode_id" field.
func QrcodeIDLTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDContains applies the Contains predicate on the "qrcode_id" field.
func QrcodeIDContains(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDHasPrefix applies the HasPrefix predicate on the "qrcode_id" field.
func QrcodeIDHasPrefix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDHasSuffix applies the HasSuffix predicate on the "qrcode_id" field.
func QrcodeIDHasSuffix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDEqualFold applies the EqualFold predicate on the "qrcode_id" field.
func QrcodeIDEqualFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldQrcodeID), v))
	})
}

// QrcodeIDContainsFold applies the ContainsFold predicate on the "qrcode_id" field.
func QrcodeIDContainsFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldQrcodeID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDContains applies the Contains predicate on the "order_id" field.
func OrderIDContains(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderID), v))
	})
}

// OrderIDHasPrefix applies the HasPrefix predicate on the "order_id" field.
func OrderIDHasPrefix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderID), v))
	})
}

// OrderIDHasSuffix applies the HasSuffix predicate on the "order_id" field.
func OrderIDHasSuffix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderID), v))
	})
}

// OrderIDEqualFold applies the EqualFold predicate on the "order_id" field.
func OrderIDEqualFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderID), v))
	})
}

// OrderIDContainsFold applies the ContainsFold predicate on the "order_id" field.
func OrderIDContainsFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderID), v))
	})
}

// OrderStatusEQ applies the EQ predicate on the "order_status" field.
func OrderStatusEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusNEQ applies the NEQ predicate on the "order_status" field.
func OrderStatusNEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusIn applies the In predicate on the "order_status" field.
func OrderStatusIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderStatus), v...))
	})
}

// OrderStatusNotIn applies the NotIn predicate on the "order_status" field.
func OrderStatusNotIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderStatus), v...))
	})
}

// OrderStatusGT applies the GT predicate on the "order_status" field.
func OrderStatusGT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusGTE applies the GTE predicate on the "order_status" field.
func OrderStatusGTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusLT applies the LT predicate on the "order_status" field.
func OrderStatusLT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusLTE applies the LTE predicate on the "order_status" field.
func OrderStatusLTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusContains applies the Contains predicate on the "order_status" field.
func OrderStatusContains(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusHasPrefix applies the HasPrefix predicate on the "order_status" field.
func OrderStatusHasPrefix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusHasSuffix applies the HasSuffix predicate on the "order_status" field.
func OrderStatusHasSuffix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusEqualFold applies the EqualFold predicate on the "order_status" field.
func OrderStatusEqualFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderStatus), v))
	})
}

// OrderStatusContainsFold applies the ContainsFold predicate on the "order_status" field.
func OrderStatusContainsFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderStatus), v))
	})
}

// BranchNameEQ applies the EQ predicate on the "branch_name" field.
func BranchNameEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBranchName), v))
	})
}

// BranchNameNEQ applies the NEQ predicate on the "branch_name" field.
func BranchNameNEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBranchName), v))
	})
}

// BranchNameIn applies the In predicate on the "branch_name" field.
func BranchNameIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBranchName), v...))
	})
}

// BranchNameNotIn applies the NotIn predicate on the "branch_name" field.
func BranchNameNotIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBranchName), v...))
	})
}

// BranchNameGT applies the GT predicate on the "branch_name" field.
func BranchNameGT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBranchName), v))
	})
}

// BranchNameGTE applies the GTE predicate on the "branch_name" field.
func BranchNameGTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBranchName), v))
	})
}

// BranchNameLT applies the LT predicate on the "branch_name" field.
func BranchNameLT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBranchName), v))
	})
}

// BranchNameLTE applies the LTE predicate on the "branch_name" field.
func BranchNameLTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBranchName), v))
	})
}

// BranchNameContains applies the Contains predicate on the "branch_name" field.
func BranchNameContains(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBranchName), v))
	})
}

// BranchNameHasPrefix applies the HasPrefix predicate on the "branch_name" field.
func BranchNameHasPrefix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBranchName), v))
	})
}

// BranchNameHasSuffix applies the HasSuffix predicate on the "branch_name" field.
func BranchNameHasSuffix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBranchName), v))
	})
}

// BranchNameEqualFold applies the EqualFold predicate on the "branch_name" field.
func BranchNameEqualFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBranchName), v))
	})
}

// BranchNameContainsFold applies the ContainsFold predicate on the "branch_name" field.
func BranchNameContainsFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBranchName), v))
	})
}

// OrderChannelEQ applies the EQ predicate on the "order_channel" field.
func OrderChannelEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelNEQ applies the NEQ predicate on the "order_channel" field.
func OrderChannelNEQ(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelIn applies the In predicate on the "order_channel" field.
func OrderChannelIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderChannel), v...))
	})
}

// OrderChannelNotIn applies the NotIn predicate on the "order_channel" field.
func OrderChannelNotIn(vs ...string) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderChannel), v...))
	})
}

// OrderChannelGT applies the GT predicate on the "order_channel" field.
func OrderChannelGT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelGTE applies the GTE predicate on the "order_channel" field.
func OrderChannelGTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelLT applies the LT predicate on the "order_channel" field.
func OrderChannelLT(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelLTE applies the LTE predicate on the "order_channel" field.
func OrderChannelLTE(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelContains applies the Contains predicate on the "order_channel" field.
func OrderChannelContains(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelHasPrefix applies the HasPrefix predicate on the "order_channel" field.
func OrderChannelHasPrefix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelHasSuffix applies the HasSuffix predicate on the "order_channel" field.
func OrderChannelHasSuffix(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelEqualFold applies the EqualFold predicate on the "order_channel" field.
func OrderChannelEqualFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderChannel), v))
	})
}

// OrderChannelContainsFold applies the ContainsFold predicate on the "order_channel" field.
func OrderChannelContainsFold(v string) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderChannel), v))
	})
}

// TotalPaidEQ applies the EQ predicate on the "total_paid" field.
func TotalPaidEQ(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalPaid), v))
	})
}

// TotalPaidNEQ applies the NEQ predicate on the "total_paid" field.
func TotalPaidNEQ(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalPaid), v))
	})
}

// TotalPaidIn applies the In predicate on the "total_paid" field.
func TotalPaidIn(vs ...int) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalPaid), v...))
	})
}

// TotalPaidNotIn applies the NotIn predicate on the "total_paid" field.
func TotalPaidNotIn(vs ...int) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalPaid), v...))
	})
}

// TotalPaidGT applies the GT predicate on the "total_paid" field.
func TotalPaidGT(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalPaid), v))
	})
}

// TotalPaidGTE applies the GTE predicate on the "total_paid" field.
func TotalPaidGTE(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalPaid), v))
	})
}

// TotalPaidLT applies the LT predicate on the "total_paid" field.
func TotalPaidLT(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalPaid), v))
	})
}

// TotalPaidLTE applies the LTE predicate on the "total_paid" field.
func TotalPaidLTE(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalPaid), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Qrcode {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qrcode(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Qrcode) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Qrcode) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
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
func Not(p predicate.Qrcode) predicate.Qrcode {
	return predicate.Qrcode(func(s *sql.Selector) {
		p(s.Not())
	})
}