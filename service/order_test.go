package service

import (
	"github.com/zhiqinkuang/easy-ecom/repository"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name    string
		args    repository.Order
		wantErr bool
	}{
		{
			name: "Valid order input",
			args: repository.Order{
				UserID:      "user123",
				OrderID:     "order456",
				OrderStatus: 1,
				UpdateTime:  time.Now(),
				TotalAmount: 299.99,
				List:        []byte(`[{"product_id": "PROD001", "quantity": 2}]`),
				UserMsg:     []byte(`{"msg": "please deliver by tomorrow"}`),
				Note:        "Urgent",
			},
			wantErr: false,
		},
		{
			name: "Invalid order input - missing OrderID",
			args: repository.Order{
				UserID:      "user123",
				OrderStatus: 1,
				TotalAmount: 299.99,
				UpdateTime:  time.Now(),
				List:        []byte(`[{"product_id": "PROD001", "quantity": 2}]`),
				UserMsg:     []byte(`{"msg": "please deliver by tomorrow"}`),
				Note:        "Urgent",
			},
			wantErr: true,
		},
		{
			name: "Invalid order input - negative TotalAmount",
			args: repository.Order{
				UserID:      "user123",
				OrderID:     "order456",
				OrderStatus: 1,
				UpdateTime:  time.Now(),
				TotalAmount: -10.99, // Invalid amount
				List:        []byte(`[{"product_id": "PROD001", "quantity": 2}]`),
				UserMsg:     []byte(`{"msg": "please deliver by tomorrow"}`),
				Note:        "Urgent",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := NewOrderService()

			err := orderService.CreateOrder(&tt.args)

			// Check if error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestUpdateOrderStatus 测试更新订单状态
func TestUpdateOrderStatus(t *testing.T) {
	tests := []struct {
		name    string
		orderId string
		status  int
		wantErr bool
	}{
		{
			name:    "Valid OrderID and status",
			orderId: "order456",
			status:  2, // Let's assume 2 is a valid status
			wantErr: false,
		},
		{
			name:    "Invalid OrderID",
			status:  2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := NewOrderService()

			// Test the UpdateOrderStatus method
			err := orderService.UpdateOrderStatus(tt.orderId, tt.status)

			// Check if the error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateOrderStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrders(t *testing.T) {
	type args struct {
		UserId string
		Status int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				UserId: "user456",
				Status: 1,
			},
			wantErr: false,
		},
		{
			name: "No orders found",
			args: args{

				Status: 1, // Assuming no orders with this status
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := NewOrderService()
			// Test the UpdateOrderStatus method
			_, err := orderService.GetOrders(tt.args.UserId, tt.args.Status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrderById(t *testing.T) {
	type args struct {
		OrderId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid OrderId",
			args: args{
				OrderId: "order456",
			},
			wantErr: false,
		},
		{
			name: "Invalid OrderId",
			args: args{
				OrderId: "invalidOrder123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := NewOrderService()
			// Test the UpdateOrderStatus method
			_, err := orderService.GetOrderById(tt.args.OrderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
