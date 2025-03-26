package ctxs

import (
	"context"
	"github.com/gin-gonic/gin"
	"testing"
)

// test Set and Get with int value
func TestSetAndGet_IntValue(t *testing.T) {
	type IntValue int
	intValue := IntValue(42)

	ctx := context.Background()
	ctx = Set(ctx, intValue)

	retrievedUserID, ok := Get[IntValue](ctx)
	if !ok {
		t.Error("Expected to find intValue in context")
	}
	if retrievedUserID != 42 {
		t.Errorf("Expected intValue to be 42, got %d", retrievedUserID)
	}
}

// test Set and Get with string value
func TestSetAndGet_StringValue(t *testing.T) {
	type StringValue string
	stringValue := StringValue("go")

	ctx := context.Background()
	ctx = Set(ctx, stringValue)

	retrievedUsername, ok := Get[StringValue](ctx)
	if !ok {
		t.Error("Expected to find stringValue in context")
	}
	if retrievedUsername != "go" {
		t.Errorf("Expected stringValue to be Alice, got %s", retrievedUsername)
	}
}

// test Set and Get with struct value
func TestSetAndGet_StructValue(t *testing.T) {
	type StringValue struct {
		Username string
		Status   int
	}

	ctx := context.Background()
	ctx = Set(ctx, StringValue{Username: "go", Status: 1})

	retrievedUsername, ok := Get[StringValue](ctx)
	if !ok {
		t.Error("Expected to find stringValue in context")
	}
	if retrievedUsername.Username != "go" || retrievedUsername.Status != 1 {
		t.Errorf("Expected stringValue to be gp, got %v", retrievedUsername)
	}
}

func TestSetAndGet_ByGinContext(t *testing.T) {

	type StringValue string
	stringValue := StringValue("go")

	ctx := &gin.Context{}
	Set(ctx, stringValue)

	func(ctx context.Context) {
		retrievedUsername, ok := Get[StringValue](ctx)
		if !ok {
			t.Error("Expected to find stringValue in context")
		}
		if retrievedUsername != "go" {
			t.Errorf("Expected stringValue to be go, got %s", retrievedUsername)
		}
	}(ctx)

}

// test Get with not found key
func TestGetNotFound(t *testing.T) {
	ctx := context.Background()

	type StringValue string
	_, ok := Get[StringValue](ctx)
	if ok {
		t.Error("Expected not to find username in context")
	}
}
