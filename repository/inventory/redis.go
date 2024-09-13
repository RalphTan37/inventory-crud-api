package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/RalphTan37/inventory-crud-api/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

// generating item key function
func itemIDKey(id uuid.UUID) string {
	return fmt.Sprintf("item: %s", id.String())
}

// insert item to inventory
func (r *RedisRepo) Insert(ctx context.Context, inventory model.Inventory) error {
	data, err := json.Marshal(inventory) //encoding inventory struct to JSON

	if err != nil {
		return fmt.Errorf("failed to encode inventory: %w", err)
	}

	key := itemIDKey(inventory.ItemID) //generates ID

	txn := r.Client.TxPipeline() //creates Go-Redis transcation pipeline

	/*
		marshal method returns byte array, cast to string
		NX = not exist - client will not override any data that already exists, instead return an error
	*/
	res := txn.SetNX(ctx, key, string(data), 0) //sets value if key DNE

	if err := res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to set: %w", err)
	}

	// item key to new set
	if err := txn.SAdd(ctx, "items", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add to items set: %w", err)
	}

	// not left in a partial state
	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

// Go-Redis Error for an item not existing
var ErrDNE = errors.New("item does not exist")

// retrieves an item from Redis by its ID
func (r *RedisRepo) FindByID(ctx context.Context, ItemID uuid.UUID) (model.Inventory, error) {
	key := itemIDKey(ItemID) //generates ID

	// handles any errors recieved
	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return model.Inventory{}, ErrDNE
	} else if err != nil {
		return model.Inventory{}, fmt.Errorf("get item: %w", err)
	}

	//convert recieved json data into model type
	var item model.Inventory
	err = json.Unmarshal([]byte(value), &item)
	if err != nil {
		return model.Inventory{}, fmt.Errorf("failed to decode item json: %w", err)
	}

	return item, nil //returns item instance
}

// delete item in inventory
func (r *RedisRepo) DeleteByID(ctx context.Context, ItemID uuid.UUID) error {
	key := itemIDKey(ItemID) //generates ID

	txn := r.Client.TxPipeline() //creates Redis transcation pipeline

	//delete cmd to remove a key from Go-Redis
	err := txn.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrDNE
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("get item : %w", err)
	}

	//remove key from the items set
	if err := txn.SRem(ctx, "items", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to remove from items set: %w", err)
	}

	//execute Redis transaction pipeline & handles errors during exec
	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

// update item in inventory
func (r *RedisRepo) Update(ctx context.Context, inventory model.Inventory) error {
	data, err := json.Marshal(inventory) //update existing records
	if err != nil {
		return fmt.Errorf("failed to encode inventory: %w", err)
	}

	key := itemIDKey(inventory.ItemID) //generates ID

	err = r.Client.SetXX(ctx, key, string(data), 0).Err() //sets value if it alreadye exists
	if errors.Is(err, redis.Nil) {
		return ErrDNE
	} else if err != nil {
		return fmt.Errorf("set item: %w", err)
	}
	return nil
}

// pagnation params
type FindAllPage struct {
	Size   uint
	Offset uint
}

// defines items & next cursor
type FindResult struct {
	Items  []model.Inventory
	Cursor uint64
}

// find all items method
func (r *RedisRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	res := r.Client.SScan(ctx, "items", uint64(page.Offset), "*", int64(page.Size)) //scanning set
	keys, cursor, err := res.Result()                                               //captures result value
	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get item ids: %w", err)
	}

	//checks key size
	if len(keys) == 0 {
		return FindResult{
			Items: []model.Inventory{},
		}, nil
	}

	//pass all keys to a single Go-Redis call
	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get items: %w", err)
	}

	//create an inventory slice w/ same len as resulting slice
	items := make([]model.Inventory, len(xs))

	//iterates over each element in result array and casts to str
	for i, x := range xs {
		x := x.(string)
		var item model.Inventory

		err := json.Unmarshal([]byte(x), &item)
		if err != nil {
			return FindResult{}, fmt.Errorf("failed to decode item json: %w", err)
		}
		items[i] = item
	}

	//result of Go-Redis scanning operation
	return FindResult{
		Items:  items,
		Cursor: cursor,
	}, nil
}
