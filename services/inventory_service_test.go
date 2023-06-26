package services

import (
	"example.com/shopping/domain"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_inventoryService_AddItems(t *testing.T) {
	type fields struct {
		lock      *sync.Mutex
		lastId    int
		items     map[int]domain.Item
		itemLocks *sync.Map
	}
	type args struct {
		items []domain.Item
	}
	type result struct {
		wantErr   bool
		itemCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    result
	}{
		{
			name: "when items are valid",
			fields: fields{
				lock:      &sync.Mutex{},
				lastId:    0,
				items:     make(map[int]domain.Item),
				itemLocks: &sync.Map{},
			},
			args: args{
				items: []domain.Item{{Name: "apple", Description: "a fruit"}, {Name: "cat", Description: "a pet"}},
			},
			res: result{
				wantErr:   false,
				itemCount: 2,
			},
		},
		{
			name: "when at least 1 item is invalid",
			fields: fields{
				lock:      &sync.Mutex{},
				lastId:    0,
				items:     make(map[int]domain.Item),
				itemLocks: &sync.Map{},
			},
			args: args{
				items: []domain.Item{{Name: "", Description: "a fruit"}, {Name: "cat", Description: "a pet"}},
			},
			res: result{
				wantErr:   true,
				itemCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iS := &inventoryService{
				lock:      tt.fields.lock,
				lastId:    tt.fields.lastId,
				items:     tt.fields.items,
				itemLocks: tt.fields.itemLocks,
			}
			if err := iS.AddItems(tt.args.items); (err != nil) != tt.res.wantErr {
				t.Errorf("AddItems() error = %v, wantErr %v", err, tt.res.wantErr)
			}
			assert.Equal(t, tt.res.itemCount, len(iS.items))
		})
	}
}

func Test_inventoryService_GetAllItems(t *testing.T) {
	type fields struct {
		lock      *sync.Mutex
		lastId    int
		items     map[int]domain.Item
		itemLocks *sync.Map
	}
	tests := []struct {
		name   string
		fields fields
		want   []domain.Item
		args   []domain.Item
	}{
		{
			name: "when there are no items",
			args: []domain.Item{},
			fields: fields{
				lock:      &sync.Mutex{},
				lastId:    0,
				items:     map[int]domain.Item{},
				itemLocks: &sync.Map{},
			},
			want: []domain.Item{},
		},
		{
			name: "when there are no items",
			args: []domain.Item{{Id: 1, Name: "name", Description: ""}},
			fields: fields{
				lock:      &sync.Mutex{},
				lastId:    0,
				items:     map[int]domain.Item{1: {Id: 1, Name: "name", Description: ""}},
				itemLocks: &sync.Map{},
			},
			want: []domain.Item{{Id: 1, Name: "name", Description: "", Available: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iS := &inventoryService{
				lock:      tt.fields.lock,
				lastId:    tt.fields.lastId,
				items:     tt.fields.items,
				itemLocks: tt.fields.itemLocks,
			}
			err := iS.AddItems(tt.args)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, iS.GetAllItems(), "GetAllItems()")
		})
	}
}

func Test_inventoryService_MarkUnavailable(t *testing.T) {
	items := []domain.Item{{Name: "apple", Description: "a fruit"}, {Name: "cat", Description: "a pet"}}
	type fields struct {
		lock      *sync.Mutex
		lastId    int
		items     map[int]domain.Item
		itemLocks *sync.Map
	}
	type args struct {
		itemId int
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "when item is not available",
			fields: fields{
				lock:      &sync.Mutex{},
				lastId:    0,
				items:     map[int]domain.Item{},
				itemLocks: &sync.Map{},
			},
			args: args{
				10,
				2,
			},
			wantErr: domain.ItemNotFoundError,
		},
		{
			name: "when item is not available",
			fields: fields{
				lock:      &sync.Mutex{},
				lastId:    0,
				items:     map[int]domain.Item{},
				itemLocks: &sync.Map{},
			},
			args: args{
				10,
				2,
			},
			wantErr: domain.ItemNotFoundError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iS := &inventoryService{
				lock:      tt.fields.lock,
				lastId:    tt.fields.lastId,
				items:     tt.fields.items,
				itemLocks: tt.fields.itemLocks,
			}
			err := iS.AddItems(items)
			assert.Nil(t, err)
			assert.Equal(t, iS.MarkUnavailable(tt.args.itemId, tt.args.userId), tt.wantErr)
		})
	}
}
