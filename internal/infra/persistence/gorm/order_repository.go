package gorm

import (
	"pizza-tracker-go/internal/domain/entity"
	"pizza-tracker-go/internal/domain/repository"
)

type OrderRepository struct {
	db *gormDB
}

func NewOrderRepository(db *gormDB) repository.OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *entity.Order) (*entity.Order, error) {
	m := toOrderModel(order)
	if err := r.db.DB.Create(m).Error; err != nil {
		return nil, err
	}
	return toOrderEntity(m), nil
}

func (r *OrderRepository) GetByID(id string) (*entity.Order, error) {
	var m OrderModel
	if err := r.db.DB.Preload("Items").First(&m, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toOrderEntity(&m), nil
}

func (r *OrderRepository) GetAll() ([]entity.Order, error) {
	var models []OrderModel
	if err := r.db.DB.Preload("Items").Order("created_at desc").Find(&models).Error; err != nil {
		return nil, err
	}
	out := make([]entity.Order, len(models))
	for i := range models {
		out[i] = *toOrderEntity(&models[i])
	}
	return out, nil
}

func (r *OrderRepository) UpdateStatus(id, status string) error {
	return r.db.DB.Model(&OrderModel{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepository) Delete(id string) error {
	return r.db.DB.Select("Items").Delete(&OrderModel{ID: id}).Error
}

func toOrderModel(o *entity.Order) *OrderModel {
	items := make([]OrderItemModel, len(o.Items))
	for i := range o.Items {
		items[i] = OrderItemModel{
			ID:           o.Items[i].ID,
			OrderID:      o.ID,
			Size:         o.Items[i].Size,
			Pizza:        o.Items[i].Pizza,
			Instructions: o.Items[i].Instructions,
		}
	}
	return &OrderModel{
		ID:           o.ID,
		Status:       o.Status,
		CustomerName: o.CustomerName,
		Phone:        o.Phone,
		Address:      o.Address,
		Items:        items,
		CreatedAt:    o.CreatedAt,
	}
}

func toOrderEntity(m *OrderModel) *entity.Order {
	items := make([]entity.OrderItem, len(m.Items))
	for i := range m.Items {
		items[i] = entity.OrderItem{
			ID:           m.Items[i].ID,
			OrderID:      m.Items[i].OrderID,
			Size:         m.Items[i].Size,
			Pizza:        m.Items[i].Pizza,
			Instructions: m.Items[i].Instructions,
		}
	}
	return &entity.Order{
		ID:           m.ID,
		Status:       m.Status,
		CustomerName: m.CustomerName,
		Phone:        m.Phone,
		Address:      m.Address,
		Items:        items,
		CreatedAt:    m.CreatedAt,
	}
}
