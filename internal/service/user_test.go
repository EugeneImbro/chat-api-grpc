package service

import (
	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_GetById(t *testing.T) {
	type mockBehavior func(m *repo_mock.MockUser, id int32)

	tt := []struct {
		name     string
		mock     mockBehavior
		input    int32
		expected *model.User
	}{
		{
			name: "OK",
			mock: func(r *repo_mock.MockUser, id int32) {
				r.EXPECT().GetById(int32(1)).Return(&model.User{Id: 1, NickName: "Richard Cheese"}, nil)
			},
			input:    1,
			expected: &model.User{Id: 1, NickName: "Richard Cheese"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			r := repo_mock.NewMockUser(c)
			tc.mock(r, tc.input)
			s := &Service{User: r}

			result, _ := s.User.GetById(int32(1))

			assert.Equal(t, result, tc.expected)
		})
	}

}
