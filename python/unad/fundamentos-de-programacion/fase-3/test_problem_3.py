import pytest
from problem_3 import *

def test_count_elements():
    testcases = [
        ["3", 3, "f", 3, "h", "3"],
        [4, 1, 2, 5, 4, 3, 1, 4],
    ]
    expected_results = [
        [
            ("3", 2),
            (3, 2),
            ("f", 1),
            ("h", 1),
        ],
        [
            (4, 3),
            (1, 2),
            (2, 1),
            (5, 1),
            (3, 1),
        ],
    ]

    for testcase, expected in zip(testcases, expected_results):
        assert list(count_elements(testcase)) == expected

def test_mean():
    testcases = [
        [4, 1, 2, 5, 4, 3, 1, 4],
        [4, 1, 2, 5, 4, 3, 5],
    ]
    expected_results = [
        3.0,
        3.42857142,
    ]

    for testcase, expected in zip(testcases, expected_results):
        assert mean(testcase) == pytest.approx(expected)
