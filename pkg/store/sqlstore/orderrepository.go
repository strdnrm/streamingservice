package sqlstore

import (
	"context"
	"database/sql"
	"reflect"
	"streamingservice/pkg/model"

	"github.com/jackskj/carta"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {
	tx, err := r.store.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx, `
	INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
		customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES (:order_uid, :track_number, :entry, :locale, :internal_signature,
		:customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard)
		`, order)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NamedExecContext(ctx, `
	INSERT INTO delivery (order_id, name, phone, zip, city, address, region, email)
	VALUES (:order_id, :name, :phone, :zip, :city, :address, :region, :email)
		`, order.Delivery)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NamedExecContext(ctx, `
	INSERT INTO payment (order_id, transaction, request_id, currency, provider, amount, payment_dt,
		bank, delivery_cost, goods_total, custom_fee)
	VALUES (:order_id, :transaction, :request_id, :currency, :provider, :amount, :payment_dt,
		:bank, :delivery_cost, :goods_total, :custom_fee)
		`, order.Payment)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, product := range order.Items {
		_, err = tx.NamedExecContext(ctx, `
		INSERT INTO order_items (order_id, chrt_id, item_track_number, price, rid, item_name,
			 sale, size, total_price, nm_id, brand, status)
		VALUES (:order_id, :chrt_id, :item_track_number, :price, :rid, :item_name,
			:sale, :size, :total_price, :nm_id, :brand, :status)`,
			product)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetById(ctx context.Context, id string) (*model.Order, error) {
	query := `
    SELECT
		o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id,
		o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
		d.name, d.phone, d.zip, d.city, d.address, d.region, d.email, p.transaction,  
		p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank,
		p.delivery_cost, p.goods_total, p.custom_fee, i.chrt_id, i.item_track_number,
		i.price, i.rid, i.item_name, i.sale, i.size, i.total_price, i.nm_id, i.brand, i.status
	FROM orders o
	LEFT JOIN delivery d ON o.order_uid = d.order_id
	LEFT JOIN payment p ON o.order_uid = p.order_id
	LEFT JOIN order_items i ON o.order_uid = i.order_id
	WHERE o.order_uid = $1
	`

	rows, err := r.store.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	order := model.Order{}

	if err := carta.Map(rows, &order); err != nil {
		return nil, err
	}

	if reflect.DeepEqual(order, model.Order{}) {
		return nil, sql.ErrNoRows
	}

	return &order, nil
}

func (r *OrderRepository) GetAll(ctx context.Context) (*[]model.Order, error) {
	query := `
    SELECT
	o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id,
	o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
	d.name, d.phone, d.zip, d.city, d.address, d.region, d.email, p.transaction,  
	p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank,
	p.delivery_cost, p.goods_total, p.custom_fee, i.chrt_id, i.item_track_number,
	i.price, i.rid, i.item_name, i.sale, i.size, i.total_price, i.nm_id, i.brand, i.status
	FROM orders o
	LEFT JOIN delivery d ON o.order_uid = d.order_id
	LEFT JOIN payment p ON o.order_uid = p.order_id
	LEFT JOIN order_items i ON o.order_uid = i.order_id
	`

	rows, err := r.store.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	orders := []model.Order{}
	if err := carta.Map(rows, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}
