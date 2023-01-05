package main

import (
	"fmt"
	"learning/interface/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInsertUser(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockUser := mocks.NewMockFigure(mockCtl)
	gomock.InOrder(
		mockUser.EXPECT().Area().Return(1.5).Times(1),
		mockUser.EXPECT().Area().Return(5.5).Times(1),
	)

	area := getSquareArea(mockUser)
	fmt.Println(area)
	
	area = getSquareArea(mockUser)
	fmt.Println(area)

	if area != 5.0 {
		t.Fatalf("Expected id: %f, got: %f", 5.0, area)
	}
	
}