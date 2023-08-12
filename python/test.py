import pytest
from main import *

def test_fib_1():
    assert fibonacci(1) == [0]

def test_fib_2():
    assert fibonacci(2) == [0,1]

def test_fib_3():
    assert fibonacci(3) == [0,1,1]

def test_fib_4():
    assert fibonacci(4) == [0,1,1,2]

def test_fib_5():
    assert fibonacci(5) == [0,1,1,2,3]

def test_fib_6():
    assert fibonacci(6) == [0,1,1,2,3,5]

def test_fib_7():
    assert fibonacci(7) == [0,1,1,2,3,5,8]
