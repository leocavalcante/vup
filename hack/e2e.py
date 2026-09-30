#!/usr/bin/env python3

import os
import subprocess
import sys
import tempfile


def main():
    """Run end-to-end tests for the vup command."""
    with tempfile.TemporaryDirectory() as tmp:
        binary = os.path.join(tmp, "vup")
        build = subprocess.run(
            ["go", "build", "-o", binary, "."],
            capture_output=True,
        )
        if build.returncode != 0:
            sys.stderr.write(f"Error building vup: {build.stderr.decode()}\n")
            sys.exit(1)

        test_cases = [
            {
                "name": "major-upgrade",
                "args": ["major", "1.2.3"],
                "expected_code": 0,
                "expected_output": "2.2.3",
            },
            {
                "name": "minor-upgrade",
                "args": ["minor", "1.2.3"],
                "expected_code": 0,
                "expected_output": "1.3.3",
            },
            {
                "name": "patch-upgrade",
                "args": ["patch", "1.2.3"],
                "expected_code": 0,
                "expected_output": "1.2.4",
            },
            {
                "name": "major-downgrade",
                "args": ["major", "-d", "2.2.3"],
                "expected_code": 0,
                "expected_output": "1.2.3",
            },
            {
                "name": "major-upgrade-value",
                "args": ["major", "-v", "2", "1.2.3"],
                "expected_code": 0,
                "expected_output": "3.2.3",
            },
            {
                "name": "major-downgrade-value",
                "args": ["major", "-d", "-v", "2", "3.2.3"],
                "expected_code": 0,
                "expected_output": "1.2.3",
            },
            {
                "name": "invalid-version",
                "args": ["major", "invalid"],
                "expected_code": 1,
                "expected_output": "invalid semantic version string",
            },
        ]

        all_passed = True
        for test in test_cases:
            result = subprocess.run([binary, *test["args"]], capture_output=True)
            actual_code = result.returncode
            if actual_code == 0:
                actual_output = result.stdout.decode().strip()
            else:
                actual_output = result.stderr.decode().strip()

            if (
                actual_code == test["expected_code"]
                and test["expected_output"] in actual_output
            ):
                print(f"PASS: {test['name']}")
            else:
                print(f"FAIL: {test['name']}")
                print(f"  Args: {test['args']}")
                print(f"  Expected code: {test['expected_code']}, got: {actual_code}")
                print(
                    f"  Expected output: {test['expected_output']}, got: {actual_output}"
                )
                all_passed = False

    if not all_passed:
        sys.exit(1)


if __name__ == "__main__":
    main()
