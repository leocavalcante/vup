#!/usr/bin/env python3

import subprocess
import os
import sys

def main():
    """
    Runs end-to-end tests for the vup command.
    """
    # Build the vup binary
    build_command = "go build -o vup ."
    try:
        subprocess.run(build_command, shell=True, check=True, capture_output=True)
    except subprocess.CalledProcessError as e:
        print(f"Error building vup: {e.stderr.decode()}")
        sys.exit(1)

    # Test cases
    test_cases = [
        {"name": "major-upgrade", "command": "./vup major 1.2.3", "expected_code": 0, "expected_output": "2.2.3"},
        {"name": "minor-upgrade", "command": "./vup minor 1.2.3", "expected_code": 0, "expected_output": "1.3.3"},
        {"name": "patch-upgrade", "command": "./vup patch 1.2.3", "expected_code": 0, "expected_output": "1.2.4"},
        {"name": "major-downgrade", "command": "./vup major -d 2.2.3", "expected_code": 0, "expected_output": "1.2.3"},
        {"name": "major-upgrade-value", "command": "./vup major -v 2 1.2.3", "expected_code": 0, "expected_output": "3.2.3"},
        {"name": "major-downgrade-value", "command": "./vup major -d -v 2 3.2.3", "expected_code": 0, "expected_output": "1.2.3"},
        {"name": "invalid-version", "command": "./vup major invalid", "expected_code": 1, "expected_output": "invalid semantic version string"},
    ]

    # Run tests
    all_passed = True
    for test in test_cases:
        result = subprocess.run(test["command"], shell=True, check=False, capture_output=True)
        actual_code = result.returncode
        if actual_code == 0:
            actual_output = result.stdout.decode().strip()
        else:
            actual_output = result.stderr.decode().strip()

        if actual_code == test["expected_code"] and test["expected_output"] in actual_output:
            print(f"PASS: {test['name']}")
        else:
            print(f"FAIL: {test['name']}")
            print(f"  Command: {test['command']}")
            print(f"  Expected code: {test['expected_code']}, got: {actual_code}")
            print(f"  Expected output: {test['expected_output']}, got: {actual_output}")
            all_passed = False

    # Clean up
    os.remove("vup")

    if not all_passed:
        sys.exit(1)

if __name__ == "__main__":
    main()
