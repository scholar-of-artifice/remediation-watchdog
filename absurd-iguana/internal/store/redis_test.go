package store

// TODO: eventually move this
/*
func TestRedisStore_SaveEvent(t *testing.T) {
	// setup: connect to the local Redis
	addr := "127.0.0.1:6379" // move this to env var
	store := NewRedisStore(addr)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// health check
	if err := store.Client.Ping(ctx).Err(); err != nil {
		t.Skipf("Skipping integration test: Redis not available at %s. (Run `docker compose up -d` first)", addr)
	}
	// define test data
	testID := "test-event-001"
	testPayload := "integration-test-payload"
	// execute the test
	t.Run("successfully saves a string payload", func(t *testing.T) {
		err := store.SaveEvent(ctx, testID, testPayload)
		if err != nil {
			t.Fatalf("Failed to save event: %v", err)
		}
		// verification
		val, err := store.Client.Get(ctx, testID).Result()
		if err != nil {
			t.Errorf("Could not retrieve key from Redis: %v", err)
		}
		if val != testPayload {
			t.Errorf("Data mismatch: got %q, expected %q", val, testPayload)
		}
		// test teardown
		store.Client.Del(ctx, testID)
	})
}
*/
