package main

import (
	"fmt"
	"testing"
)

func TestMarkAndToys1(t *testing.T) {
	k := int32(50)
	prices := []int32{1, 12, 5, 111, 200, 1000, 10}

	numOfToys := maximumToys(prices, k)

	fmt.Println(numOfToys)
	if numOfToys != 4 {
		t.Errorf("got %d instead of 4", numOfToys)
	}
}

func TestMarkAndToys2(t *testing.T) {
	k := int32(1)
	prices := []int32{1, 12, 5, 111, 200, 1000, 10}

	numOfToys := maximumToys(prices, k)

	fmt.Println(numOfToys)
	if numOfToys != 1 {
		t.Errorf("got %d instead of 1", numOfToys)
	}
}

func TestMarkAndToys3(t *testing.T) {
	k := int32(1)
	prices := []int32{4, 12, 5, 111, 200, 1000, 10}

	numOfToys := maximumToys(prices, k)

	fmt.Println(numOfToys)
	if numOfToys != 0 {
		t.Errorf("got %d instead of 0", numOfToys)
	}
}

func TestMarkAndToys4(t *testing.T) {
	k := int32(10000)
	prices := []int32{4, 12, 5, 111, 200, 1000, 10}

	numOfToys := maximumToys(prices, k)

	fmt.Println(numOfToys)
	if numOfToys != 7 {
		t.Errorf("got %d instead of 7", numOfToys)
	}
}

func TestMarkAndToys5(t *testing.T) {
	k := int32(7)
	prices := []int32{1, 2, 3, 4}

	numOfToys := maximumToys(prices, k)

	fmt.Println(numOfToys)
	if numOfToys != 3 {
		t.Errorf("got %d instead of 3", numOfToys)
	}
}
