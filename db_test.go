package main

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UUID_InsertSelect(t *testing.T) {
	db, err := ConnectDB()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()

	_, err = db.Exec(ctx, "DELETE FROM test_table")
	require.NoError(t, err)

	testUUID := uuid.New()
	_, err = db.Exec(ctx, "INSERT INTO test_table (id) VALUES ($1)", testUUID)
	require.NoError(t, err)

	var retrievedUUID uuid.UUID
	err = db.QueryRow(ctx, "SELECT id FROM test_table LIMIT 1").Scan(&retrievedUUID)
	require.NoError(t, err)

	assert.Equal(t, testUUID, retrievedUUID)
}

func Test_UUID_Array_InsertSelect(t *testing.T) {
	db, err := ConnectDB()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	uuidSlice := []uuid.UUID{uuid1, uuid2}

	var uuidArray pgtype.UUIDArray
	uuidArray.Set(uuidSlice)

	var result []uuid.UUID
	err = db.QueryRow(ctx, "SELECT $1::uuid[]", uuidArray).Scan(&result)
	require.NoError(t, err)

	assert.ElementsMatch(t, uuidSlice, result)
}

func Test_UUID_String_Array_InsertSelect(t *testing.T) {
	db, err := ConnectDB()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()

	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuidSlice := []string{uuid1.String(), uuid2.String()}

	var result []string
	err = db.QueryRow(ctx, "SELECT $1::uuid[]", uuidSlice).Scan(&result)
	require.NoError(t, err)

	assert.ElementsMatch(t, uuidSlice, result)
}
