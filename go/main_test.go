package main

import (
    "testing"
    "reflect"
)


func TestFib2(t *testing.T){

    got := Fibonacci(2)
    want := []int{0,1,1}

    if !reflect.DeepEqual(got,want) {
        t.Errorf("got %+q, wanted %+q", got, want)
    }
}

func TestFib1(t *testing.T){

    got := Fibonacci(1)
    want := []int{0,1}

    if !reflect.DeepEqual(got,want) {
        t.Errorf("got %+q, wanted %+q", got, want)
    }
}

func TestFib3(t *testing.T){

    got := Fibonacci(3)
    want := []int{0,1,1,2}

    if !reflect.DeepEqual(got,want) {
        t.Errorf("got %+q, wanted %+q", got, want)
    }
}

func TestFib4(t *testing.T){

    got := Fibonacci(4)
    want := []int{0,1,1,2,3}

    if !reflect.DeepEqual(got,want) {
        t.Errorf("got %+q, wanted %+q", got, want)
    }
}

func TestFib5(t *testing.T){

    got := Fibonacci(5)
    want := []int{0,1,1,2,3,5}

    if !reflect.DeepEqual(got,want) {
        t.Errorf("got %+q, wanted %+q", got, want)
    }
}
