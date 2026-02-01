#!/usr/bin/env python3
"""
Piscine Test Results Formatter
Formats Go test output in a beautiful, readable way
"""

import sys
import re
from datetime import datetime

# ANSI color codes
COLORS = {
    'GREEN': '\033[92m',
    'RED': '\033[91m',
    'YELLOW': '\033[93m',
    'BLUE': '\033[94m',
    'MAGENTA': '\033[95m',
    'CYAN': '\033[96m',
    'WHITE': '\033[97m',
    'BOLD': '\033[1m',
    'RESET': '\033[0m'
}

def color(text, color_name):
    """Add color to text"""
    return f"{COLORS.get(color_name, '')}{text}{COLORS['RESET']}"

def print_header():
    """Print a nice header"""
    print("\n" + "="*70)
    print(color("🧪 PISCINE GO TEST RESULTS", 'BOLD'))
    print(color(f"📅 {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}", 'CYAN'))
    print("="*70 + "\n")

def print_summary(total, passed, failed, skipped):
    """Print test summary"""
    print("\n" + "="*70)
    print(color("📊 TEST SUMMARY", 'BOLD'))
    print("="*70)
    
    print(f"\n{color('Total Tests:', 'BOLD')} {total}")
    print(f"{color('✅ Passed:', 'GREEN')} {passed}")
    
    if failed > 0:
        print(f"{color('❌ Failed:', 'RED')} {failed}")
    
    if skipped > 0:
        print(f"{color('⏭️  Skipped:', 'YELLOW')} {skipped}")
    
    # Calculate success rate
    if total > 0:
        success_rate = (passed / total) * 100
        if success_rate == 100:
            emoji = "🎉"
            msg_color = 'GREEN'
        elif success_rate >= 80:
            emoji = "👍"
            msg_color = 'YELLOW'
        else:
            emoji = "⚠️"
            msg_color = 'RED'
        
        print(f"\n{emoji} {color(f'Success Rate: {success_rate:.1f}%', msg_color)}")
    
    print("\n" + "="*70 + "\n")

def format_test_output(line):
    """Format a single test output line"""
    # Test running
    if "=== RUN" in line:
        test_name = line.split("RUN")[-1].strip()
        return color(f"🏃 Running: {test_name}", 'BLUE')
    
    # Test passed
    elif "--- PASS:" in line:
        test_name = line.split("PASS:")[-1].split("(")[0].strip()
        time_match = re.search(r'\(([^)]+)\)', line)
        time_str = time_match.group(1) if time_match else "0.00s"
        return color(f"✅ PASSED: {test_name} ({time_str})", 'GREEN')
    
    # Test failed
    elif "--- FAIL:" in line:
        test_name = line.split("FAIL:")[-1].split("(")[0].strip()
        time_match = re.search(r'\(([^)]+)\)', line)
        time_str = time_match.group(1) if time_match else "0.00s"
        return color(f"❌ FAILED: {test_name} ({time_str})", 'RED')
    
    # Test skipped
    elif "--- SKIP:" in line:
        test_name = line.split("SKIP:")[-1].strip()
        return color(f"⏭️  SKIPPED: {test_name}", 'YELLOW')
    
    # Error message
    elif line.strip().startswith("Error:") or "error:" in line.lower():
        return color(f"⚠️  {line.strip()}", 'RED')
    
    # Coverage
    elif "coverage:" in line:
        return color(f"📊 {line.strip()}", 'CYAN')
    
    # Benchmark
    elif "Benchmark" in line and "ns/op" in line:
        return color(f"🏎️  {line.strip()}", 'MAGENTA')
    
    return line

def main():
    """Main function to process test output"""
    print_header()
    
    total = 0
    passed = 0
    failed = 0
    skipped = 0
    
    for line in sys.stdin:
        # Count tests
        if "--- PASS:" in line:
            passed += 1
            total += 1
        elif "--- FAIL:" in line:
            failed += 1
            total += 1
        elif "--- SKIP:" in line:
            skipped += 1
            total += 1
        
        # Format and print
        formatted = format_test_output(line)
        print(formatted, end='')
    
    print_summary(total, passed, failed, skipped)

if __name__ == "__main__":
    main()