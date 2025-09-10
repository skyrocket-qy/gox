package redisx_test

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/skyrocket-qy/gox/redisx"
)

func TestCuckooFilter(t *testing.T) {
	ctx := context.Background()
	rdb, mock := redismock.NewClientMock()

	key := "test_cuckoo_filter"
	value := "test_value"
	value2 := "test_value_2"

	// 1. Reserve the filter
	mock.ExpectDo("CF.RESERVE", key, int64(1000)).SetVal("OK")

	err := redisx.CuckooFilterReserve(ctx, rdb, key, 1000)
	if err != nil {
		t.Fatalf("CuckooFilterReserve failed: %v", err)
	}

	// 2. Add a value to the filter
	mock.ExpectDo("CF.ADD", key, value).SetVal(true)

	added, err := redisx.CuckooFilterAdd(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("CuckooFilterAdd failed: %v", err)
	}

	if !added {
		t.Errorf("CuckooFilterAdd should return true for a new value")
	}

	// 3. Check if the value exists
	mock.ExpectDo("CF.EXISTS", key, value).SetVal(true)

	exists, err := redisx.CuckooFilterExists(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("CuckooFilterExists failed: %v", err)
	}

	if !exists {
		t.Errorf("CuckooFilterExists should return true for a value that was added")
	}

	// 4. Check for a non-existent value
	mock.ExpectDo("CF.EXISTS", key, "non_existent_value").SetVal(false)

	exists, err = redisx.CuckooFilterExists(ctx, rdb, key, "non_existent_value")
	if err != nil {
		t.Fatalf("CuckooFilterExists failed for non-existent value: %v", err)
	}

	if exists {
		t.Errorf("CuckooFilterExists should return false for a value that was not added")
	}

	// 5. Add a value that already exists
	mock.ExpectDo("CF.ADD", key, value).SetVal(false)

	added, err = redisx.CuckooFilterAdd(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("CuckooFilterAdd failed for existing value: %v", err)
	}

	if added {
		t.Errorf("CuckooFilterAdd should return false for an existing value")
	}

	// 6. Count the items
	mock.ExpectDo("CF.COUNT", key, value).SetVal(int64(1))

	count, err := redisx.CuckooFilterCount(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("CuckooFilterCount failed: %v", err)
	}

	if count != 1 {
		t.Errorf("CuckooFilterCount should return 1, but got %d", count)
	}

	// 7. Add a new value with AddNX
	mock.ExpectDo("CF.ADDNX", key, value2).SetVal(true)

	added, err = redisx.CuckooFilterAddNX(ctx, rdb, key, value2)
	if err != nil {
		t.Fatalf("CuckooFilterAddNX failed: %v", err)
	}

	if !added {
		t.Errorf("CuckooFilterAddNX should return true for a new value")
	}

	// 8. Add an existing value with AddNX
	mock.ExpectDo("CF.ADDNX", key, value2).SetVal(false)

	added, err = redisx.CuckooFilterAddNX(ctx, rdb, key, value2)
	if err != nil {
		t.Fatalf("CuckooFilterAddNX failed for existing value: %v", err)
	}

	if added {
		t.Errorf("CuckooFilterAddNX should return false for an existing value")
	}

	// 9. Delete a value
	mock.ExpectDo("CF.DEL", key, value).SetVal(true)

	deleted, err := redisx.CuckooFilterDel(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("CuckooFilterDel failed: %v", err)
	}

	if !deleted {
		t.Errorf("CuckooFilterDel should return true for a value that was deleted")
	}

	// 10. Check if the deleted value exists
	mock.ExpectDo("CF.EXISTS", key, value).SetVal(false)

	exists, err = redisx.CuckooFilterExists(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("CuckooFilterExists failed for deleted value: %v", err)
	}

	if exists {
		t.Errorf("CuckooFilterExists should return false for a deleted value")
	}

	// 11. Delete a non-existent value
	mock.ExpectDo("CF.DEL", key, "non_existent_value").SetVal(false)

	deleted, err = redisx.CuckooFilterDel(ctx, rdb, key, "non_existent_value")
	if err != nil {
		t.Fatalf("CuckooFilterDel failed for non-existent value: %v", err)
	}

	if deleted {
		t.Errorf("CuckooFilterDel should return false for a non-existent value")
	}

	// Clean up after test
	rdb.Del(ctx, key)
}
