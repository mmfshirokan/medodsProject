package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mmfshirokan/medodsProject/internal/model"
	mocks "github.com/mmfshirokan/medodsProject/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestAdd(t *testing.T) {
	testTable := []struct {
		input  model.RefreshToken
		addErr error
	}{
		{
			input: model.RefreshToken{
				ID:         uuid.New(),
				UserID:     uuid.New(),
				Expiration: time.Now(),
			},
			addErr: nil,
		},
		{
			input: model.RefreshToken{
				ID:         uuid.New(),
				UserID:     uuid.New(),
				Expiration: time.Now(),
			},
			addErr: nil,
		},
		{
			input: model.RefreshToken{
				ID:         uuid.New(),
				UserID:     uuid.New(),
				Expiration: time.Now(),
			},
			addErr: nil,
		},
	}

	rp := mocks.NewRepository(t)
	sr := New(rp)

	for _, test := range testTable {
		cl := rp.EXPECT().Add(mock.Anything, mock.Anything).Return(test.addErr)

		actual, err := sr.Add(context.Background(), test.input)

		assert.Nil(t, err)
		err = bcrypt.CompareHashAndPassword([]byte(actual.Hash), returnRftHashKey(actual.ID.String()))
		assert.Nil(t, err)

		rp.AssertExpectations(t)
		cl.Unset()

	}
}

func TestValidateRFT(t *testing.T) {
	testTable := []struct {
		input uuid.UUID
		hash  string

		HasGetErr bool
		result    bool
	}{
		{
			HasGetErr: false,
			result:    true,
		},
		{
			HasGetErr: false,
			result:    true,
		},
		{
			HasGetErr: true,
			result:    false,
		},
	}

	for i, test := range testTable {
		if !test.HasGetErr {
			id := uuid.New()

			hash, err := bcrypt.GenerateFromPassword(returnRftHashKey(id.String()), bcrypt.MinCost)
			if err != nil {
				t.FailNow()
			}

			testTable[i].input = id
			testTable[i].hash = string(hash)
		}
	}

	rp := mocks.NewRepository(t)
	sr := New(rp)

	for _, test := range testTable {

		var cl *mocks.Repository_GetHash_Call
		if test.HasGetErr {
			cl = rp.EXPECT().GetHash(mock.Anything, mock.Anything).Return("", errors.New("some error"))
		} else {
			cl = rp.EXPECT().GetHash(mock.Anything, test.input).Return(test.hash, nil)
		}

		res, _ := sr.ValidateRFT(context.Background(), test.input)

		assert.Equal(t, test.result, res)
		rp.AssertExpectations(t)
		cl.Unset()
	}
}

func TestValidatePwd(t *testing.T) {
	testTable := []struct {
		input uuid.UUID
		pwd   string
		hash  string

		HasPwdGetErr bool
		result       bool
	}{
		{
			pwd:          "test1",
			HasPwdGetErr: false,
			result:       true,
		},
		{
			pwd:          "test2",
			HasPwdGetErr: false,
			result:       true,
		},
		{
			pwd:          "test3",
			HasPwdGetErr: true,
			result:       false,
		},
	}

	for i, test := range testTable {
		if !test.HasPwdGetErr {
			id := uuid.New()

			hash, err := bcrypt.GenerateFromPassword(returnPwdHashKey(test.pwd), bcrypt.MinCost)
			if err != nil {
				t.FailNow()
			}

			testTable[i].input = id
			testTable[i].hash = string(hash)
		}
	}

	rp := mocks.NewRepository(t)
	sr := New(rp)

	for _, test := range testTable {

		var cl *mocks.Repository_GetPwd_Call
		if test.HasPwdGetErr {
			cl = rp.EXPECT().GetPwd(mock.Anything, mock.Anything).Return("", errors.New("some error"))
		} else {
			cl = rp.EXPECT().GetPwd(mock.Anything, test.input).Return(test.hash, nil)
		}

		res, _ := sr.ValidatePWD(context.Background(), test.input, test.pwd)

		assert.Equal(t, test.result, res)
		rp.AssertExpectations(t)
		cl.Unset()
	}
}

// AddUsr is test function (no test for test func)
