package impl

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/EugeneImbro/chat-backend/internal/repository"
	"github.com/EugeneImbro/chat-backend/internal/repository/mock"
	"github.com/EugeneImbro/chat-backend/internal/service"
)

func TestUserService_GetById(t *testing.T) {
	type mockBehavior func(m *repo_mock.MockRepository, id int32)

	tt := []struct {
		name     string
		mock     mockBehavior
		input    int32
		expected *service.User
		err error
	}{
		{
			name: "OK",
			mock: func(r *repo_mock.MockRepository, id int32) {
				r.EXPECT().GetUserByID(context.Background(), int32(1)).Return(&repository.User{Id: 1, NickName: "Richard Cheese"}, nil)
			},
			input:    1,
			expected: &service.User{Id: 1, NickName: "Richard Cheese"},
		},
		{
			name: "NOT FOUND",
			mock: func(r *repo_mock.MockRepository, id int32) {
				r.EXPECT().GetUserByID(context.Background(), int32(1)).Return(nil, repository.ErrNotFound)
			},
			input: 1,
			err:   service.ErrNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			r := repo_mock.NewMockRepository(c)
			tc.mock(r, tc.input)
			s := &us{repo: r}

			result, err := s.GetById(context.Background(), int32(1))

			assert.Equal(t, result, tc.expected)
			assert.Equal(t, err, tc.err)
		})
	}

}
